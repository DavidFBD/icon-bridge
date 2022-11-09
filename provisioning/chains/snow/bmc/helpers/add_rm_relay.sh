#!/bin/bash

source ../../evm/bmc/helpers/add_rm_relay.sh

function snow_add_relay() {
    evm_add_relay $args
}

function snow_remove_relay() {
    evm_remove_relay $args
}