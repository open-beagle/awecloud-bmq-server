package conf

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Alias    string `json:"alias"`
	Role     string `json:"role"`
}
