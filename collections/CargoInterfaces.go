package collections

// interface required for cargo with string key
type cargoString interface {
	GetKey() string
}

// interface required for cargo with int key
type cargoInt interface {
	GetKey() int
}

// interface required for cargo with int64 key
type cargoInt64 interface {
	GetKey() int64
}

// interface required for cargo with int32 key
type cargoInt32 interface {
	GetKey() int32
}

// interface required for cargo with int16 key
type cargoInt16 interface {
	GetKey() int16
}

// interface required for cargo with int8 key
type cargoInt8 interface {
	GetKey() int8
}

// interface required for cargo with uint64 key
type cargoUInt64 interface {
	GetKey() int64
}

// interface required for cargo with uint32 key
type cargoUInt32 interface {
	GetKey() int32
}

// interface required for cargo with uint16 key
type cargoUInt16 interface {
	GetKey() int16
}

// interface required for cargo with uint8 key
type cargoUInt8 interface {
	GetKey() uint8
}

// interface required for cargo with rune key
type cargoRune interface {
	GetKey() rune
}

// interface required for cargo with byte key
type cargoByte interface {
	GetKey() byte
}

// interface required for cargo with float32 key
type cargoFloat32 interface {
	GetKey() float32
}

// interface required for cargo with float64 key
type cargoFloat64 interface {
	GetKey() float64
}

// interface required for cargo with complex64 key
type cargoComplex64 interface {
	GetKey() complex64
}

// interface required for cargo with complex128 key
type cargoComplex128 interface {
	GetKey() complex128
}

// checks if passed cargo implements any of the required interfaces
func implementsLinkedListCargo(cargo interface{}) bool {
	if _, ok := cargo.(cargoString); ok {
		return true
	}
	if _, ok := cargo.(cargoInt); ok {
		return true
	}
	if _, ok := cargo.(cargoInt64); ok {
		return true
	}
	if _, ok := cargo.(cargoInt32); ok {
		return true
	}
	if _, ok := cargo.(cargoInt16); ok {
		return true
	}
	if _, ok := cargo.(cargoInt8); ok {
		return true
	}
	if _, ok := cargo.(cargoUInt64); ok {
		return true
	}
	if _, ok := cargo.(cargoUInt32); ok {
		return true
	}
	if _, ok := cargo.(cargoUInt16); ok {
		return true
	}
	if _, ok := cargo.(cargoUInt8); ok {
		return true
	}
	if _, ok := cargo.(cargoRune); ok {
		return true
	}
	if _, ok := cargo.(cargoByte); ok {
		return true
	}
	if _, ok := cargo.(cargoFloat32); ok {
		return true
	}
	if _, ok := cargo.(cargoFloat64); ok {
		return true
	}
	if _, ok := cargo.(cargoComplex64); ok {
		return true
	}
	if _, ok := cargo.(cargoComplex128); ok {
		return true
	}
	return false
}
