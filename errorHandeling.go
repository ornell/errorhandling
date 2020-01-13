package errorHandeling

import "log"

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ErrorPrint(err error){
	if err != nil {
		log.Print(err.Error())
	}
}
func ErrorPanic(err error){
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}