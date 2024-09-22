package database

import "log"

func CreateUsersTable() {
	schema := `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
	role VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    ethereum_address VARCHAR(42) UNIQUE NOT NULL,
    keystore_file VARCHAR(255) NOT NULL,
    encrypted_mnemonic TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

	_, err := DB.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}

func CreateTokensTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS  tokens (
   id SERIAL PRIMARY KEY,
    token_name VARCHAR(100) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    total_supply NUMERIC NOT NULL,
	decimals INTEGER NOT NULL DEFAULT 18,
    owner_id INTEGER NOT NULL REFERENCES users(id),
	contract_address VARCHAR(64) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)
}
