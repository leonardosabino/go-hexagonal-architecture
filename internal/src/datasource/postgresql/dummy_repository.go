package datasource

import (
	"hexagonal/template/internal/src/domain"

	"gorm.io/gorm"
)

type IDummyRepository interface {
	GetDummy(dummy domain.Dummy, limit int, offset int) ([]domain.Dummy, error)
	CountDummy(dummy domain.Dummy) (int64, error)
}

type DummyRepository struct {
	database Database
}

func ConstructorDummyRepository() *DummyRepository {
	return &DummyRepository{
		database: GetDatabase(),
	}
}

func (c *DummyRepository) GetDummy(dummy domain.Dummy, limit int, offset int) ([]domain.Dummy, error) {
	var dummies []DummyTable

	query := c.createQuery(dummy)

	query.Limit(int(limit)).Offset(int(offset) * int(limit))

	database := query.Find(&dummies)

	return toDummies(dummies), database.Error
}

func (c *DummyRepository) CountDummy(dummy domain.Dummy) (int64, error) {
	var count int64

	query := c.createQuery(dummy)

	database := query.Count(&count)

	return count, database.Error
}

func (c *DummyRepository) createQuery(dummy domain.Dummy) *gorm.DB {
	query := c.database.connection.Table("dummy")

	if dummy.Id != nil {
		query.Where("id = ?", dummy.Id)
	}

	if dummy.Description != nil {
		query.Where("description = ?", dummy.Description)
	}

	return query
}
