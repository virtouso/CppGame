package Subs

import (
	"github.com/virtouso/GoGame/Game/Model/Components/Subs"
	"github.com/virtouso/GoGame/Game/Model/Systems"
)

type GraphicsSystem struct {
	images []*Subs.ImageRenderer
}

type IGraphicsSystem interface {
	Systems.ISystem
	AddImageRenderer(image *Subs.ImageRenderer)
}
