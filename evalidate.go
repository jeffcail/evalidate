package evalidate

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type ECustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *ECustomValidator) Validate(i interface{}) error {
	c.lazy()
	return c.validate.Struct(i)
}

func (c *ECustomValidator) lazy() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}

func NewECustomValidator() *ECustomValidator {
	return new(ECustomValidator)
}

func Check(err error) (msg string) {
	if err == nil {
		return
	}

	errs := err.(validator.ValidationErrors)
	for _, er := range errs {
		msg = er.Error()
	}
	return
}
