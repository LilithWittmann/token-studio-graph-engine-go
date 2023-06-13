package nodes

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ArrayindexResolver struct {
}

func (r ArrayindexResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	index, _ := strconv.Atoi(state["index"].(string))
	return map[string]interface{}{
		"output": state["array"].([]interface{})[index],
	}, nil
}

func (r ArrayindexResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["array"]; !ok {
		return errors.New("Missing required field 'array'")
	}
	if _, ok := state["index"]; !ok {
		return errors.New("Missing required field 'index'")
	}
	return nil
}

type ArrifyResolver struct {
}

func (r ArrifyResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {

	outputList := make([]interface{}, 0)

	for k, _ := range state {
		outputList = append(outputList, state[k])
	}
	return map[string]interface{}{
		"output": outputList,
	}, nil
}

func (r ArrifyResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	return nil
}

type ReverseResolver struct {
}

func reverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (r ReverseResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	result := state["array"]
	return map[string]interface{}{
		"output": result,
	}, nil
}

func (r ReverseResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["array"]; !ok {
		return errors.New("Missing required field 'input'")
	}
	return nil
}

type SliceResolver struct {
}

func (r SliceResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println(state["array"].([]interface{})[1:3])
	start, _ := strconv.Atoi(state["start"].(string))
	end, _ := strconv.Atoi(state["end"].(string))
	return map[string]interface{}{
		"output": state["array"].([]interface{})[start:end],
	}, nil
}

func (r SliceResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["array"]; !ok {
		return errors.New("Missing required field 'array'")
	}

	if _, ok := state["start"]; !ok {
		return errors.New("Missing required field 'start'")
	}

	if _, ok := state["end"]; !ok {
		return errors.New("Missing required field 'end'")
	}

	return nil
}

type JoinResolver struct {
}

func (r JoinResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	stringList := make([]string, 0)

	for _, v := range state["array"].([]interface{}) {
		stringList = append(stringList, fmt.Sprintf("%v", v))
	}

	return map[string]interface{}{
		"output": strings.Join(stringList, state["delimiter"].(string)),
	}, nil
}

func (r JoinResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["array"]; !ok {
		return errors.New("Missing required field 'array'")
	}

	if _, ok := state["delimiter"]; !ok {
		return errors.New("Missing required field 'separator'")
	}
	return nil
}
