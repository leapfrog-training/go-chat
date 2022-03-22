package auth

/**
 * Document Schema for Todo table.
 * @type {struct} Todo
 */
type User struct {
	Id       int
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

/**
 * Document Schema for Todo table.
 * @type {struct} Todo
 */
type Chat struct {
	Id       int
	Chat    string `json:"email"`
	Username string `json:"username"`
}