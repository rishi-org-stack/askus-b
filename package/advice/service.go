package advice

import "github.com/lib/pq"

type (
	DB interface {
	}
	Service interface {
	}
	Advice struct {
		ID             int            `gorm:"primary" json:"ID"`
		Heading        string         `json:"heading"`
		Body           string         `json:"body"`
		LikedByPatient []Like         `json:"likedByPatient"`
		LikedByDoc     []Like         `json:"likedByDoc"`
		Tags           pq.StringArray `json:"tags"`
		PostedBy       int            `json:"postedBy"`
		Type           string         `json:"type"`
		PostedFor      int            `json:"postedFor"`
	}

	Like struct {
		ID       int `gorm:"primary" json:"ID"`
		LikedBy  int `json:"likedBy"`
		AdviceID int `json:"adviceID"`
		Advice   *Advice
	}
)

const (
	GLOBAL   = "GLOBAL"
	PERSONEL = "PERSONEL"
)
