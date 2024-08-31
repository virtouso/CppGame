package GameObjects

import "github.com/virtouso/GoGame/Game/Model/Components"

type GameObjectFunctions interface {
	Start()
	Update()
	Draw()
	Finish()
}

type GameObjectBase struct {
	Components map[string]*Components.Component
}

func (g GameObjectBase) Start() {
	//TODO implement me
	panic("implement me")
}

func (g GameObjectBase) Update() {
	//TODO implement me
	panic("implement me")
}

func (g GameObjectBase) Draw() {
	//TODO implement me
	panic("implement me")
}

func (g GameObjectBase) Finish() {
	//TODO implement me
	panic("implement me")
}

func implement() {
	var _ GameObjectFunctions = GameObjectBase{}

}
