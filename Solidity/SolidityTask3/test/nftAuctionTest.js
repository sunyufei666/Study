const {ethers, deployments} = require("hardhat");

describe("NftAuction", async function() {
    it("should test nft auction successfully", async function() {
        await main();
    });
});

async function main() {
    deployments.fixture(["deployNftAuction"]);

}