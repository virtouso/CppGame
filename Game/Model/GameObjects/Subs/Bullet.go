package Subs

import "github.com/virtouso/GoGame/Game/Model/GameObjects"

type Bullet struct {
	GameObjects.GameObjectBase
}

func (b Bullet) Start() {
	b.GameObjectBase.Start()
	b.GameObjectBase.Components = append(b.GameObjectBase.Components)
}

func (b Bullet) Update() {
	//TODO implement me
	panic("implement me")
}

func (b Bullet) Draw() {
	//TODO implement me
	panic("implement me")
}

func (b Bullet) Finish() {
	//TODO implement me
	panic("implement me")
}

func init() {
	var _ GameObjects.GameObjectFunctions = Bullet{}
}
