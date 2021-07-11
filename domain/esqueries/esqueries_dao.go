package esqueries

import "github.com/olivere/elastic/v7"

func (q *EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	equals := make([]elastic.Query, 0)
	for _, eq := range q.Equals {
		equals = append(equals, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	query.Must(equals...)
	return query
}
