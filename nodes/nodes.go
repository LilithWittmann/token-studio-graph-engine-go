package nodes

type NodeResolver interface {
	Resolve(map[string]interface{}, map[string]interface{}) (map[string]interface{}, error)
	Validate(map[string]interface{}, map[string]interface{}) error
}

// create a node resolver interface

//export NodeTypes
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

func GetSupportedNodes() map[NodeTypes]NodeResolver {
	return map[NodeTypes]NodeResolver{
		ADD: AdditionResolver{}, SUBTRACT: SubtractResolver{}, MULTIPLY: MultiplyResolver{}, DIV: DivideResolver{}, MOD: ModuloResolver{}, ABS: AbsoluteValueResolver{}, ROUND: RoundResolver{}, SIN: SineResolver{}, COS: CosineResolver{}, TAN: TangentResolver{}, LERP: LerpResolver{}, CLAMP: ClampResolver{}, RANDOM: RandomResolver{}, COUNT: CountResolver{},
		CONSTANT: ConstantResolver{}, ENUMERATED_INPUT: EnumeratedConstantResolver{}, SLIDER: SliderResolver{},
		IF: IfResolver{}, NOT: NotResolver{}, AND: AndResolver{}, OR: OrResolver{}, SWITCH: SwitchResolver{}, COMPARE: CompareResolver{},
		INPUT: InputResolver{}, OUTPUT: OutputResolver{},
		UPPERCASE: UppercaseResolver{}, LOWER: LowercaseResolver{}, REGEX: RegexResolver{}, PASS_UNIT: PassUnitResolver{},
	}
}
