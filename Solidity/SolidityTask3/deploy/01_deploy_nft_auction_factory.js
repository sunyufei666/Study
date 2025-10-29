const {ethers, upgrades} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({deployments, getNamedAccounts}) => {
    const {save} = deployments;
    const {deployer} = await getNamedAccounts();
    console.log("工厂合约部署用户的地址：",deployer);

    const nftAuctionFactory = await ethers.getContractFactory("NftAuctionFactory");
    const nftAuctionProxy = await upgrades.deployProxy(nftAuctionFactory, [], {initializer: 'initialize'});
    await nftAuctionProxy.waitForDeployment();

    const proxyAddress = await nftAuctionProxy.getAddress();
    console.log("代理合约地址：", proxyAddress);
    const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
    console.log("实现合约地址：", implAddress);

    const filePath = path.join(__dirname, "./.cache/proxyNftAuctionFactory.json");
    fs.writeFileSync(filePath, JSON.stringify({
        proxyAddress,
        implAddress,
        abi: nftAuctionFactory.interface.format("json")
    }));

    await save("NftAuctionFactoryProxy", {abi: nftAuctionFactory.interface.format("json"), address: proxyAddress});
};

module.exports.tags = ["deployNftAuctionFactory"];