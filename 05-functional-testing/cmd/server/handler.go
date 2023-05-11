package server

import (
	"functional/prey"
	"functional/shark"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	shark shark.Shark
	prey  prey.Prey
}

func NewHandler(shark shark.Shark, prey prey.Prey) *Handler {
	return &Handler{shark: shark, prey: prey}
}

// PUT: /v1/shark

func (h *Handler) ConfigureShark() gin.HandlerFunc {
	type request struct {
		XPosition float64 `json:"x_position"`
		YPosition float64 `json:"y_position"`
		Speed     float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		var sharkRequest request
		if err := context.ShouldBindJSON(&sharkRequest); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		h.shark.Configure([2]float64{sharkRequest.XPosition, sharkRequest.YPosition}, sharkRequest.Speed)

		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

// PUT: /v1/prey

func (h *Handler) ConfigurePrey() gin.HandlerFunc {
	type request struct {
		Speed float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		var preyRequest request
		if err := context.ShouldBindJSON(&preyRequest); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		h.prey.SetSpeed(preyRequest.Speed)

		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

// POST: /v1/simulate

func (h *Handler) SimulateHunt() gin.HandlerFunc {
	type response struct {
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Time    float64 `json:"time"`
	}

	return func(context *gin.Context) {
		err, res := h.shark.Hunt(h.prey)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
				"time":    0,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "could catch it",
			"time":    res,
		})

	}
}
