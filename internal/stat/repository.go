package stat

import (
	"13-AdvancedDB/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	Db *db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(LinkId uint) {
	//если нет статистики - создаем
	//если есть увеличиваем на 1
	CurrentDate := datatypes.Date(time.Now())
	var stat Stat
	repo.Db.Find(&stat, "link_id = ? and date = ?", LinkId, CurrentDate)
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: LinkId,
			Clicks: 1,
			Date:   CurrentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}

func (repo *StatRepository) GetStat(FromD time.Time, ToD time.Time, by string) (Stat, error) {
	var res Stat

	return res, nil
}
