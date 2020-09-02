package services

// Errors indicate an abnormal condition occurring in the program.
// As we need to make a lot of error check in our project, it will be code repetition.
// So we are making this functin, to minimize the amount of code repetition.
// Now, we will need ony 1 line (call this function)
//  instead of writting 3 line
// if err != nil {
// 	  ......
// }

import "log"

// PanicIfError will make program stop with error message if error found.
func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

// LogIfError will just print error if error is found.
// We need this so that program do not stop
// and we will have a log for inspection.
// We are printing function name also, so that it will be
// easy to track back this error.
func LogIfError(function string, e error) {
	if e != nil {
		log.Println("Error in function: ", function, " and error is: ", e)
	}
}

// LogAndPanicIfError will print log and after that stop program by panic.
// We just need the for easy read.
func LogAndPanicIfError(function string, e error) {
	if e != nil {
		log.Println("Error in function: ", function, " and error is: ", e)
		panic(e)
	}
}
