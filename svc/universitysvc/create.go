package universitysvc

import (
	"context"
	"studentHelper/repo/postgredb/universitydb"
)

func (s *service) Add(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	s.postgredb.Universitydb.Add(&universitydb.University{
		Title:                    req.Title,
		Address:                  req.Address,
		Country:                  req.Country,
		Region:                   req.Region,
		Scholarships:             req.Scholarships,
		MaleToFemale:             req.MaleToFemale,
		NumberOfStudents:         req.NumberOfStudents,
		TuitionFee:               req.TuitionFee,
		PrcInternationalStudents: req.PrcInternationalStudents,
		Description:              req.Description,
		StudentsPerStaff:         req.StudentsPerStaff,
		AcceptanceRate:           req.AcceptanceRate,
		AvgACTComposite:          req.AvgACTComposite,
		AvgACTEnglish:            req.AvgACTEnglish,
		Rank:                     req.Rank,
		AvgACTMath:               req.AvgACTMath,
		AvgSATReadingWriting:     req.AvgSATReadingWriting,
		AvgSATMath:               req.AvgSATMath,
	})

	return &CreateResponse{}, nil
}
