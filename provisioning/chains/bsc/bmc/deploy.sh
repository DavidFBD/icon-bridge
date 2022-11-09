#!/bin/bash

source ../../evm/bmc/deploy.sh

function bsc_deploy_bmc() {
    evm_deploy_bmc $args
}

function bsc_extract_bmc_management_address() {
    bsc_extract_bmc_management_address $args
}

function bsc_extract_bmc_periphery_address() {
    evm_extract_bmc_management_address $args
}