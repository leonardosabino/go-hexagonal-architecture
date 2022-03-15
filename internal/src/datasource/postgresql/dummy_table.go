package datasource

import (
	"hexagonal/template/internal/src/domain"
)

type DummyTable struct {
	Id          string
	Description string
}

func (dummyTable DummyTable) ToDummy() domain.Dummy {
	return domain.Dummy{
		Id:          &dummyTable.Id,
		Description: &dummyTable.Description,
	}
}

func toDummies(dummiesTable []DummyTable) []domain.Dummy {
	var dummies []domain.Dummy

	for _, dummy := range dummiesTable {
		dummies = append(dummies, dummy.ToDummy())
	}
	return dummies
}
