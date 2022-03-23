package env

import (
	`os`
)

func hostname() (env string) {
	if name, err := os.Hostname(); nil == err {
		env = name
	} else {
		env = defaultEnvErrorValue
	}

	return
}
