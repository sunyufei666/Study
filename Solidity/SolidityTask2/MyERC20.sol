// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract MyERC20 {

    mapping(address account => uint256) private _balances;

    mapping(address account => mapping(address spender => uint256)) private _allowances;

    uint256 private _totalSupply;

    // 使用event记录转账和授权操作
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    // 1. 查询账户余额
    function balanceOf(address account) public view virtual returns (uint256) {
        return _balances[account];
    }

    // 2. 转账
    function transfer(address to, uint256 value) public virtual returns (bool) {
        address from = msg.sender;
        _transfer(from, to, value);
        return true;
    }

    function _transfer(address from, address to, uint256 value) internal {
        if (from == address(0)) {
            revert("invalid sender!");
        }
        if (to == address(0)) {
            revert("invalid receiver!");
        }
        _update(from, to, value);
    }

    function _update(address from, address to, uint256 value) internal virtual {
        if (from == address(0)) {
            _totalSupply += value;
        } else {
            uint256 fromBalance = _balances[from];
            if (fromBalance < value) {
                revert("insufficient balance!");
            }
            _balances[from] = fromBalance - value;
        }

        if (to == address(0)) {
            _totalSupply -= value;
        } else {
            _balances[to] += value;
        }
        emit Transfer(from, to, value);
    }

    // 3.1 授权（是指代币所有者（owner）调用ERC20合约的approve函数，授权给另一个地址（spender）一定数量的代币，允许spender代表owner使用这些代币）
    function approve(address spender, uint256 value) public virtual returns (bool) {
        address owner = msg.sender;
        _approve(owner, spender, value);
        return true;
    }

    function _approve(address owner, address spender, uint256 value) internal virtual {
        if (owner == address(0)) {
            revert("invalid approver!");
        }
        if (spender == address(0)) {
            revert("invalid spender!");
        }
        // 授权的代币数量
        _allowances[owner][spender] = value;
        emit Approval(owner, spender, value);
    }

    // 3.2 代扣转账（是指被授权的地址（spender）调用transferFrom函数，从所有者（owner）账户中转移指定数量的代币到另一个地址（to））
    function transferFrom(address from, address to, uint256 value) public virtual returns (bool) {
        address spender = msg.sender;
        _spendAllowance(from, spender, value);
        _transfer(from, to, value);
        return true;
    }

    function _spendAllowance(address owner, address spender, uint256 value) internal virtual {
        uint256 currentAllowance = _allowances[owner][spender];
        if (currentAllowance < type(uint256).max) {
            // 转账数量不能超过剩余授权的代币数量
            if (currentAllowance < value) {
                revert("insufficient allowance!");
            }
            unchecked {
                // 更新授权代币数量
                _approve(owner, spender, currentAllowance - value);
            }
        }
    }

    // 4. 提供 mint 函数，允许合约所有者增发代币
    function _mint(address account, uint256 value) internal {
        if (account == address(0)) {
            revert("invalid receiver!");
        }
        _update(address(0), account, value);
    }

    // 销毁代币
    function _burn(address account, uint256 value) internal {
        if (account == address(0)) {
            revert("invalid sender!");
        }
        _update(account, address(0), value);
    }
}