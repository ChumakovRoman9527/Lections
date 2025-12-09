package stat

import (
	"14-TestingAPI/configs"
	"14-TestingAPI/pkg/middleware"
	"14-TestingAPI/pkg/req"
	"14-TestingAPI/pkg/res"
	"fmt"
	"net/http"
	"slices"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type statHandler struct {
	StatRepository *StatRepository
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &statHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), &deps.Config.Auth))

}

func (handler *statHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		layout := "2006-01-02"

		body, err := req.HandleQuery[GetStatRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		/* все что ниже должна сделать библиотека валидейт
		requiredParams := []string{"from", "to", "by"}

		for key := range requiredParams {
			if !params.Has(requiredParams[key]) {
				http.Error(w, fmt.Sprintf("Отсутствует параметр: %s", requiredParams[key]), http.StatusBadRequest)
				return
			}
		}
		*/
		// fmt.Println(body)
		requiredOrder := []string{GroupByDay, GroupByMonth}
		q_FromD := body.FromD
		q_ToD := body.ToD
		OrderBy := body.By

		FromD, err := time.Parse(layout, q_FromD)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ToD, err := time.Parse(layout, q_ToD)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !slices.Contains(requiredOrder, OrderBy) {
			http.Error(w, fmt.Sprintf("Параметр by может быть или day или month"), http.StatusBadRequest)
			return
		}

		result := handler.StatRepository.GetStats(FromD, ToD, OrderBy)
		res.Json(w, result, 200)

	}
}
