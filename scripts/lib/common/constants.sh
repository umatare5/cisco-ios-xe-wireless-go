#!/usr/bin/env bash

# Cisco WNC Script Constants Library
# Common constants and default values used across all WNC scripts

# Function to get default protocol
get_default_protocol() {
    echo "https"
}

# Function to get default controller
get_default_controller() {
    echo "wnc1.example.internal"
}

# Function to get default output format
get_default_output_format() {
    echo "pretty"
}

# Function to get default timeout
get_default_timeout() {
    echo "30"
}

# Function to get RESTCONF data path
get_restconf_data_path() {
    echo "/restconf/data"
}

# Function to get RESTCONF modules path
get_restconf_modules_path() {
    echo "/restconf/tailf/modules"
}

# Function to get accept header for YANG data
get_yang_accept_header() {
    echo "application/yang-data+json"
}

# Function to get network error message
get_network_error_message() {
    echo "Failed to connect to the controller"
}

# Function to get standard exit codes
get_exit_success() {
    echo "0"
}

get_exit_error() {
    echo "1"
}

get_exit_invalid_args() {
    echo "2"
}

get_exit_auth_error() {
    echo "3"
}

get_exit_network_error() {
    echo "4"
}
