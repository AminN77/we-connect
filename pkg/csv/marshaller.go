package csv

type Marshaller interface {
	Unmarshal(record []string) error
}
