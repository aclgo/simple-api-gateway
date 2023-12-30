package auth

type ErrInvalidToken struct {
}

func (e ErrInvalidToken) Error() string {
	return "invalid token"
}
