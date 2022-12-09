package universitysvc

import (
	"context"
)

type GetRequest struct {
	ID int `json:"id"`
}

type GetResponse struct {
	JsonEntity
}

func (s *service) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	res, err := s.postgredb.Universitydb.Get(req.ID)
	if err != nil {
		return nil, err
	}

	return &GetResponse{JsonEntity: JsonEntity{
		ID:                       res.ID,
		Title:                    res.Title,
		Address:                  res.Address,
		Country:                  res.Country,
		Region:                   res.Region,
		Scholarships:             res.Scholarships,
		MaleToFemale:             res.MaleToFemale,
		NumberOfStudents:         res.NumberOfStudents,
		TuitionFee:               res.TuitionFee,
		PrcInternationalStudents: res.PrcInternationalStudents,
		Description:              res.Description,
		StudentsPerStaff:         res.StudentsPerStaff,
		AcceptanceRate:           res.AcceptanceRate,
		AvgACTComposite:          res.AvgACTComposite,
		AvgACTEnglish:            res.AvgACTEnglish,
		AvgACTMath:               res.AvgACTMath,
		AvgSATReadingWriting:     res.AvgSATReadingWriting,
		AvgSATMath:               res.AvgSATReadingWriting,
		Rank:                     res.Rank,
	}}, nil
}
