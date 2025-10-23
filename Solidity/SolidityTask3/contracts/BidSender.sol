// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {IRouterClient} from "@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IRouterClient.sol";
import {Client} from "@chainlink/contracts-ccip/src/v0.8/ccip/libraries/Client.sol";

contract BidSender {

    IRouterClient public ccipRouter;
    address public mainAuction;
    uint64 public mainChainSelector;
    address public paymentToken;

    constructor(
        address router, 
        address auction, 
        uint64 chainSelector,
        address _paymentToken
    ) {
        ccipRouter = IRouterClient(router);
        mainAuction = auction;
        mainChainSelector = chainSelector;
        paymentToken = _paymentToken;
    }

    // 用户调用这个函数进行跨链出价
    function placeCrossChainBid(uint256 amount) external {
        require(amount > 0, "Bid amount must be greater than 0");
        
        // 转移代币到合约（用于支付CCIP费用和可能的资金锁定）
        // 注意：这里需要用户先授权这个合约可以转移他们的代币
        IERC20(paymentToken).transferFrom(msg.sender, address(this), amount);
        
        // 构建CCIP消息
        Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
            receiver: abi.encode(mainAuction),
            data: abi.encodeWithSignature(
                "receiveCrossChainBid(address,uint256,uint64)",
                msg.sender,
                amount,
                getCurrentChainSelector()
            ),
            tokenAmounts: new Client.EVMTokenAmount[](0),
            extraArgs: "",
            feeToken: address(0)
        });

        // 计算CCIP费用（需要用原生代币支付）
        uint256 fee = ccipRouter.getFee(mainChainSelector, message);
        
        // 发送跨链消息
        ccipRouter.ccipSend{value: fee}(mainChainSelector, message);
        
        emit BidSent(msg.sender, amount);
    }

    // 获取当前链的selector
    function getCurrentChainSelector() public pure returns (uint64) {
        // 需要根据实际部署的链设置正确的selector
        return 16015286601757825753; // 示例：Sepolia测试网
    }

    // 管理员可以提取合约中的代币（用于调试或处理异常情况）
    function withdrawTokens(address token, uint256 amount) external {
        // 在实际应用中应该添加权限控制
        IERC20(token).transfer(msg.sender, amount);
    }

    // 合约需要原生代币来支付CCIP费用
    receive() external payable {}
}