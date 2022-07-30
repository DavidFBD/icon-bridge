const BTSPeriphery = artifacts.require("BTSPeriphery");
const BTSCore = artifacts.require("BTSCore");
const { deployProxy } = require("@openzeppelin/truffle-upgrades");

module.exports = async function (deployer, network) {
    if (network !== "development") {
        console.log('Start deploy Proxy BTSCore ');

        await deployProxy(
            BTSCore,
            [
                process.env.BSH_COIN_NAME,
                process.env.BSH_COIN_FEE,
                process.env.BSH_FIXED_FEE,
            ],
            { deployer }
        );

        console.log('Start deploy Proxy BTSPeriphery ');
        await deployProxy(
            BTSPeriphery,
            [process.env.BMC_PERIPHERY_ADDRESS, BTSCore.address],
            { deployer }
        );

        console.log('Start deploy BTSCore ');
        const btsCore = await BTSCore.deployed();

        console.log('Updating BTS peripher');
        await btsCore.updateBTSPeriphery(BTSPeriphery.address);
    }
};
