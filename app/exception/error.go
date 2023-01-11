package exception

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func PanicResponse(msg string) {
	panic(ValidationError{
		Message: msg,
	})
}
