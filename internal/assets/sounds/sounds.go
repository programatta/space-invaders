package sounds

import _ "embed"

//go:embed shoot.wav
var ShootWav []byte

//go:embed invaderkilled.wav
var InvaderKilledWav []byte

//go:embed cannonexplosion.wav
var CannonExplosionWav []byte

//go:embed ufo_highpitch.wav
var UfoHighpitchWav []byte
