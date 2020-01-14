package TocoFormula

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_Fif(t *testing.T) {
	result := "~10012A00307600000ECD4CC1430"
	data := []byte(result)
	form1, err := CompileFormula("fif((fif(ad(1,4)>1000,(ax(1,4)*10+ad(1,5)/2),-1)+22-ab(1,4))>45996,3,4)")
	if err != nil {
		t.Error(err)
	}
	res, err := GetAttributeValue(form1, data)
	if res != 3 {
		t.Error("0) ", " get ", res, " Expecting: ", 3, " error: ", err)
	}
	form2, err := CompileFormula("fif((fif(ad(1,4)>1000,(ax(1,4)*10+ad(1,5)/2),-1)+22-ab(1,4))>45998,3,4)")
	if err != nil {
		t.Error(err)
	}
	res, err = GetAttributeValue(form2, data)
	if res != 4 {
		t.Error("0) ", " get ", res, " Expecting: ", 3, " error: ", err)
	}
}

type isFormulaRight struct {
	formula string
	result  int
}

func Test_Func_Hx(t *testing.T) {
	origin := "01030401234567797F"
	data, _ := hex.DecodeString(origin)
	var tests = []isFormulaRight{
		{"hx(3,2)", 291},
		{"hx(3)", 291},
		{"hx(3,2,0)", 8961},
		{"hx(5,2)", 17767},
		{"hx(3,3)", 74565},
		{"hx(3,3,0)", 4530945},
		//{"hx(3,2,2)", 8961},
		//{"hx(5,2)", 1767},
		//{"hx(6,4)", 17767},
		//{"hx(a,4)", 17767},
		//{"hx(3,5)", 17767},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Hxu(t *testing.T) {
	origin := "01030401234589797F"
	data, _ := hex.DecodeString(origin)
	var tests = []isFormulaRight{
		{"hxu(3,2)", 291},
		{"hxu(3)", 291},
		{"hxu(3,2,0)", 8961},
		{"hxu(5,2)", 17801},
		{"hxu(5,2,0)", -30395},
		{"hxu(3,3)", 74565},
		{"hxu(4,3,0)", 8996131},
		{"hxu(4,3)", 2311561},
		{"hxu(3,4,0)", -1991957759},
		{"hxu(3,4)", 19088777},
		//{"hx(3,2,2)", 8961},
		//{"hx(5,2)", 1767},
		//{"hx(6,4)", 17767},
		//{"hx(a,4)", 17767},
		//{"hx(3,5)", 17767},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Hb(t *testing.T) {
	origin := "010300000002C40C"
	data, _ := hex.DecodeString(origin)
	var tests = []isFormulaRight{
		{"hb(6,0)", 0},
		{"hb(6,1)", 0},
		{"hb(6,2)", 1},
		{"hb(6,3)", 0},
		{"hb(6,4)", 0},
		{"hb(6,5)", 0},
		{"hb(6,6)", 1},
		{"hb(6,7)", 1},
		{"hb(1,7)", 0},
		{"hb(1,5)", 0},
		//{"hb(9,2)", 0},
		//{"hb(1,8)", 0},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Ht(t *testing.T) {
	origin := "010304000002AD36424C"
	data, _ := hex.DecodeString(origin)
	var tests = []isFormulaRight{
		{"ht(0,4)", 103},
		{"ht(2,4)", 304},
		{"ht(3,5)", 30400},
		{"ht(14,2)", 36},
		//{"ht(14,4)", 3643},
		//{"ht(12,2)", 3642},
		//{"ht(16,6)", 3642},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Hc(t *testing.T) {
	origin := "AAAAC0A800F4022045100738FFFFFF00C0A80101FFFFFFFFFFFF32322E320030372E310000302D31322E380000000100E2FF000000"
	data, _ := hex.DecodeString(origin)
	var tests = []isFormulaRight{
		{"hc(38,6)*10", -128},
		{"hc(32,6)*10", 71},
		{"hc(26,6)*10", 222},
		{"hc(32,6)*10", 71},
		{"hc(32,6)", 7},
		//{"hc(50,6)*10", 222},
		//{"hc(0,6)*10", 222},
		//{"hc(38,10)*10", 222},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Hp(t *testing.T) {
	//origin := "51032042480000435D2EE0435E099D435DD740435DAFE943BFBA3E43C058FB43BFDD80CEF1"
	origin := "2803404940800042A4000046A29E00442BA94544E9B91E42F8E8183E40A5E63FCC80CD48CB0A200000000000000000460696CC4479FFFF4479FFFF0000000043480000986F"
	//origin := "2803404880020043820000479DB58043F6CA2B44B080AC44547F693E9597853FEBE43F47F81780000000000000000045436FFF4479FFFF4479FFFF0000000043480000AEF7"
	data, _ := hex.DecodeString(origin)
	var tests = []isFormulaRight{
		{"hp(3)", 788480},
		{"hp(7)*1000", 82000},
		{"hp(11)*1000", 20815000},
		{"hp(15)*1000", 686640},
		{"hp(19)*1000", 1869780},
		{"hp(23)*1000", 124450},
		{"hp(27)*1000", 190},
		{"hp(31)*1000", 1600},
		{"hp(7)*100", 8200},
		//{"hp(35)", 221},
		//{"hp(16)", 221},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Ap(t *testing.T) {
	origin := "~10012A00307600000ECD4CC14300C0C14333F3C143333313403333F33F3333F33F00PP4842F6281C3F713D0A3F14AE073F00005F4300805E4333335F4300004842E491"
	data := []byte(origin)
	var tests = []isFormulaRight{
		{"ap(19)*10", 3866},
		{"ap(19)", 386},
		//{"ap(19,0)", 386},
		//{"ap(29)", 386},
		//{"ap(133)", 386},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		fmt.Println(res, v.result)
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Ax(t *testing.T) {
	origin := "~10012A0020680133B35A4300005F43CDCC6043CDCC5C4300005D4300005D43000060410000C0400000B0409AD9CC439A9947420000009AD9CC43E77D"
	data := []byte(origin)
	var tests = []isFormulaRight{
		{"ax(1,4)", 4097},
		{"ax(15,4)", 13235},
		{"ax(1,4,0)", 272},
		//{"ax(1,4,0)", 273},
		//{"ax(120,2)", 13235},
		//{"ax(1,3)", 13235},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Ad(t *testing.T) {
	origin := ":F00+10.10209-463FFFFFF5100080000E30000000000017A014C04971FQ"
	data := []byte(origin)
	var tests = []isFormulaRight{
		{"ad(10,3)", 209},
		{"ad(13,4)", -463},
		{"ad(5,4)*10", 101},
		//{"ad(53,8)*10", 101},
		//{"ad(51,4)*10", 101},
		//{"ad(4,4)*10", 101},
		//{"ad(5,9)*10", 101},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Ac(t *testing.T) {
	origin := "01A091023131314D5330312E30342E313331302E34322E34362D3830333832333832323138323139323230303031313030303630303036303030362B303030322B303030312B303030312B302E3035302E303337373337373337343432352B303030303030303230303031303030312A2A2A2E2A2A2A2A2E2A2A2A2A2E2A2B30302E303030393030353030353039393931303003BF"
	data := []byte(origin)
	var tests = []isFormulaRight{
		{"ac(50,6)", -80},
		{"ac(148,10)*100", 5},
		{"ac(148,10)", 0},
		//{"ac(148,20)", 0},
		//{"ac(288,10)", 0},
		//{"ac(288,12)", 0},
		//{"ac(50,7)", -80},
		//{"ac(50,0)", -80},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Ab(t *testing.T) {
	origin := ":F0061A0102090463FFFFFF5100080000E30000000000017A014C04971FQ"
	data := []byte(origin)
	var tests = []isFormulaRight{
		{"ab(33,4)", 0},
		{"ab(33,5)", 1},
		{"ab(1,3)", 0},
		//{"ab(59,3)", 0},
		//{"ab(33,8)", 0},
		//{"ab(58,3)", 0},
	}
	for i, v := range tests {
		form, err := CompileFormula(v.formula)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}
}

func Test_Func_Av(t *testing.T) {
	var origins = []isFormulaRight{
		{"-32.6", -32},
		{"0006", 6},
		{"+32.6", 32},
		{".6", 0},
		{"6.", 6},
		//{"-3a.6", 6},
		//{"+-3", 6},
		//{"30.6.3", 6},
	}
	form, _ := CompileFormula("av()")
	for i, v := range origins {
		data := []byte(v.formula)
		res, err := GetAttributeValue(form, data)
		if err != nil {
			t.Error(i, ") ", v.formula, " err: ", err)
			continue
		}
		if res != v.result {
			t.Error(i, ") ", v.formula, " get ", res, " Expecting: ", v.result)
		}
	}

}

func BenchmarkHx(b *testing.B) {
	origin := "01030401234567797F"
	data, _ := hex.DecodeString(origin)
	fString := "hx(3,2,0)"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkHb(b *testing.B) {
	origin := "01030401234567797F"
	data, _ := hex.DecodeString(origin)
	fString := "hb(6,7)"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkHt(b *testing.B) {
	origin := "01030401234567797F"
	data, _ := hex.DecodeString(origin)
	fString := "ht(2,4)"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkHc(b *testing.B) {
	origin := "01030432322E320030372E310000302D31322E3800"
	data, _ := hex.DecodeString(origin)
	fString := "hc(3,6)*10"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkHp(b *testing.B) {
	origin := "51032042480000435D2EE0435E099D435DD740435DAFE943BFBA3E43C058FB43BFDD80CEF1"
	data, _ := hex.DecodeString(origin)
	fString := "hp(3)"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkAp(b *testing.B) {
	origin := "~10012A00307600000ECD4CC14300C0C14333F3C143333313403333F33F3333F33F00PP4842F6281C3F713D0A3F14AE073F00005F4300805E4333335F4300004842E491"
	data := []byte(origin)
	fString := "ap(19)"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkAx(b *testing.B) {
	origin := "~10012A00307600000ECD4CC14300C0C14333F3C143333313403333F33F3333F33F00PP4842F6281C3F713D0A3F14AE073F00005F4300805E4333335F4300004842E491"
	data := []byte(origin)
	fString := "ax(15,4)"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkAd(b *testing.B) {
	origin := ":F00+10.10209-463FFFFFF5100080000E30000000000017A014C04971FQ"
	data := []byte(origin)
	fString := "ad(5,4)*10"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkAc(b *testing.B) {
	origin := "01A091023131314D5330312E30342E313331302E34322E34362D3830333832333832323138323139323230303031313030303630303036303030362B303030322B303030312B303030312B302E3035302E303337373337373337343432352B303030303030303230303031303030312A2A2A2E2A2A2A2A2E2A2A2A2A2E2A2B30302E303030393030353030353039393931303003BF"
	data := []byte(origin)
	fString := "ac(148,10)*100"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkAb(b *testing.B) {
	origin := ":F0061A0102090463FFFFFF5100080000E30000000000017A014C04971FQ"
	data := []byte(origin)
	fString := "ab(33,4)"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkAv(b *testing.B) {
	origin := "-32.6"
	data := []byte(origin)
	fString := "av()"
	form, _ := CompileFormula(fString)
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form, data)
	}
}

func BenchmarkFif(b *testing.B) {
	result := "~10012A00307600000ECD4CC1430"
	data := []byte(result)
	form1, _ := CompileFormula("fif((fif(ad(1,4)>1000,(ax(1,4)*10+ad(1,5)/2),-1)+22-ab(1,4))>45996,3,4)")
	for i := 0; i < b.N; i++ {
		_, _ = GetAttributeValue(form1, data)
	}
}
