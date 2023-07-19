package services

import (
	"github.com/ppinnamani/chandu/login/pkg/rest/server/daos"
	"github.com/ppinnamani/chandu/login/pkg/rest/server/models"
)

type TestfieldsService struct {
	testfieldsDao *daos.TestfieldsDao
}

func NewTestfieldsService() (*TestfieldsService, error) {
	testfieldsDao, err := daos.NewTestfieldsDao()
	if err != nil {
		return nil, err
	}
	return &TestfieldsService{
		testfieldsDao: testfieldsDao,
	}, nil
}

func (testfieldsService *TestfieldsService) CreateTestfields(testfields *models.Testfields) (*models.Testfields, error) {
	return testfieldsService.testfieldsDao.CreateTestfields(testfields)
}

func (testfieldsService *TestfieldsService) UpdateTestfields(id int64, testfields *models.Testfields) (*models.Testfields, error) {
	return testfieldsService.testfieldsDao.UpdateTestfields(id, testfields)
}

func (testfieldsService *TestfieldsService) DeleteTestfields(id int64) error {
	return testfieldsService.testfieldsDao.DeleteTestfields(id)
}

func (testfieldsService *TestfieldsService) ListTestfields() ([]*models.Testfields, error) {
	return testfieldsService.testfieldsDao.ListTestfields()
}

func (testfieldsService *TestfieldsService) GetTestfields(id int64) (*models.Testfields, error) {
	return testfieldsService.testfieldsDao.GetTestfields(id)
}
