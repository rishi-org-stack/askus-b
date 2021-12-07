package user

//TODO: request service is almost done need:
//TODO: a route to get all followed doctors following doctors followed by patients
//TODO: needa rbac for user request module Done
//TODO: need a  new caching system
import (
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"

	"github.com/lib/pq"
)

type (
	DB interface {
		FindOrCreateDoctor(ctx context.Context, email string) (*Doctor, error)
		FindOrCreatePatient(ctx context.Context, email string) (*Patient, error)
		// CreateConnection(ctx context.Context, conn *Connections) (*Connections, error)
		GetDoctorByID(ctx context.Context, id float64) (*Doctor, error)
		GetDoctorByName(ctx context.Context, name string) (*Doctor, error)
		GetPatientByID(ctx context.Context, id float64) (*Patient, error)
		UpdatePatientByID(ctx context.Context, pt *Patient, id float64) (*Patient, error)
		UpdateDoctorByID(ctx context.Context, doc *Doctor, id float64) (*Doctor, error)
		CreateReq(ctx context.Context, req *Request) (*Request, error)
		GetReqsWithSenderID(ctx context.Context, id float64) (*[]Request, error)
		GetReqsWithDocID(ctx context.Context, id float64) (*[]Request, error)
		GetReqWithSenderandDocID(ctx context.Context, senderId float64, docId, typeOfClient string) (*Request, error)
		GetReq(ctx context.Context, id string) (*Request, error)
		UpdateReq(ctx context.Context, req *Request) (*Request, error)
		CreateFBD(ctx context.Context, req *FollowedByDoctor) (*FollowedByDoctor, error)
		GetFBD(ctx context.Context, docId float64) (*[]FollowedByDoctor, error)
		CreateFD(ctx context.Context, fd *FollowingDoctor) (*FollowingDoctor, error)
		GetFD(ctx context.Context, userID float64) (*[]FollowingDoctor, error)
		CreateFBP(ctx context.Context, req *FollowedByPatient) (*FollowedByPatient, error)
		GetFBP(ctx context.Context, docID float64) (*[]FollowedByPatient, error)
		CreateFDBP(ctx context.Context, req *FollowedDoctorsByPatient) (*FollowedDoctorsByPatient, error)
		GetFDBP(ctx context.Context, patientId float64) (*[]FollowedDoctorsByPatient, error)
	}
	Service interface {
		FindOrCreateDoctor(ctx context.Context, email string) (*Doctor, utilError.ApiErrorInterface)
		FindOrCreatePatient(ctx context.Context, email string) (*Patient, utilError.ApiErrorInterface)
		GetDoctorByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetDoctorByName(ctx context.Context, name string) (*response.Response, utilError.ApiErrorInterface)
		GetPatientByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetUserByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetMyRequests(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetRequestForMe(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		UpdatePatientByID(ctx context.Context, pt *Patient) (*response.Response, utilError.ApiErrorInterface)
		UpdateDoctortByID(ctx context.Context, doc *Doctor) (*response.Response, utilError.ApiErrorInterface)
		CreateReq(ctx context.Context, id string) (*response.Response, utilError.ApiErrorInterface)
		UpdateStatusOfReq(ctx context.Context, reqId string, status string) (*response.Response, utilError.ApiErrorInterface)
		GetFBD(context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetFD(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetFBP(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetFDBP(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		IsFriend(context.Context, int) (bool, utilError.ApiErrorInterface)
	}
	ReportDependence interface {
		IsFriend(context.Context, int) (bool, utilError.ApiErrorInterface)
	}
	//TODO:User ID needs to be of type Object Id
	Doctor struct {
		ID int `json:"id" gorm:"primary"`
		Info
		// Address
		Specialities       pq.StringArray      `gorm:"type:text[]" json:"specialities"`
		ExpInYears         string              `json:"exp_in_years"`
		Experiences        *[]Experience       `json:"experience"`
		DegreeID           int                 `json:"degree_id"`
		Degree             pq.StringArray      `gorm:"type:text[]" json:"degree"`
		FollowingDoctors   []FollowingDoctor   `json:"following_doctors"`
		FollowedByDoctors  []FollowedByDoctor  `json:"followed_by_doctors"`
		FollowedByPatients []FollowedByPatient `json:"followed_by_patients"`
		Requests           []Request           `json:"requests"`
	}
	//FD
	FollowingDoctor struct {
		ID       int     `json:"id" gorm:"primary"`
		DoctorID int     `json:"doctor_id"`
		Doctor   *Doctor `json:"doctor"`
		UserID   int     `json:"user_id"`
		User     *Doctor `json:"user"`
	}
	//FDBP
	FollowedDoctorsByPatient struct {
		ID        int      `json:"id" gorm:"primary"`
		PatientID int      `json:"patient_id"`
		Patient   *Patient `json:"patient"`
		UserID    int      `json:"user_id"`
		User      *Doctor  `json:"user"`
	}
	//FBD
	FollowedByDoctor struct {
		ID       int     `json:"id" gorm:"primary"`
		DoctorID int     `json:"doctor_id"`
		Doctor   *Doctor `json:"doctor"`
		UserID   int     `json:"user_id"`
		User     *Doctor `json:"user"`
	}
	//FBP
	FollowedByPatient struct {
		ID       int      `json:"id" gorm:"primary"`
		DoctorID int      `json:"doctor_id"`
		Doctor   *Doctor  `json:"doctor"`
		UserID   int      `json:"user_id"`
		User     *Patient `json:"user"`
	}
	Experience struct {
		ID            int          `json:"id" gorm:"primary"`
		Institution   *Institution `json:"Institution"`
		WorkedBetween string       `json:"worked_between"`
		Department    string       `json:"department"`
		Title         string       `json:"title"`
		DoctorID      int          `gorm:"not null" json:"doctor_id"`
		Doctor        *Doctor
	}
	Institution struct {
		ID int `json:"id" gorm:"primary"`
		// Address
		Name         string      `json:"name"`
		ExperienceID int         `json:"experience_id"`
		Experience   *Experience `json:"experience"`
	}
	Address struct {
		// ID      int `json:"id" gorm:"primary"`
		City    string
		Country string
		PinCode string
		State   string
		Street  string
	}
	Patient struct {
		ID int `json:"id" gorm:"primary"`
		Info
		// Address   *Address
		Symptoms pq.StringArray `gorm:"type:text[]" json:"symptoms"`
		// ConnectionsID int            `json:"connectionsID"`
		// Connections   *Connections
	}
	Degree struct {
	}
	Info struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
		Name  string `json:"name" gorm:"not null"`
		Age   int    `json:"age"`
		State string `json:"state"`
		Sex   string `json:"sex"`
		// Address
	}
	// Connections struct {
	// ID int `json:"id" gorm:"primary"`
	// Doctors       []Doctor  `json:"doctors"`
	// Patients      []Patient `json:"patients"`
	// ConnectedWith string    `json:"connectedWith"`
	// }
	Request struct {
		ID         int     `json:"id" gorm:"primary"`
		SenderID   int     `json:"sender_id"`
		DoctorID   int     `json:"doctor_id"`
		Status     string  `json:"status"`
		GenratedBy string  `json:"genrated_by"`
		Doctor     *Doctor `json:"doctor"`
	}
)

const (
	DoctorClient  = "doctor"
	PatientClient = "patient"
)

const (
	Onboarded = "ONBOARDED"
	NEW       = "NEW"
	Dropped   = "DROPPED"
)
const (
	ACCEPTED = "ACCEPTED"
	REJECTED = "REJECTED"
	PENDING  = "PENDING"
)

const Schema = "usr"

func (Doctor) TableName() string {
	return Schema + ".doctors"
}

func (Patient) TableName() string {
	return Schema + ".patients"
}

func (Institution) TableName() string {
	return Schema + ".institutions"
}

func (Experience) TableName() string {
	return Schema + ".experiences"
}

func (Request) TableName() string {
	return Schema + ".requests"
}

func (FollowedByDoctor) TableName() string {
	return Schema + ".followed_by_doctors"
}
func (FollowedByPatient) TableName() string {
	return Schema + ".followed_by_patients"
} // func (Connections) TableName() string

func (FollowingDoctor) TableName() string {
	return Schema + ".following_doctors"
}

func (FollowedDoctorsByPatient) TableName() string {
	return Schema + ".followed_doctors_by_patient"
}
