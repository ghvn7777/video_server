package defs

type UserCredential struct { // 自动转 json
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

// response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comments struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string
	TTL int64
}