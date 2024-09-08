package model

type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Email    string `db:"email" json:"email"`
	Role     string `db:"role" json:"role"`
	Wallet   string `db:"ethereum_address json:wallet"`
}
