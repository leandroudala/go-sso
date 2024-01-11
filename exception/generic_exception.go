package exception

func InvalidIDException(param string) ApplicationException {
	return ApplicationException{
		StatusCode: 400,
		Message:    "The value '" + param + "' id not a valid ID.",
	}
}
