package universitysvc

import (
	"context"
	"fmt"
)

type CreateRequest struct {
	ID                                int
	Title                             string
	Address                           string
	Country                           string
	Region                            string
	Scholarships                      string
	MaleToFemale                      string
	NumberOfStudents                  string
	TuitionFee                        string
	PercentageOfInternationalStudents float64
	Description                       string
	StudentsPerStaff                  float64
	AcceptanceRate                    float64
	AvgACTComposite                   int
	AvgACTEnglish                     int
	Rank                              string
	AvgACTMath                        int
	AvgSATReadingWriting              int
	AvgSATMath                        int
}

type CreateResponse struct {
}

func (s *service) Add(context.Context, *CreateRequest) (*CreateResponse, error) {
	res := &CreateResponse{}

	fmt.Println("amogus")

	return res, nil
}
