package token_studio_graph_engine

import (
	"fmt"
	"os"
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
	execute(content)
}

func TestExecuteNoInput(t *testing.T) {
	content, _ := os.ReadFile("fixtures/noInput.json")
	_, err := execute(content)
	if err == nil {
		t.Fatalf(`excecute without input didn't trigger an error (%v)`, err)
	}
}

func TestExecuteNoOutput(t *testing.T) {
	content, _ := os.ReadFile("fixtures/noOutput.json")
	_, err := execute(content)
	if err == nil {
		t.Fatalf(`excecute without output didn't trigger an error (%v)`, err)
	}
}

func TestUnknownNode(t *testing.T) {
	content, _ := os.ReadFile("fixtures/unknownNode.json")
	_, err := execute(content)
	if err == nil {
		t.Fatalf(`unknown node didn't trigger an error (%v)`, err)
	}
}

func TestComplexJSON(t *testing.T) {
	content, _ := os.ReadFile("fixtures/math.json")
	_, err := execute(content)
	if err != nil {
		t.Fatalf(`should not trigger an error but did (%v)`, err)
	}
}

func TestMultipleOutputs(t *testing.T) {
	content, _ := os.ReadFile("fixtures/multipleOutputs.json")
	result, _ := execute(content)
	fmt.Println(result)
}
