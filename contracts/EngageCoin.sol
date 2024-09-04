// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract EngageCoin is ERC20, Ownable {
    // Constructor to set the initial supply and token details
    constructor() ERC20("EngageCoin", "ENGC") {
        // Initial minting to the owner of the contract (can be a multisig wallet or platform address)
        _mint(msg.sender, 1000000 * 10 ** decimals()); // Mint 1 million EngageCoins initially
    }

    // Function to mint new tokens as rewards
    function mintReward(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }
}
