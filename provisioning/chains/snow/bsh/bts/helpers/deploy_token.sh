#!/bin/bash

source ../../evm/bsh/bts/helpers/deploy_token.sh

function snow_deploy_token() {
    evm_deploy_token $args
}