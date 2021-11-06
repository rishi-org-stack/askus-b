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
		Specialities       pq.StringArray `gorm:"type:text[]"`
		ExpInYears         string
		Experiences        *[]Experience
		DegreeID           int
		Degree             pq.StringArray `gorm:"type:text[]"`
		FollowingDoctors   []FollowingDoctor
		FollowedByDoctors  []FollowedByDoctor
		FollowedByPatients []FollowedByPatient
		Requests           []Request
	}
	//FD
	FollowingDoctor struct {
		ID       int
		DoctorID int
		Doctor   *Doctor
		UserID   int
		User     *Doctor
	}
	//FDBP
	FollowedDoctorsByPatient struct {
		ID        int
		PatientID int
		Patient   *Patient
		UserID    int
		User      *Doctor
	}
	//FBD
	FollowedByDoctor struct {
		ID       int
		DoctorID int
		Doctor   *Doctor
		UserID   int
		User     *Doctor
	}
	//FBP
	FollowedByPatient struct {
		ID       int
		DoctorID int
		Doctor   *Doctor
		UserID   int
		User     *Patient
	}
	Experience struct {
		ID            int `json:"id" gorm:"primary"`
		Institution   *Institution
		WorkedBetween string
		Department    string
		Title         string
		DoctorID      int `gorm:"not null"`
		Doctor        *Doctor
	}
	Institution struct {
		ID int `json:"id" gorm:"primary"`
		// Address
		Name         string
		ExperienceID int
		Experience   *Experience
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
		Symptoms pq.StringArray `gorm:"type:text[]"`
		// ConnectionsID int            `json:"connectionsID"`
		// Connections   *Connections
	}
	Degree struct {
	}
	Info struct {
		Email string
		Phone string
		Name  string `json:"name" gorm:"not null"`
		Age   int
		State string
		Sex   string
		// Address
	}
	// Connections struct {
	// ID int `json:"id" gorm:"primary"`
	// Doctors       []Doctor  `json:"doctors"`
	// Patients      []Patient `json:"patients"`
	// ConnectedWith string    `json:"connectedWith"`
	// }
	Request struct {
		ID         int `json:"id" gorm:"primary"`
		SenderID   int
		DoctorID   int
		Status     string
		GenratedBy string
		Doctor     *Doctor
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
