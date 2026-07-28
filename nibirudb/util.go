package nibirudb

import (
	"fmt"

	"gorm.io/gorm"
)

func IsEmpty(s string) bool {
	return s == ""
}

func FormatLike(s string) string {
	return fmt.Sprintf("%%%s%%", s)
}

func FormatQueryLike(query *gorm.DB, field string, val string) *gorm.DB {
	return query.Where(fmt.Sprintf("%s ILIKE ?", field), FormatLike(val))
}
