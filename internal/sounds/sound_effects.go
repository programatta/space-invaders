package sounds

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/programatta/spaceinvaders/internal/assets/sounds"
)

type SoundEffects struct {
	shootPlayer *audio.Player
}

func NewSoundEffects() *SoundEffects {
	soundEffects := &SoundEffects{}

	const sampleRate = 44100
	audioContext := audio.NewContext(sampleRate)

	wavStream, decodeErr := wav.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(sounds.ShootWav))
	if decodeErr != nil {
		panic(decodeErr)
	}

	player, playerErr := audioContext.NewPlayer(wavStream)
	if playerErr != nil {
		panic(playerErr)
	}

	soundEffects.shootPlayer = player

	return soundEffects
}

func (se SoundEffects) PlayShoot() {
	se.shootPlayer.Rewind()
	se.shootPlayer.Play()
}
