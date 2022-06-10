package dto

type Video struct {
	Id       int64
	UserInfo *UserInfo `json:"author"`
	PlayUrl  string
	CoverUrl string
	title    string
}
