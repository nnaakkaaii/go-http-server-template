package crypto

import "encoding/base64"

func Decrypto(pw string) (string, error) {
	s, err := base64.StdEncoding.DecodeString(pw)
	if err != nil {
		return "", err
	}
	return string(s), err
}
