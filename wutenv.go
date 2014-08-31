package wutenv

import (
	"os"
)

var Env = "development"

var IsDev = false
var IsProd = false
var IsStaging = false

func init() {
	for _, v := range []string{"GO_ENV", "GO_FLAVOR"} {
		if val := os.Getenv(v); val != "" {
			Env = val
			break
		}
	}

	if Env == "dev" || Env == "development" {
		IsDev = true

		return
	}

	if Env == "prod" || Env == "production" {
		IsProd = true

		return
	}

	if Env == "staging" {
		IsStaging = true

		return
	}
}
