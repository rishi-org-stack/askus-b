package auth

import (
	util "askUs/v1/util"
	cache "askUs/v1/util/cache"
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type OTP struct {
	Otp    string
	Expiry time.Time
}

var (
	DB *cache.BadgerDB
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
		Otp:    otp,
		Expiry: time.Now().Add(time.Second * time.Duration(sec)),
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
		btArray, err := util.Serialize(otp)
		if err != nil {
			return err
		}
		// DB, err := cache.Connect()
		// if err != nil {
		// 	return err
		// }
		err = DB.Set(Key, btArray)
		return err
	}
	return errors.New("otp expired")
}

func (otp *OTP) Get(key int) error {
	Key := strconv.Itoa(key)
	// DB, err := cache.Connect()
	// if err != nil {
	// 	return err
	// }
	array, err := DB.Get(Key)
	if err != nil {
		return err
	}
	err = util.DeSerialize(array, otp)
	return err
}

func (otp *OTP) Delete(key int) error {
	Key := strconv.Itoa(key)
	DB, err := cache.Connect()
	if err != nil {
		return err
	}
	err = DB.Delete(Key)
	return err
}
