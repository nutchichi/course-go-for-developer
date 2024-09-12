package flight

type FtResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []FtData `json:"data"`
}

type FtData struct {
	ID          int    `json:"id"`
	Number      int    `json:"airline_number"`
	AirlineCode string `json:"airline_code"`
	Destination string `json:"destination"`
	Arrival     string `json:"arrival"`
}

type FtRequest struct {
	Number      int    `json:"airline_number"`
	AirlineCode string `json:"airline_code"`
	Destination string `json:"destination"`
	Arrival     string `json:"arrival"`
}

type FtService interface {
	GetFlights() (FtResponse, error)
	GetFlight(int) (*FtResponse, error)
	NewFlight(FtRequest) error
	UpdateFlight(int, FtRequest) (*FtResponse, error)
	DeleteFlight(int) (*FtResponse, error)
}

type ftService struct {
	ftRepo FtRepository
}

func NewFlightService(flightRepo FtRepository) FtService {
	return ftService{ftRepo: flightRepo}
}

func (s ftService) GetFlights() (FtResponse, error) {
	fts, err := s.ftRepo.GetAll()
	if err != nil {
		return FtResponse{
			Code:    500,
			Message: err.Error(),
		}, err
	}

	ftRes := FtResponse{}
	dataAll := []FtData{}
	for _, ft := range fts {

		data := FtData{
			ID:          ft.ID,
			Number:      ft.Number,
			AirlineCode: ft.AirlineCode,
			Destination: ft.Destination,
			Arrival:     ft.Arrival,
		}

		dataAll = append(dataAll, data)

	}

	ftRes = FtResponse{
		Code:    0,
		Message: "Success",
		Data:    dataAll,
	}

	return ftRes, nil

}

func (s ftService) GetFlight(id int) (*FtResponse, error) {
	ftDb, err := s.ftRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	d := []FtData{}
	d = append(d, FtData{
		ID:          ftDb.ID,
		Number:      ftDb.Number,
		AirlineCode: ftDb.AirlineCode,
		Destination: ftDb.Destination,
		Arrival:     ftDb.Arrival,
	})

	ftRes := FtResponse{
		Code:    0,
		Message: "Success",
		Data:    d}

	return &ftRes, nil
}

func (s ftService) NewFlight(fr FtRequest) error {

	fdb := FtDB{
		Number:      fr.Number,
		AirlineCode: fr.AirlineCode,
		Destination: fr.Destination,
		Arrival:     fr.Arrival,
	}

	err := s.ftRepo.Create(fdb)

	return err
}

func (s ftService) UpdateFlight(id int, ftReq FtRequest) (*FtResponse, error) {

	ftDb := FtDB{
		Number:      ftReq.Number,
		AirlineCode: ftReq.AirlineCode,
		Destination: ftReq.Destination,
		Arrival:     ftReq.Arrival,
	}

	ftDb, err := s.ftRepo.UpdateById(id, ftDb)
	if err != nil {
		return nil, err
	}

	ds := []FtData{}
	d := FtData{
		ID:          ftDb.ID,
		Number:      ftDb.Number,
		AirlineCode: ftDb.AirlineCode,
		Destination: ftDb.Destination,
		Arrival:     ftDb.Arrival,
	}

	ds = append(ds, d)
	ftRes := &FtResponse{
		Code:    0,
		Message: "Success",
		Data:    ds,
	}

	return ftRes, nil
}

func (s ftService) DeleteFlight(id int) (*FtResponse, error) {
	err := s.ftRepo.DeleteById(id)
	if err != nil {
		return nil, err
	}

	ftRes := &FtResponse{
		Code:    0,
		Message: "Success",
		Data:    []FtData{},
	}

	return ftRes, err
}
