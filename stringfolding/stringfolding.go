package stringfolding

// read here https://opendsa-server.cs.vt.edu/ODSA/Books/Everything/html/HashFuncExamp.html

type IStringFolding interface {
	Generate(string) int64
}

type StringFolding struct {
	max int64 // max int folding generate
}

func CreateStringFolding(max int64) *StringFolding {
	return &StringFolding{max}
}

func (f *StringFolding) Generate(in string) int64 {
	return Generate(in, f.max)
}

func Generate(in string, max int64) int64 {
	if in == "" {
		return 0
	}
	var mul int64 = 1
	var sum int64
	for i := 0; i < len(in); i++ {
		if i%4 == 0 {
			mul = 1
		} else {
			mul = mul * 256
		}
		sum += int64(in[i]) * mul
	}
	return sum % max
}
