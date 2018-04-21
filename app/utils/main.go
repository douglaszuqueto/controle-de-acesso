package utils

import "log"

// CheckErr checkerr
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
