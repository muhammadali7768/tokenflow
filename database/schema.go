package database

import "log"

func CreateUsersTable() {
	schema := `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
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

func CreatePostsTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)
}

func CreateLiksTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS likes (
   id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    post_id INT REFERENCES posts(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)
}

func CreateCommentsTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS comments (
   id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    post_id INT REFERENCES posts(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    parent_comment_id INT REFERENCES comments(id) ON DELETE CASCADE, -- Allows for nested comments
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)
}

func CreateFollowersTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS followers  (
    id SERIAL PRIMARY KEY,
    follower_id INT REFERENCES users(id) ON DELETE CASCADE,
    following_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)
}

func CreateReferralsTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS  referrals (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    referred_user_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)
}

func CreateTokensTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS  tokens (
   id SERIAL PRIMARY KEY,
    token_name VARCHAR(100) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    total_supply BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)
}

func CreateTransactionsTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS   transactions (
    id SERIAL PRIMARY KEY,
    from_user_id INT REFERENCES users(id),
    to_user_id INT REFERENCES users(id),
    amount BIGINT NOT NULL,
    transaction_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)
}

func CreateUserEngagementsTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS   user_engagements (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    engagement_type VARCHAR(50) NOT NULL, -- e.g., "post", "like", "share", "referral"
    points_awarded INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)
}
