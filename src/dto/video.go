package dto

type Video struct {
	Id       uint
	UserInfo *UserInfo `json:"author"`
	PlayUrl  string
	CoverUrl string
	title    string
}
