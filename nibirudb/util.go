package nibirudb

import "fmt"

func IsEmpty(s string) bool {
	return s == ""
}

func FormatLike(s string) string {
	return fmt.Sprintf("%%%s%%", s)
}
