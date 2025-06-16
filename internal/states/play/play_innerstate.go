package play

type playInnerStateId int

const (
	playing playInnerStateId = iota
	gameOver
)
