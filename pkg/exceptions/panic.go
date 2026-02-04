package exceptions

import "log"

func PanicLogging(err error) {
	if err != nil {
		log.Panic(err)
	}
}
