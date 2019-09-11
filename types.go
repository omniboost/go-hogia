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

	s := fmt.Sprintf("%.2f", float64(a))
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
