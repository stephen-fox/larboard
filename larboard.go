package larboard

type Researcher interface {
	IsHalo2() error
	IsMap() error
	Name() (string, error)
	Scenario() (string, error)
	Signature() (string, error)
}