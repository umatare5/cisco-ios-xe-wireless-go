package controller

// SaveConfigRPCOutput is the reply of cisco-ia:save-config, whose declaration is byte-identical on
// cisco-ia@2021-11-01 and cisco-ia@2024-07-01.
//
// Output is a pointer because an answer with no body decodes into a zero value of this type, and a
// non-pointer field would make that zero indistinguishable from an account the controller sent.
type SaveConfigRPCOutput struct {
	Output *SaveConfigRPCOutputData `json:"cisco-ia:output"` // Controller's reply container
}

// SaveConfigRPCOutputData holds the one leaf of grouping cisco-ia-output, which all eight cisco-ia
// RPCs answer with.
type SaveConfigRPCOutputData struct {
	Result string `json:"result"` // Controller's account of the save (Live: IOS-XE 17.12.8, 17.15.6, 17.18.4a)
}
