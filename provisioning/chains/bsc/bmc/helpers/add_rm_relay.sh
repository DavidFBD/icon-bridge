#!/bin/bash

source ../../evm/bmc/helpers/add_rm_relay.sh

function bsc_add_relay() {
    evm_add_relay $args
}

function bsc_remove_relay() {
    evm_remove_relay $args
}