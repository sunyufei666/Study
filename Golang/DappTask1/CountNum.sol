// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

contract CountNum {

    uint256 public count;

    function add() external {
        count += 1;
    }
}