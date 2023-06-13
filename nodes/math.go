package nodes

import "C"
import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type AdditionResolver struct {
	NodeResolver
}

func InputsToFloat(data map[string]interface{}) map[string]float64 {
	// convert every entry in the map to float64
	floatMap := make(map[string]float64)
	for k, v := range data {
		// check if it is already float64
		if _, ok := v.(float64); !ok {
			floatMap[k], _ = strconv.ParseFloat(v.(string), 64)
		} else {
			floatMap[k] = v.(float64)
		}

	}
	return floatMap

}

func (r AdditionResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	return map[string]interface{}{"output": (inputs["1"] + inputs["2"])}, nil
}

func (r AdditionResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	// check if 1 and 2 exist and can be converted to float64
	if _, ok := state["1"]; !ok {
		return errors.New("Missing required field '1'")
	}

	if _, ok := state["2"]; !ok {
		return errors.New("Missing required field '2'")
	}
	return nil
}

type SubtractResolver struct {
	NodeResolver
}

func (r SubtractResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)

	return map[string]interface{}{
		"output": inputs["1"] - inputs["2"],
	}, nil
}

func (r SubtractResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["1"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	if _, ok := data["2"]; !ok {
		return errors.New("Missing required field 'b'")
	}
	return nil
}

type MultiplyResolver struct {
	NodeResolver
}

func (r MultiplyResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	return map[string]interface{}{
		"output": inputs["1"] * inputs["2"],
	}, nil
}

func (r MultiplyResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["a"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	if _, ok := data["b"]; !ok {
		return errors.New("Missing required field 'b'")
	}
	return nil
}

type DivideResolver struct {
	NodeResolver
}

func (r DivideResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println("DivideResolver.Resolve")
	inputs := InputsToFloat(state)

	return map[string]interface{}{
		"output": inputs["1"] / inputs["2"],
	}, nil
}

func (r DivideResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["1"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	if _, ok := data["2"]; !ok {
		return errors.New("Missing required field 'b'")
	}
	return nil
}

type ModuloResolver struct {
	NodeResolver
}

func (r ModuloResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	fmt.Println(inputs["a"])
	fmt.Println(inputs["b"])
	fmt.Println(math.Mod(inputs["a"], inputs["b"]))
	return map[string]interface{}{
		"output": math.Mod(inputs["a"], inputs["b"]),
	}, nil
}

func (r ModuloResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["a"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	if _, ok := data["b"]; !ok {
		return errors.New("Missing required field 'b'")
	}
	return nil
}
