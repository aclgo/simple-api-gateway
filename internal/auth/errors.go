package auth

type ErrInvalidToken struct {
}

func (e ErrInvalidToken) Error() string {
	return "invalid token"
}

type ErrRevogedToken struct {
}

func (e ErrRevogedToken) Error() string {
	return "revoged token"
}
