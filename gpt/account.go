package gpt

type AccountType string

const (
	ACCOUNT_TYPE_SECRET_KEY   AccountType = "secretkey"
	ACCOUNT_TYPE_ACCESS_TOKEN AccountType = "accesstoken"
)

type SecretKeyAccount struct {
	AccountName string
	SecretKey   string
	AccountType AccountType
}

func (s *SecretKeyAccount) String() string {
	return s.AccountName
}

func (s *SecretKeyAccount) Type() AccountType {
	return s.AccountType
}

func (s *SecretKeyAccount) Name() string {
	return s.AccountName
}

func (s *SecretKeyAccount) Token() string {
	return s.SecretKey
}

// todo
type UserPwdAccount struct {
	UserName    string
	Pwd         string
	accessToken string
}

func (s *UserPwdAccount) Name() string {
	return s.UserName
}

func (s *UserPwdAccount) Token() string {
	// todo login and get access token
	panic("not implement")
}
