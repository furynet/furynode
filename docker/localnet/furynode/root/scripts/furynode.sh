#!/bin/sh
#
# Furynet: Furynode Genesis Entrypoint.
#

#
# Configure the node.
#
setup() {
  furygen node create "$CHAINNET" "$MONIKER" "$MNEMONIC" --bind-ip-address "$BIND_IP_ADDRESS" --standalone --keyring-backend test
}

#
# Run the node under cosmovisor.
#
run() {
  furynoded start --rpc.laddr=tcp://0.0.0.0:26657 --minimum-gas-prices="$GAS_PRICE"
}

setup
run
