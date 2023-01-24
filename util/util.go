package util

func OptStr(Value any, Optional string) string {
	Val, ok := Value.(string)
	if ok && Val != "" {
		return Val
	} else {
		return Optional
	}
}
func OptInt(Value any, Optional int) int {
	Val, ok := Value.(int)
	if ok && Val != 0 {
		return Val
	} else {
		return Optional
	}
}
func OptInt8(Value any, Optional int8) int8 {
	Val, ok := Value.(int8)
	if ok && Val != 0 {
		return Val
	} else {
		return Optional
	}
}
func OptInt16(Value any, Optional int16) int16 {
	Val, ok := Value.(int16)
	if ok && Val != 0 {
		return Val
	} else {
		return Optional
	}
}
func Opt32(Value any, Optional int32) int32 {
	Val, ok := Value.(int32)
	if ok && Val != 0 {
		return Val
	} else {
		return Optional
	}
}
func OptInt64(Value any, Optional int64) int64 {
	Val, ok := Value.(int64)
	if ok && Val != 0 {
		return Val
	} else {
		return Optional
	}
}

func OptBool(Value any, Optional bool) bool {
	Val, ok := Value.(bool)
	if ok && (!Val || Val) {
		return Val
	} else {
		return Optional
	}
}
