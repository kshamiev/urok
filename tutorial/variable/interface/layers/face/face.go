package face

type DB interface {
	Name()
	Get() string
	Select() []string
	Insert(s string)
	Update(id int64, s string)
}

type RD interface {
	Name()
	IsIndex(index string) bool
	Members(index string, value string) bool
	Set(index string, value interface{}) error
}
