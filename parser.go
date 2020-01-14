package TocoFormula

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
)

// 根据[]byte计算十六进制数
// 参数1：起始地址，默认3，必须
// 参数2：长度，默认2（int16），支持1（uint8）、4（int32），可忽略
// 参数3：字节序， 默认1:BigEndian 0:LittleEndian，可忽略
func Hx(data []byte, args ...int) (float64, error) {
	start := 3 //起始地址
	size := 2  //长度，默认2（int16）
	order := 1 //字节序，siz>1,  默认1:BigEndian 0:LittleEndian
	argLength := len(args)

	if argLength == 0 || argLength > 3 {
		return 0, errors.New("wrong formula")
	}

	if argLength > 0 {
		start = args[0]
	}
	if argLength > 1 {
		size = args[1]
	}
	if argLength > 2 {
		order = args[2]
	}

	if start+size > len(data) {
		return 0, errors.New("wrong formula")
	}

	switch size {
	case 1:
		return float64(data[start]), nil
	case 2:
		if order == 1 {
			value := int16(binary.BigEndian.Uint16(data[start : start+2]))
			return float64(value), nil
		} else if order == 0 {
			value := int16(binary.LittleEndian.Uint16(data[start : start+2]))
			return float64(value), nil
		} else {
			return 0, errors.New("wrong formula")
		}
	case 3, 4:
		dt := data[start : start+size]

		if order == 1 {
			if size == 3 {
				dt = append([]byte{0}, dt...)
			}
			return float64(binary.BigEndian.Uint32(dt)), nil
		} else if order == 0 {
			if size == 3 {
				dt = append(dt, []byte{0}...)
			}
			return float64(binary.LittleEndian.Uint32(dt)), nil
		} else {
			return 0, errors.New("wrong formula")
		}
	default:
		return 0, errors.New("wrong formula")
	}
}

// 根据[]byte计算十六进制数（返回带符号的数）
// 参数1：起始地址，默认3，必须
// 参数2：长度，默认2（int16），支持1（uint8）、4（int32），可忽略
// 参数3：字节序， 默认1:BigEndian 0:LittleEndian，可忽略
func Hxu(data []byte, args ...int) (float64, error) {
	start := 3 //起始地址
	size := 2  //长度，默认2（int16）
	order := 1 //字节序，siz>1,  默认1:BigEndian 0:LittleEndian
	argLength := len(args)

	if argLength == 0 || argLength > 3 {
		return 0, errors.New("wrong formula")
	}

	if argLength > 0 {
		start = args[0]
	}
	if argLength > 1 {
		size = args[1]
	}
	if argLength > 2 {
		order = args[2]
	}

	if start+size > len(data) {
		return 0, errors.New("wrong formula")
	}
	//var bytesBuffer bytes.Buffer // 数据buffer
	switch size {
	case 1:
		return float64(int8(data[start])), nil
	case 2:
		bytesBuffer := bytes.NewBuffer(data[start : start+2])
		var tmp int16
		if order == 1 {
			err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
			return float64(tmp), err
		} else if order == 0 {
			err := binary.Read(bytesBuffer, binary.LittleEndian, &tmp)
			return float64(tmp), err
		} else {
			return 0, errors.New("wrong formula")
		}
	case 3, 4:
		bytesBuffer := bytes.NewBuffer(data[start : start+size])
		var tmp int32
		if order == 1 {
			if size == 3 {
				bytesBuffer = bytes.NewBuffer(append([]byte{0}, data[start:start+size]...))
			}
			err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
			return float64(tmp), err
		} else if order == 0 {
			if size == 3 {
				bytesBuffer = bytes.NewBuffer(append(data[start:start+size], []byte{0}...))
			}
			err := binary.Read(bytesBuffer, binary.LittleEndian, &tmp)
			return float64(tmp), err
		} else {
			return 0, errors.New("wrong formula")
		}
	default:
		return 0, errors.New("wrong formula")
	}
}

//  Hb 根据[]byte计算bit值；
func Hb(data []byte, args ...int) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("wrong formula")
	}
	index := args[0]
	order := uint(args[1])
	if order >= 8 || order < 0 || index >= len(data) {
		return 0, errors.New("wrong formula")
	}
	temp := int(data[index])
	temp = temp & (1 << order) >> order
	return float64(temp), nil
}

// Ht 将十六进制表示的bcd码转换成十进制整数
func Ht(data []byte, args ...int) (float64, error) {
	l := len(args)
	if l < 1 {
		return 0, errors.New("wrong formula")
	}
	index := args[0]
	length := 2
	if l > 1 {
		length = args[1]
	}
	if length > 8 || length < 2 {
		return 0, errors.New("wrong formula")
	}
	hexStr := fmt.Sprintf("%x", data)
	if (index + length) > len(hexStr) {
		return 0, errors.New("wrong formula")
	}
	result, err := strconv.ParseInt(hexStr[index:index+length], 10, 64)
	if err != nil {
		return 0, err
	}
	return float64(result), nil
}

// Hc 将十六进制表示的ascii码数字转换成浮点数
func Hc(data []byte, args ...int) (float64, error) {
	l := len(args)
	if l < 1 {
		return 0, errors.New("wrong formula")
	}
	index := args[0]
	length := 2
	if l > 1 {
		length = args[1]
	}
	if length > 8 || length < 2 {
		return 0, errors.New("wrong formula")
	}
	if (index + length) > len(data) {
		return 0, errors.New("wrong formula")
	}
	dt := data[index : index+length]
	for id, ch := range dt {
		if ch == 0 {
			dt[id] = 48
		}
	}
	return strconv.ParseFloat(string(dt), 64)
}

// Hp 将十六进制表示的4个Byte转换成浮点数（默认长度为4）
func Hp(data []byte, args ...int) (float64, error) {
	if len(args) < 1 {
		return 0, errors.New("wrong formula")
	}
	index := args[0]
	if (index + 4) > len(data) {
		return 0, errors.New("wrong formula")
	}
	binString := ""
	dt := data[index : index+4]
	for _, ch := range dt {
		binString += fmt.Sprintf("%08b", ch)
	}
	// 正负
	mark := 1.0
	if binString[0] == 0 {
		mark = -1
	}

	eMark, _ := strconv.ParseInt(binString[1:9], 2, 64)
	mMark, _ := strconv.ParseInt(binString[9:], 2, 64)
	r1 := math.Pow(2, float64(eMark-127))
	r2 := mark * r1 * float64(mMark) * math.Pow(2, -23)
	return math.Floor((r1+r2)*100+0.5) / 100.0, nil
}

// Ap 截取电总协议中的字符串，并转换成浮点数
func Ap(data []byte, args ...int) (float64, error) {
	l := len(args)
	if l < 1 {
		return 0, errors.New("wrong formula")
	}
	index := args[0]
	order := 1
	if l > 1 {
		order = args[1]
	}
	if order > 1 || order < 0 {
		return 0, errors.New("wrong formula")
	}
	if (index + 8) > len(data) {
		return 0, errors.New("wrong formula")
	}
	binString := ""
	dt := string(data[index : index+8])
	if order == 1 {
		for i := 3; i >= 0; i-- {
			binString += dt[2*i : 2*i+2]
		}
	} else {
		binString = dt
	}
	ndt, err := hex.DecodeString(binString)
	if err != nil {
		return 0, errors.New("wrong formula")
	}
	resultString := ""
	for _, ch := range ndt {
		resultString += fmt.Sprintf("%08b", ch)
	}
	// 正负
	mark := 1.0
	if resultString[0] == 0 {
		mark = -1
	}

	eMark, _ := strconv.ParseInt(resultString[1:9], 2, 64)
	mMark, _ := strconv.ParseInt(resultString[9:], 2, 64)
	r1 := math.Pow(2, float64(eMark-127))
	r2 := mark * r1 * float64(mMark) * math.Pow(2, -23)
	return math.Floor((r1+r2)*100+0.5) / 100.0, nil
}

// Ax 把字符串作为十六进制字符串，截取并转换成整数
func Ax(data []byte, args ...int) (float64, error) {
	l := len(args)
	if l < 1 {
		return 0, errors.New("wrong formula")
	}
	index := args[0]

	length := 2 //截取字符串长度，默认2，1，2，4
	if l > 1 {
		length = args[1]
	}

	order := 1
	if l > 2 {
		order = args[2]
	}

	if (index + length) > len(data) {
		return 0, errors.New("wrong formula")
	}

	dt := string(data[index : index+length])

	ndt, err := hex.DecodeString(dt)
	if err != nil {
		return 0, errors.New("wrong formula")
	}

	switch length {
	case 2:
		return float64(ndt[0]), nil
	case 4:
		if order == 1 {
			value := int16(binary.BigEndian.Uint16(ndt[:]))
			return float64(value), nil
		} else if order == 0 {
			value := int16(binary.LittleEndian.Uint16(ndt[:]))
			return float64(value), nil
		} else {
			return 0, errors.New("wrong formula")
		}
	case 8:
		if order == 1 {
			value := int32(binary.BigEndian.Uint32(ndt[:]))
			return float64(value), nil
		} else if order == 0 {
			value := int32(binary.LittleEndian.Uint32(ndt[:]))
			return float64(value), nil
		} else {
			return 0, errors.New("wrong formula")
		}
	default:
		return 0, errors.New("wrong formula")
	}
}

// Ad 将字符串转换成浮点数
func Ad(data []byte, args ...int) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("wrong formula")
	}
	index := args[0]
	length := args[1] //截取字符串长度，最大8
	if length > 8 || length < 0 {
		return 0, errors.New("wrong formula")
	}

	if (index + length) > len(data) {
		return 0, errors.New("wrong formula")
	}

	dt := string(data[index : index+length])
	return strconv.ParseFloat(dt, 64)
}

// Ac 将bcd字符串转换成浮点数
func Ac(data []byte, args ...int) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("wrong formula")
	}
	index := args[0]
	length := args[1] //截取字符串长度，最大8
	if length > 16 || length < 0 || length%2 == 1 {
		return 0, errors.New("wrong formula")
	}

	if (index + length) > len(data) {
		return 0, errors.New("wrong formula")
	}
	dt := string(data[index : index+length])
	ndt, err := hex.DecodeString(dt)
	if err != nil {
		return 0, errors.New("wrong formula")
	}
	result := string(ndt[:])
	return strconv.ParseFloat(result, 64)
}

// Ab 将字符串按照十六进制方式，每次取2位转换成布尔值
func Ab(data []byte, args ...int) (float64, error) {
	if len(args) != 2 {
		return 0, errors.New("wrong formula")
	}
	index := args[0]
	order := uint8(args[1]) // 字节位置0-7
	if order > 7 || order < 0 {
		return 0, errors.New("wrong formula")
	}

	if (index + 2) > len(data) {
		return 0, errors.New("wrong formula")
	}
	dt := string(data[index : index+2])
	ndt, err := hex.DecodeString(dt)
	if err != nil {
		return 0, errors.New("wrong formula")
	}
	temp := uint(ndt[0])
	temp = temp & (1 << order) >> order
	return float64(temp), nil
}

// Av 将字符串按照十六进制方式，每次取2位转换成布尔值
func Av(data []byte) (float64, error) {
	dt := string(data)
	return strconv.ParseFloat(dt, 64)
}
