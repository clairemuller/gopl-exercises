// Pacakge weightconv converts its numeric argument to pounds and kilograms

package weightconv

import (
	"fmt"
)

// declare two named types with underlying type of float64
type Pound float64
type Kilogram float64

// PToK converts pounds to kilograms
func PToK(p Pound) Kilogram {
	return Kilogram(p / 2.205)
}

// KToP converts kilograms to pounds
func KToP(k Kilogram) Pound {
	return Pound(k * 2.205)
}

// declare String methods to control how
func (p Pound) String() string {
	return fmt.Sprintf("%g pounds", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kilograms", k)
}
