package user

type UserDB struct {
	Id   int64  `json:"id"`
	TgId int64  `json:"telegram_id"`
	Name string `json:"name"`
}
