package params

import (
	"errors"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Cplx struct {
	Name    string  `toml:"name"`
	Cx0     float64 `toml:"cx0"`
	Cy0     float64 `toml:"cy0"`
	Maxiter uint    `toml:"maxiter"`
	Xmin    float64 `toml:"xmin"`
	Ymin    float64 `toml:"ymin"`
	Xmax    float64 `toml:"xmax"`
	Ymax    float64 `toml:"ymax"`
}

type CplxParams struct {
	Parameters []Cplx `toml:"complex_fractal"`
}

func LoadCplxParameters(filename string) ([]Cplx, error) {
	var parameters CplxParams

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return parameters.Parameters, errors.New("Parameter file does not exist.")
	}
	if err != nil {
		return parameters.Parameters, err
	}

	_, err = toml.DecodeFile(filename, &parameters)
	if err != nil {
		log.Fatal(err)
		return parameters.Parameters, err
	}
	return parameters.Parameters, nil
}
