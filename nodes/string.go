package nodes

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type UppercaseResolver struct {
}

func (r UppercaseResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"output": strings.ToUpper(state["input"].(string)),
	}, nil
}

func (r UppercaseResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["input"]; !ok {
		return errors.New("Missing required field 'input'")
	}
	return nil
}

type LowercaseResolver struct {
}

func (r LowercaseResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"output": strings.ToLower(state["input"].(string)),
	}, nil
}

func (r LowercaseResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["input"]; !ok {
		return errors.New("Missing required field 'input'")
	}
	return nil
}

type RegexResolver struct {
}

func (r RegexResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	rx := regexp.MustCompile(state["match"].(string))
	// TODO: add flags from state["flags"]
	return map[string]interface{}{
		"output": rx.ReplaceAllString(state["input"].(string), state["replace"].(string)),
	}, nil
}

func (r RegexResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["input"]; !ok {
		return errors.New("Missing required field 'input'")
	}
	return nil
}

type PassUnitResolver struct {
}

func (r PassUnitResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println(state)
	//TODO implement this properly
	if strings.Contains(state["value"].(string), state["fallback"].(string)) {
		return map[string]interface{}{
			"output": state["value"],
		}, nil
	} else {
		return map[string]interface{}{
			"output": state["value"].(string) + state["fallback"].(string),
		}, nil
	}

}

func (r PassUnitResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["value"]; !ok {
		return errors.New("Missing required field 'input'")
	}
	return nil
}
