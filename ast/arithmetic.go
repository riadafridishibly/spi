package ast

// WARNING: VERY VERY UGLY CODE AHEAD
// I'll try to refactor this in future

func Plus(left, right ValueType) ValueType {
	var lvalInt int64
	var lvalFloat float64
	var isLvalueFloat bool
	var ok bool

	lvalInt, ok = left.(int64)

	if !ok {
		lvalFloat, ok = left.(float64)
		isLvalueFloat = true
	} else {
		isLvalueFloat = false
	}

	var rvalInt int64
	var rvalFloat float64

	rvalInt, ok = right.(int64)

	if !ok {
		rvalFloat, ok = right.(float64)
		if isLvalueFloat {
			return float64(lvalFloat) + rvalFloat
		} else {
			return float64(lvalInt) + rvalFloat
		}
	} else { // rvalue int
		if isLvalueFloat { // lvalue float
			return float64(lvalFloat) + float64(rvalInt)
		} else { // lvalue int
			return lvalInt + rvalInt
		}
	}
}

func Minus(left, right ValueType) ValueType {
	var lvalInt int64
	var lvalFloat float64
	var isLvalueFloat bool
	var ok bool

	lvalInt, ok = left.(int64)

	if !ok {
		lvalFloat, ok = left.(float64)
		isLvalueFloat = true
	} else {
		isLvalueFloat = false
	}

	var rvalInt int64
	var rvalFloat float64

	rvalInt, ok = right.(int64)

	if !ok {
		rvalFloat, ok = right.(float64)
		if isLvalueFloat {
			return float64(lvalFloat) - rvalFloat
		} else {
			return float64(lvalInt) - rvalFloat
		}
	} else { // rvalue int
		if isLvalueFloat { // lvalue float
			return float64(lvalFloat) - float64(rvalInt)
		} else { // lvalue int
			return lvalInt - rvalInt
		}
	}
}

func Multiply(left, right ValueType) ValueType {
	var lvalInt int64
	var lvalFloat float64
	var isLvalueFloat bool
	var ok bool

	lvalInt, ok = left.(int64)

	if !ok {
		lvalFloat, ok = left.(float64)
		isLvalueFloat = true
	} else {
		isLvalueFloat = false
	}

	var rvalInt int64
	var rvalFloat float64

	rvalInt, ok = right.(int64)

	if !ok { // rvalue float
		rvalFloat, ok = right.(float64)
		if isLvalueFloat { // lvalue float
			return lvalFloat * rvalFloat
		} else { // lvalue int
			return float64(lvalInt) * rvalFloat
		}
	} else { // rvalue int
		if isLvalueFloat { // lvalue float
			return lvalFloat * float64(rvalInt)
		} else { // lvalue int
			return lvalInt * rvalInt
		}
	}
}

func IntegerDiv(left, right ValueType) ValueType {
	var lvalInt int64
	var lvalFloat float64
	var isLvalueFloat bool
	var ok bool

	lvalInt, ok = left.(int64)

	if !ok {
		lvalFloat, ok = left.(float64)
		isLvalueFloat = true
	} else {
		isLvalueFloat = false
	}

	var rvalInt int64
	var rvalFloat float64

	rvalInt, ok = right.(int64)

	if !ok { // rvalue float
		rvalFloat, ok = right.(float64)
		if isLvalueFloat { // lvalue float
			return lvalFloat / rvalFloat
		} else { // lvalue int
			return float64(lvalInt) / rvalFloat
		}
	} else { // rvalue int
		if isLvalueFloat { // lvalue float
			return lvalFloat / float64(rvalInt)
		} else { // lvalue int
			return lvalInt / rvalInt
		}
	}
}

func FloatDiv(left, right ValueType) ValueType {
	var lvalInt int64
	var lvalFloat float64
	var isLvalueFloat bool
	var ok bool

	lvalInt, ok = left.(int64)

	if !ok {
		lvalFloat, ok = left.(float64)
		isLvalueFloat = true
	} else {
		isLvalueFloat = false
	}

	var rvalInt int64
	var rvalFloat float64

	rvalInt, ok = right.(int64)

	if !ok { // rvalue float
		rvalFloat, ok = right.(float64)
		if isLvalueFloat { // lvalue float
			return lvalFloat / rvalFloat
		} else { // lvalue int
			return float64(lvalInt) / rvalFloat
		}
	} else { // rvalue int
		if isLvalueFloat { // lvalue float
			return lvalFloat / float64(rvalInt)
		} else { // lvalue int
			return float64(lvalInt) / float64(rvalInt)
		}
	}
}
