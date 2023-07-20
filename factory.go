package gopool

type PoolFactory interface {
	CreatePool(size int) Pool
}

type poolFactory struct{}

func NewPoolFactory() PoolFactory {
	return &poolFactory{}
}

func (pf *poolFactory) CreatePool(size int) Pool {
	return NewPool(size)
}
