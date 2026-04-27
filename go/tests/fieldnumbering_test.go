package apitests

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/metal-stack/tenant-api/go/tests/protoparser"
	"github.com/stretchr/testify/require"
)

func Test_FieldNumbering(t *testing.T) {
	files, err := getProtos("../../proto")
	require.NoError(t, err)
	var errs []error

	for _, filename := range files {
		fd, err := protoparser.Parse(filename)
		require.NoError(t, err)

		for _, mt := range fd.GetMessageType() {
			var (
				firstField = true
				lastNumber int32
			)

			for _, field := range mt.GetField() {
				if field.Number != nil {
					if firstField {
						if *field.Number != 1 {
							errs = append(errs, fmt.Errorf("%s %s %s %d != %d", filename, *mt.Name, *field.Name, 1, *field.Number))
						}
						firstField = false
					} else {
						if lastNumber+1 != *field.Number {
							errs = append(errs, fmt.Errorf("%s %s %s %d != %d", filename, *mt.Name, *field.Name, lastNumber+1, *field.Number))
						}
					}
					lastNumber = *field.Number
				}
			}
		}

		for _, et := range fd.GetEnumType() {
			var (
				firstField = true
				lastNumber int32
			)

			for _, value := range et.GetValue() {
				if value.Number != nil {
					if firstField {
						if *value.Number != 0 {
							errs = append(errs, fmt.Errorf("%s %s %s %d != %d", filename, *et.Name, *value.Name, 0, *value.Number))
						}
						firstField = false
					} else {
						if lastNumber+1 != *value.Number {
							errs = append(errs, fmt.Errorf("%s %s %s %d != %d", filename, *et.Name, *value.Name, lastNumber+1, *value.Number))
						}
					}
					lastNumber = *value.Number
				}
			}
		}
	}
	require.NoError(t, errors.Join(errs...))
}

func getProtos(root string) ([]string, error) {
	var (
		walk = func(root string) ([]string, error) {
			var files []string
			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if strings.HasSuffix(info.Name(), ".proto") {
					files = append(files, path)
				}
				return nil
			})
			return files, err
		}
	)

	files, err := walk(root)
	if err != nil {
		return nil, err
	}
	return files, nil
}
