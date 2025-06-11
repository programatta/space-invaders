package sounds

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/programatta/spaceinvaders/internal/assets/sounds"
)

type SoundEffects struct {
	shootPlayer         *audio.Player
	invaderKilledPlayer *audio.Player
}

func NewSoundEffects() *SoundEffects {
	soundEffects := &SoundEffects{}

	const sampleRate = 44100
	audioContext := audio.NewContext(sampleRate)

	soundEffects.shootPlayer = loadSound(audioContext, sounds.ShootWav)
	soundEffects.invaderKilledPlayer = loadSound(audioContext, sounds.InvaderKilledWav)

	return soundEffects
}

func (se SoundEffects) PlayShoot() {
	se.shootPlayer.Rewind()
	se.shootPlayer.Play()
}

func (se SoundEffects) PlayAlienKilled() {
	se.invaderKilledPlayer.Rewind()
	se.invaderKilledPlayer.Play()
}

func loadSound(audioContext *audio.Context, sourceSound []byte) *audio.Player {
	wavStream, decodeErr := wav.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(sourceSound))
	if decodeErr != nil {
		panic(decodeErr)
	}

	player, playerErr := audioContext.NewPlayer(wavStream)
	if playerErr != nil {
		panic(playerErr)
	}
	return player
}
