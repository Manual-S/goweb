package demo

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserModel struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}
