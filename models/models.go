package models

const (
	CollegeAve   = "01"
	Livingston   = "03"
	Busch        = "04"
	CookDouglass = "05"
)

type ErrorResponse struct {
	StatusCode int
	Status     string
}

type MenuResponse struct {
	StatusCode int
	Status     string
	URL        string
	Menu       []MenuSection
}

type MenuSection struct {
	Name  string
	Items []MenuItem
}

type MenuItem struct {
	Name string
	Info string
}
