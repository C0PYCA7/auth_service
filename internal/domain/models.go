package domain

type Users struct {
	Id          int64
	Name        string
	Surname     string
	Login       string
	Password    string
	Email       string
	EnableTwoFa bool
}
