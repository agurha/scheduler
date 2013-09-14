package models

type CronTime struct {
	Second     int
	Minute     int
	Hour       int
	DayOfMonth int
	Month      int
	DayOfWeek  int
}
