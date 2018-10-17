package defs

type UserCredential struct { // 自动转 json
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

