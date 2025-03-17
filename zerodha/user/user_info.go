package user

type UserInfo struct {
	UserId string
	//portfolio string
	//watchlist string
	FundsRemaining float64
}

func NewUser(userId string, fundsRemaining float64) *UserInfo {
	return &UserInfo{
		UserId:         userId,
		FundsRemaining: fundsRemaining,
	}
}
