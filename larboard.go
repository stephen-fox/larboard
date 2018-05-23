package larboard

type Researcher interface {
	IsHalo2() error
	IsMap() error
	Scenario() (string, error)
}