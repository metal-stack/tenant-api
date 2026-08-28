package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"

	sprig "github.com/go-task/slim-sprig/v3"
	"github.com/metal-stack/tenant-api/go/tests/protoparser"

	_ "embed"
)

var (
	//go:embed go_mock_client.tpl
	mockClientTpl string
	//go:embed go_client.tpl
	clientTpl string
)

type api struct {
	Name     string
	Services []struct {
		Name     string
		FileName string
	}
	Path string
}

func main() {
	fmt.Println("generating clients")

	svcs, err := svcs("../proto")
	if err != nil {
		panic(err)
	}

	err = writeTemplate("../go/client/client.go", clientTpl, svcs)
	if err != nil {
		panic(err)
	}

	err = writeTemplate("../go/tests/mock_clients.go", mockClientTpl, svcs)
	if err != nil {
		panic(err)
	}

}

func svcs(root string) (map[string]api, error) {
	var (
		result = map[string]api{}
		walk   = func(root string) ([]string, error) {
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

	for _, filename := range files {
		fd, err := protoparser.Parse(filename)
		if err != nil {
			return nil, err
		}
		_, name, _ := strings.Cut(*fd.Package, "tenant.")
		name = strings.ReplaceAll(name, ".", "")

		a, ok := result[name]
		if !ok {
			a = api{
				Name: name,
				Path: path.Dir(strings.TrimPrefix(filename, root)),
			}
		}
		for _, serviceDesc := range fd.GetService() {
			a.Services = append(a.Services, struct {
				Name     string
				FileName string
			}{
				Name:     *serviceDesc.Name,
				FileName: path.Base(filename),
			})
		}
		result[name] = a
	}

	return result, nil
}

func writeTemplate(dest, text string, data any) error {
	t, err := template.New("").Funcs(sprig.FuncMap()).Parse(text)
	if err != nil {
		return fmt.Errorf("unable to parse template %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return fmt.Errorf("unable to execute template %w", err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("unable to format source \n%s\n %w", buf.String(), err)
	}

	fmt.Println("wrote " + dest)

	return os.WriteFile(dest, p, 0644)
}
