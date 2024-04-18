package bicep

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/parser/bicep/antlr/parser"
	"github.com/antlr4-go/antlr/v4"
)

type Parser struct {
}

const kicsPrefix = "_kics_"
const kicsLine = kicsPrefix + "line"
const kicsLines = kicsPrefix + "lines"
const kicsArray = kicsPrefix + "arr"

const CloseParenthesis = "')"

type BicepVisitor struct {
	parser.BasebicepVisitor
	paramList    map[string]interface{}
	varList      map[string]interface{}
	resourceList []interface{}
}

type JSONBicep struct {
	Parameters map[string]interface{} `json:"parameters"`
	Variables  map[string]interface{} `json:"variables"`
	Resources  []interface{}          `json:"resources"`
}

type KicsObjectProperty struct {
	objectProperty map[string]interface{}
	line           int
}

func NewBicepVisitor() *BicepVisitor {
	paramList := map[string]interface{}{}
	varList := map[string]interface{}{}
	resourceList := []interface{}{}
	return &BicepVisitor{paramList: paramList, varList: varList, resourceList: resourceList}
}

func convertVisitorToJSONBicep(visitor *BicepVisitor) *JSONBicep {
	return &JSONBicep{
		Parameters: visitor.paramList,
		Variables:  visitor.varList,
		Resources:  visitor.resourceList,
	}
}

// Parse - parses bicep to BicepVisitor template (json file)
func (p *Parser) Parse(file string, _ []byte) ([]model.Document, []int, error) {
	bicepVisitor := NewBicepVisitor()
	stream, _ := antlr.NewFileStream(file)
	lexer := parser.NewbicepLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	bicepParser := parser.NewbicepParser(tokenStream)
	bicepParser.RemoveErrorListeners()
	bicepParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	bicepParser.Program().Accept(bicepVisitor)

	var doc model.Document

	jBicep := convertVisitorToJSONBicep(bicepVisitor)
	bicepBytes, err := json.Marshal(jBicep)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(bicepBytes, &doc)
	if err != nil {
		return nil, nil, err
	}

	return []model.Document{doc}, nil, nil
}

func (s *BicepVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, val := range ctx.AllStatement() {
		val.Accept(s)
	}

	return nil
}

func (s *BicepVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.ParameterDecl() != nil {
		return ctx.ParameterDecl().Accept(s)
	}
	if ctx.VariableDecl() != nil {
		return ctx.VariableDecl().Accept(s)
	}
	if ctx.ResourceDecl() != nil {
		return ctx.ResourceDecl().Accept(s)
	}

	return nil
}

func parseDecorators(decorators []parser.IDecoratorContext, s *BicepVisitor) map[string]interface{} {
	decoratorsMap := map[string]interface{}{}

	for _, val := range decorators {
		decorator, ok := val.Accept(s).(map[string][]interface{})
		if !ok {
			return nil
		}
		for name, values := range decorator {
			if name == "description" {
				metadata := map[string]interface{}{}
				metadata["description"] = values[0]
				decoratorsMap["metadata"] = metadata
			} else if name == "maxLength" || name == "minLength" || name == "minValue" || name == "maxValue" {
				decoratorsMap[name] = values[0]
			} else {
				decoratorsMap[name] = values
			}
		}
	}

	return decoratorsMap
}

func (s *BicepVisitor) VisitParameterDecl(ctx *parser.ParameterDeclContext) interface{} {
	param := map[string]interface{}{}
	identifier, ok := ctx.Identifier().Accept(s).(string)
	if !ok {
		return nil
	}
	if ctx.ParameterDefaultValue() != nil {
		paramVal := ctx.ParameterDefaultValue().Accept(s)
		switch paramVal := paramVal.(type) {
		case map[string][]interface{}:
			stringifiedFunction := parseFunctionCall(paramVal)
			param["defaultValue"] = "[" + stringifiedFunction + "]"
		case interface{}:
			if isDotFunction(paramVal) {
				paramVal = "[" + paramVal.(string) + "]"
			}
			param["defaultValue"] = paramVal
		default:
			param["defaultValue"] = nil
		}
	}
	if ctx.TypeExpression() != nil {
		typeExpression := ctx.TypeExpression().Accept(s)
		param["type"] = typeExpression
	}
	decoratorsMap := parseDecorators(ctx.AllDecorator(), s)
	for name, values := range decoratorsMap {
		if name == "secure" {
			if param["type"] == "string" {
				param["type"] = "secureString"
			} else if param["type"] == "object" {
				param["type"] = "secureObject"
			}
		} else {
			if name == "allowed" {
				param["allowedValues"] = decoratorsMap["allowed"]
			} else {
				param[name] = values
			}
		}
	}

	line := map[string]int{kicsLine: ctx.GetStop().GetLine() + 1}
	lines := map[string]map[string]int{
		kicsPrefix + "defaultValue": line,
		kicsPrefix + "type":         line,
	}

	param[kicsLines] = lines

	s.paramList[identifier] = param
	return nil
}

func (s *BicepVisitor) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	var variable = map[string]interface{}{}
	identifier, ok := ctx.Identifier().Accept(s).(string)
	if !ok {
		return nil
	}
	expression := ctx.Expression().Accept(s)
	decoratorsMap := parseDecorators(ctx.AllDecorator(), s)
	for name, values := range decoratorsMap {
		variable[name] = values
	}

	variable["value"] = expression
	s.varList[identifier] = variable

	return nil
}

func (s *BicepVisitor) VisitResourceDecl(ctx *parser.ResourceDeclContext) interface{} {
	resource := map[string]interface{}{}
	interpString := ctx.InterpString().Accept(s).(string)
	identifier, ok := ctx.Identifier().Accept(s).(string)
	if !ok {
		return nil
	}
	resourceType := strings.Split(interpString, "@")[0]
	apiVersion := strings.Split(interpString, "@")[1]
	resource["type"] = resourceType
	resource["apiVersion"] = apiVersion
	decoratorsMap := parseDecorators(ctx.AllDecorator(), s)
	for name, values := range decoratorsMap {
		resource[name] = values
	}

	resource["identifier"] = identifier
	if ctx.Object() != nil {
		object, ok := ctx.Object().Accept(s).(map[string]interface{})
		if !ok {
			return nil
		}
		for key, val := range object {
			resource[key] = val
		}
	}

	lines, ok := resource[kicsLines].(map[string]interface{})
	if !ok {
		lines = map[string]interface{}{}
	}

	line := map[string]int{kicsLine: ctx.GetStart().GetLine()}
	lines[kicsPrefix+"apiVersion"] = line

	line = map[string]int{kicsLine: ctx.GetStart().GetLine()}
	lines[kicsPrefix+"type"] = line

	s.resourceList = append(s.resourceList, resource)

	return nil
}

func (s *BicepVisitor) VisitParameterDefaultValue(ctx *parser.ParameterDefaultValueContext) interface{} {
	param := ctx.Expression().Accept(s)
	return param
}

/*
Converts functioncall data (map of identifying string to slice of arguments) into a string

	Example: "FunctionName": ["arg1", 2, "arg3", map[Function2: [arg4, arg5]]] becomes
	"FunctionName(arg1, 2, arg3, Function2(arg4, arg5))"
*/
func parseFunctionCall(functionData map[string][]interface{}) string {
	stringifiedFunctionCall := ""

	for functionName, argumentList := range functionData {
		stringifiedFunctionCall += functionName + "("
		for index, argument := range argumentList {
			switch argument := argument.(type) {
			case string:
				stringifiedFunctionCall += argument
			case int:
				convertedArgument := strconv.Itoa(argument)
				stringifiedFunctionCall += convertedArgument
			case map[string][]interface{}:
				stringifiedFunctionCall += parseFunctionCall(argument)
			}

			if index < len(argumentList)-1 {
				stringifiedFunctionCall += ", "
			}
		}
	}
	stringifiedFunctionCall += ")"

	return stringifiedFunctionCall
}

func (s *BicepVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.GetChildCount() > 1 {
		if ctx.Identifier() != nil {
			identifier, ok := ctx.Identifier().Accept(s).(string)
			if !ok {
				return nil
			}
			for variable := range s.varList {
				if variable == identifier {
					identifier = "variables('" + identifier + CloseParenthesis
				}
			}
			for parameter := range s.paramList {
				if parameter == identifier {
					identifier = "parameters('" + identifier + CloseParenthesis
				}
			}
			exp := ctx.Expression(0).Accept(s)
			if ctx.DOT() != nil {
				switch exp := exp.(type) {
				case map[string][]interface{}:
					return parseFunctionCall(exp) + "." + identifier
				case string:
					return exp + "." + identifier
				default:
					return nil
				}
			}
		} else if ctx.LogicCharacter() == nil {
			for _, val := range ctx.AllExpression() {
				val.Accept(s)
			}
		}
	}

	if ctx.PrimaryExpression() != nil {
		return ctx.PrimaryExpression().Accept(s)
	}

	return ""
}

func (s *BicepVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.LiteralValue() != nil {
		return ctx.LiteralValue().Accept(s)
	}
	if ctx.FunctionCall() != nil {
		return ctx.FunctionCall().Accept(s)
	}
	if ctx.InterpString() != nil {
		return ctx.InterpString().Accept(s)
	}
	if ctx.MULTILINE_STRING() != nil {
		finalString := strings.ReplaceAll(ctx.MULTILINE_STRING().GetText(), "'''", "")
		finalString = strings.ReplaceAll(finalString, "\r", "")
		finalString = strings.ReplaceAll(finalString, "\n", "")
		return finalString
	}
	if ctx.Array() != nil {
		return ctx.Array().Accept(s)
	}
	if ctx.Object() != nil {
		return ctx.Object().Accept(s)
	}
	if ctx.ParenthesizedExpression() != nil {
		return ctx.ParenthesizedExpression().Accept(s)
	}

	return nil
}

func (s *BicepVisitor) VisitLiteralValue(ctx *parser.LiteralValueContext) interface{} {
	if ctx.NUMBER() != nil {
		number, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 32)
		return number
	}
	if ctx.TRUE() != nil {
		return true
	}
	if ctx.FALSE() != nil {
		return false
	}
	if ctx.NULL() != nil {
		return nil
	}
	if ctx.Identifier() != nil {
		identifier, ok := ctx.Identifier().Accept(s).(string)
		if !ok {
			return nil
		}
		for variable := range s.varList {
			if variable == identifier {
				identifier = "variables('" + identifier + CloseParenthesis
			}
		}
		for parameter := range s.paramList {
			if parameter == identifier {
				identifier = "parameters('" + identifier + CloseParenthesis
			}
		}
		return identifier
	}

	return nil
}

func (s *BicepVisitor) VisitInterpString(ctx *parser.InterpStringContext) interface{} {
	if ctx.GetChildCount() > 1 {
		interpString := []interface{}{}
		interpString = append(interpString, ctx.STRING_LEFT_PIECE().GetText())
		if ctx.AllSTRING_MIDDLE_PIECE() != nil && (len(ctx.AllSTRING_MIDDLE_PIECE()) > 0) {
			for idx, val := range ctx.AllSTRING_MIDDLE_PIECE() {
				interpString = append(interpString, ctx.Expression(idx).Accept(s), val.GetText())
			}
		}
		// Last expression with string right piece
		interpString = append(interpString,
			ctx.Expression(len(ctx.AllSTRING_MIDDLE_PIECE())).Accept(s),
			ctx.STRING_RIGHT_PIECE().GetText())
		str := ""
		for _, v := range interpString {
			switch v := v.(type) {
			case string:
				str += v
			case map[string][]interface{}:
				for identifier, argumentList := range v {
					resStr := "[" + identifier + "("
					for idx, arg := range argumentList {
						stringArg, ok := arg.(string)
						if !ok {
							return ""
						}
						resStr += stringArg
						if idx < len(argumentList)-1 {
							resStr += ", "
						}
					}

					resStr += ")]"
					str += resStr
				}
			}
		}
		return str
	}

	unformattedString := ctx.STRING_COMPLETE().GetText()
	finalString := strings.ReplaceAll(unformattedString, "'", "")
	return finalString
}

func (s *BicepVisitor) VisitArray(ctx *parser.ArrayContext) interface{} {
	array := []interface{}{}
	for _, val := range ctx.AllArrayItem() {
		expression := val.Accept(s)
		if isParameter(expression) || isDotFunction(expression) {
			expression = "[" + expression.(string) + "]"
		}
		array = append(array, expression)
	}
	return array
}

func (s *BicepVisitor) VisitArrayItem(ctx *parser.ArrayItemContext) interface{} {
	return ctx.Expression().Accept(s)
}

func (s *BicepVisitor) VisitObject(ctx *parser.ObjectContext) interface{} {
	object := map[string]interface{}{}
	propertiesLines := map[string]interface{}{}
	for _, val := range ctx.AllObjectProperty() {
		objectProperty, ok := val.Accept(s).(KicsObjectProperty)
		if !ok {
			return object
		}
		for key, val := range objectProperty.objectProperty {
			object[key] = val
			line := map[string]interface{}{kicsLine: objectProperty.line}

			arr, isArray := val.([]interface{})
			if isArray {
				for range arr {
					arrLine := map[string]int{kicsLine: objectProperty.line}
					kicsDefault := map[string]interface{}{kicsPrefix + "_default": arrLine}
					kicsArr := []interface{}{kicsDefault}
					line[kicsArray] = kicsArr
				}
			}
			propertiesLines[kicsPrefix+key] = line
		}
	}

	defaultLine := map[string]int{kicsLine: ctx.GetStart().GetLine()}
	propertiesLines[kicsPrefix+"_default"] = defaultLine

	object[kicsLines] = propertiesLines

	return object
}

func isParameter(expression interface{}) bool {
	exp, ok := expression.(string)
	if !ok {
		return false
	}

	return strings.Contains(exp, "parameters(") || strings.Contains(exp, "variables(")
}

func isDotFunction(expression interface{}) bool {
	exp, ok := expression.(string)
	if !ok {
		return false
	}

	return strings.Contains(exp, ").")
}

func (s *BicepVisitor) VisitObjectProperty(ctx *parser.ObjectPropertyContext) interface{} {
	objectProperty := map[string]interface{}{}
	if ctx.Expression() != nil {
		objectValue := ctx.Expression().Accept(s)
		if isParameter(objectValue) || isDotFunction(objectValue) {
			objectValue = "[" + objectValue.(string) + "]"
		}

		if ctx.Identifier() != nil {
			identifier, ok := ctx.Identifier().Accept(s).(string)
			if !ok {
				return KicsObjectProperty{objectProperty: map[string]interface{}{}, line: 0}
			}
			objectProperty[identifier] = objectValue
		}
		if ctx.InterpString() != nil {
			interpString, ok := ctx.InterpString().Accept(s).(string)
			if !ok {
				return KicsObjectProperty{objectProperty: map[string]interface{}{}, line: 0}
			}
			objectProperty[interpString] = objectValue
		}
	}

	return KicsObjectProperty{objectProperty: objectProperty, line: ctx.GetStart().GetLine()}
}

func (s *BicepVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	if ctx.IDENTIFIER() != nil {
		identifier := ctx.IDENTIFIER().GetText()
		return identifier
	}
	if (ctx.PARAM()) != nil {
		return ctx.PARAM().GetText()
	}
	if (ctx.RESOURCE()) != nil {
		return ctx.RESOURCE().GetText()
	}
	if (ctx.VAR()) != nil {
		return ctx.VAR().GetText()
	}
	if (ctx.TRUE()) != nil {
		return ctx.TRUE().GetText()
	}
	if (ctx.FALSE()) != nil {
		return ctx.FALSE().GetText()
	}
	if (ctx.NULL()) != nil {
		return ctx.NULL().GetText()
	}
	if (ctx.STRING()) != nil {
		return ctx.STRING().GetText()
	}
	if (ctx.INT()) != nil {
		return ctx.INT().GetText()
	}
	if (ctx.BOOL()) != nil {
		return ctx.BOOL().GetText()
	}
	if (ctx.OBJECT()) != nil {
		return ctx.OBJECT().GetText()
	}
	return nil
}

func (s *BicepVisitor) VisitParenthesizedExpression(ctx *parser.ParenthesizedExpressionContext) interface{} {
	return ctx.Expression().Accept(s)
}

func (s *BicepVisitor) VisitDecorator(ctx *parser.DecoratorContext) interface{} {
	decorator := ctx.DecoratorExpression().Accept(s)
	return decorator
}

func (s *BicepVisitor) VisitDecoratorExpression(ctx *parser.DecoratorExpressionContext) interface{} {
	return ctx.FunctionCall().Accept(s)
}

func (s *BicepVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	identifier, ok := ctx.Identifier().Accept(s).(string)
	if !ok {
		return map[string]interface{}{}
	}
	var argumentList []interface{}
	if ctx.ArgumentList() != nil {
		argumentList, ok = ctx.ArgumentList().Accept(s).([]interface{})
		if !ok {
			return map[string]interface{}{}
		}
	}
	functionCall := map[string][]interface{}{
		identifier: argumentList,
	}

	return functionCall
}

func (s *BicepVisitor) VisitArgumentList(ctx *parser.ArgumentListContext) interface{} {
	var argumentList []interface{}
	for _, val := range ctx.AllExpression() {
		argument := val.Accept(s)
		argumentList = append(argumentList, argument)
	}
	return argumentList
}

func (s *BicepVisitor) VisitTypeExpression(ctx *parser.TypeExpressionContext) interface{} {
	return ctx.Identifier().Accept(s)
}

// GetKind returns the kind of the parser
func (p *Parser) GetKind() model.FileKind {
	return model.KindBICEP
}

// SupportedExtensions returns Bicep extensions
func (p *Parser) SupportedExtensions() []string {
	return []string{".bicep"}
}

// SupportedTypes returns types supported by this parser, which are bicep files
func (p *Parser) SupportedTypes() map[string]bool {
	return map[string]bool{"bicep": true, "azureresourcemanager": true}
}

// GetCommentToken return the comment token of Bicep files - #
func (p *Parser) GetCommentToken() string {
	return "//"
}

// StringifyContent converts original content into string formatted version
func (p *Parser) StringifyContent(content []byte) (string, error) {
	return string(content), nil
}

// Resolve resolves bicep files variables
func (p *Parser) Resolve(fileContent []byte, _ string, _ bool) ([]byte, error) {
	return fileContent, nil
}

// GetResolvedFiles returns the list of files that are resolved
func (p *Parser) GetResolvedFiles() map[string]model.ResolvedFile {
	return make(map[string]model.ResolvedFile)
}
