package govaluate

type OperatorSymbol int

const (
	VALUE          OperatorSymbol = iota // 用于表示值的操作符标志
	LITERAL                              // 用于表示文字字面量的操作符标志
	NOOP                                 // 用于表示空操作的操作符标志
	EQ                                   // 相等比较操作符标志
	NEQ                                  // 不等比较操作符标志
	GT                                   // 大于比较操作符标志
	LT                                   // 小于比较操作符标志
	GTE                                  // 大于等于比较操作符标志
	LTE                                  // 小于等于比较操作符标志
	IN                                   // 包含操作符标志
	NOTIN                                // 包含操作符标志
	REQ                                  // 正则表达式匹配操作符标志
	NREQ                                 // 正则表达式不匹配操作符标志
	AND                                  // 逻辑与操作符标志
	OR                                   // 逻辑或操作符标志
	MODULUS                              // 取余操作符标志
	EXPONENT                             // 指数操作符标志
	PLUS                                 // 加法操作符标志
	MINUS                                // 减法操作符标志
	MULTIPLY                             // 乘法操作符标志
	DIVIDE                               // 除法操作符标志
	BITWISE_AND                          // 位与操作符标志
	BITWISE_OR                           // 位或操作符标志
	BITWISE_XOR                          // 位异或操作符标志
	BITWISE_LSHIFT                       // 位左移操作符标志
	BITWISE_RSHIFT                       // 位右移操作符标志
	BITWISE_NOT                          // 位非操作符标志
	INVERT                               // 位反操作符标志
	NEGATE                               // 取反(非)操作符标志
	TERNARY_TRUE                         // 28: 三元条件操作符中的真条件标志
	TERNARY_FALSE                        // 29: 三元条件操作符中的假条件标志
	COALESCE                             // 30: 合并操作符标志
	FUNCTIONAL                           // 31: 函数调用操作符标志
	ACCESS                               // 32: 访问操作符标志
	SEPARATE                             // 33: 分隔操作符标志
)

type operatorPrecedence int // 操作符的优先级

const (
	noopPrecedence           operatorPrecedence = iota // 无操作的优先级
	valuePrecedence                                    // 值的优先级
	functionalPrecedence                               // 函数的优先级
	prefixPrecedence                                   // 前缀操作符的优先级
	exponentialPrecedence                              // 指数操作符的优先级
	additivePrecedence                                 // 加法操作符的优先级
	bitwisePrecedence                                  // 位操作符的优先级
	bitwiseShiftPrecedence                             // 位移操作符的优先级
	multiplicativePrecedence                           // 乘法操作符的优先级
	comparatorPrecedence                               // 比较操作符的优先级
	ternaryPrecedence                                  // 三元条件操作符的优先级
	logicalAndPrecedence                               // 逻辑与操作符的优先级
	logicalOrPrecedence                                // 逻辑或操作符的优先级
	separatePrecedence                                 // 分割操作符的优先级

)

func findOperatorPrecedenceForSymbol(symbol OperatorSymbol) operatorPrecedence {

	switch symbol {
	case NOOP:
		return noopPrecedence
	case VALUE:
		return valuePrecedence
	case EQ:
		fallthrough
	case NEQ:
		fallthrough
	case GT:
		fallthrough
	case LT:
		fallthrough
	case GTE:
		fallthrough
	case LTE:
		fallthrough
	case REQ:
		fallthrough
	case NREQ:
		fallthrough
	case IN:
		return comparatorPrecedence
	case AND:
		return logicalAndPrecedence
	case OR:
		return logicalOrPrecedence
	case BITWISE_AND:
		fallthrough
	case BITWISE_OR:
		fallthrough
	case BITWISE_XOR:
		return bitwisePrecedence
	case BITWISE_LSHIFT:
		fallthrough
	case BITWISE_RSHIFT:
		return bitwiseShiftPrecedence
	case PLUS:
		fallthrough
	case MINUS:
		return additivePrecedence
	case MULTIPLY:
		fallthrough
	case DIVIDE:
		fallthrough
	case MODULUS:
		return multiplicativePrecedence
	case EXPONENT:
		return exponentialPrecedence
	case BITWISE_NOT:
		fallthrough
	case NEGATE:
		fallthrough
	case INVERT:
		return prefixPrecedence
	case COALESCE:
		fallthrough
	case TERNARY_TRUE:
		fallthrough
	case TERNARY_FALSE:
		return ternaryPrecedence
	case ACCESS:
		fallthrough
	case FUNCTIONAL:
		return functionalPrecedence
	case SEPARATE:
		return separatePrecedence
	}

	return valuePrecedence
}

// 比较
var comparatorSymbols = map[string]OperatorSymbol{
	"==": EQ,
	"!=": NEQ,
	">":  GT,
	">=": GTE,
	"<":  LT,
	"<=": LTE,
	"=~": REQ,
	"!~": NREQ,
	"in": IN,
}

var logicalSymbols = map[string]OperatorSymbol{
	"&&": AND,
	"||": OR,
}

var bitwiseSymbols = map[string]OperatorSymbol{
	"^": BITWISE_XOR,
	"&": BITWISE_AND,
	"|": BITWISE_OR,
}

var bitwiseShiftSymbols = map[string]OperatorSymbol{
	">>": BITWISE_RSHIFT,
	"<<": BITWISE_LSHIFT,
}

var additiveSymbols = map[string]OperatorSymbol{
	"+": PLUS,
	"-": MINUS,
}

var multiplicativeSymbols = map[string]OperatorSymbol{
	"*": MULTIPLY,
	"/": DIVIDE,
	"%": MODULUS,
}

var exponentialSymbolsS = map[string]OperatorSymbol{
	"**": EXPONENT,
}

var prefixSymbols = map[string]OperatorSymbol{
	"-": NEGATE,
	"!": INVERT,
	"~": BITWISE_NOT,
}

var ternarySymbols = map[string]OperatorSymbol{
	"?":  TERNARY_TRUE,
	":":  TERNARY_FALSE,
	"??": COALESCE,
}

// 计算
var modifierSymbols = map[string]OperatorSymbol{
	"+":  PLUS,
	"-":  MINUS,
	"*":  MULTIPLY,
	"/":  DIVIDE,
	"%":  MODULUS,
	"**": EXPONENT,
	"&":  BITWISE_AND,
	"|":  BITWISE_OR,
	"^":  BITWISE_XOR,
	">>": BITWISE_RSHIFT,
	"<<": BITWISE_LSHIFT,
}

var separatorSymbols = map[string]OperatorSymbol{
	",": SEPARATE,
}

// IsModifierType 判断运算符是否已定义
func (this OperatorSymbol) IsModifierType(candidate []OperatorSymbol) bool {

	for _, symbolType := range candidate {
		if this == symbolType {
			return true
		}
	}

	return false
}

/*
通常用于格式化类型检查错误时。
我们可以将字符串化的符号存储在其他地方，而不需要重复的代码块来翻译
操作符符号到字符串，但这将需要更多的内存，并在某处的另一个字段。
添加操作符很少见，所以我们在这里只对其进行字符串化。
*/
func (this OperatorSymbol) String() string {

	switch this {
	case NOOP:
		return "NOOP"
	case VALUE:
		return "VALUE"
	case EQ:
		return "="
	case NEQ:
		return "!="
	case GT:
		return ">"
	case LT:
		return "<"
	case GTE:
		return ">="
	case LTE:
		return "<="
	case REQ:
		return "=~"
	case NREQ:
		return "!~"
	case AND:
		return "&&"
	case OR:
		return "||"
	case IN:
		return "in"
	case BITWISE_AND:
		return "&"
	case BITWISE_OR:
		return "|"
	case BITWISE_XOR:
		return "^"
	case BITWISE_LSHIFT:
		return "<<"
	case BITWISE_RSHIFT:
		return ">>"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case MULTIPLY:
		return "*"
	case DIVIDE:
		return "/"
	case MODULUS:
		return "%"
	case EXPONENT:
		return "**"
	case NEGATE:
		return "-"
	case INVERT:
		return "!"
	case BITWISE_NOT:
		return "~"
	case TERNARY_TRUE:
		return "?"
	case TERNARY_FALSE:
		return ":"
	case COALESCE:
		return "??"
	}
	return ""
}
