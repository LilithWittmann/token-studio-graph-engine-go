package token_studio_graph_engine

import (
	"errors"
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/lilithwittmann/token-studio-graph-engine-go/nodes"
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
			graph.EdgeAttribute("SourceHandle", edge.SourceHandle),
			graph.EdgeAttribute("TargetHandle", edge.TargetHandle))
	}
	return g

}

func findTerminals(inputGraph Graph) (Terminals, error) {
	terminals := Terminals{}
	// Check and map input and output and validate if there are non-existing nodes
	for _, node := range inputGraph.Nodes {
		switch node.Type {
		case nodes.INPUT:
			terminalInput := node
			terminals.Input = &terminalInput
		case nodes.OUTPUT:
			terminalOutput := node
			terminals.Output = &terminalOutput
		default:
			/*
				if !(nodes.GetSupportedNodes()[node.Type]) {
					return terminals, errors.New("Unkonwn node type '" + string(node.Type) + "'")
				}*/
			fmt.Print("Unknown node type: ")
		}
	}

	return terminals, nil
}

//export execute
func execute(json_graph []byte) (map[string]interface{}, error) {
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

		// create a source target map
		sourceTargetMap := map[string][]graph.Edge[string]{}
		edges, _ := connectedGraph.Edges()
		for _, edge := range edges {
			sourceTargetMap[edge.Source] = append(sourceTargetMap[edge.Source], edge)
		}

		fmt.Println(node.Type)
		result, err := nodes.GetSupportedNodes()[node.Type].Resolve(node.Data, inputGraph.State[node.ID])
		if err != nil {
			return nil, err
		}

		if node.ID == terminals.Output.ID {
			return result, nil
		}

		// update the state for all nodes that have this one as input
		for _, edge := range sourceTargetMap[node.ID] {
			// get the node that is connected to the current node
			connectedNode, _ := connectedGraph.Vertex(edge.Target)
			// update the state of the connected node
			if inputGraph.State[connectedNode.ID] == nil {
				inputGraph.State[connectedNode.ID] = make(map[string]interface{})
			}
			inputGraph.State[connectedNode.ID][edge.Properties.Attributes["TargetHandle"]] = result[edge.Properties.Attributes["SourceHandle"]]
			fmt.Println("Updated state of the handle " + edge.Properties.Attributes["TargetHandle"] + " of the node " + connectedNode.ID + " with value " + fmt.Sprint(result))
		}

	}
	return nil, nil

}
