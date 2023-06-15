package token_studio_graph_engine

import (
	"encoding/json"
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

type GraphCache struct {
	Graphs           map[string]*Graph
	ExecutionResults map[string]map[string]string
}

//export NewGraphCache
func NewGraphCache() (*GraphCache, error) {
	var g *GraphCache
	return g, nil
}

//export InitializeGraph
func (cache *GraphCache) InitializeGraph(graphID string, jsonInput []byte) error {
	graph, err := NewGraph(jsonInput)
	if err != nil {
		return err
	}
	cache.Graphs[graphID] = graph
	return nil
}

//export ExecuteGraph
func (cache *GraphCache) ExecuteGraph(graphID string, graphInputState []byte) ([]byte, error) {
	input := make(map[string]interface{})
	err := json.Unmarshal(graphInputState, &input)

	//check if graph with the ID graphID is loaded
	if cache.Graphs[graphID] == nil {
		return nil, errors.New("graph with ID " + graphID + " not loaded")
	}

	result, err := cache.Graphs[graphID].ExecuteToJSON(input)

	// add graph result to cache

	return result, err

}

var stateTracker StateTracker

func (inputGraph *Graph) convertGraphToGraphlib() graph.Graph[string, Node] {

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

func (inputGraph *Graph) findTerminals() (Terminals, error) {
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
		}
	}

	return terminals, nil
}

//export NewGraph
func NewGraph(jsonInput []byte) (*Graph, error) {
	var g *Graph
	err := json.Unmarshal(jsonInput, &g)

	return g, err
}

// Export ExecuteToJSON
func (inputGraph *Graph) ExecuteToJSON(input map[string]interface{}) ([]byte, error) {
	result, err := inputGraph.Execute(input)
	jsonResult, _ := json.Marshal(result)
	return jsonResult, err
}

//export Execute
func (inputGraph *Graph) Execute(input map[string]interface{}) (map[string]interface{}, error) {

	connectedGraph := inputGraph.convertGraphToGraphlib()
	fmt.Println(inputGraph)

	// find the start and endpoint of our Graph and ensure that it actually exist
	terminals, err := inputGraph.findTerminals()
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
	// sort the graph topologically so we can Execute it
	topologicSortedGraph, _ := graph.TopologicalSort(connectedGraph)

	fmt.Println("Topologic:")

	// create a source target map
	sourceTargetMap := map[string][]Edge{}
	for _, edge := range inputGraph.Edges {
		fmt.Println(edge)
		sourceTargetMap[edge.Source] = append(sourceTargetMap[edge.Source], edge)
	}

	// go through the topologic sorted graph and Execute the nodes
	for _, nodeID := range topologicSortedGraph {

		fmt.Println(nodeID)
		node, _ := connectedGraph.Vertex(nodeID)

		fmt.Println(node.Type)
		// check if resolver exists
		if nodes.GetSupportedNodes()[node.Type] == nil {
			return nil, errors.New("No resolver found for node type '" + string(node.Type) + "'")
		}

		// if nod is input node add the graph input state to the node
		if node.ID == terminals.Input.ID && inputGraph.State[node.ID]["values"] != nil {

			for k, v := range input {
				inputGraph.State[node.ID]["values"].(map[string]interface{})[k] = v
			}

			fmt.Println(inputGraph.State[node.ID]["values"])

		}

		// resolve the node
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

			inputGraph.State[connectedNode.ID][edge.TargetHandle] = result[edge.SourceHandle]
			fmt.Println("Updated state of the handle " + edge.TargetHandle + " of the node " + connectedNode.ID + " with value " + fmt.Sprint(result))
		}

	}
	return nil, nil

}
