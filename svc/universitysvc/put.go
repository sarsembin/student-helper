package universitysvc

import (
	"context"
	"studentHelper/repo/postgredb/universitydb"
)

type PutRequest struct {
	JsonEntity
}

type PutResponse struct {
}

func (s *service) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	err := s.postgredb.Universitydb.Put(&universitydb.Entity{
		ID:                       req.ID,
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
	if err != nil {
		return nil, err
	}

	return &PutResponse{}, nil
}
