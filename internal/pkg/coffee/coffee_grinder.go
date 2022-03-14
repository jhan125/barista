package coffee

type Grinder interface {
	SetGrindSize(GrindSize)
	Grind(CoffeeBean) CoffeePowder
}

type BasicGrinder struct {
	size GrindSize
}

func NewBasicGrinder() *BasicGrinder {
	return &BasicGrinder{
		size: GrindSizeMedium,
	}
}

func (g *BasicGrinder) SetGrindSize(size GrindSize) {
	g.size = size
}

func (g *BasicGrinder) Grind(bean CoffeeBean) CoffeePowder {
	return CoffeePowder{
		CoffeeBean: bean,
		GrindSize:  g.size,
	}
}