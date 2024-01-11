package exception

func InvalidLoginDataException() ApplicationException {
	return ApplicationException{
		StatusCode: 400,
		Message:    "Invalid credentials",
	}
}
