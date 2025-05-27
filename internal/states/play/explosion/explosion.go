package explosion

import "github.com/programatta/spaceinvaders/internal/states/play/common"

type Explosioner interface {
	common.Manageer
	common.Eraser
}
