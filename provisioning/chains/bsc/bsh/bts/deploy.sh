#!/bin/bash

source ../../evm/bsh/bts/deploy.sh

function bsc_deploy_bts() {
    evm_deploy_bts $args
}

function bsc_extract_btscore_address() {
    evm_extract_btscore_address $args
}

function bsc_extract_btsperiphery_address() {
    evm_extract_btsperiphery_address $args
}