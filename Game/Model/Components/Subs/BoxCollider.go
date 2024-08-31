package Subs

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/virtouso/GoGame/Game/Model/Components"
)

type BoxCollider struct {
}

func (b BoxCollider) Start() {
	//TODO implement me
	panic("implement me")
}

func (b BoxCollider) Update() {
	//TODO implement me
	panic("implement me")
}

func (b BoxCollider) Draw(screen *ebiten.Image) {
	//TODO implement me
	panic("implement me")
}

func (b BoxCollider) End() {
	//TODO implement me
	panic("implement me")
}

func (b BoxCollider) GetName() string {
	//TODO implement me
	panic("implement me")
}

func implement() {
	var _ Components.Component = BoxCollider{}
}
