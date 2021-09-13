package config

import (
	"os"
	"strconv"
)

type (
	Env struct {
		DB                     string
		Port                   string
		DatabaseContextTimeout int
		JWTDurtaion            int
		Algo                   string
		Key                    string
		OTPExpiry              int
	}
)

func Init() *Env {
	dbTimeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		dbTimeout = 10
	}
	otpExp, err := strconv.Atoi(os.Getenv("OTP_EXPIRY"))
	if err != nil {
		otpExp = 60
	}
	return &Env{
		DB:                     os.Getenv("DB"),
		Port:                   ":" + os.Getenv("PORT"),
		DatabaseContextTimeout: dbTimeout,
		Algo:                   "HS256",
		Key:                    "RishiStack!1709",
		JWTDurtaion:            60,
		OTPExpiry:              otpExp,
	}
}
