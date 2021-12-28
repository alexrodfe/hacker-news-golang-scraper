package errorutility

import "fmt"

// ErrorWrapper will wrap error messages inside a desired format for proper error management and tracking
func ErrorWrapper(message string) func(params ...interface{}) error {
	return func(params ...interface{}) error {
		return fmt.Errorf(message, params...)
	}
}
