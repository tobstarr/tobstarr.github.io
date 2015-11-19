package wrap

import (
	"fmt"
	"os"

	"github.com/go-errors/errors"
)

func handle(err error) {
	if c, ok := err.(interface {
		ErrorStack() string
	}); ok {
		fmt.Println(c.ErrorStack())
	} else {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

func e(i interface{}, args ...interface{}) error {
	if c, ok := i.(string); ok {
		return errors.Wrap(fmt.Errorf(c, args...), 1)
	} else if i == nil {
		return nil
	}
	return errors.Wrap(i, 1)
}
