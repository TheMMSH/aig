package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Controller struct {
	Service Service
}

func NewController(s Service) Controller {
	return Controller{Service: s}
}

func (cr Controller) SetRoutes(e *gin.Engine) {
	e.GET("/health", cr.health)
	e.POST("/api/users", cr.createUser)
	e.POST("/api/users/generateotp", cr.generateOtp)
	e.POST("/api/users/verifyotp", cr.verifyOtp)
}

func (cr Controller) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (cr Controller) createUser(c *gin.Context) {
	var userReq CreateUserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user, err := cr.Service.CreateUser(c, userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (cr Controller) generateOtp(c *gin.Context) {
	var otpReq GenerateOtpRequest

	if err := c.ShouldBindJSON(&otpReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	expTime, err := cr.Service.GenerateOtp(c, otpReq.PhoneNumber)
	if err != nil {
		status := http.StatusBadRequest
		if strings.Contains(err.Error(), "not found") {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"expires_at": expTime.Format(time.RFC3339)})
}

func (cr Controller) verifyOtp(c *gin.Context) {
	var verifyOtpReq VerifyOtpRequest

	if err := c.ShouldBindJSON(&verifyOtpReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err := cr.Service.VerifyOTP(c, verifyOtpReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "otp is correct!"})
}
