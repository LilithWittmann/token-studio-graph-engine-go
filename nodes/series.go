package nodes

import (
	"errors"
	"fmt"
	"math"
)

type ArithmeticSeriesResolver struct {
}

func (r ArithmeticSeriesResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println(state)
	outputState := make(map[string]interface{})
	inputs := InputsToFloat(state)
	// generate the series
	series := make([]float64, 0)
	base, _ := inputs["base"]
	steps, _ := inputs["steps"]
	increment, _ := inputs["increment"]
	stepsDown, _ := inputs["stepsDown"]

	for i := stepsDown; i > 0; i-- {
		series = append(series, base-float64(i)*increment)
	}

	for i := 0; i < int(steps); i++ {
		series = append(series, base+float64(i)*increment)
	}

	fmt.Println(series)
	outputState["array"] = series
	for i := stepsDown * -1; i < steps; i++ {
		fmt.Println(i)
		outputState[fmt.Sprintf("%d", int(i))] = series[int(i)+int(stepsDown)]
	}

	return outputState, nil
}

func (r ArithmeticSeriesResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["base"]; !ok {
		return errors.New("Missing required field 'base'")
	}
	if _, ok := state["steps"]; !ok {
		return errors.New("Missing required field 'steps'")
	}
	if _, ok := state["increment"]; !ok {
		return errors.New("Missing required field 'increment'")
	}
	if _, ok := state["stepsDown"]; !ok {
		return errors.New("Missing required field 'stepsDown'")
	}
	return nil
}

type HarmonicSeriesResolver struct {
}

type HarmonicScaleValue struct {
	Size      float64
	Frequency int
	Note      int
}

func calculateHarmonicScale(options struct {
	StepsDown int
	Steps     int
	Notes     int
	Base      float64
	Ratio     float64
}) []HarmonicScaleValue {
	scale := make([]HarmonicScaleValue, 0)
	fmt.Println(options)
	for i := 0 - options.StepsDown; i < options.Steps; i++ {
		for j := 0; j < options.Notes; j++ {
			scaleItem := options.Base * math.Pow(options.Ratio, float64(i*options.Notes+j)/float64(options.Notes))
			scale = append(scale, HarmonicScaleValue{
				Size:      scaleItem,
				Frequency: i,
				Note:      j,
			})
		}
	}

	return scale
}

func (r HarmonicSeriesResolver) Resolve(data map[string]interface{}, state map[string]interface{}) (map[string]interface{}, error) {
	outputState := make(map[string]interface{})
	// generate the series
	inputs := InputsToFloat(state)
	base, _ := inputs["base"]
	steps, _ := inputs["steps"]
	stepsDown, _ := inputs["stepsDown"]
	ratio, _ := inputs["ratio"]
	notes, _ := inputs["notes"]

	scale := calculateHarmonicScale(struct {
		StepsDown int
		Steps     int
		Notes     int
		Base      float64
		Ratio     float64
	}{
		StepsDown: int(stepsDown),
		Steps:     int(steps),
		Notes:     int(notes),
		Base:      base,
		Ratio:     ratio,
	})

	outputState["array"] = scale
	for _, item := range scale {
		outputState[fmt.Sprintf("%d-%d", item.Frequency, item.Note)] = item.Size
	}

	return outputState, nil
}

func (r HarmonicSeriesResolver) Validate(data map[string]interface{}, state map[string]interface{}) error {
	if _, ok := state["base"]; !ok {
		return errors.New("Missing required field 'base'")
	}
	if _, ok := state["steps"]; !ok {
		return errors.New("Missing required field 'steps'")
	}
	if _, ok := state["stepsDown"]; !ok {
		return errors.New("Missing required field 'stepsDown'")
	}
	if _, ok := state["ratio"]; !ok {
		return errors.New("Missing required field 'ratio'")
	}
	return nil
}
