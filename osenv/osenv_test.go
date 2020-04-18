package osenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestDataFileTestConf = "testdata/test_conf.env"
)

func TestLoadFile(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		err := LoadFile(TestDataFileTestConf)
		assert.NoError(t, err)
		assert.Equal(t, "CONF_1", os.Getenv("CONF_1"))
		assert.Equal(t, "CONF_2", os.Getenv("CONF_2"))
	})

	t.Run("FAIL_TO_LOAD_CONF_FILE", func(t *testing.T) {
		err := LoadFile("testdata/file_doesnt_exist.env")
		assert.Error(t, err)
	})
}

func TestGetInt(t *testing.T) {
	const testIntEnvKey = "TEST_INT_ENV"

	err := os.Setenv(testIntEnvKey, "123")
	assert.NoError(t, err)

	t.Run("OK", func(t *testing.T) {
		confVal := GetInt(testIntEnvKey)
		assert.Equal(t, 123, confVal)
	})

	t.Run("GET_DEFAULT_CONF", func(t *testing.T) {
		confVal := GetInt("DOESNT_EXIST_CONF", 333)
		assert.Equal(t, 333, confVal)
	})
}

func TestGetString(t *testing.T) {
	const testStringEnvKey = "TEST_STRING_ENV"

	err := os.Setenv(testStringEnvKey, "string")
	assert.NoError(t, err)

	t.Run("OK", func(t *testing.T) {
		confVal := GetString(testStringEnvKey)
		assert.Equal(t, "string", confVal)
	})

	t.Run("GET_DEFAULT_CONF", func(t *testing.T) {
		confVal := GetString("DOESNT_EXIST_CONF", "string")
		assert.Equal(t, "string", confVal)
	})
}

func TestGetBool(t *testing.T) {
	const testBoolEnvKey = "TEST_BOOL_ENV"

	err := os.Setenv(testBoolEnvKey, "true")
	assert.NoError(t, err)

	t.Run("OK", func(t *testing.T) {
		confVal := GetBool(testBoolEnvKey)
		assert.Equal(t, true, confVal)
	})

	t.Run("GET_DEFAULT_CONF", func(t *testing.T) {
		confVal := GetBool("DOESNT_EXIST_CONF", true)
		assert.Equal(t, true, confVal)
	})
}
