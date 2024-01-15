package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 // QQ
	SignEmail SignStatus = 2 //	Gitee
	SignGitee SignStatus = 3 // Email
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	switch s {
	case SignQQ:
		return "QQ"
	case SignEmail:
		return "Gitee"
	case SignGitee:
		return "Email"
	default:
		return "其他"
	}
}
