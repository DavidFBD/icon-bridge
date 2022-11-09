#!/bin/bash

source ../../evm/bmc/helpers/add_rm_bsh.sh

function bsc_add_service() {
    evm_add_service $args
}

function bsc_remove_service() {
    evm_remove_service $args
}