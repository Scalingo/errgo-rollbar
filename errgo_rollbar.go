package errgorollbar

import (
	"github.com/stvp/rollbar"
	"gopkg.in/errgo.v1"
)

func BuildStack(err error) rollbar.Stack {
	stack := rollbar.Stack{}
	for err != nil {
		if errgoErr, ok := err.(*errgo.Err); !ok {
			break
		} else {
			frame := rollbar.Frame{
				Filename: errgoErr.File,
				Line:     errgoErr.Line,
			}
			stack = append(stack, frame)
			err = errgoErr.Underlying()
		}
	}
	return stack
}
