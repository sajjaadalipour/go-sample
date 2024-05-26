package error

type AppError struct {
	Message string
	Cause   error
	Code    string
	Status  int
}

func (e AppError) Error() string {
	return e.Message
}
