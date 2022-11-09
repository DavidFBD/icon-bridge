#!/bin/bash

function icon_deploy_bmc() {
    # Input: bmc contract folder
    # Check: that the folder-name is bmc
    # clear existing artifacts (if any) 
    # compile contracts
    # truffle migrate to deploy contract (catch error and retry)
    # extract and return bmc address
}