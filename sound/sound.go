package sound

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	"os"
	"time"
)

func Sound(path string) error {
	f, _ := os.Open(path)
	s, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	done := make(chan struct{})
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))
	_ = <-done

	return nil
}
