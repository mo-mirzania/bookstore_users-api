package users

type User struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"fname"`
	LastName   string `json:"lname"`
	Email      string `json:"email"`
	DateCreate string `json:"date"`
}
