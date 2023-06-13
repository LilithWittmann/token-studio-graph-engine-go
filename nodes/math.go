package nodes

import "C"
import (
	"errors"
	"math"
	"math/rand"
	"strconv"
)

type AdditionResolver struct {
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
}

func (r SubtractResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)

	return map[string]interface{}{
		"output": inputs["1"] - inputs["2"],
	}, nil
}

func (r SubtractResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["1"]; !ok {
		return errors.New("Missing required field '1'")
	}
	if _, ok := state["2"]; !ok {
		return errors.New("Missing required field '2'")
	}
	return nil
}

type MultiplyResolver struct {
}

func (r MultiplyResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	return map[string]interface{}{
		"output": inputs["1"] * inputs["2"],
	}, nil
}

func (r MultiplyResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["1"]; !ok {
		return errors.New("Missing required field '1'")
	}
	if _, ok := state["2"]; !ok {
		return errors.New("Missing required field '2'")
	}
	return nil
}

type DivideResolver struct {
}

func (r DivideResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	return map[string]interface{}{
		"output": inputs["1"] / inputs["2"],
	}, nil
}

func (r DivideResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["1"]; !ok {
		return errors.New("Missing required field '1'")
	}
	if _, ok := state["2"]; !ok {
		return errors.New("Missing required field '2'")
	}
	return nil
}

type ModuloResolver struct {
}

func (r ModuloResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	return map[string]interface{}{
		"output": math.Mod(inputs["a"], inputs["b"]),
	}, nil
}

func (r ModuloResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["a"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	if _, ok := state["b"]; !ok {
		return errors.New("Missing required field 'b'")
	}
	return nil
}

type AbsoluteValueResolver struct {
}

func (r AbsoluteValueResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	return map[string]interface{}{
		"output": math.Abs(inputs["input"]),
	}, nil
}

func (r AbsoluteValueResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["input"]; !ok {
		return errors.New("Missing required field 'input'")
	}
	return nil
}

type RoundResolver struct {
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (r RoundResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	precision := uint(inputs["precision"])
	return map[string]interface{}{
		"output": roundFloat(inputs["value"], precision),
	}, nil
}

func (r RoundResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["value"]; !ok {
		return errors.New("Missing required field 'value'")
	}
	if _, ok := state["precision"]; !ok {
		return errors.New("Missing required field 'precision'")
	}
	return nil
}

type SineResolver struct {
}

func (r SineResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	return map[string]interface{}{
		"output": math.Sin(inputs["value"]),
	}, nil
}

func (r SineResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["value"]; !ok {
		return errors.New("Missing required field '1'")
	}
	return nil
}

type CosineResolver struct {
}

func (r CosineResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	return map[string]interface{}{
		"output": math.Cos(inputs["value"]),
	}, nil
}

func (r CosineResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["value"]; !ok {
		return errors.New("Missing required field '1'")
	}
	return nil
}

type TangentResolver struct {
}

func (r TangentResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	return map[string]interface{}{
		"output": math.Tan(inputs["value"]),
	}, nil
}

func (r TangentResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["value"]; !ok {
		return errors.New("Missing required field 'value'")
	}
	return nil
}

type LerpResolver struct {
}

func (r LerpResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	a := inputs["a"]
	b := inputs["b"]
	t := inputs["t"]
	return map[string]interface{}{
		"output": a + (b-a)*t,
	}, nil
}

func (r LerpResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["a"]; !ok {
		return errors.New("Missing required field 'a'")
	}
	if _, ok := state["b"]; !ok {
		return errors.New("Missing required field 'b'")
	}
	if _, ok := state["t"]; !ok {
		return errors.New("Missing required field 't'")
	}
	return nil
}

type ClampResolver struct {
}

func (r ClampResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	inputs := InputsToFloat(state)
	x := inputs["value"]
	min := inputs["min"]
	max := inputs["max"]
	return map[string]interface{}{
		"output": math.Max(math.Min(x, max), min),
	}, nil
}

func (r ClampResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["value"]; !ok {
		return errors.New("Missing required field 'value'")
	}
	if _, ok := state["min"]; !ok {
		return errors.New("Missing required field 'min'")
	}
	if _, ok := state["max"]; !ok {
		return errors.New("Missing required field 'max'")
	}
	return nil
}

type RandomResolver struct {
}

func (r RandomResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {

	return map[string]interface{}{
		"output": rand.Float64(),
	}, nil
}

func (r RandomResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	return nil
}

type CountResolver struct {
}

func (r CountResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {

	return map[string]interface{}{
		"output": len(state["input"].([]interface{})),
	}, nil
}

func (r CountResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	return nil
}
