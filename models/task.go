package models

import "time";

type Task struct {
	ID          int
	Ticket      int
	Title       string
	Description string
	StartTime 	time.Time
	EndTime   	time.Time
}
