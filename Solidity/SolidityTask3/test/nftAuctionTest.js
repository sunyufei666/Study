const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");

describe("NftAuction", async function() {
    it("should test nft auction successfully", async function() {
        await main();
    });
});

async function main() {
    await deployments.fixture(["deployNftAuction"]);
    const NftAuctionProxy = await deployments.get("NftAuctionProxy");

    const [signer, buyer] = await ethers.getSigners();

    // 1. 部署NFT合约
    const nftFactory = await ethers.getContractFactory("MyNFT");
    const myNFT = await nftFactory.deploy();
    await myNFT.waitForDeployment();
    const myNFTAddress = await myNFT.getAddress();
    console.log("myNFT deployed to:", myNFTAddress);

    // mint10个NFT
    for (let i = 0; i <= 10; i++) {
        await myNFT.mint(signer.address, i + 1);
    }

    const tokenId = 1;
    // 2. 调用 createAuction 方法创建拍卖
    const nftAuction = await ethers.getContractAt("NftAuction", NftAuctionProxy.address);

    // 给代理合约授权
    await myNFT.connect(signer).setApprovalForAll(NftAuctionProxy.address, true);

    await nftAuction.createAuction(
        10,
        ethers.parseEther("0.01"),
        erc721Address,
        tokenId
    );

    const auction = await nftAuction.auction;
    console.log("创建拍卖成功：", auction);

    // 3. 购买者参与拍卖
    nftAuction.connect(buyer).placeBid(0, { value: ethers.parseEther("0.1") });

    // 4. 结束拍卖
    await new Promise(resolve => setTimeout(resolve, 10 * 1000));
    await nftAuction.connect(signer).endAuction(0);

    // 验证结果
    const auctionResult = await nftAuction.auction;
    console.log("结束拍卖后读取拍卖成功: ", auctionResult);
    expect(auctionResult.highestBidder).to.equal(buyer.address);
    expect(auctionResult.highestBid).to.equal(ethers.parseEther("0.01"));

    // 验证NFT归属
    const owner = await myNFT.ownerOf(tokenId);
    console.log("NFT归属:", owner);
    expect(owner).to.equal(buyer.address);
}