package yaml

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	Foo struct {
		Bar string `yaml:"bar"`
	} `yaml:"foo"`
}

func TestYaml_Load(t *testing.T) {
	assert := assert.New(t)

	t.Run("when the file does not exist", func(t *testing.T) {
		filePath := "test_data/non_existent.yml"
		v := new(testData)

		expectedErr := errors.New("yaml_load: error reading file. err: open test_data/non_existent.yml: no such file or directory")

		t.Run("it returns an error", func(t *testing.T) {
			err := Load(filePath, v)

			assert.Equal(expectedErr, err)
		})
	})

	t.Run("when fails unmarshalling data", func(t *testing.T) {
		filePath := "test_data/data_invalid.yml"
		v := new(testData)

		expectedErr := errors.New("yaml_load: error unmarshalling YAML data into *yaml.testData. err: yaml: line 1: mapping values are not allowed in this context")

		t.Run("it returns an error", func(t *testing.T) {
			err := Load(filePath, v)

			assert.Equal(expectedErr, err)
		})
	})

	t.Run("when the loading is succesful", func(t *testing.T) {
		filePath := "test_data/data.yml"
		v := new(testData)

		t.Run("it returns a nil error", func(t *testing.T) {
			err := Load(filePath, v)

			assert.NoError(err)
			assert.NotEmpty(v.Foo.Bar)
		})
	})
}
