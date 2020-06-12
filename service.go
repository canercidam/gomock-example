package example

import (
	"fmt"

	"github.com/canercidam/gomock-example/repo"

	"github.com/canercidam/gomock-example/api"
)

type API interface {
	MakeSomeQuery(param string) (*api.QueryResult, error)
}

type Repository interface {
	StoreData(data *repo.StoredData) error
}

type Service struct {
	api API
	r   Repository
}

func NewService(api API, r Repository) *Service {
	return &Service{api: api, r: r}
}

func (s *Service) DoSomething(id int) error {
	param := fmt.Sprintf("id=%d", id)

	result, err := s.api.MakeSomeQuery(param)
	var data string

	switch err {
	case nil:
		data = result.Data
	case api.ErrInvalidParam:
		data = "failed"
	default:
		return fmt.Errorf("failed to make some query: %v", err)
	}

	if err := s.r.StoreData(&repo.StoredData{ID: id, Data: data}); err != nil {
		return fmt.Errorf("failed to store data: %v", err)
	}

	return nil
}
