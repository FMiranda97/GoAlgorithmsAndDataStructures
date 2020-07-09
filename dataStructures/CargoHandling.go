package dataStructures

import (
	"errors"
	"reflect"
)

type cargoString interface {
	GetKey() string
}

type cargoInt64 interface {
	GetKey() int64
}

type cargoInt interface {
	GetKey() int
}

func comparator(this interface{}, target interface{}) (int8, error) {
	if reflect.TypeOf(this) != reflect.TypeOf(target) {
		return 0, errors.New("compared parameters implement different interfaces")
	}
	switch this.(type) {
	case cargoString:
		a := this.(cargoString)
		b := target.(cargoString)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case cargoInt:
		a := this.(cargoInt)
		b := target.(cargoInt)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case cargoInt64:
		a := this.(cargoInt)
		b := target.(cargoInt)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, errors.New("compared parameters implement unsupported interface")
	}
}

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
	return false
}
