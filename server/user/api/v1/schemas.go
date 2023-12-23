package v1

type MeModel struct {
	Token string `json:"token"`
}

func (m *MeModel) IsValid() int {
	if m.Token == "" {
		return 1
	}
	return 0
}
