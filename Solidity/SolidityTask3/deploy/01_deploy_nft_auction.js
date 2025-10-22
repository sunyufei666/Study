const {ethers, upgrades} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({deployments, getNamedAccounts}) => {
    const {save} = deployments;
    const {deployer} = await getNamedAccounts();
    console.log("部署用户的地址是：", deployer);

    const auctionFactory = await ethers.getContractFactory("NftAuction");
    const auctionProxy = await upgrades.deployProxy(auctionFactory, [], {initializer: 'initialize'});
    await auctionProxy.waitForDeployment();

    const proxyAddress = await auctionProxy.getAddress();
    console.log("代理合约地址：", proxyAddress);
    const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
    console.log("实现合约地址：", implAddress);

    const filePath = path.join(__dirname, "./.cache/proxyNftAuction.json");
    fs.writeFileSync(filePath, JSON.stringify({
        proxyAddress,
        implAddress,
        abi: auctionFactory.interface.format("json")
    }));

    await save("NftAuctionProxy", {abi: auctionFactory.interface.format("json"), address: proxyAddress});
};

module.exports.tags = ["deployNftAuction"];