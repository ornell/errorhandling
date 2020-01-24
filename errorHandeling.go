package errorHandeling

import (
	"fmt"
)

func ErrorPrint(err error, comment string){
	if err != nil {
		fmt.Errorf(comment+" %v", err)
	}
}

func ErrorPanic(err error, comment string){
	if err != nil {
		fmt.Errorf(comment+" %v", err)
		panic(err.Error())
	}
}
