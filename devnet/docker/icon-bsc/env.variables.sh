#!/bin/sh
export CONFIG_DIR=${CONFIG_DIR:-${ICONBRIDGE_CONFIG_DIR}}
export CONTRACTS_DIR=${CONTRACTS_DIR:-${ICONBRIDGE_CONTRACTS_DIR}}
export SCRIPTS_DIR=${SCRIPTS_DIR:-${ICONBRIDGE_SCRIPTS_DIR}}

# GOLOOP Config
GOLOOPCHAIN=${GOLOOPCHAIN:-'goloop'}
export GOLOOP_RPC_URI=$ICON_ENDPOINT
export GOLOOP_RPC_ADMIN_URI=http://$GOLOOPCHAIN/admin/system
export GOLOOP_RPC_KEY_STORE=${ICON_KEY_STORE:-"$CONFIG_DIR/$GOLOOPCHAIN.keystore.json"}
export GOLOOP_RPC_KEY_SECRET=${ICON_SECRET:-"$CONFIG_DIR/$GOLOOPCHAIN.keysecret"}
export GOLOOP_RPC_STEP_LIMIT=${GOLOOP_RPC_STEP_LIMIT:-5000000000}
export GOLOOP_CHAINSCORE=cx000000000000000000000000000000000000000
export GOLOOP_RPC_NID=${GOLOOP_RPC_NID:-$(cat $CONFIG_DIR/nid.icon)}

#BSC Config
export BSC_NID=${BSC_NID:-'97'}
export BSC_BMC_NET=${BSC_BMC_NET:-'0x61.bsc'}
export BSC_RPC_URI=${BSC_ENDPOINT:-'http://binancesmartchain:8545'}

#Token Config
export TOKEN_NAME=ETH
export TOKEN_SYM=ETH
export TOKEN_SUPPLY=0x186A0
export TOKEN_DECIMALS=0x12
export SVC_NAME=TokenBSH
