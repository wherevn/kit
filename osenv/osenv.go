package osenv

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadFile env configuration from file.
func LoadFile(fileNames ...string) error {
	err := godotenv.Load(fileNames...)
	if err != nil {
		return err
	}
	return nil
}

func GetInt(key string, defaultValues ...int) int {
	var val int
	if len(defaultValues) > 0 {
		val = defaultValues[0]
	}
	if envVal, isExist := getEnv(key); isExist {
		if v, err := strconv.Atoi(envVal); err == nil {
			val = v
		}
	}
	return val
}

func GetString(key string, defaultValues ...string) string {
	var val string
	if len(defaultValues) > 0 {
		val = defaultValues[0]
	}
	if envVal, isExist := getEnv(key); isExist {
		val = envVal
	}
	return val
}

func GetBool(key string, defaultValues ...bool) bool {
	var val bool
	if len(defaultValues) > 0 {
		val = defaultValues[0]
	}
	if envVal, isExist := getEnv(key); isExist {
		if v, err := strconv.ParseBool(envVal); err == nil {
			val = v
		}
	}
	return val
}

func getEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}
