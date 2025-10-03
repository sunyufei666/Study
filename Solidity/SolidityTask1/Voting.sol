// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*
1. 创建一个名为Voting的合约，包含以下功能：
    一个mapping来存储候选人的得票数
    一个vote函数，允许用户投票给某个候选人
    一个getVotes函数，返回某个候选人的得票数
    一个resetVotes函数，重置所有候选人的得票数
*/
contract Voting {
    mapping(string => uint256) private voteNum;
    string[] private candidates;

    function vote(string calldata candidate) external {
        voteNum[candidate] += 1;
        candidates.push(candidate);
    }

    function getVotes(string calldata candidate) external view returns (uint256) {
        return voteNum[candidate];
    }

    function resetVotes() external {
        for (uint256 i = 0; i < candidates.length; i++) {
            voteNum[candidates[i]] = 0;
        }
        delete candidates;
    }
}