package vrchatapi

type LoginResult int

const (
	Failed LoginResult = iota
	Wait2FA
	Succeeded
)
