# Steps to follow by Investor/Distributor 

## Requirements
1 . List of address to distribute to :  For ValidatorSubsidy and LiquidityMining this list will be provided by furynet ,based on the users who have claimed rewards during the whole week .

2 . Enough funds to distribute : The account balance for the distributor needs to greater than the sum of all the outputs from step 1

3 . Authorized runner address : After creation ,the authorized runner would be responsible for triggering the reward transfers(10 per block)

## Steps
### Create Dispensation
Use the CLI to create a dispensations
```shell
furynoded tx dispensation create [Distribution Type Airdrop/ValidatorSubsidy/LiquidityMining] [List of output addresses in JSON format] [Address of authorized runner] --from [Address of invester/This private keys is used to sign the tx] --yes --gas auto --gas-adjustment=1.5 --gas-prices 1.0fury --node tcp://rpc.furynet.finance:80 --chain-id furynet
```
sample command
```shell
furynoded tx dispensation create LiquidityMining output.json fury1l7hypmqk2yc334vc6vmdwzp5sdefygj2ad93p5 --from fury1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd --yes --gas auto --gas-adjustment=1.5 --gas-prices 1.0fury --node tcp://rpc.furynet.finance:80 --chain-id furynet
```
sample output
```json
{
  "height": "0",
  "txhash": "A9D019E1080ECD6A012B20B3058534AC6643BD17634F181FBE7F8F5C43B94D8E",
  "raw_log": "[]"
}
```
The Tx hash can then be used to query the blockchain and get the distribution name
```shell
furynoded q tx A9D019E1080ECD6A012B20B3058534AC6643BD17634F181FBE7F8F5C43B94D8E --node tcp://rpc.furynet.finance:80 --chain-id furynet
```
The output from the command would contain the relevant event 
```json
{
        "type": "distribution_started",
        "attributes": [
        {
        "key": "module_account",
        "value": "fury1zvwfuvy3nh949rn68haw78rg8jxjevgm2c820c"
        },
        {
        "key": "distribution_name",
        "value": "1158855_fury1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd"
        },
        {
        "key": "distribution_type",
        "value": "LiquidityMining"
        }
]
}
```

### Next Steps 
- The next step is to run this dispensation. The run transaction needs to be signed by th authorized runner.
- The authorized runner address would normally be an account owned by furynet.
- The steps to be followed by the authorised runner are in the next document