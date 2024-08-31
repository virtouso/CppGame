package Systems

type ISystem interface {
	Start()
	Update()
	Draw()
	Finish()
}
