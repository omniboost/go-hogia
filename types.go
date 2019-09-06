package hogia

import "time"

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
