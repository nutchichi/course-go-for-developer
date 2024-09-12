package flight

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ftHandler struct {
	ftSrv FtService
}

func NewFlightHandler(ftSrv FtService) ftHandler {
	return ftHandler{ftSrv: ftSrv}
}

func (h ftHandler) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h ftHandler) GetFlightsHandler(c *gin.Context) {
	ftRes, err := h.ftSrv.GetFlights()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, ftRes)

}

func (h ftHandler) GetFlightByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": "Invalid id",
		})
		return
	}

	ftRes, err := h.ftSrv.GetFlight(id)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    9999,
			"message": "Flight not found",
		})
		return
	}

	c.JSON(200, ftRes)
}

func (h ftHandler) CreateFlightHandler(c *gin.Context) {
	var request FtRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}

	err = h.ftSrv.NewFlight(request)
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

func (h ftHandler) UpdateFlightHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": "Invalid id",
		})
		return
	}

	var ftReq FtRequest
	err = c.ShouldBindJSON(&ftReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}

	ftRes, err := h.ftSrv.UpdateFlight(id, ftReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ftRes)

}

func (h ftHandler) DeleteFlightHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    9999,
			"message": "Invalid id",
		})
		return
	}

	ftRes, err := h.ftSrv.DeleteFlight(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    9999,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ftRes)
}
