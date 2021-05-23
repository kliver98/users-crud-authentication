package error

//Struct that represents an Error Object for json response of API
type GeneralError struct {
	Message string
}

func (e *GeneralError) Error() string {
	m := "Unknwon error"
	if e.Message != "" {
		m = e.Message
	}
	return m
}
