package coffee

// CoffeeBean is the object of coffee bean.
type CoffeeBean struct {
	Origin    string
	RoastType RoastType
}

// CoffeePowder is derived from coffee beans
type CoffeePowder struct {
	CoffeeBean
	GrindSize GrindSize
}

// Milk has different types and
type Milk struct {
	Type   MilkType
	Volume volume
}

// RoastType represents how the coffee bean is roasted.
type RoastType string

// GrindSize represents how the coffee grinder grinds the coffee beans.
type GrindSize string

// CoffeeSize is the size of a cup of coffee.
type CoffeeSize string

// MilkType is the type of milk.
type MilkType string

// SyrupType The type of syrup.
type SyrupType string

// volume is the volume of any liquid eg. Milk, Water, Expresso.
// To simplify, we use a integer to represent it.
type volume int

const (
	GrindSizeCoarse GrindSize = "Coarse"
	GrindSizeMedium GrindSize = "Medium"
	GrindSizeFine   GrindSize = "Fine"

	CoffeeSizeLarge  CoffeeSize = "Large"
	CoffeeSizeMedium CoffeeSize = "Medium"
	CoffeeSizeSmall  CoffeeSize = "Small"

	Light  RoastType = "Light"
	Medium RoastType = "Medium"
	Dark   RoastType = "Dark"

	Foam       MilkType = "Foam"
	HalfHalf   MilkType = "Half&Half"
	Cream      MilkType = "Cream"
	Whole      MilkType = "Whole Milk"
	Half       MilkType = "50% Milk"
	Zero       MilkType = "0% Milk"
	AlmondMilk MilkType = "Almond Milk"
	SoyMilk    MilkType = "Soy Milk"

	Choccolate SyrupType = "Chocolate"
	Caramel    SyrupType = "Caramel"
	Vanile     SyrupType = "Vanile"
	Hazelnut   SyrupType = "Hazelnut"
)

// type Flavor struct {
// 	Sweet   string
// 	Floral  string
// 	Fruity  string
// 	Sour    string
// 	Green   string
// 	Roasted string
// 	Spices  string
// 	Nutty   string
// }

type CoffeeDrink struct {
	Size           CoffeeSize
	ExpressoVolume volume
	Milk           *Milk
	Syrup          SyrupType
	Others         map[string]int
	powder         CoffeePowder
}

// NewCoffeeDrink creates basic coffee drink.
// ExpressoVolume, CoffeePowder and size are required.
func NewCoffeeDrink(powder CoffeePowder, size CoffeeSize) *CoffeeDrink {
	return &CoffeeDrink{
		powder: powder,
		Size:   size,
	}
}





