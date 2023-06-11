package types

import (
	"fmt"
	"github.com/dominikbraun/graph"
	"sync"
)

type ExternalLoader func(request interface{}) (interface{}, error)

type FlowGraph struct {
	Nodes []Node
	Edges []Edge
	State map[string]interface{}
}

type MinimizedNode struct {
	ID   string
	Type string
	Data interface{}
}

type MinimizedEdge struct {
	ID           string
	Source       string
	Target       string
	SourceHandle string
	TargetHandle string
}

type MinimizedFlowGraph struct {
	Nodes []MinimizedNode
	Edges []MinimizedEdge
}

type Lookup map[string]NodeDefinition

type Terminals struct {
	Input  *MinimizedNode
	Output *MinimizedNode
}

type ExecuteOptions struct {
	Graph          MinimizedFlowGraph
	InputValues    map[string]interface{}
	Nodes          []NodeDefinition
	ExternalLoader ExternalLoader
}

type StateTracker struct {
	sync.Mutex
	Data map[string]map[string]interface{}
}

var stateTracker StateTracker

func convertFlowGraphToGraphlib(inputGraph MinimizedFlowGraph) graph.Graph[string, MinimizedNode] {

	nodeHash := func(n MinimizedNode) string {
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

func findTerminals(inputGraph MinimizedFlowGraph) Terminals {
	terminals := Terminals{}

	for _, node := range inputGraph.Nodes {
		switch node.Type {
		case "INPUT":
			terminals.Input = &node
		case "OUTPUT":
			terminals.Output = &node
		}
	}

	return terminals
}

func flowGraphToNodeLookup(inputGraph MinimizedFlowGraph) map[string]MinimizedNode {
	nodeLookup := make(map[string]MinimizedNode)

	for _, node := range inputGraph.Nodes {
		nodeLookup[node.ID] = node
	}

	return nodeLookup
}

func flowGraphToEdgeLookup(inputGraph MinimizedFlowGraph) map[string]MinimizedEdge {
	edgeLookup := make(map[string]MinimizedEdge)

	for _, edge := range inputGraph.Edges {
		edgeLookup[edge.ID] = edge
	}

	return edgeLookup
}

//export minimizeFlowGraph
func minimizeFlowGraph(inputGraph FlowGraph) MinimizedFlowGraph {
	state := inputGraph.State
	nodeLookup := make(map[string]bool)
	nodes := make([]MinimizedNode, len(inputGraph.Nodes))

	for i, node := range inputGraph.Nodes {
		nodeLookup[node.ID] = true
		nodes[i] = MinimizedNode{
			ID:   node.ID,
			Data: state[node.ID],
			Type: string(node.Type),
		}
	}

	edges := make([]MinimizedEdge, 0)

	for _, edge := range inputGraph.Edges {
		if !nodeLookup[edge.Source] || !nodeLookup[edge.Target] {
			fmt.Println("Edge is not connected to a node. Ignoring...", edge)
			continue
		}

		edges = append(edges, MinimizedEdge{
			ID:           edge.ID,
			Target:       edge.Target,
			Source:       edge.Source,
			SourceHandle: edge.SourceHandle,
			TargetHandle: edge.TargetHandle,
		})
	}

	return MinimizedFlowGraph{
		Nodes: nodes,
		Edges: edges,
	}
}

func defaultMapOutput(input, state, processed interface{}) interface{} {
	return map[string]interface{}{
		"output": processed,
	}
}

/*
func execute(opts ExecuteOptions) (interface{}, error) {
	currentGraph := opts.Graph
	inputValues := opts.InputValues
	nodes := opts.Nodes
	externalLoader := opts.ExternalLoader

	g := convertFlowGraphToGraphlib(currentGraph)
	lookup := make(Lookup)

	for _, node := range nodes {
		lookup[node.Type] = node
	}

	nodeLookup := flowGraphToNodeLookup(currentGraph)
	edgeLookup := flowGraphToEdgeLookup(currentGraph)
	terminals := findTerminals(currentGraph)

	if terminals.Input == nil {
		return nil, errors.New("No input node found")
	}

	if terminals.Output == nil {
		return nil, errors.New("No output node found")
	}

	topologic, _ := graph.TopologicalSort(g)
	stateTracker := StateTracker{Data: make(map[string]map[string]interface{})}

	for _, nodeID := range topologic {
		node, ok := nodeLookup[nodeID.ID]

		if !ok {
			continue
		}

		nodeType, ok := lookup[node.Type]

		if !ok {
			return nil, errors.New("No node type found for " + node.Type)
		}

		if node.Type == "INPUT" {
			stateTracker.Lock()
			stateTracker.Data[nodeID.ID] = map[string]interface{}{
				"input": inputValues,
			}
			stateTracker.Unlock()
		}

		stateTracker.Lock()
		input := stateTracker.Data[nodeID.ID]["input"]
		stateTracker.Unlock()
		mappedInput := input

		if nodeType.MapInput != nil {
			mappedInput = nodeType.MapInput(input, node.Data)
		}

		if nodeType.ValidateInputs != nil {
			if err := nodeType.ValidateInputs(mappedInput, node.Data); err != nil {
				return nil, errors.New("Validation failed for node " + nodeID.ID + " of type " + node.Type + " with error: " + err)
				// Todo: Add flow trace
			}
		}

		if nodeType.External && externalLoader == nil {
			return nil, errors.New("Node " + nodeID + " of type " + node.Type + " requires an external loader")
		}

		ephemeral := map[string]interface{}{}

		if nodeType.External && externalLoader != nil {
			ephemeralRequest := nodeType.External(mappedInput, node.Data)
			var err error
			ephemeral, err = externalLoader(ephemeralRequest)

			if err != nil {
				return nil, err
			}
		}

		output, err := nodeType.Process(mappedInput, node.Data, ephemeral)

		if err != nil {
			return nil, err
		}

		if outputPromise, ok := output.(Promise); ok {
			output, err = outputPromise.Await()

			if err != nil {
				return nil, err
			}
		}

		mapOutput := nodeType.MapOutput

		if mapOutput == nil {
			mapOutput = defaultMapOutput
		}

		output = mapOutput(mappedInput, node.Data, output, ephemeral)

		stateTracker.Lock()
		stateTracker.Data[nodeID.ID]["output"] = output
		stateTracker.Unlock()

		outgoing, _ := g.PredecessorMap()

		for _, edge := range outgoing[nodeID] {
			edgeID := g.Edge(edge.V, edge.W, edge.Name)
			edgeData, ok := edgeLookup[edgeID]

			if !ok {
				return nil, errors.New("No edge data found for " + edgeID.Properties.Attributes["id"])
			}

			stateTracker.Lock()
			outputValue := stateTracker.Data[nodeID.ID]["output"].(map[string]interface{})[edgeData.SourceHandle]
			affectedNode := stateTracker.Data[edge.W]

			if affectedNode != nil {
				affectedNode["input"].(map[string]interface{})[edgeData.TargetHandle] = outputValue
			}

			stateTracker.Unlock()
		}
	}

	//stateTracker.RLock()
	output := stateTracker.Data[terminals.Output.ID]["output"]
	//stateTracker.RUnlock()

	return output, nil
}
*/
