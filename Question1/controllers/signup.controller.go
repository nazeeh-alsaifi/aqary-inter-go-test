package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nazeeh-alsaifi/aqary-inter-go-test/db/sqlc"
)

type AuthController struct {
	db *db.Queries
}

func NewAuthController(db *db.Queries) *AuthController {
	return &AuthController{db}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var credentials *db.User

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check if phone unique
	count, err := ac.db.CheckPhoneNumberUnique(ctx, credentials.PhoneNumber)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	if count != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": "we have this phone number in our records!"})
		return
	}
	// end check

	// create user
	args := &db.CreateUserParams{
		Name:        credentials.Name,
		PhoneNumber: credentials.PhoneNumber,
	}

	user, err := ac.db.CreateUser(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": user}})
	// end create
}

func (ac *AuthController) GenerateOtp(ctx *gin.Context) {
	var credentials *db.User

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check if phone number exists
	count, err := ac.db.CheckPhoneNumberUnique(ctx, credentials.PhoneNumber)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	if count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "msg": "user not found!"})
		return
	}
	// end check

	// create user
	args := &db.UpdateUserByPhoneNumberParams{
		PhoneNumber:       credentials.PhoneNumber,
		Otp:               pgtype.Text{String: generateRandomDigits(4), Valid: true},
		OtpExpirationTime: pgtype.Timestamp{Time: time.Now().UTC().Local().Add(time.Minute), Valid: true},
	}

	user, err := ac.db.UpdateUserByPhoneNumber(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": user}})
	// end create
}

func (ac *AuthController) VerifyOtp(ctx *gin.Context) {
	var credentials *db.User

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check if phone unique
	user, err := ac.db.GetUserByPhoneNumber(ctx, credentials.PhoneNumber)

	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "msg": "user not found!"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	// check if otp correct
	otpValid := credentials.Otp.String == user.Otp.String

	timeNow := time.Now()
	time1 := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), 0, time.UTC)

	time2 := user.OtpExpirationTime.Time

	otpTimedOut := time1.After(time2)
	if !otpValid {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "failed", "msg": "otp not valid!"})
		return
	}

	if otpTimedOut {
		ctx.JSON(http.StatusForbidden, gin.H{"fail": "failed", "msg": "otp expired!"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": user, "time.Now()": time1, "user.OtpExpirationTime.Time.UTC()": time2, "otpTimedOut": otpTimedOut}})
	// end check
}

func generateRandomDigits(count int) string {
	digits := make([]byte, count)
	for i := 0; i < count; i++ {
		digits[i] = byte(rand.Intn(10)) + '0' // Generates a random digit (0-9)
	}
	return string(digits)
}
