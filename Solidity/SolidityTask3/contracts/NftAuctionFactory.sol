// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./NftAuction.sol";

contract NftAuctionFactory is Initializable {

    address public admin; // 管理员

    address[] public allAuctions; // 所有拍卖的合约地址

    mapping(address => address[]) public userAuctions; // 用户创建的拍卖合约地址

    function initialize() public initializer {
        admin = msg.sender;
    }

    /**
     * @dev 创建新的拍卖合约
     */
    function createAuction() external returns (address auction) {
        // 创建新的拍卖合约
        bytes memory bytecode = type(NftAuction).creationCode;
        bytes32 salt = keccak256(abi.encodePacked(msg.sender, allAuctions.length, block.timestamp));
        
        assembly {
            auction := create2(0, add(bytecode, 32), mload(bytecode), salt)
        }
        
        // 初始化拍卖合约
        NftAuction(auction).initialize();
        
        // 记录拍卖合约
        allAuctions.push(auction);
        userAuctions[msg.sender].push(auction);
    }
    
    /**
     * @dev 获取所有拍卖合约数量
     */
    function allAuctionsLength() external view returns (uint256) {
        return allAuctions.length;
    }
    
    /**
     * @dev 获取用户创建的拍卖合约数量
     */
    function userAuctionsLength(address user) external view returns (uint256) {
        return userAuctions[user].length;
    }
    
    /**
     * @dev 预测拍卖合约地址
     */
    function predictAuctionAddress(address creator, uint256 nonce) external view returns (address) {
        bytes32 salt = keccak256(abi.encodePacked(creator, nonce, block.timestamp));
        bytes memory bytecode = type(NftAuction).creationCode;
        bytes32 hash = keccak256(
            abi.encodePacked(
                bytes1(0xff),
                address(this),
                salt,
                keccak256(bytecode)
            )
        );
        return address(uint160(uint256(hash)));
    }
}