package crypto

import "encoding/base64"

func Encrypto(pw string) string {
	return base64.StdEncoding.EncodeToString([]byte(pw))
}
