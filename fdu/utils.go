package fdu

func isValidCredential(uid string, pwd string) bool {
	return len(uid) == 11 && len(pwd) >= 6
}
