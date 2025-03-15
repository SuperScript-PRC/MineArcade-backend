package accounts

func IsAccountOK(account string, password string) (bool, string) {
	return IsPasswordCorrect(account, password)
}
