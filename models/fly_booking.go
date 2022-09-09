package models

type DataComConfig struct {
	HeaderUser    string
	HeaderPass    string
	ProductKey    string
	Language      string
	AgentAccount  string
	AgentPassword string
}

type FlightInfoRequest struct {
	Adt int
	Chd int
	Inf int
	ViewMode string
	ListFlight []*ListFlightChoosing
  }
  
  type ListFlightChoosing struct {
	StartPoint string
	EndPoint string
	DepartDate string
	Airline string
	Month int
	Year int
  }