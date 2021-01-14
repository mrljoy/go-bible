package lengconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%gM", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%gF", f)
}

func MToF(m Meter) Foot { return Foot(m * 10000 / 3048) }

func FToM(f Foot) Meter { return Meter(3048 / f * 10000) }
