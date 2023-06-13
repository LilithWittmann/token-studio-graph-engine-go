package nodes

import "C"
import (
	"errors"
)

type ConstantResolver struct {
}

func (r ConstantResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"output": state["input"],
	}, nil
}

func (r ConstantResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["input"]; !ok {
		return errors.New("Missing required field 'input'")
	}
	return nil
}

type EnumeratedConstantResolver struct {
}

func (r EnumeratedConstantResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"output": state["current"],
	}, nil
}

func (r EnumeratedConstantResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["current"]; !ok {
		return errors.New("Missing required field 'current'")
	}
	return nil
}

type SliderResolver struct {
}

func (r SliderResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	result, _ := state["value"].([]interface{})[0].(float64)
	return map[string]interface{}{
		// TODO: fixme
		"output": result,
	}, nil
}

func (r SliderResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["value"]; !ok {
		return errors.New("Missing required field 'value'")
	}
	return nil
}
