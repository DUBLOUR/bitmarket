package userModel

type ICredentials interface {
	SetEmail(string)
	GetEmail() string
	SetPassword(string)
	GetPassword() string
}

type IUser interface {
	RenewToken() IToken
}

type IToken interface {
	Str() string
	Renew()
}
