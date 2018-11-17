package conf

import (
	"os"
)

func getEnv(variableName, defaultVariable string) string {
	var result = os.Getenv(variableName)
	if result != "" {
		return result
	} else {
		return defaultVariable
	}
}

var ModelFolder = getEnv("MODEL_FOLDER", "../test")
