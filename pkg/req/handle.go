package req

import (
	"14-TestingAPI/pkg/res"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)

	if err != nil {
		res.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		res.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &body, nil
}
func HandleQuery[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	// 1. Получаем все GET‑параметры
	queryParams := r.URL.Query()

	// 2. Собираем в map[string]string (или в структуру, если схема известна)
	params := make(map[string]string)
	for key, values := range queryParams {
		// Если параметр встречается несколько раз, берём первое значение
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		res.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	body, err := Decode[T](io.NopCloser(bytes.NewReader(jsonParams)))

	if err != nil {
		res.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		res.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &body, nil
}
