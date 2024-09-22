# Blockchain Project (Go & Go-Ethereum)

## Project Overview

This is a **Go** and **Go-Ethereum** blockchain project that features two types of roles:

1. **Owner**:  
   - The owner has the authority to add ERC20 tokens.  
   - The owner can transfer these custom tokens and Ethers to different wallets.

2. **User**:  
   - Users can sign up to the platform.  
   - Upon signup, the system automatically generates wallets and keys for the user, which are stored in the `wallet` folder.  
   - Users can transfer tokens and Ethers but **cannot** create custom tokens.

The client-side of this project is built with **Next.js** and can be found in the `client` folder.

## Key Technologies

- **Go**: Backend logic and server implementation.
- **Go-Fiber**: Web framework for creating high-performance APIs.
- **Go-Ethereum**: Ethereum interaction, including smart contracts and token management.
- **PostgreSQL**: Database for storing user data, transactions, and wallet information.
- **Next.js**: Frontend framework used for building the client-side of the project, located in the `client` folder.

---

