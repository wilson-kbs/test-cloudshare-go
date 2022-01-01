package web

import (
	"bytes"
	_ "embed"
	"html/template"
	"log"
)

var (
	//go:embed tmpl/index.tmpl
	indexTemplateFile []byte

	indexFile *bytes.Buffer
)

type Config struct {
	Base string
}

func ParseIndexFile(config Config) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	if config.Base == "/" {
		config.Base = ""
	}

	tmpl, err := template.New("app").Parse(string(indexTemplateFile))
	check(err)

	indexFile = bytes.NewBuffer([]byte{})

	err = tmpl.Execute(indexFile, config)
	check(err)

}

func GetIndex() []byte {
	if indexFile == nil {
		return []byte{}
	}

	return indexFile.Bytes()
}
