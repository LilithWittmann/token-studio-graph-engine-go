package nodes

import (
	"errors"
	"fmt"
)

type InputResolver struct {
}

func (r InputResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	outputState := make(map[string]interface{})
	// iterate over the definitions in the data and assign them to the state
	for k, _ := range state["definition"].(map[string]interface{}) {
		outputState[k] = state["values"].(map[string]interface{})[k]
	}
	return outputState, nil
}

func (r InputResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["name"]; !ok {
		return errors.New("Missing required field 'name'")
	}
	return nil
}

type OutputResolver struct {
}

func (r OutputResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {

	outputState := make(map[string]interface{})

	// iterate over the mapping in the state map and assign them to the output
	for _, v := range state["mappings"].([]interface{}) {
		// check if the state exist and if yes apply to output
		if _, ok := state[v.(map[string]interface{})["key"].(string)]; ok {
			outputState[v.(map[string]interface{})["name"].(string)] = state[v.(map[string]interface{})["key"].(string)]
		}
	}
	fmt.Println(outputState)

	return outputState, nil
}

func (r OutputResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := data["mappings"]; !ok {
		return errors.New("Missing required field 'mappings'")
	}
	return nil
}
