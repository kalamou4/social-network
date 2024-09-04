package user

type User struct {
    ID         string `json:"id"`
    First_name string `json:"first_name"`
    Last_name  string `json:"last_name"`
    Nick_name  string `json:"nick_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
    Age        int    `json:"age"`
    Gender     string `json:"gender"`
    About_me   string `json:"about_me"`
    Visibility bool   `json:"visibility"`
}