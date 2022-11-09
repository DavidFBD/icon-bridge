#!/bin/bash

source ../../evm/bsh/bts/deploy.sh

function snow_deploy_bts() {
    evm_deploy_bts $args
}

function snow_extract_btscore_address() {
    evm_extract_btscore_address $args
}

function snow_extract_btsperiphery_address() {
    evm_extract_btsperiphery_address $args
}