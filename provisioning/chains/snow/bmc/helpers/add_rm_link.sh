#!/bin/bash

source ../../evm/bmc/helpers/add_rm_bsh.sh

function snow_add_link() {
    source evm_add_link $args
}

function snow_remove_link() {
    source evm_remove_link $args
}