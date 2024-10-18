//
//  (C) Copyright 2019, 2020, 2021, 2022, 2023, 2024  Pavel Tisnovsky
//
//  All rights reserved. This program and the accompanying materials
//  are made available under the terms of the Eclipse Public License v1.0
//  which accompanies this distribution, and is available at
//  http://www.eclipse.org/legal/epl-v10.html
//
//  Contributors:
//      Pavel Tisnovsky
//

package configuration

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Configuration struct{}

func LoadConfiguration(configFileName string) error {
	var configuration Configuration
	blob := ``
	_, err := toml.Decode(blob, &configuration)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
