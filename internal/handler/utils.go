package handler

import (
	"github.com/google/uuid"
	"golang.org/x/text/encoding/charmap"
	"regexp"
	"strings"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func IsValidPHONE(p string) bool {
	r, _ := regexp.Compile("^(?:\\+?7|8)?(?:[\\s\\-(_]+)?(\\d{3})(?:[\\s\\-_)]+)?(\\d{3})(?:[\\s\\-_]+)?(\\d{2})(?:[\\s\\-_]+)?(\\d{2})$")
	return r.MatchString(p)
}

func EncodeWindows1251(ba []uint8) []uint8 {
	enc := charmap.Windows1251.NewEncoder()
	out, _ := enc.String(string(ba))
	return []uint8(out)
}

func GetFullName(FirstName, LastName, MiddleName string) string {
	var fullName []string
	if LastName != "" {
		fullName = append(fullName, LastName)
	}
	if FirstName != "" {
		fullName = append(fullName, FirstName)
	}
	if MiddleName != "" {
		fullName = append(fullName, MiddleName)
	}
	if len(fullName) > 0 {
		return strings.Join(fullName, " ")
	}
	return ""
}
