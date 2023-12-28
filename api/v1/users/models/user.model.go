package users

// User represents the user data structure
type User struct {
	ID        int    `json:"user_id"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}
