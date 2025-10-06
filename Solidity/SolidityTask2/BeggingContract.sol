// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BeggingContract {

    // 合约所有者
    address public owner;

    mapping(address donors => uint256 amount) private _donations;

    // 排行榜
    address[] public _donationRank;

    // 捐赠开始时间
    uint256 public startTime;
    // 捐赠结束时间
    uint256 public endTime;

    event Donation(address donors, uint256 value);

    constructor () {
        owner = msg.sender;
        startTime = block.timestamp;
        endTime = startTime + 7 days;
    }

    // 捐赠
    function donate() external payable {
        require(block.timestamp > startTime && block.timestamp < endTime, "Donation period is incorrect!");
        
        _donations[msg.sender] += msg.value;

        if (!onDonationRank(msg.sender)) {
            _donationRank.push(msg.sender);
        }

        // 重新排序
        updateDonationRank();

        emit Donation(msg.sender, msg.value);
    }

    // 合约所有者提取所有资金
    function withdraw() external {
        require(msg.sender == owner, "Only the owner can withdraw funds!");
        uint256 balance = address(this).balance;
        require(balance > 0, "Insufficient funds!");
        payable(owner).transfer(balance);

        // 重置捐赠数量
        for (uint256 i = 0; i < _donationRank.length; i++) {
            _donations[_donationRank[i]] = 0;
        }
    }

    // 查询某个地址的捐赠金额
    function getDonation(address from) public view returns (uint256) {
        return _donations[from];
    }

    // 重新计算金额排序
    function updateDonationRank() private {
        for (uint256 i = 0; i < 3; i++) {
            for (uint256 j = i; j < _donationRank.length - 1; j++) {
                if (_donations[_donationRank[j]] < _donations[_donationRank[j + 1]]) {
                    address temp = _donationRank[j];
                    _donationRank[j] = _donationRank[j + 1];
                    _donationRank[j + 1] = temp;
                }
            }
        }
    }

    // 判断捐赠者是否在排行榜中
    function onDonationRank(address donors) private view returns(bool) {
        for (uint256 i = 0; i < _donationRank.length; i++) {
            if (_donationRank[i] == donors) {
                return true;
            }
        }
        return false;
    }
}