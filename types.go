package hogia

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) String() string {
	return d.Format("2006-01-02")
}

type IntBool bool

func (b IntBool) String() string {
	if b {
		return "0"
	}

	return "-1"
}

type Amount float64

func (a Amount) String() string {
	if a == 0.0 {
		return ""
	}

	return a.Round2()
}

func (a Amount) Round2() string {
	s := fmt.Sprintf("%.2f", float64(a))
	return strings.Replace(s, ".", ",", -1)
}

func (a Amount) Round4() string {
	s := fmt.Sprintf("%.4f", float64(a))
	return strings.Replace(s, ".", ",", -1)
}

type Quantity int

func (q Quantity) String() string {
	if q == 0 {
		return ""
	}

	return fmt.Sprint(int(q))
}

type Dimension int

func (d Dimension) String() string {
	if d == 0 {
		return ""
	}

	return fmt.Sprint(int(d))
}

type Project int

func (p Project) String() string {
	if p == 0 {
		return ""
	}

	return fmt.Sprint(int(p))
}

type FloatOpt float64

func (f FloatOpt) String() string {
	if f == 0.0 {
		return ""
	}

	s := fmt.Sprintf("%.4f", float64(f))
	return strings.Replace(s, ".", ",", -1)
}

type IntOpt int

func (i IntOpt) String() string {
	if i == 0 {
		return ""
	}

	return fmt.Sprint(int(i))
}
