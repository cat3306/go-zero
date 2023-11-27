package gen

type Generator interface {
	StartFromDDL(filename string, withCache, strict bool, database string) error
}
