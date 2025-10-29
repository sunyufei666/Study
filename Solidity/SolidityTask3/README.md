# NFT拍卖场

## 项目结构

nft-auction/
├── contracts/
│   ├── MyNFT.sol
│   ├── NftAuction.sol
│   ├── NftAuctionFactory.sol
│   └── NftAuctionV2.sol
├── deploy/
│   ├── .cache/
│   │   ├── proxyNftAuctionFactory.json
│   │   └── proxyNftAuction.json
│   ├── NftAuction.sol
│   ├── NftAuctionFactory.sol
│   └── NftAuctionV2.sol
├── test/
└── hardhat.config.js

## 功能说明

1. NFT合约
使用 ERC721 标准实现一个 NFT 合约。
支持 NFT 的铸造和转移。

2. 拍卖合约
实现一个拍卖合约，支持以下功能：
创建拍卖：允许用户将 NFT 上架拍卖。
出价：允许用户以 ERC20 或以太坊出价。
结束拍卖：拍卖结束后，NFT 转移给出价最高者，资金转移给卖家。

3. 工厂模式
使用类似于 Uniswap V2 的工厂模式，管理每场拍卖。
工厂合约负责创建和管理拍卖合约实例。
集成 Chainlink 预言机

4. 价格计算
使用 Chainlink 的 feedData 预言机，获取 ERC20 和以太坊到美元的价格。
在拍卖合约中，将出价金额转换为美元，方便用户比较。

5. 合约升级
UUPS/透明代理：
使用 UUPS 或透明代理模式实现合约升级。
确保拍卖合约和工厂合约可以安全升级。

## 部署步骤

1. 初始化项目
```
npx hardhat --init
```

2. 安装需要的依赖包
```
npm install dotenv
npm install hardhat-deploy
npm install @openzeppelin/contracts-upgradeable
npm install @openzeppelin/hardhat-upgrades
npm install @nomiclabs/hardhat-ethers
npm install @openzeppelin/contracts
npm install @chainlink/contracts
```
3. 部署合约

```
npx hardhat deploy --network sepolia
```