#!/bin/bash

source ../../evm/bsh/bts/helpers/register_token.sh

function bsc_register_native_token() {
    evm_register_native_token $args
}

function bsc_register_wrapped_coin() {
    evm_register_wrapped_coin $args
}

function bsc_get_token_address() {
    evm_get_token_address $args
}