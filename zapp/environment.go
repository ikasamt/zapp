package zapp

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
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

func AdminBasicauth(env *Environment) gin.HandlerFunc {
	adminBasicauthName := (*env)[`admin_basicauth_name`].(string)
	adminBasicauthPassword := (*env)[`admin_basicauth_password`].(string)
	adminAccount := gin.Accounts{adminBasicauthName: adminBasicauthPassword}
	return gin.BasicAuth(adminAccount)
}
