package universitysvc

import (
	"context"
)

type GetAllRequest struct {
}

type GetAllResponse struct {
	Data []Data `json:"data"`
}

func (s *service) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	res, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var data []Data

	for i := range res {
		data = append(data, Data{
			ID: res[i].ID,
			JsonEntity: JsonEntity{
				Title:                    res[i].Title,
				Address:                  res[i].Address,
				Country:                  res[i].Country,
				Region:                   res[i].Region,
				Scholarships:             res[i].Scholarships,
				MaleToFemale:             res[i].MaleToFemale,
				NumberOfStudents:         res[i].NumberOfStudents,
				TuitionFee:               res[i].TuitionFee,
				PrcInternationalStudents: res[i].PrcInternationalStudents,
				Description:              res[i].Description,
				StudentsPerStaff:         res[i].StudentsPerStaff,
				AcceptanceRate:           res[i].AcceptanceRate,
				AvgACTComposite:          res[i].AvgACTComposite,
				AvgACTEnglish:            res[i].AvgACTEnglish,
				AvgACTMath:               res[i].AvgACTMath,
				AvgSATReadingWriting:     res[i].AvgSATReadingWriting,
				AvgSATMath:               res[i].AvgSATReadingWriting,
				Rank:                     res[i].Rank,
			},
		},
		)
	}

	return &GetAllResponse{Data: data}, nil
}
