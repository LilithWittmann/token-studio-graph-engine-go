package token_studio_graph_engine

import "C"

import nodes "github.com/lilithwittmann/token-studio-graph-engine-go/nodes"

//export Node
type Node struct {
	ID   string                 `json:"id"`
	Type nodes.NodeTypes        `json:"type"`
	Data map[string]interface{} `json:"data"`
}

//export Edge
type Edge struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Source       string `json:"source"`
	SourceHandle string `json:"sourceHandle"`
	Target       string `json:"target"`
	TargetHandle string `json:"targetHandle"`
}

//export Graph
type Graph struct {
	Nodes []Node                            `json:"nodes"`
	Edges []Edge                            `json:"edges"`
	State map[string]map[string]interface{} `json:"state"`
}
