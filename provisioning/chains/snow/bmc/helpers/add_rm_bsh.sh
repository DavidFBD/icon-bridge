#!/bin/bash

source ../../evm/bmc/helpers/add_rm_bsh.sh

function snow_add_service() {
    evm_add_service $args
}

function snow_remove_service() {
    evm_remove_service $args
}