package protocol

type ILogger interface {

	Error(err error)
}

type ICaptcha interface {

	Validate(response string) bool;
}

type IPassword interface {
	
	GenerateHashPassword(password string) (string, error)

	CompareHashAndPassword(hashedPassword, password string) error
}

type IToken interface {

	GenerateToken(sub, role string, exp int64) (string, error)
}

type ITwoFactor interface {

	ValidatePasscode(passcode, secret string) bool
}