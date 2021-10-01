package advice

type AdviceService struct {
	AdviceData DB
}

func Init(db DB) Service {
	return &AdviceService{
		AdviceData: db,
	}
}
