#!/bin/bash
. ../credentials.sh

# if we don't sleep there are issues
sleep 10

echo "create liquidity pool from catk:fury"


# nativeAmount 10000000 catk
# externalAmount 10000000 fury
furynoded tx clp create-pool \
 --chain-id=furynet-local \
 --keyring-backend=test \
 --from akasha \
 --symbol catk \
 --fees 100000fury \
 --nativeAmount   10000000000000000000000000 \
 --externalAmount 10000000000000000000000000 \
 --yes

# if we don't sleep there are issues
sleep 5

echo "create liquidity pool from cbtk:fury"
# create liquidity pool from cbtk:fury
# nativeAmount 10000000 cbtk
# externalAmount 10000000 fury
furynoded tx clp create-pool \
 --chain-id=furynet-local \
 --keyring-backend=test \
 --from akasha \
 --symbol cbtk \
 --fees 100000fury \
 --nativeAmount   10000000000000000000000000 \
 --externalAmount 10000000000000000000000000 \
 --yes

# should now be able to swap from catk:cbtk

sleep 5

echo "create liquidity pool from ceth:fury"
# nativeAmount 8300 ceth
# externalAmount 10000000 fury
furynoded tx clp create-pool \
 --chain-id=furynet-local \
 --keyring-backend=test \
 --from akasha \
 --symbol ceth \
 --fees 100000fury \
 --nativeAmount   10000000000000000000000000 \
 --externalAmount 8300000000000000000000 \
 --yes

 # should now be able to swap from x:ceth

sleep 5

echo "create liquidity pool from cusdc:fury"
furynoded tx clp create-pool \
 --chain-id=furynet-local \
 --keyring-backend=test \
 --from akasha \
 --symbol cusdc \
 --fees 100000fury \
 --nativeAmount   10000000000000000000000000 \
 --externalAmount 10000000000000000000000000 \
 --yes

sleep 5

echo "create liquidity pool from clink:fury"
furynoded tx clp create-pool \
 --chain-id=furynet-local \
 --keyring-backend=test \
 --from akasha \
 --symbol clink \
 --fees 100000fury \
 --nativeAmount   10000000000000000000000000 \
 --externalAmount 588235000000000000000000 \
 --yes

sleep 5

echo "create liquidity pool from ctest:fury"
furynoded tx clp create-pool \
 --chain-id=furynet-local \
 --keyring-backend=test \
 --from akasha \
 --symbol ctest \
 --fees 100000fury \
 --nativeAmount   10000000000000000000000000 \
 --externalAmount 10000000000000 \
 --yes
