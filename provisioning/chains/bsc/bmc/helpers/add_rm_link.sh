#!/bin/bash

source ../../evm/bmc/helpers/add_rm_bsh.sh

function bsc_add_link() {
    source evm_add_link $args
}

function bsc_remove_link() {
    source evm_remove_link $args
}