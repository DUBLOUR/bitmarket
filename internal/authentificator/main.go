package authentificator

import "bitmarket/internal/userModel"

type Auth struct {
	Hash       IHash
	controller IUserController
}

type Credentials struct {
	email    string
	password string
}

func (c *Credentials) SetEmail(email string) {
	c.email = email
}

func (c Credentials) GetEmail() string {
	return c.email
}

func (c *Credentials) SetPass(password string) {
	c.password = password
}
func (c Credentials) GetPass() string {
	return c.password
}

type IAuth interface {
	Register(Credentials) error
	Login(Credentials) (userModel.IToken, error)
	LoginByToken(userModel.IToken) error
}

func (a Auth) Register(user userModel.ICredentials) error {
	return nil
}

func (a Auth) Login(user userModel.ICredentials) (userModel.IToken, error) {
	return StringToken{""}, nil
}

func (a Auth) LoginByToken(t userModel.IToken) error {
	return nil
}

type IUserController interface {
	Append(userModel.IUser) error
	Delete(userModel.IUser) error
	FindByEmail(string) (userModel.IUser, error)
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

func (t StringToken) Renew() {

}
