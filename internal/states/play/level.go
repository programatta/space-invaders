package play

type ConfigLevel = struct {
	alienFireDelay float32
	alienMoveDelay float32
}

type Level struct {
	current int
	levels  []ConfigLevel
}

func NewLevel() *Level {
	level := &Level{
		current: 0,
	}

	level.levels = []ConfigLevel{
		{
			alienFireDelay: 0.400,
			alienMoveDelay: 0.35,
		},
		{
			alienFireDelay: 0.370,
			alienMoveDelay: 0.3,
		},
		{
			alienFireDelay: 0.300,
			alienMoveDelay: 0.26,
		},
		{
			alienFireDelay: 0.260,
			alienMoveDelay: 0.2,
		},
		{
			alienFireDelay: 0.220,
			alienMoveDelay: 0.15,
		},
	}

	return level
}

func (l *Level) Next() bool {
	if l.current+1 < len(l.levels) {
		l.current++
		return true
	}
	return false
}

func (l *Level) Current() ConfigLevel {
	return l.levels[l.current]
}
