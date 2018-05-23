package larboard

type MapDescriber interface {
	IsHalo2() (bool, error)
	IsMap() (bool, error)
}