package internal

type playInnerStateId int

const (
	playing playInnerStateId = iota
	gameOver
)
