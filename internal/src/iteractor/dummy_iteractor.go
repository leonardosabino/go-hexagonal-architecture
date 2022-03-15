package iteractor

import (
	database "hexagonal/template/internal/src/datasource/postgresql"
	"hexagonal/template/internal/src/domain"
)

type IDummyIteractor interface {
	GetDummy(dummy domain.Dummy, limit int, offset int) ([]domain.Dummy, int64, error)
}

type DummyIteractor struct {
	dummyRepository database.IDummyRepository
}

func ConstructorDummyIteractor() *DummyIteractor {
	return &DummyIteractor{
		dummyRepository: database.ConstructorDummyRepository(),
	}
}

func (ps *DummyIteractor) GetDummy(dummy domain.Dummy, limit int, offset int) ([]domain.Dummy, int64, error) {

	dummiesListChannel := make(chan []domain.Dummy, 1)
	countChannel := make(chan int64, 1)
	errorChannel := make(chan error, 2)

	go func() {
		dummiesList, repositoryError := ps.dummyRepository.GetDummy(dummy, limit, offset)
		dummiesListChannel <- dummiesList
		errorChannel <- repositoryError
	}()

	go func() {
		count, repositoryError := ps.dummyRepository.CountDummy(dummy)
		countChannel <- count
		errorChannel <- repositoryError
	}()

	dummiesList := <-dummiesListChannel
	count := <-countChannel
	repositoryError := <-errorChannel

	if repositoryError != nil {
		return nil, 0, repositoryError
	}

	return dummiesList, count, nil
}
