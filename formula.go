package TocoFormula

import (
	"errors"
	"github.com/Knetic/govaluate"
	"strconv"
)

// Functions 自定义函数数组，用于计算属性和流程控制函数（fif）
var Functions = map[string]govaluate.ExpressionFunction{
	"max": func(args ...interface{}) (interface{}, error) {
		result := 0.0
		for id, v := range args {
			if value, ok := v.(float64); ok {
				if id == 0 || value > result {
					result = value
				}
			} else {
				return result, errors.New("格式错误")
			}
		}
		return result, nil
	},
	"avg": func(args ...interface{}) (interface{}, error) {
		total := 0.0
		length := len(args)
		for _, v := range args {
			if value, ok := v.(float64); ok {
				total = total + value
			} else {
				return 0.0, errors.New("格式错误")
			}
		}
		return total / float64(length), nil
	},
	"min": func(args ...interface{}) (interface{}, error) {
		result := 0.0
		for id, v := range args {
			if value, ok := v.(float64); ok {
				if id == 0 || value < result {
					result = value
				}
			} else {
				return result, errors.New("格式错误")
			}
		}
		return result, nil
	},
	"fif": func(args ...interface{}) (interface{}, error) {
		if len(args) != 3 {
			return nil, errors.New("wrong parameter")
		}
		if result, ok := args[0].(bool); ok {
			if result {
				return args[1], nil
			} else {
				return args[2], nil
			}
		}
		return nil, errors.New("wrong parameter")
	},
}

// GetAttributeValue 根据公式和返回串获取属性值（整数）
func GetAttributeValue(f Formula, data []byte) (int, error) {
	if !f.Ready {
		return 0, errors.New("formula error")
	}
	switch f.Type {
	case 0:
		if len(f.Functions) == 1 {
			value, err := CalculateValue(f.Functions[0], data)
			if err != nil {
				return 0, err
			}
			return int(value), nil
		}
		return 0, errors.New("formula error")
	case 1, 2:
		args := map[string]interface{}{}
		if len(f.Functions) > 0 {
			for _, fun := range f.Functions {
				value, err := CalculateValue(fun, data)
				if err != nil {
					return 0, err
				}
				args[fun.Index] = value
			}
		}
		result, err := f.Expression.Evaluate(args)
		if err != nil {
			return 0, err
		}
		return int(result.(float64)), nil
	default:
		return 0, errors.New("unknown type")
	}
}

// CalculateValue 根据子函数计算属性值
func CalculateValue(fun Function, data []byte) (float64, error) {
	switch fun.Name {
	case "hx":
		return Hx(data, fun.Args...)
	case "hxu":
		return Hxu(data, fun.Args...)
	case "hb":
		return Hb(data, fun.Args...)
	case "ht":
		return Ht(data, fun.Args...)
	case "hc":
		return Hc(data, fun.Args...)
	case "hp":
		return Hp(data, fun.Args...)
	case "ap":
		return Ap(data, fun.Args...)
	case "ax":
		return Ax(data, fun.Args...)
	case "ad":
		return Ad(data, fun.Args...)
	case "ac":
		return Ac(data, fun.Args...)
	case "ab":
		return Ab(data, fun.Args...)
	case "av":
		return Av(data)
	default:
		return 0, nil
	}
}

// CalculateFormula 通过公式和数据，获得计算属性的值
func CalculateFormula(formula string, list []float64) (interface{}, error) {
	expression, _ := govaluate.NewEvaluableExpressionWithFunctions(formula, Functions)
	parameters := map[string]interface{}{}
	for id, v := range list {
		pIndex := "p" + strconv.Itoa(id+1)
		parameters[pIndex] = v
	}

	return expression.Evaluate(parameters)
}
