package processor

import "fmt"

type CustomErrorType struct {
	error
}

func (c *CustomErrorType) CustomError(fn func() error) {
	if err:=fn() ; err != nil {
		fmt.Println(err.Error())
	}
}