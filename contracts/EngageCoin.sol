// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "./openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./openzeppelin/contracts/access/Ownable.sol";

contract EngageCoin is ERC20, Ownable {
    constructor(address originalOwner) Ownable(originalOwner) ERC20("EngageCoin", "ENGC") {
        // Initial minting to the owner of the contract 
        _mint(msg.sender, 1000000 * 10 ** decimals()); // Mint 1 million EngageCoins initially
    }

    // Function to mint new tokens as rewards
    function mintReward(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }
}
