package user

type AuthModel struct {
	Id          int      `json:"id"`
	OwnerId     int      `json:"owner_id"`
	Permissions []string `json:"permissions"`
}
