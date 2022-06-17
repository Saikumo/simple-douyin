package dto

type Video struct {
	Id            uint      `json:"id"`
	UserInfo      *UserInfo `json:"author" gorm:"-"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	Title         string    `json:"title"`
	FavoriteCount uint      `json:"favorite_count"`
	CommentCount  uint      `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
}
