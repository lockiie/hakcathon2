package models

import (
	"strconv"
)

type Pagination struct {
	OffSet string
	Limit  string
}

func (pag Pagination) Pag() string {
	offSet, err := strconv.Atoi(pag.OffSet)
	if err != nil {
		offSet = 0
	}
	limit, err := strconv.Atoi(pag.Limit)
	if err != nil {
		limit = 20
	}
	return " LIMIT " + strconv.Itoa(offSet) + ", " + strconv.Itoa(limit) + ""
}
