package Components

import "github.com/hajimehoshi/ebiten/v2"

type Component interface {
	Start()
	Update()
	Draw(screen *ebiten.Image)
	End()

	GetName() string
}
