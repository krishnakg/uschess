package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const CONFIG_FILE = "/etc/uschess/.config"

func getDatabaseConnectionString() (databaseType string, connectionString string) {
	content, err := ioutil.ReadFile(CONFIG_FILE)
	checkErr(err)

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
