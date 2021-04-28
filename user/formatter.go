package user

import . "sidu/entity"

type UserFormatter struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}

func FormatUser(user User, token string) (userResponse UserFormatter) {
	userResponse = UserFormatter{
		ID:      user.ID,
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
		Token:   token,
	}
	return
}
