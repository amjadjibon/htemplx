package main

import "htemplx/cmd"

//	@title			Swagger Example API
//	@version		1.0
//	@description	HTEMPLX API SERVER
//	@termsOfService	http://localhost:8080/terms-and-conditions/

//	@contact.name	API Support
//	@contact.url	http://localhost:8080/support
//	@contact.email	support@htemplx.io

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	cmd.Execute()
}
