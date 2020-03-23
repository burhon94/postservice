package auth

type Auth struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Roles   []string `json:"roles"`
	Expired int64    `json:"exp"`
}
