package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ConfigFile is the file that stores configuration parameters for the server.
const ConfigFile = "/etc/uschess/.config"

// GetDatabaseConnectionString constructs the database connection string based on config
// information from ConfigFile.
func GetDatabaseConnectionString() (databaseType string, connectionString string) {
	content, err := ioutil.ReadFile(ConfigFile)
	CheckErr(err)

	lines := strings.Split(string(content), "\n")

	vars := make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, "=")
		vars[parts[0]] = parts[1]
	}
	databaseType = vars["DB_TYPE"]
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		vars["DB_USER"], vars["DB_PASS"], vars["DB_SERVER"], vars["DB_PORT"], vars["DB_NAME"])
	return
}
