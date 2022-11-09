#!/bin/bash

function evm_deploy_bmc() {
    # Input: bmc contract folder
    # Check: that the folder-name is bmc
    # clear existing artifacts (if any) 
    # compile contracts
    # truffle migrate to deploy contract (catch error and retry)
    # return success/failure
}

function evm_extract_bmc_management_address() {
    # Input: bmc contract folder
    # Check: if deployment was successful
    # Return BMC Management from truffle deployment artifacts
}

function evm_extract_bmc_periphery_address() {
    # Input: bmc contract folder
    # Check: if deployment was successful
    # Return BMC Periphery from truffle deployment artifacts
}