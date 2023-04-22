const BN = require('bn.js');

module.exports = async (cb) => {
    const Web3 = require("web3");

    const furynetUtilities = require('./furynetUtilities')
    const contractUtilites = require('./contractUtilities');

    const logging = furynetUtilities.configureLogging(this);

    const argv = furynetUtilities.processArgs(this, {
        ...furynetUtilities.sharedYargOptions,
        ...furynetUtilities.transactionYargOptions
    });

    try {
        logging.debug(`sendLockTx arguments: ${JSON.stringify(argv, undefined, 2)}`);

        const bridgeBankContract = await contractUtilites.buildContract(this, argv, logging, "BridgeBank", argv.bridgebank_address);

        let cosmosRecipient = Web3.utils.utf8ToHex(argv.furynet_address);
        let coinDenom = argv.symbol;
        let amount = new BN(argv.amount);

        let transactionParameters = {
            from: argv.ethereum_address,
            value: coinDenom === furynetUtilities.NULL_ADDRESS ? amount : 0,
        };

        await contractUtilites.setAllowance(this, argv.symbol, argv.amount, argv, logging, transactionParameters);

        const bridgeBankContractLockArgs = {
            recipient: cosmosRecipient,
            token: coinDenom,
            amount,
            transactionParameters
        }
        logging.debug(`bridgeBankContract.lock arguments: ${JSON.stringify(bridgeBankContractLockArgs, undefined, 2)}`);
        const lockResult = await bridgeBankContract.lock(cosmosRecipient, coinDenom, amount, transactionParameters);

        console.log(JSON.stringify(lockResult, undefined, 0))
    } catch (e) {
        console.error(`lock error: ${e} ${e.message}`);
        throw(e);
    }

    return cb();
};