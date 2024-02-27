package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get All News
// @Tags news
// @Description GetAllNews
// @ID GetAllNews
// @Accept  json
// @Produce  json
// @Success 200 {array} model.GetAllNews
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news [get]
func (h *Handler) GetAllNews(c *gin.Context) {
	news, err := h.services.News.GetAllNews()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "can't find news")
		return
	}

	c.JSON(http.StatusOK, news)
}

// @Summary get News By Id
// @Tags news
// @Description getNewsById
// @ID getNewsById
// @Accept  json
// @Produce  json
// @Param  id  path  int  true  "News ID"
// @Success 200 {object} model.GetNewsById
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/{id} [get]
func (h *Handler) getNewsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	news, err := h.services.News.FindNewsById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, news)
}
