package v1

type SignUpModel struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type SignInModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


func (m *SignInModel) IsValid() int {
	if m.Email == "" {
		return 1
	}
	if m.Password == "" {
		return 2
	}
	return 0
}

// Функция для проверки заполненности полей в SignUpModel
func (signUpModel *SignUpModel) IsValid() int {
	if signUpModel.FirstName == "" {
		return 1
	}
	if signUpModel.LastName == "" {
		return 2
	}
	if signUpModel.UserName == "" {
		return 3
	}
	if signUpModel.Email == "" {
		return 4
	}
	if signUpModel.Password == "" {
		return 5
	}
	return 0
}
