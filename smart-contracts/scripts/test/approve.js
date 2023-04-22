const BN = require('bn.js');

module.exports = async (cb) => {
    const Web3 = require("web3");

    const furynetUtilities = require('./furynetUtilities')
    const contractUtilites = require('./contractUtilities');

    const logging = furynetUtilities.configureLogging(this);

    const argv = furynetUtilities.processArgs(this, {
        ...furynetUtilities.sharedYargOptions,
        ...furynetUtilities.amountYargOption,
        ...furynetUtilities.ethereumAddressYargOption,
        ...furynetUtilities.symbolYargOption,
        'spender_address': {
            type: "string",
            demandOption: true
        },
    });

    const amount = new BN(argv.amount, 10);

    const web3instance = contractUtilites.buildWeb3(this, argv, logging);
    const tokenContract = await contractUtilites.buildContract(this, argv, logging,"BridgeToken", argv.symbol.toString());

    const result = await tokenContract.approve(argv.spender_address, argv.amount, {
        from: argv.ethereum_address,
        value: 0,
        gas: argv.gas
    });

    console.log(JSON.stringify(result));

    return cb();
};
