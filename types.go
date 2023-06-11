package token_studio_graph_engine

import "C"

type Node struct {
	ID   string      `json:"id"`
	Type NodeTypes   `json:"type"`
	Data interface{} `json:"data"`
}

type Edge struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Source       string `json:"source"`
	SourceHandle string `json:"sourceHandle"`
	Target       string `json:"target"`
	TargetHandle string `json:"targetHandle"`
}

type JSONGraph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type NodeDefinition struct {
	Type           string
	Defaults       interface{}
	External       func(mappedInput interface{}, state interface{}) interface{}
	MapInput       func(input interface{}, state interface{}) interface{}
	ValidateInputs func(input interface{}, state interface{})
	Process        func(input interface{}, state interface{}, ephemeral interface{}) interface{}
	MapOutput      func(input interface{}, state interface{}, output interface{}, ephemeral interface{}) interface{}
}

type NodeTypes string

const (
	INPUT  NodeTypes = "studio.tokens.generic.input"
	OUTPUT NodeTypes = "studio.tokens.generic.output"

	ENUMERATED_INPUT NodeTypes = "studio.tokens.input.enumerated-constant"
	CONSTANT         NodeTypes = "studio.tokens.input.constant"
	SLIDER           NodeTypes = "studio.tokens.input.slider"

	CSS_MAP NodeTypes = "studio.tokens.css.map"

	// Logic
	IF      NodeTypes = "studio.tokens.logic.if"
	NOT     NodeTypes = "studio.tokens.logic.not"
	AND     NodeTypes = "studio.tokens.logic.and"
	OR      NodeTypes = "studio.tokens.logic.or"
	SWITCH  NodeTypes = "studio.tokens.logic.switch"
	COMPARE NodeTypes = "studio.tokens.logic.compare"

	// Array
	ARRAY_INDEX NodeTypes = "studio.tokens.array.index"
	ARRIFY      NodeTypes = "studio.tokens.array.arrify"
	REVERSE     NodeTypes = "studio.tokens.array.reverse"
	SLICE       NodeTypes = "studio.tokens.array.slice"
	JOIN        NodeTypes = "studio.tokens.array.join"

	// Math
	ADD      NodeTypes = "studio.tokens.math.add"
	SUBTRACT NodeTypes = "studio.tokens.math.subtract"
	MULTIPLY NodeTypes = "studio.tokens.math.multiply"
	DIV      NodeTypes = "studio.tokens.math.divide"
	ABS      NodeTypes = "studio.tokens.math.abs"
	ROUND    NodeTypes = "studio.tokens.math.round"
	SIN      NodeTypes = "studio.tokens.math.sin"
	COS      NodeTypes = "studio.tokens.math.cos"
	TAN      NodeTypes = "studio.tokens.math.tan"
	LERP     NodeTypes = "studio.tokens.math.lerp"
	CLAMP    NodeTypes = "studio.tokens.math.clamp"
	MOD      NodeTypes = "studio.tokens.math.mod"
	RANDOM   NodeTypes = "studio.tokens.math.random"
	COUNT    NodeTypes = "studio.tokens.math.count"

	// Color
	SCALE           NodeTypes = "studio.tokens.color.scale"
	BLEND           NodeTypes = "studio.tokens.color.blend"
	ADVANCED_BLEND  NodeTypes = "studio.tokens.color.blendAdv"
	CREATE_COLOR    NodeTypes = "studio.tokens.color.create"
	EXTRACT         NodeTypes = "studio.tokens.color.extract"
	TRANSFORM_COLOR NodeTypes = "studio.tokens.color.transform"

	//Sets
	FLATTEN    NodeTypes = "studio.tokens.sets.flatten"
	ALIAS      NodeTypes = "studio.tokens.sets.alias"
	REMAP      NodeTypes = "studio.tokens.sets.remap"
	INLINE_SET NodeTypes = "studio.tokens.sets.inline"
	SET        NodeTypes = "studio.tokens.sets.external"
	INVERT_SET NodeTypes = "studio.tokens.sets.invert"

	//Series
	ARITHMETIC_SERIES NodeTypes = "studio.tokens.sets.arithmetic"
	HARMONIC_SERIES   NodeTypes = "studio.tokens.sets.harmonic"

	//String
	UPPERCASE NodeTypes = "studio.tokens.string.uppercase"
	LOWER     NodeTypes = "studio.tokens.string.lowercase"
	REGEX     NodeTypes = "studio.tokens.string.regex"
	PASS_UNIT NodeTypes = "studio.tokens.typing.passUnit"

	//Accessibility
	CONTRAST        NodeTypes = "studio.tokens.accessibility.contrast"
	COLOR_BLINDNESS NodeTypes = "studio.tokens.accessibility.colorBlindness"
)
