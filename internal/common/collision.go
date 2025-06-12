package common

type Collider interface {
	Rect() (float32, float32, float32, float32)
	OnCollide()
}

func CheckCollision(sourceObj, targetObj Collider) bool {
	sx0, sy0, sw, sh := sourceObj.Rect()
	tx0, ty0, tw, th := targetObj.Rect()

	hasCollision := sx0 < tx0+tw && sx0+sw > tx0 && sy0 < ty0+th && sh+sy0 > ty0

	return hasCollision
}
