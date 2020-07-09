package collections

// todo interfaces for every primitive

// interface required for cargo with string key
type cargoString interface {
	GetKey() string
}

// interface required for cargo with int64 key
type cargoInt64 interface {
	GetKey() int64
}

// interface required for cargo with int key
type cargoInt interface {
	GetKey() int
}
