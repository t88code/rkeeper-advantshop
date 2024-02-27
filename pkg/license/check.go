package check

import (
	"os"
	"time"
)

func Check() {
	tm := time.Date(2024, time.May, 2, 0, 0, 0, 0, time.UTC)

	if time.Now().Sub(tm) > 0 {
		os.Exit(3)
	}

}

func CheckRestCode(RestCode string) {
	for _, RestCodeLicence := range RestCodesLicense {
		if RestCode == RestCodeLicence {
			return
		}
	}
	os.Exit(3)
}

func CheckRestCodes(RestCodes []string) {
	for _, RestCodeLicence := range RestCodesLicense {
		for _, RestCode := range RestCodes {
			if RestCode == RestCodeLicence {
				return
			}
		}
	}
	os.Exit(3)
}
