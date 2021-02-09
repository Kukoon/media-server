package runtime

import (
	"os"
	"time"

	"github.com/naoina/toml"
)

// TOMLDuration a time.Duration inside toml files
type TOMLDuration time.Duration

// UnmarshalText implements encoding.TextUnmarshaler
func (d *TOMLDuration) UnmarshalText(data []byte) error {
	duration, err := time.ParseDuration(string(data))
	if err == nil {
		*d = TOMLDuration(duration)
	}
	return err
}

// MarshalText implements encoding.TextMarshaler
func (d TOMLDuration) MarshalText() ([]byte, error) {
	return []byte(time.Duration(d).String()), nil
}

// ReadTOML reads a config model from path of a toml file
func ReadTOML(file string, data interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewDecoder(f).Decode(data)
}

// SaveTOML to path
func SaveTOML(outputFile string, data interface{}) error {
	tmpFile := outputFile + ".tmp"

	file, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	err = toml.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}

	file.Close()
	return os.Rename(tmpFile, outputFile)
}
