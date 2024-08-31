package Subs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/virtouso/GoGame/Game/Model/Components"
)

type ImageRenderer struct {
	image     *ebiten.Image
	ImageName string
	name      string
}

func (i ImageRenderer) GetName() string {
	return "image_renderer"
}

func (i ImageRenderer) Start() {
	//TODO implement me
	panic("implement me")
}

func (i ImageRenderer) Update() {
	//TODO implement me
	panic("implement me")
}

func (i ImageRenderer) Draw(screen *ebiten.Image) {
	screen.DrawImage(i.image, nil)
}

func (i ImageRenderer) End() {
	//TODO implement me
	panic("implement me")
}

func Implement() {
	var _ Components.Component = ImageRenderer{}

}
