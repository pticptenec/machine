package super

type HasLamp interface {
	On()
	Off()
}

type IntI interface {
	Int() int
}

type DoesAction interface {
	Do(IntI) Beverage
	Check() bool
}

type Beverage interface {
	String()
}

type BlendTank struct {
}

type GrindTank struct {
}

type WaterTank struct {
}

type SyrupTank struct {
}

type Tea struct{}
type Coffee struct{}
type Mint struct{}

type Lavender struct{}
type Nut struct{}
