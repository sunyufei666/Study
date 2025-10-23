// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
// import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract NftAuctionV2 is Initializable {

    struct Auction {
        address seller; // 卖家
        uint256 startPrice; // 起拍价格
        address highestBidder; // 最高出价者
        uint256 highestBid; // 最高出价

        bool ended; // 拍卖是否结束
        uint256 startTime; // 开始时间
        uint256 duration; // 持续时间
        
        address nftAddress; // NFT合约地址
        uint256 tokenId; // 
        address tokenAddress; // 出价所用的代币地址，address(0)表示ETH
    }

    Auction public auction; // 拍卖信息

    uint256 public nextAuctionId; // 下一个拍卖ID

    address public admin; // 管理员地址

    mapping(address => AggregatorV3Interface) public priceFeeds; // 代币价格预言机地址映射

    function initialize() public initializer {
        admin = msg.sender;
    }

    function setPriceFeed(address tokenAddress, address feedAddress) external {
        require(msg.sender == admin, "Only admin can set price feed");
        priceFeeds[tokenAddress] = AggregatorV3Interface(feedAddress);

    }

    function getChainlinkDataFeedLatestPrice(address tokenAddress) public view returns (int256) {
        AggregatorV3Interface priceFeed = priceFeeds[tokenAddress];
        (, int256 price, , , ) = priceFeed.latestRoundData();
        return price;
    }

    // 创建交易
    function createAuction(uint256 startPrice, uint256 duration, address nftAddress, uint256 tokenId) public {
        require(startPrice > 0, "Start price must be greater than 0");
        require(duration > 0, "Duration must be greater than 0");
        require(auction.seller == address(0), "Auction already exists for this NFT");

        // 将NFT授权给合约进行拍卖
        IERC721(nftAddress).approve(address(this), tokenId);

        auction = Auction({
            seller: msg.sender,
            startPrice: startPrice,
            highestBidder: address(0),
            highestBid: 0,
            ended: false,
            startTime: block.timestamp,
            duration: duration,
            nftAddress: nftAddress,
            tokenId: tokenId,
            tokenAddress: address(0)
        });
        nextAuctionId++;
    }

    // 买家参与买单
    function placeBid(uint256 amount, address tokenAddress) external payable {
        // 判断拍卖是否结束
        require(!auction.ended && block.timestamp < auction.startTime + auction.duration, "Auction has ended");
        // 判断出价是否高于当前最高价
        uint payValue;
        if (tokenAddress == address(0)) {
            // 计算ETH出价
            amount = msg.value;
            payValue = amount * uint(getChainlinkDataFeedLatestPrice(address(0)));
        } else {
            // 计算ERC20代币出价
            payValue = amount * uint(getChainlinkDataFeedLatestPrice(tokenAddress));
        }

        require(payValue > auction.highestBid * uint(getChainlinkDataFeedLatestPrice(auction.nftAddress)) 
            && payValue >= auction.startPrice * uint(getChainlinkDataFeedLatestPrice(auction.nftAddress)), "Bid amount is too low");
        
        if (tokenAddress == address(0)) {
            // 退回之前的ETH出价
            payable(auction.highestBidder).transfer(auction.highestBid);
        } else {
            // 处理ERC20代币出价
            IERC20(tokenAddress).transferFrom(msg.sender, address(this), amount);
            // 退回之前的出价
            IERC20(tokenAddress).transfer(auction.highestBidder, auction.highestBid);
        }

        auction.tokenAddress = tokenAddress;
        auction.highestBidder = msg.sender;
        auction.highestBid = amount;
    }

    // 结束拍卖
    function endAuction() external {
        require(!auction.ended && block.timestamp >= auction.startTime + auction.duration, "Auction duration has not yet passed or already ended");
        // 将NFT转移给最高出价者
        IERC721(auction.nftAddress).transferFrom(address(this), auction.highestBidder, auction.tokenId);
        // 将资金转移给卖家
        if (auction.tokenAddress == address(0)) {
            payable(auction.seller).transfer(auction.highestBid);
        } else {
            IERC20(auction.tokenAddress).transfer(auction.seller, auction.highestBid);
        }
        auction.ended = true;
    }

}