package books

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type book_ctrl struct {
	repo *book_repo
}

func NewCtrl(rep *book_repo) *book_ctrl {
	return &book_ctrl{rep}
}

func (rep *book_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := rep.repo.FindAll()

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (rep *book_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data Book
	json.NewDecoder(r.Body).Decode(&data)

	result, err := rep.repo.Add(&data)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}
