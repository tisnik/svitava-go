package params

import (
	"errors"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Cplx struct {
	Name    string  `toml:"name"`
	Type    string  `toml:"type"`
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

func LoadCplxParameters(filename string) (map[string]Cplx, error) {
	var parameters CplxParams
	asMap := map[string]Cplx{}

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return asMap, errors.New("Parameter file does not exist.")
	}
	if err != nil {
		return asMap, err
	}

	_, err = toml.DecodeFile(filename, &parameters)
	if err != nil {
		log.Fatal(err)
		return asMap, err
	}
	for _, parameter := range parameters.Parameters {
		asMap[parameter.Name] = parameter
	}
	return asMap, nil
}
