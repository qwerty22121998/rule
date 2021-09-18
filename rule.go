package rule

import "fmt"

type Rule interface {
	Validate() error
}

func Validate(rules ...Rule) error {
	for _, rule := range rules {
		if err := rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type rule struct {
	validFunc func() error
}

func (r rule) Validate() error {
	return r.validFunc()
}

func Require(name string, value interface{}) Rule {
	return rule{
		validFunc: func() error {
			if isZero(value) {
				return fmt.Errorf("field [%v] is required", name)
			}
			return nil
		},
	}
}

func MaxLength(name string, value string, length int) Rule {
	return rule{
		validFunc: func() error {
			if len(value) > length {
				return fmt.Errorf("field [%v] length exceed the limit of %v", name, length)
			}
			return nil
		},
	}
}

func None() Rule {
	return rule{
		validFunc: func() error {
			return nil
		},
	}
}

func Group(rules ...Rule) Rule {
	return rule{validFunc: func() error {
		for _, rule := range rules {
			if err := rule.Validate(); err != nil {
				return err
			}
		}
		return nil
	}}
}
