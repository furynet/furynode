#!/usr/bin/env bash


rm -rf ~/.furynoded
rm -rf ~/.furynoded
rm -rf furynode.log
rm -rf testlog.log

cd "$(dirname "$0")"

./init.sh
sleep 8
furynoded start >> furynode.log 2>&1  &
sleep 8

yes Y | furynoded tx clp create-pool --from akasha --symbol catk --nativeAmount 1000 --externalAmount 1000
sleep 8
yes Y | furynoded tx clp add-liquidity --from fury --symbol catk --nativeAmount 5000000000000000000000 --externalAmount 5000000000000000000
sleep 8
echo "Getting from CLI"
furynoded query clp lplist catk

echo "Getting from REST"
curl --request GET -sL \
     --url 'http://localhost:1317/clp/getLpList?symbol=catk'\
