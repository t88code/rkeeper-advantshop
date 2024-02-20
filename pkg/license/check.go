package check

import (
	"os"
	"time"
)

func Check() {
	tm := time.Date(2024, time.May, 2, 0, 0, 0, 0, time.UTC)

	if time.Now().Sub(tm) > 0 {
		os.Exit(1)
	}

}

func CheckRestCode(RestCode string) {
	for _, code := range RestCodes {
		if RestCode == code {
			return
		}
	}
	os.Exit(1)
}
