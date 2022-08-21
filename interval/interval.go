package interval

import "time"

type Interval struct {
	Start time.Time
	End   time.Time
}

func NewInterval() Interval {

	return Interval{
		Start: time.Now().AddDate(0, 0, -10),
		End:   time.Now().AddDate(0, 0, 1),
	}
}

func (i Interval) Past(days int) *Interval {
	i.Start = time.Now().AddDate(0, 0, -days)
	i.End = time.Now()
	return &i
}

func (i Interval) Future(days int) *Interval {
	i.Start = time.Now()
	i.End = time.Now().AddDate(0, 0, days)
	return &i
}

func (i Interval) StartDate() string {
	return i.Start.Format("2006-01-02")
}
