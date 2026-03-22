package equipment

const (
	PriceKirka        = 1
	PriceVentilyaciya = 2
	PriceVagonetka    = 3
)

type Equipment struct {
	Kirka        bool
	Ventilyaciya bool
	Vagonetka    bool
}

func (e *Equipment) BuyKirka() {
	e.Kirka = true
}

func (e *Equipment) BuyVentilyaciya() {
	e.Ventilyaciya = true
}

func (e *Equipment) BuyVagonetka() {
	e.Vagonetka = true
}
