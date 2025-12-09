package stat

import (
	"14-TestingAPI/pkg/db"
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

func (repo *StatRepository) GetStats(FromD time.Time, ToD time.Time, by string) []GetStatResponce {
	var res []GetStatResponce
	var reqGroup string
	switch by {
	case GroupByDay:
		reqGroup = "to_char(date, 'yyyy-mm-dd') as period, sum(clicks)"
	case GroupByMonth:
		reqGroup = "to_char(date, 'yyyy-mm') as period, sum(clicks)"
	}
	repo.Db.Table("stats").
		Select(reqGroup).
		Where("date between ? and ?", FromD, ToD).
		Group("period").
		Order("period").
		Scan(&res)

	return res
}
