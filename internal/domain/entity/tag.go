package entity

type Tag struct {
	TagId   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
	TagBody string `json:"tag_body"`
}
