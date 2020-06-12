package example_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/canercidam/gomock-example/api"
	mock_example "github.com/canercidam/gomock-example/mocks"
	"github.com/canercidam/gomock-example/repo"
	"github.com/golang/mock/gomock"

	example "github.com/canercidam/gomock-example"
	"github.com/stretchr/testify/suite"
)

const (
	testID            = 1
	testParam         = "id=1"
	testData          = "somedata"
	testDataOnFailure = "failed"
)

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

// matcher is a custom matcher only for demonstration.
type matcher struct {
	*repo.StoredData
}

// Matches check if given struct matches with wrapped struct.
func (m *matcher) Matches(x interface{}) bool {
	d1 := m.StoredData
	d2, ok := x.(*repo.StoredData)
	if !ok {
		return false
	}

	return d1.Data == d2.Data && d1.ID == d2.ID
}

// String returns wanted value representation in string.
func (m *matcher) String() string {
	return fmt.Sprintf("%v", m.StoredData)
}

type Suite struct {
	service *example.Service
	api     *mock_example.MockAPI
	repo    *mock_example.MockRepository
	suite.Suite
}

func (s *Suite) SetupTest() {
	s.api = mock_example.NewMockAPI(gomock.NewController(s.T()))
	s.repo = mock_example.NewMockRepository(gomock.NewController(s.T()))
	s.service = example.NewService(s.api, s.repo)
}

func (s *Suite) TestSuccess() {
	s.api.EXPECT().MakeSomeQuery(testParam).Return(&api.QueryResult{Data: testData}, nil)
	s.repo.EXPECT().StoreData(&matcher{&repo.StoredData{ID: testID, Data: testData}}).Return(nil)

	err := s.service.DoSomething(testID)
	s.Nil(err)
}

func (s *Suite) TestInvalidParam() {
	s.api.EXPECT().MakeSomeQuery(testParam).Return(nil, api.ErrInvalidParam)
	s.repo.EXPECT().StoreData(&matcher{&repo.StoredData{ID: testID, Data: testDataOnFailure}}).Return(nil)

	err := s.service.DoSomething(testID)
	s.Nil(err)
}

func (s *Suite) TestStorageFailed() {
	s.api.EXPECT().MakeSomeQuery(testParam).Return(&api.QueryResult{Data: testData}, nil)
	s.repo.EXPECT().StoreData(&matcher{&repo.StoredData{ID: testID, Data: testData}}).Return(errors.New("some internal error"))

	err := s.service.DoSomething(testID)
	s.NotNil(err)
}
