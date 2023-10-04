package authentication

type LoginUserEntity struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponseEntity struct {
	UserInfo    *UserEntity `json:"user"`
	AccessToken string      `json:"access_token"`
}

type RegisterUserEntity struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	RoleID    int    `json:"role_id"`
	Password  string `json:"password"`
}

type UserEntity struct {
	Id         string `json:"id" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Email      string `json:"email" db:"email"`
	Access     int    `json:"access" db:"access"`
	AccessName string `json:"access_name"`
}
