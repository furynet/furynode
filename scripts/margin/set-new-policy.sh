#!/usr/bin/env bash

set -x

furynoded tx clp pmtp-params \
  --pmtp_start=22811 \
  --pmtp_end=224410 \
  --epochLength=14400 \
  --rGov=0.05 \
  --from=$FURY_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --node ${FURYNODE_NODE} \
  --chain-id=$FURYNODE_CHAIN_ID \
  --broadcast-mode=block \
  -y