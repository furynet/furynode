#!/bin/zsh

echo "initializing furynoded and doing initial IBC transfers"
scripts/init-multichain.sh > /dev/null 2>&1
scripts/do_ibc_transfers.sh > /dev/null 2>&1
echo "killing furynoded"
ps auxww | grep furynode-1 | egrep -v grep | awk '{ print $2 }' | xargs kill
sleep 60
echo "killing hermes"
killall hermes

# start hermes
echo "starting hermes"
hermes start > hermes.log 2>&1 &

echo "Sleeping to let hermes boot"
sleep 10

echo "second round of ibc transfers"
scripts/do_ibc_transfers.sh

echo "sleeping to let the localnet-1 expire"
sleep 850

echo "restarting localnet-1"
furynoded start --home ~/.furynode-1 --p2p.laddr 0.0.0.0:27655 --grpc.address 0.0.0.0:9090 --grpc-web.address 0.0.0.0:9093 --address tcp://0.0.0.0:27659 --rpc.laddr tcp://127.0.0.1:27665 >> abci_1.log 2>&1 &
sleep 10

echo "bouncing hermes"
killall hermes

# start hermes
echo "starting hermes"
hermes start > hermes.log 2>&1 &

echo "Sleeping to let hermes boot"
sleep 10

echo "Checking channels"
hermes query packet unreceived-packets localnet-1 transfer channel-0
hermes query packet unreceived-packets localnet-1 transfer channel-1
hermes query packet unreceived-packets localnet-2 transfer channel-0
hermes query packet unreceived-packets localnet-2 transfer channel-1
hermes query packet unreceived-packets localnet-3 transfer channel-0
hermes query packet unreceived-packets localnet-3 transfer channel-1

echo "Creating channels to clone"
scripts/create_clonable_hermes_channels.sh > /dev/null 2>&1

echo "Current Fury balances"
echo "localnet-1"
furynoded q bank balances $(furynoded keys show fury -a --keyring-backend=test --home ~/.furynode-1) --node tcp://127.0.0.1:27665
echo ""
echo "localnet-2"
furynoded q bank balances $(furynoded keys show fury -a --keyring-backend=test --home ~/.furynode-1) --node tcp://127.0.0.1:27666
echo ""
echo "localnet-3"
furynoded q bank balances $(furynoded keys show fury -a --keyring-backend=test --home ~/.furynode-1) --node tcp://127.0.0.1:27667
echo ""

echo "Current Akasha balances"
echo "localnet-1"
furynoded q bank balances $(furynoded keys show akasha -a --keyring-backend=test --home ~/.furynode-1) --node tcp://127.0.0.1:27665
echo ""
echo "localnet-2"
furynoded q bank balances $(furynoded keys show akasha -a --keyring-backend=test --home ~/.furynode-1) --node tcp://127.0.0.1:27666
echo ""
echo "localnet-3"
furynoded q bank balances $(furynoded keys show akasha -a --keyring-backend=test --home ~/.furynode-1) --node tcp://127.0.0.1:27667
echo ""

echo "voting on proposals"
scripts/vote.sh > /dev/null 2>&1
scripts/vote_localnet_1.sh > /dev/null 2>&1
