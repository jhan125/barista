package coffee

type CoffeeMaker interface {
	MakeCoffee(CoffeePowder) CoffeeDrink
}

type CoffeeOption func(*CoffeeDrink)

func WithMilk(milk *Milk) CoffeeOption {
	return func(c *CoffeeDrink) {
		c.Milk = milk
	}
}
