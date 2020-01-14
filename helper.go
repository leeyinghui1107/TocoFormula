package TocoFormula

import (
	"bytes"
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/leeyinghui1107/TocoLibrary/TocoString"
	"regexp"
	"strings"
)

// 普通函数不能包含f作为名称，否则将作为函数类型进行计算
var FuncName = map[string]string{
	"hx":  "hx",
	"hxu": "hxu",
	"hb":  "hb",
	"ht":  "ht",
	"hc":  "hc",
	"hp":  "hp",
	"ap":  "ap",
	"ax":  "ax",
	"ad":  "ad",
	"ac":  "ac",
	"ab":  "ab",
	"av":  "av",
	"fif": "fif",
}

// Formula 公式结构
type Formula struct {
	Ready      bool                           // 判定公式是否准备就绪
	Type       int                            // 类型 默认0：single  1：expression  2：function
	Expression *govaluate.EvaluableExpression // 计算表达式
	Functions  []Function                     // 子函数列表
}

// Function 子函数结构
type Function struct {
	Index string // 索引，a-z，根据顺序递增
	Name  string // 子函数名称，小写
	Args  []int  // 子函数参数，正整数数组
}

// CompileFormula 解析公式字符串
func CompileFormula(source string) (Formula, error) {
	var result = Formula{
		Ready: false,
	}
	// 删除空格并转换为小写字符串
	fString := strings.ToLower(TocoString.DeleteSpace(source))

	// 判断公式是否包含非Ascii字符
	if !TocoString.IsASCII(fString) {
		return result, errors.New("wrong formula format")
	}

	// 表达式
	var expression bytes.Buffer

	// 子函数列表
	var functions []Function

	// 公式
	r, _ := regexp.Compile(`^[a-z]\w+\([0-9,]*\)$`)
	if r.MatchString(fString) {
		f, err := resolveSingleFunction(fString)
		if err != nil {
			return result, err
		}
		functions = append(functions, *f)
		return Formula{true, 0, nil, functions}, nil
	}

	// 公式类型
	sort := 1
	if strings.Contains(fString, "f") {
		sort = 2
	}

	r, _ = regexp.Compile(`[a-z]\w+\([0-9,]*\)`)
	fIndex := byte('a') // 子函数参数名

	lastIndex := 0
	// 获取所有匹配的子函数
	for _, mat := range r.FindAllStringIndex(fString, -1) {
		funcString := fString[mat[0]:mat[1]]
		f, err := resolveSingleFunction(funcString)
		if err != nil {
			return result, err
		}
		f.Index = string(fIndex)
		functions = append(functions, *f)
		fIndex++
		expression.WriteString(fString[lastIndex:mat[0]])
		expression.WriteString(f.Index)
		lastIndex = mat[1]
	}
	if lastIndex < len(fString) {
		expression.WriteString(fString[lastIndex:])
	}

	// 表达式
	if sort == 1 {
		exp, err := govaluate.NewEvaluableExpression(expression.String())
		if err != nil {
			return result, err
		}
		return Formula{true, sort, exp, functions}, nil
	}
	// 函数
	exp, err := govaluate.NewEvaluableExpressionWithFunctions(expression.String(), Functions)
	if err != nil {
		return result, err
	}
	return Formula{true, sort, exp, functions}, nil

}

// resolveSingleFunction 根据表达式，转换成子函数返回
func resolveSingleFunction(fString string) (*Function, error) {
	end := strings.Index(fString, "(")
	fName := fString[0:end]
	if _, ok := FuncName[fName]; ok {
		paramString := fString[end+1 : len(fString)-1]
		var args []int
		f := Function{"a", fName, args}
		if len(paramString) > 0 {
			args, err := TocoString.SplitStrToInt(paramString, ",")
			if err != nil {
				return nil, err
			}
			f.Args = args
		}
		return &f, nil
	}
	return nil, errors.New("wrong function")
}
