package core

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"text/template"
)

func Gen(plugin *protogen.Plugin, f *protogen.File) *protogen.GeneratedFile {
	err := genMessage(plugin, f)
	if err != nil {
		panic(err)
	}
	return nil
}

func genMessage(plugin *protogen.Plugin, f *protogen.File) error {
	w := plugin.NewGeneratedFile(
		fmt.Sprintf("%s.pb.protojson.go", f.GeneratedFilenamePrefix),
		f.GoImportPath,
	)

	err := template.Must(
		template.New("importTpl").Parse(importTpl),
	).Execute(w, f)

	if err != nil {
		return err
	}

	err = template.Must(
		template.New("initTp;").Parse(initTpl),
	).Execute(w, f)

	if err != nil {
		return err
	}

	msgTpl := template.Must(
		template.New("codecTpl").Parse(codecTpl),
	)

	for _, msg := range f.Messages {
		if msg.Desc.IsMapEntry() {
			continue
		}

		err = msgTpl.Execute(w, struct {
			*protogen.Message
			*protogen.File
		}{msg, f})

		if err != nil {
			return err
		}
	}

	return nil
}
