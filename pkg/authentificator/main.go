package authentificator

type Auth struct {
	Hash       IHash
	controller IUserController
}

type IUserController interface {
	Append(IUser) error
	Delete(IUser) error
	FindByEmail(string) (IUser, error)
}

type IToken interface {
	Str() string
	Renew() IToken
}

type IUser interface {
	RenewToken() IToken
}

type IDataHolder interface {
	Append(string) error
	Delete(string) error
	FindByKey(string) (string, error)
}

type IHash interface {
	Encode(string) string
}

type ShaHasher struct{}

func (*ShaHasher) Encode(string) string {
	return ""
}

type StringToken struct {
	Token string
}

func (t StringToken) Str() string {
	return t.Token
}

func (t StringToken) Renew() StringToken {
	return t
}
