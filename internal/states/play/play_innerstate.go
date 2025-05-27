package play

type playInnerStateId int

const (
	starting playInnerStateId = iota
	playing
	pauseRequest
	pause
	gameOver
)
