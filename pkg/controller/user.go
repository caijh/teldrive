package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgdrive/teldrive/pkg/httputil"
)

func (uc *Controller) GetStats(c *gin.Context) {
	res, err := uc.UserService.GetStats(c)
	if err != nil {
		httputil.NewError(c, err.Code, err.Error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *Controller) UpdateChannel(c *gin.Context) {
	res, err := uc.UserService.UpdateChannel(c)
	if err != nil {
		httputil.NewError(c, err.Code, err.Error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *Controller) ListChannels(c *gin.Context) {
	res, err := uc.UserService.ListChannels(c)
	if err != nil {
		httputil.NewError(c, err.Code, err.Error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *Controller) AddBots(c *gin.Context) {
	res, err := uc.UserService.AddBots(c)
	if err != nil {
		httputil.NewError(c, err.Code, err.Error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *Controller) ListSessions(c *gin.Context) {
	res, err := uc.UserService.ListSessions(c)
	if err != nil {
		httputil.NewError(c, err.Code, err.Error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *Controller) RemoveSession(c *gin.Context) {
	res, err := uc.UserService.RemoveSession(c)
	if err != nil {
		httputil.NewError(c, err.Code, err.Error)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (uc *Controller) RemoveBots(c *gin.Context) {
	res, err := uc.UserService.RemoveBots(c)
	if err != nil {
		httputil.NewError(c, err.Code, err.Error)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *Controller) GetProfilePhoto(c *gin.Context) {
	if c.Query("photo") != "" {
		uc.UserService.GetProfilePhoto(c)
	}
}
