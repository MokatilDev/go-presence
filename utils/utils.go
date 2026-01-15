package utils

import (
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PrintErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
