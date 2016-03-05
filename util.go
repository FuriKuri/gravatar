package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Map(list []string, fn func(string) string) []string {
	result := make([]string, len(list))
	for index, element := range list {
		result[index] = fn(element)
	}
	return result
}

func Hash(value string) string {
	hasher := md5.New()
	hasher.Write([]byte(value))
	return fmt.Sprintf("%v", hex.EncodeToString(hasher.Sum(nil)))
}
