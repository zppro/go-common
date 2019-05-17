package data

type QueryResult interface {
	Total() int64
	Rows() []Row
}
