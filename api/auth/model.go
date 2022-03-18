package auth

/**
 * Document Schema for Todo table.
 * @type {struct} Todo
 */
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
