package manti

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type Search struct {
	search      string
	columns     []string
	fromIndexes []string
	where       string
	order       string
	limit       int
	offset      int
	options     string
}

func NewSearch(search string, fromIndexes ...string) Search {
	return Search{
		search:      search,
		fromIndexes: fromIndexes,
	}
}

func (s *Search) Select(columns ...string) {
	s.columns = columns
}

func (s *Search) Where(where string) {
	s.where = where
}

func (s *Search) Order(order string) {
	s.order = order
}

func (s *Search) Limit(offset, limit int) {
	s.limit = limit
	s.offset = offset
}

func (s *Search) Options(options string) {
	s.options = options
}

func (s *Search) Fetch(ctx context.Context, result Parser) error {
	query := "SELECT " + strings.Join(s.columns, ", ") + " FROM " + strings.Join(s.fromIndexes, ", ")
	// where
	switch {
	case s.search != "" && s.where != "":
		query += " WHERE MATCH('" + s.search + "') AND " + s.where
	case s.search == "" && s.where != "":
		query += " WHERE " + s.where
	case s.search != "" && s.where == "":
		query += " WHERE MATCH('" + s.search + "') "
	}
	// order by
	if s.order != "" {
		query += " ORDER BY " + s.order
	}
	// limit
	if s.limit > 0 {
		query += " LIMIT " + strconv.Itoa(s.offset) + ", " + strconv.Itoa(s.limit)
	}
	// options
	if s.options != "" {
		query += " OPTION " + s.options
	} else {
		query += " OPTION ranker=proximity, cutoff=0, retry_count=0, retry_delay=0;"
	}
	fmt.Printf(query)
	return SearchCustom(ctx, result, query)
}
