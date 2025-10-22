const {ethers, upgrades} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({ getNamedAccounts, deployments}) => {
    const {save} = deployments;
    const {deployer} = getNamedAccounts();
    console.log("升级合约用户的地址：", deployer);

    // 获取保存的地址数据
    const filePath = path.join(__dirname, "./.cache/proxyNftAuction.json");
    const fileData = fs.readFileSync(filePath, "utf-8");
    const {proxyAddress} = JSON.parse(fileData);

    const auctionV2Factory = await ethers.getContractFactory("NftAuctionV2");
    const auctionV2Proxy = await upgrades.upgradeProxy(proxyAddress, auctionV2Factory);
    await auctionV2Proxy.waitForDeployment();

    const proxyAddressV2 = await auctionV2Proxy.getAddress();
    console.log("升级后合约的地址：", proxyAddressV2);

    await save("NftAuctionV2Proxy", {abi: auctionV2Factory.interface.format("json"), address: proxyAddressV2});
};

module.exports.tags = ["upgradeNftAuction"];