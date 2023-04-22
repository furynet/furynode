#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean install
furynoded init test --chain-id=localnet

echo "Generating deterministic account - fury"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | furynoded keys add fury --recover

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | furynoded keys add akasha --recover


furynoded keys add mkey --multisig fury,akasha --multisig-threshold 2
furynoded add-genesis-account $(furynoded keys show fury -a) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink
furynoded add-genesis-account $(furynoded keys show akasha -a) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink
furynoded add-genesis-account $(furynoded keys show mkey -a) 500000000000000000000000fury

furynoded add-genesis-clp-admin $(furynoded keys show fury -a)
furynoded add-genesis-clp-admin $(furynoded keys show akasha -a)

furynoded add-genesis-validators $(furynoded keys show fury -a --bech val)

furynoded gentx fury 1000000000000000000000000stake --keyring-backend test

echo "Collecting genesis txs..."
furynoded collect-gentxs

echo "Validating genesis file..."
furynoded validate-genesis


mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/release-20210414000000/bin

cp $GOPATH/bin/old/furynoded $DAEMON_HOME/cosmovisor/genesis/bin
cp $GOPATH/bin/furynoded $DAEMON_HOME/cosmovisor/upgrades/release-20210414000000/bin/

#contents="$(jq '.gov.voting_params.voting_period = 10' $DAEMON_HOME/config/genesis.json)" && \
#echo "${contents}" > $DAEMON_HOME/config/genesis.json
