package auth

import (
	cache "askUs/v1/util/cache"
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type OTP struct {
	Otp      string
	Expiry   time.Time
	duration time.Duration
}

var (
	DB *cache.RedisDB
)

func init() {
	DB, _ = cache.Connect()
}
func GenrateOtp(sec int) *OTP {
	rand.Seed(time.Now().UnixNano())
	s := "1234567890"
	otp := ""
	for i := 0; i < 6; i++ {
		index := rand.Intn(len(s))
		otp += string(s[index])
	}
	return &OTP{
		Otp:      otp,
		Expiry:   time.Now().Add(time.Second * time.Duration(sec)),
		duration: time.Second * time.Duration(sec),
	}
}

func isExpired(otp *OTP) bool {
	if time.Now().Add(time.Second*1) == otp.Expiry {
		return true
	}
	return false
}

func (otp *OTP) Set(key int) error {
	if !isExpired(otp) {
		Key := strconv.Itoa(key)
		err := DB.Set(Key, otp.Otp, otp.duration)
		return err
	}
	return errors.New("otp expired")
}

func (otp *OTP) Get(key int) (string, error) {
	Key := strconv.Itoa(key)
	array, err := DB.Get(Key)
	if err != nil {
		return "", err
	}
	return array, nil
}

func (otp *OTP) Delete() error {
	// Key := strconv.Itoa(key)
	DB, err := cache.Connect()
	if err != nil {
		return err
	}
	err = DB.Delete()
	return err
}
