// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract MyNFT is ERC721URIStorage {

    using Strings for uint256;

    uint256 private _tokenIds;

    constructor() ERC721("Ynfe", "YYYYY") {}

    function mintMFT(address recipient, string memory tokenURI) public returns (uint256) {
        _tokenIds++;
        _mint(recipient, _tokenIds);
        _setTokenURI(_tokenIds, tokenURI);
        return _tokenIds;
    }
}