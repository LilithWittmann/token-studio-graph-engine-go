package nodes

import (
	"errors"
	"fmt"
	"strconv"
)

func InputsTooBool(data map[string]interface{}) map[string]bool {
	boolMap := make(map[string]bool)
	for k, v := range data {
		// check if it is already float64
		if _, ok := v.(bool); !ok {
			boolMap[k], _ = strconv.ParseBool(v.(string))
		} else {
			boolMap[k] = v.(bool)
		}
	}
	return boolMap
}

type IfResolver struct {
}

func (r IfResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	if _, ok := state["condition"].(bool); !ok {
		state["condition"], _ = strconv.ParseBool(state["condition"].(string))
	} else {
		state["condition"] = state["condition"].(bool)
	}

	if state["condition"].(bool) {
		return map[string]interface{}{
			"output": state["a"],
		}, nil
	} else {
		return map[string]interface{}{
			"output": state["b"],
		}, nil
	}
}

func (r IfResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["condition"]; !ok {
		return errors.New("Missing required field 'condition'")
	}
	if _, ok := data["a"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	if _, ok := data["b"]; !ok {
		return errors.New("Missing required field 'b'")
	}
	return nil
}

type AndResolver struct {
}

func (r AndResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println(state)
	inputs := InputsTooBool(state)
	n := len(inputs)

	result := inputs[strconv.Itoa(0)]

	for i := 1; i < n; i++ {
		result = result && inputs[strconv.Itoa(i)]
	}

	return map[string]interface{}{
		"output": result,
	}, nil
}

func (r AndResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["a"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	if _, ok := data["b"]; !ok {
		return errors.New("Missing required field 'b'")
	}
	return nil
}

type OrResolver struct {
}

func (r OrResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println(state)
	inputs := InputsTooBool(state)
	n := len(inputs)

	result := inputs[strconv.Itoa(0)]
	for i := 1; i < n; i++ {
		result = result || inputs[strconv.Itoa(i)]
	}

	return map[string]interface{}{
		"output": inputs["1"] || inputs["2"],
	}, nil
}

func (r OrResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	inputs := InputsTooBool(state)
	n := len(inputs)
	if n < 2 {
		return errors.New("Not enough inputs")
	}
	return nil
}

type NotResolver struct {
}

func (r NotResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println(state)
	inputs := InputsTooBool(state)
	return map[string]interface{}{
		"output": !inputs["a"],
	}, nil
}

func (r NotResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["a"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	return nil
}

type SwitchResolver struct {
}

func (r SwitchResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println(state)
	// iterate over state[order] and check if state[condition] matches; if it matches return the key as output
	for _, v := range state["order"].([]interface{}) {
		if state["condition"] == state[v.(string)] {
			return map[string]interface{}{
				"output": v,
			}, nil
		}
	}

	return map[string]interface{}{
		"output": state["default"],
	}, nil
}

func (r SwitchResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["a"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	return nil
}

type CompareResolver struct {
}

const (
	CompareEqual            = "=="
	CompareNotEqual         = "!="
	CompareLessThan         = "<"
	CompareLessThanEqual    = "<="
	CompareGreaterThan      = ">"
	CompareGreaterThanEqual = ">="
)

func (r CompareResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {

	switch data["operator"] {
	case CompareEqual:
		return map[string]interface{}{
			"output": state["a"] == state["b"],
		}, nil
	case CompareNotEqual:
		return map[string]interface{}{
			"output": state["a"] != state["b"],
		}, nil
	case CompareLessThan:
		return map[string]interface{}{
			"output": state["a"].(float64) < state["b"].(float64),
		}, nil
	case CompareLessThanEqual:
		return map[string]interface{}{
			"output": state["a"].(float64) <= state["b"].(float64),
		}, nil
	case CompareGreaterThan:
		return map[string]interface{}{
			"output": state["a"].(float64) > state["b"].(float64),
		}, nil
	case CompareGreaterThanEqual:
		return map[string]interface{}{
			"output": state["a"].(float64) >= state["b"].(float64),
		}, nil
	default:
		return map[string]interface{}{
			"output": false,
		}, nil
	}
}

func (r CompareResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["a"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	if _, ok := data["b"]; !ok {
		return errors.New("Missing required field 'b'")
	}
	if _, ok := data["operator"]; !ok {
		return errors.New("Missing required field 'operator'")
	}
	return nil
}
