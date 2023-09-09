package valueobjects

type ScheduleType int

const (
	Minutely ScheduleType = iota
	Hourly
	Daily
	Weekly
	Monthly
)

type Schedule struct {
	Interval int
	Type     ScheduleType
}
