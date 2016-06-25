package workshop

import "time"

type Registration struct {
	ID        int
	CourseID  int
	FirstName string
	LastName  string
	Address   string
	City      string
	State     string
	Zip       string
}

type Course struct {
	ID          int
	Name        string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Location    string
}
