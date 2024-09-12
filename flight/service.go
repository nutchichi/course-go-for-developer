package flight

type FlightData struct {
	ID          int    `json:"id"`
	Number      int    `json:"airline_number"`
	AirlineCode string `json:"airline_code"`
	Destination string `json:"destination"`
	Arrival     string `json:"arrival"`
}

type FlightResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []FlightData `json:"data"`
}

type NewFlightRequest struct {
	Number      int    `json:"airline_number"`
	AirlineCode string `json:"airline_code"`
	Destination string `json:"destination"`
	Arrival     string `json:"arrival"`
}

type FlightService interface {
	GetFlights() (FlightResponse, error)
	GetFlight(int) (*FlightResponse, error)
	NewFlight(NewFlightRequest) error
	UpdateFlight(int, NewFlightRequest) (*FlightResponse, error)
	DeleteFlight(int) (*FlightResponse, error)
}

type flightService struct {
	flightRepo FlightRepository
}

func NewFlightService(flightRepo FlightRepository) FlightService {
	return flightService{flightRepo: flightRepo}
}

func (s flightService) GetFlights() (FlightResponse, error) {
	flights, err := s.flightRepo.GetAll()
	if err != nil {
		return FlightResponse{
			Code:    500,
			Message: err.Error(),
		}, err
	}

	filghtReponses := FlightResponse{}
	dataAll := []FlightData{}
	for _, f := range flights {

		data := FlightData{
			ID:          f.ID,
			Number:      f.Number,
			AirlineCode: f.AirlineCode,
			Destination: f.Destination,
			Arrival:     f.Arrival,
		}

		dataAll = append(dataAll, data)

	}

	filghtReponses = FlightResponse{
		Code:    0,
		Message: "Success",
		Data:    dataAll,
	}

	return filghtReponses, nil

}

func (s flightService) GetFlight(id int) (*FlightResponse, error) {
	flight, err := s.flightRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	dataAll := []FlightData{}
	dataAll = append(dataAll, FlightData{
		ID:          flight.ID,
		Number:      flight.Number,
		AirlineCode: flight.AirlineCode,
		Destination: flight.Destination,
		Arrival:     flight.Arrival,
	})

	f := FlightResponse{
		Code:    0,
		Message: "Success",
		Data:    dataAll}

	return &f, nil
}

func (s flightService) NewFlight(request NewFlightRequest) error {

	flight := FlightDB{
		Number:      request.Number,
		AirlineCode: request.AirlineCode,
		Destination: request.Destination,
		Arrival:     request.Arrival,
	}

	err := s.flightRepo.Create(flight)

	return err
}

func (s flightService) UpdateFlight(id int, request NewFlightRequest) (*FlightResponse, error) {

	flightDB := FlightDB{
		Number:      request.Number,
		AirlineCode: request.AirlineCode,
		Destination: request.Destination,
		Arrival:     request.Arrival,
	}

	flightDB, err := s.flightRepo.UpdateById(id, flightDB)
	if err != nil {
		return nil, err
	}

	dataAll := []FlightData{}
	flightData := FlightData{
		ID:          flightDB.ID,
		Number:      flightDB.Number,
		AirlineCode: flightDB.AirlineCode,
		Destination: flightDB.Destination,
		Arrival:     flightDB.Arrival,
	}

	dataAll = append(dataAll, flightData)
	response := &FlightResponse{
		Code:    0,
		Message: "Success",
		Data:    dataAll,
	}

	return response, nil
}

func (s flightService) DeleteFlight(id int) (*FlightResponse, error) {
	err := s.flightRepo.DeleteById(id)
	if err != nil {
		return nil, err
	}

	response := &FlightResponse{
		Code:    0,
		Message: "Success",
		Data:    []FlightData{},
	}

	return response, err
}
