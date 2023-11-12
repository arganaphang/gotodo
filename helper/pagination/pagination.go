package pagination

import (
	"errors"
	"math"
	"net/url"
	"strconv"
)

type Pagination struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
	Limit     int `json:"-"`
	Offset    int `json:"-"`
}

func Transform(query url.Values) (*Pagination, error) {
	pageQuery := query.Get("page")
	if pageQuery == "" {
		pageQuery = "1"
	}
	perPageQuery := query.Get("per_page")
	if perPageQuery == "" {
		perPageQuery = "10"
	}
	page, err := strconv.ParseInt(pageQuery, 10, 32)
	if err != nil {
		return nil, errors.New("page is not a number")
	}

	perPage, err := strconv.ParseInt(perPageQuery, 10, 32)
	if err != nil {
		return nil, errors.New("per_page is not a number")
	}

	limit := perPage
	offset := (page - 1) * perPage

	return &Pagination{
		Page:    int(page),
		PerPage: int(perPage),
		Limit:   int(limit),
		Offset:  int(offset),
	}, nil
}

func (p *Pagination) Finish(count int) {
	p.Total = count
	p.TotalPage = int(math.Ceil(float64(count) / float64(p.PerPage)))
}
