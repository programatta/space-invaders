package config

const WindowWidth int = 642
const WindowHeight int = 642
const PixelSize int = 1

var Ratio float32 = 1.5
var DesignWidth int = int(float32(WindowWidth/2) / Ratio)
var DesignHeight int = int(float32(WindowHeight/2) / Ratio)

var Dt float32 = 1.0 / 60
var Speed float32 = 200
