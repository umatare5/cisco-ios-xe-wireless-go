package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// Generic HTTP Operation Functions for Service Layer
// These functions provide a consistent interface for HTTP operations across all services.

// Get is a generic helper reducing boilerplate in service GET methods.
//
// The response is validated before it is decoded: a body whose sole top-level key is not the
// node the endpoint asked for, or for which T declares no field, is an error rather than a
// zero value of T. A read answered with no body is a successful read of a node that holds
// nothing: T comes back at its zero value, which for a list node is a nil slice.
func Get[T any](ctx context.Context, c *Client, endpoint string, opts ...GetOption) (*T, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	endpoint = applyGetOptions(endpoint, opts)

	body, err := c.do(ctx, http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		var out T
		return &out, nil
	}

	return decodeSoleKey[T](body, endpoint)
}

// GetRaw performs a GET and returns the response body as the controller sent it.
//
// This is the seam behind the root client's GetData: it applies GetOption values the
// same way Get does, and leaves the envelope check to the caller.
//
// A read answered with no body is a successful read of a node that holds nothing: the
// returned slice is non-nil and empty, and the error is nil.
func GetRaw(ctx context.Context, c *Client, endpoint string, opts ...GetOption) ([]byte, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	return c.do(ctx, http.MethodGet, applyGetOptions(endpoint, opts))
}

// Faults in an untyped request, reported before anything is sent.
var (
	// errEmptyMethod reports the one method value net/http reinterprets rather than rejects: it
	// substitutes GET, which would read where a write was asked for.
	errEmptyMethod = errors.New("HTTP method cannot be empty")
	// errInvalidJSONPayload reports payload bytes this package would send under a JSON content
	// type without them being JSON. Checking here keeps the message ours rather than the
	// encoder's, which names its own internal types.
	errInvalidJSONPayload = errors.New("payload bytes are not a JSON document")
)

// EditRaw performs a request with the caller's method on a RESTCONF data path and returns the
// response body as the controller sent it.
//
// This is the seam behind the root client's PostData, PutData, PatchData and DeleteData. A nil
// payload sends no body and no Content-Type, which is what DELETE wants.
func EditRaw(ctx context.Context, c *Client, method, endpoint string, payload any) ([]byte, error) {
	body, err := prepareUntypedRequest(c, method, payload)
	if err != nil {
		return nil, err
	}

	return c.doWithPayload(ctx, method, endpoint, body)
}

// CallRPCRaw posts an RPC input to a RESTCONF operations path and returns the output body as the
// controller sent it.
//
// This is the seam behind the root client's PostRPC. It differs from EditRaw in the RESTCONF root
// the path resolves against, /restconf/operations rather than /restconf/data, and in fixing the
// method to POST.
func CallRPCRaw(ctx context.Context, c *Client, rpcEndpoint string, payload any) ([]byte, error) {
	body, err := prepareUntypedRequest(c, http.MethodPost, payload)
	if err != nil {
		return nil, err
	}

	return c.doRPC(ctx, rpcEndpoint, body)
}

// RequestRaw performs a request with the caller's method on the path the caller wrote, routed to
// the RESTCONF root that path names.
//
// This is the seam behind the root client's Request, which exists because the controller's schema
// moves between releases: a node or an operation this package has no accessor and no verb method
// for still has to be reachable without waiting for a release. A path already under
// /restconf/operations goes to the operations root; anything else goes to the data root, which
// passes a /restconf/data-prefixed path through and prefixes a bare one.
func RequestRaw(ctx context.Context, c *Client, method, path string, payload any) ([]byte, error) {
	if strings.HasPrefix(path, routes.RESTCONFOperationsPath) {
		body, err := prepareUntypedRequest(c, method, payload)
		if err != nil {
			return nil, err
		}

		return c.doRPC(ctx, path, body)
	}

	return EditRaw(ctx, c, method, path, payload)
}

// prepareUntypedRequest checks what every untyped route checks and returns the payload in the shape
// the request builder has to marshal.
//
// Bytes are carried as json.RawMessage rather than marshaled as a value, because a []byte handed to
// encoding/json encodes as a base64 string: a body read with GetData could not be edited and sent
// back. Their validity is checked here, so the JSON content type this package sets is never a lie.
func prepareUntypedRequest(c *Client, method string, payload any) (any, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}
	if method == "" {
		return nil, errEmptyMethod
	}

	switch body := payload.(type) {
	case nil:
		return nil, nil
	case []byte:
		return rawJSONPayload(body)
	case json.RawMessage:
		return rawJSONPayload(body)
	}

	return payload, nil
}

// rawJSONPayload carries bytes to the wire as written, refusing what is not JSON.
func rawJSONPayload(body []byte) (any, error) {
	if len(body) == 0 {
		return nil, nil
	}
	if !json.Valid(body) {
		return nil, errInvalidJSONPayload
	}

	return json.RawMessage(body), nil
}

// decode unmarshals a write response body into T, treating an empty body as a zero T.
//
// A write is answered with the changed node, an RPC output, or nothing at all, so there is no
// single top-level key to hold against the endpoint; the read path validates in decodeSoleKey.
func decode[T any](body []byte) (*T, error) {
	var out T
	if len(body) == 0 {
		return &out, nil
	}

	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &out, nil
}

// Post is a generic helper for sending POST requests with payload.
func Post[T any](ctx context.Context, c *Client, endpoint string, payload any) (*T, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	body, err := c.doWithPayload(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return nil, err
	}

	return decode[T](body)
}

// PostVoid is a generic helper for POST operations without expecting a response body.
func PostVoid(ctx context.Context, c *Client, endpoint string, payload any) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.doWithPayload(ctx, http.MethodPost, endpoint, payload)
	return err
}

// PostRPCVoid is a generic helper for RPC POST operations without expecting a response body.
func PostRPCVoid(ctx context.Context, c *Client, rpcEndpoint string, payload any) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.doRPC(ctx, rpcEndpoint, payload)
	return err
}

// Put is a generic helper for PUT operations that expect a response body.
func Put[T any](ctx context.Context, c *Client, endpoint string, payload any) (*T, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	body, err := c.doWithPayload(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return nil, err
	}

	return decode[T](body)
}

// PutVoid is a generic helper for PUT operations without expecting a response body.
func PutVoid(ctx context.Context, c *Client, endpoint string, payload any) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.doWithPayload(ctx, http.MethodPut, endpoint, payload)
	return err
}

// Patch is a generic helper for PATCH operations that expect a response body.
func Patch[T any](ctx context.Context, c *Client, endpoint string, payload any) (*T, error) {
	if c == nil {
		return nil, errors.New(ierrors.ErrClientNil)
	}

	body, err := c.doWithPayload(ctx, http.MethodPatch, endpoint, payload)
	if err != nil {
		return nil, err
	}

	return decode[T](body)
}

// PatchVoid is a generic helper for PATCH operations without expecting a response body.
func PatchVoid(ctx context.Context, c *Client, endpoint string, payload any) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.doWithPayload(ctx, http.MethodPatch, endpoint, payload)
	return err
}

// Delete is a generic helper for DELETE operations without expecting a response body.
func Delete(ctx context.Context, c *Client, endpoint string) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	_, err := c.do(ctx, http.MethodDelete, endpoint)
	return err
}
