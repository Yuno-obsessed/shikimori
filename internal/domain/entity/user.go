package entity

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Level    int    `json:"level"`
	RoleId   int    `json:"role_id"`
}
