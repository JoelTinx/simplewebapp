package error

import (
	"log"
)

func catchError(err error) {
	if err != nil {
		log.Println(err.Error())
		util.SendEmailError(err.Error())
	}
}