#!/bin/bash

source ../../evm/bmc/deploy.sh

function snow_deploy_bmc() {
    evm_deploy_bmc $args
}

function snow_extract_bmc_management_address() {
    snow_extract_bmc_management_address $args
}

function snow_extract_bmc_periphery_address() {
    evm_extract_bmc_management_address $args
}