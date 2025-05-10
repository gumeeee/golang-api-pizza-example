package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReview(c *gin.Context) {
	pizzaId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var newReview models.Review

	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := service.ValidateReviewRating(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	for i, pizza := range data.Pizzas {
		if pizza.ID == pizzaId {
			pizza.Review = append(pizza.Review, newReview)
			data.Pizzas[i] = pizza
			data.SavePizzas()

			c.JSON(http.StatusCreated, gin.H{
				"message": "Review added successfully",
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Pizza not found",
	})
}
