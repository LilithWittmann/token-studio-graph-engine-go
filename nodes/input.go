package nodes

import "C"
import (
	"errors"
)

type ConstantResolver struct {
	NodeResolver
}

func (r ConstantResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"output": state["input"],
	}, nil
}

func (r ConstantResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["value"]; !ok {
		return errors.New("Missing required field 'value'")
	}
	return nil
}
