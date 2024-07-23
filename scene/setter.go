package scene

import "time"

type Setter interface {
	SetScene(scene string) error
}

type BaseSetter struct {
	Set chan<- string
}

func (setter *BaseSetter) SetScene(scene string) error {
	setter.Set <- scene
	return nil
}

// Sets scene when requested via API
type ApiSetter struct {
	BaseSetter
}

// Sets scene when observing change in MusicCast scene
type MusicCastSetter struct {
	BaseSetter
}

type DummySetter struct {
	BaseSetter
}

func (s *DummySetter) Run() {
	for {
		time.Sleep(500 * time.Millisecond)
		s.SetScene(time.Now().String())
	}

}