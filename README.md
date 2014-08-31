## wutenv

Simple `development`, `production`, `staging` application environment detection by GO_ENV or GO_FLAVOR environment variables.

## Usage

~~~ go
// main.go
package main

import (
  "github.com/8protons/wutenv"
)

func main() {
  if wutenv.IsDev {
  	// development environment == GO_ENV=development or GO_ENV=dev
  }

  if wutenv.IsProd {
  	// production environment == GO_ENV=production or GO_ENV=prod
  }

  if wutenv.IsStaging {
  	// staging environment == GO_ENV=staging
  }
}

~~~

## Authors
* [Anton Sekatski](https://github.com/antonsekatski)