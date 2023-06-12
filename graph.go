package token_studio_graph_engine

import (
	"errors"
	"fmt"
	"github.com/dominikbraun/graph"
	"sync"
)

type ExternalLoader func(request interface{}) (interface{}, error)

type Terminals struct {
	Input  *Node
	Output *Node
}

type StateTracker struct {
	sync.Mutex
	Data map[string]map[string]interface{}
}

var stateTracker StateTracker

func convertGraphToGraphlib(inputGraph Graph) graph.Graph[string, Node] {

	nodeHash := func(n Node) string {
		return n.ID
	}
	g := graph.New(nodeHash, graph.Directed(), graph.Acyclic())

	for _, node := range inputGraph.Nodes {
		g.AddVertex(node)
	}

	for _, edge := range inputGraph.Edges {
		g.AddEdge(edge.Source, edge.Target,
			graph.EdgeAttribute("ID", edge.ID),
			graph.EdgeAttribute("SourceHandle", edge.SourceHandle))
	}
	return g

}

func findTerminals(inputGraph Graph) (Terminals, error) {
	terminals := Terminals{}
	// Check and map input and output and validate if there are non-existing nodes
	for _, node := range inputGraph.Nodes {
		switch node.Type {
		case INPUT:
			terminalInput := node
			terminals.Input = &terminalInput
		case OUTPUT:
			terminalOutput := node
			terminals.Output = &terminalOutput
		default:
			if !(getSupportedNodes()[node.Type]) {
				return terminals, errors.New("Unkonwn node type '" + string(node.Type) + "'")
			}
		}
	}

	return terminals, nil
}

//export execute
func execute(json_graph []byte) ([]byte, error) {
	inputGraph, err := NewGraph(json_graph)
	fmt.Println("State")
	fmt.Println(inputGraph.State)
	// this usually means that we don't support the input graph
	if err != nil {
		return nil, err
	}

	connectedGraph := convertGraphToGraphlib(inputGraph)
	fmt.Println(inputGraph)

	// find the start and endpoint of our Graph and ensure that it actually exist
	terminals, err := findTerminals(inputGraph)
	if err != nil {
		return nil, err
	}

	if terminals.Input == nil {
		return nil, errors.New("No input node found")
	}

	if terminals.Output == nil {
		return nil, errors.New("No output node found")
	}

	fmt.Println(connectedGraph.Edges())
	fmt.Println("Input")
	fmt.Println(terminals.Input)
	fmt.Println("Output")
	fmt.Println(terminals.Output)
	// sort the graph topologically so we can execute it
	topologicSortedGraph, _ := graph.TopologicalSort(connectedGraph)

	fmt.Println("Topologic:")
	// go through the topologic sorted graph and execute the nodes
	for _, nodeID := range topologicSortedGraph {
		fmt.Println(nodeID)
		node, _ := connectedGraph.Vertex(nodeID)

		if inputGraph.State[node.ID] != nil {
			fmt.Println(inputGraph.State[node.ID])
		}
		
		fmt.Println(node.Type)
	}
	return json_graph, nil

}
