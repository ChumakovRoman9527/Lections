package link

import (
	"14-TestingAPI/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepositoryDeps struct {
	DataBase *db.Db
}

type LinkRepository struct {
	DataBase *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		DataBase: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.DataBase.DB.First(&link, "hash=?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) CheckByHash(hash string) (bool, error) {
	var link Link
	var exists bool
	result := repo.DataBase.DB.Where(&link, "hash=?", hash).Find(&exists)
	if result.Error != nil {
		return false, result.Error
	}
	return exists, nil
}

func (repo *LinkRepository) GetByID(id uint) (*Link, error) {
	var link Link
	result := repo.DataBase.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {

	result := repo.DataBase.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) GetCount() int64 {

	var count int64

	repo.DataBase.
		Table("links").
		Where("deleted_at is null").
		Count(&count)

	return count
}

func (repo *LinkRepository) GetAll(limit, offset int) []Link {
	var links []Link

	repo.DataBase.
		Table("links").
		Where("deleted_at is null").
		Order("id asc").
		Limit(limit).
		Offset(offset).
		Scan(&links)

	return links
}

func (repo *LinkRepository) GetAllLinksResponse(limit, offset int) GetAllLinksResponse {
	links := repo.GetAll(limit, offset)
	count := repo.GetCount()
	return GetAllLinksResponse{
		Links: links,
		Count: int(count),
	}
}
