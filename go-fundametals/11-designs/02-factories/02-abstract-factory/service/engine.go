package service

type Engine interface {
	Assemble() string
	Capacity() int
}
