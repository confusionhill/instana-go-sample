package authentication

import "coffe-shop/database"

func RegisterNewUser(newUser RegisterUserEntity) (*UserEntity, error) {
	db := database.GetDB()
	resp := UserEntity{}
	query := `
			INSERT INTO users
			    (first_name, last_name, email, password, access) 
			VALUES ($1,$2,$3,$4, $5) RETURNING id, access;
			`
	err := db.Get(&resp, query, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password, newUser.RoleID)
	resp.FirstName = newUser.FirstName
	resp.LastName = newUser.LastName
	resp.Email = newUser.Email
	return &resp, err
}

func PerformLoginUser(user LoginUserEntity) (*UserEntity, error) {
	query := "SELECT id, first_name, last_name, email, access from users where email=$1 and password=$2"
	db := database.GetDB()
	resp := UserEntity{}
	err := db.Get(&resp, query, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
