// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "./openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ENGCStaking {
    IERC20 public engcToken;
    uint public rewardRate; // Reward rate per second, i,e 10000000 Wei

    struct Stake {
        uint amount;
        uint rewardDebt;
        uint depositTime;
    }

    mapping(address => Stake) public stakes;

    // Total staked amount
    uint public totalStaked;

    constructor(IERC20 _engcToken, uint _rewardRate) {
        engcToken = _engcToken;
        rewardRate = _rewardRate;
    }

    // Deposit ENGC tokens for staking
    function deposit(uint _amount) public {
        require(_amount > 0, "Cannot stake 0 tokens");

        // Transfer ENGC tokens to the contract
        engcToken.transferFrom(msg.sender, address(this), _amount);

        // Update staking data
        Stake storage stake = stakes[msg.sender];
        stake.amount += _amount;
        stake.depositTime = block.timestamp;

        totalStaked += _amount;
    }

    // Withdraw ENGC tokens and rewards
    function withdraw(uint _amount) public {
        Stake storage stake = stakes[msg.sender];
        require(stake.amount >= _amount, "Insufficient staked balance");

        // Calculate rewards
        uint rewards = calculateReward(msg.sender);

        // Reset staked amount
        stake.amount -= _amount;

        // Update total staked
        totalStaked -= _amount;

        // Transfer staked tokens and rewards
        engcToken.transfer(msg.sender, _amount + rewards);
    }

    // View function to see pending rewards for a staker
    function pendingReward(address _staker) public view returns (uint) {
        return calculateReward(_staker);
    }

      // Withdraw only rewards without unstaking
    function withdrawRewards() public {
        uint rewards = calculateReward(msg.sender);
        require(rewards > 0, "No rewards to withdraw");

        // Reset the user's reward debt
        Stake storage stake = stakes[msg.sender];
        stake.depositTime = block.timestamp; // Reset the deposit time to current time

        // Transfer the rewards
        engcToken.transfer(msg.sender, rewards);
    }

    // Calculate rewards based on staked amount and time
    function calculateReward(address _staker) internal view returns (uint) {
        Stake storage stake = stakes[_staker];
        uint stakedDuration = block.timestamp - stake.depositTime;

        // Reward = (amount staked) * (reward rate per second) * (staked duration in seconds)
        uint reward = (stake.amount * rewardRate * stakedDuration) / 1e18;

        return reward;
    }
}
