package entities

type Resume struct {
	ResumeID int
	Name     string
}

type Process struct {
	ProcessID int
	ResumeID  int
	JobID     int
}
