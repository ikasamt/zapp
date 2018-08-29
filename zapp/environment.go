package zapp

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"
)

type Environment map[string]interface{}

func ReadEnvironments() map[string]Environment {

	fn := "./config/environments.yml"
	value, _ := ioutil.ReadFile(fn)

	ret1 := make(map[string]Environment)
	yaml.Unmarshal(value, &ret1)

	return ret1
}
