package service

type StreetCarEngine struct{}

func (s StreetCarEngine) Assemble() string {
	return "Assembling components of Street Car Engine"
}

func (s StreetCarEngine) Capacity() int {
	return 1000
}
