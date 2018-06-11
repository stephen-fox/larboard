package larboard

type Cartographer interface {
	SetMap(haloMap HaloMap) error
	IsHalo2() error
	IsMap() error
	Name() (string, error)
	Scenario() (string, error)
	Signature() (string, error)
	Sign() (string, error)
}

type HaloMap struct {
	FilePath string `json:"file_path"`
}