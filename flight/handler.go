package flight

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type flightHandler struct {
	flightSrv FlightService
}

func NewFlightHandler(flightSrv FlightService) flightHandler {
	return flightHandler{flightSrv: flightSrv}
}

func (h flightHandler) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h flightHandler) GetFlightsHandler(c *gin.Context) {
	f, err := h.flightSrv.GetFlights()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, f)

}

func (h flightHandler) GetFlightByIDHandler(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": "Invalid id",
		})
		return
	}

	f, err := h.flightSrv.GetFlight(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": "Flight not found",
		})
		return
	}

	c.JSON(200, f)
}

func (h flightHandler) CreateFlightHandler(c *gin.Context) {
	var request NewFlightRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}

	err = h.flightSrv.NewFlight(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "Flight created successfully",
	})
}

func (h flightHandler) UpdateFlightHandler(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": "Invalid id",
		})
		return
	}

	var request NewFlightRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}

	f, err := h.flightSrv.UpdateFlight(i, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, f)

}

func (h flightHandler) DeleteFlightHandler(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": "Invalid id",
		})
		return
	}

	f, err := h.flightSrv.DeleteFlight(i)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, f)
}
