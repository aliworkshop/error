package errorslib

func HandleError(err error) ErrorModel {
	if err == nil {
		return nil
	}
	if m, ok := err.(ErrorModel); ok {
		return m
	}
	e := Internal(err)
	return e
}
