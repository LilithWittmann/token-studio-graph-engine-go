package token_studio_graph_engine

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestInitGraphFromJSON(t *testing.T) {
	content, _ := os.ReadFile("fixtures/math.json")
	g, _ := NewGraph(content)
	fmt.Println(g.Nodes)
	fmt.Println(g.Edges)
}

func TestExecute(t *testing.T) {
	content, _ := os.ReadFile("fixtures/math.json")
	graph, _ := NewGraph(content)
	input := make(map[string]interface{})
	result, _ := graph.Execute(input)
	fmt.Println(result)
}

func TestExecuteNoInput(t *testing.T) {
	content, _ := os.ReadFile("fixtures/noInput.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	_, err := g.Execute(input)
	if err == nil {
		t.Fatalf(`excecute without input didn't trigger an error (%v)`, err)
	}
}

func TestExecuteNoOutput(t *testing.T) {
	content, _ := os.ReadFile("fixtures/noOutput.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	_, err := g.Execute(input)
	if err == nil {
		t.Fatalf(`excecute without output didn't trigger an error (%v)`, err)
	}
}

func TestUnknownNode(t *testing.T) {
	content, _ := os.ReadFile("fixtures/unknownNode.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	_, err := g.Execute(input)
	if err == nil {
		t.Fatalf(`unknown node didn't trigger an error (%v)`, err)
	}
}

func TestComplexJSON(t *testing.T) {
	content, _ := os.ReadFile("fixtures/math.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	_, err := g.Execute(input)
	if err != nil {
		t.Fatalf(`should not trigger an error but did (%v)`, err)
	}
}

func TestMultipleOutputs(t *testing.T) {
	content, _ := os.ReadFile("fixtures/multipleOutputs.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	result, _ := g.Execute(input)
	if !(result["number"] == 0.33333333333333215 && result["second"] == 18.333333333333332) {
		t.Fatalf(`Wrong result (%v)`, result)
	}
}

func TestLogicNodes(t *testing.T) {
	content, _ := os.ReadFile("fixtures/randomLogic.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	result, _ := g.Execute(input)
	foo, _ := result["foo"].(bool)
	output_1, _ := result["output_1"]
	output_2, _ := strconv.ParseInt(result["output_2"].(string), 10, 64)
	if !(foo == true && output_1 == false && output_2 == 1) {
		t.Fatalf(`Wrong result (%v)`, result)
	}
}

func TestInputNodes(t *testing.T) {
	content, _ := os.ReadFile("fixtures/InputNodes.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	result, _ := g.Execute(input)
	if !(result["enumerated"] == "bar" && result["constant"] == "4" && result["slider"] == 2.5) {
		t.Fatalf(`Wrong result (%v)`, result)
	}
}

func TestStringNodes(t *testing.T) {
	content, _ := os.ReadFile("fixtures/StringNodes.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	result, _ := g.Execute(input)
	if !(result["regex_out"] == "aZbZcZ" && result["upper"] == "UPPER" && result["lower"] == "lower" && result["pixels"] == "42px") {
		t.Fatalf(`Wrong result (%v)`, result)
	}
}

func TestArrayNodes(t *testing.T) {
	content, _ := os.ReadFile("fixtures/ArrayNodes.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	result, _ := g.Execute(input)
	if !(result["count"] == 3) {
		t.Fatalf(`Wrong result (%v)`, result)
	}
}

func TestSeriesNodes(t *testing.T) {
	content, _ := os.ReadFile("fixtures/SeriesNodes.json")
	g, _ := NewGraph(content)
	input := make(map[string]interface{})
	result, _ := g.Execute(input)
	fmt.Println(result)
	if result["first_step_arithmetic"] != 14.0 || result["first_step_harmonic"] != 7.111111111111111 || result["last_step_arithmetic"] != 17.0 || result["last_step_harmonic"] != 29.393876913398135 {
		t.Fatalf(`Wrong result (%v)`, result)
	}

}

func TestExecuteWithInput(t *testing.T) {
	fmt.Println("Test with input")
	content, _ := os.ReadFile("fixtures/math.json")
	graph, _ := NewGraph(content)
	input := make(map[string]interface{})
	input["number"] = 12
	result, _ := graph.Execute(input)
	if result["number"] != 1.3333333333333357 {
		t.Fatalf(`Wrong result (%v)`, result)
	}
}
