package gen

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"text/template"
)

func Gen(plugin *protogen.Plugin, f *protogen.File) *protogen.GeneratedFile {
	err := genCodec(plugin, f)
	if err != nil {
		panic(err)
	}
	err = genMessage(plugin, f)
	if err != nil {
		panic(err)
	}
	return nil
}

func genCodec(plugin *protogen.Plugin, f *protogen.File) error {
	w := plugin.NewGeneratedFile(
		fmt.Sprintf("%s.pb_codec.go", f.GoPackageName),
		f.GoImportPath,
	)
	return template.Must(
		template.New("initTpl").Parse(initTpl),
	).Execute(w, f)
}

func genMessage(plugin *protogen.Plugin, f *protogen.File) error {
	w := plugin.NewGeneratedFile(
		fmt.Sprintf("%s.pb_json.go", f.GoPackageName),
		f.GoImportPath,
	)
	err := template.Must(
		template.New("importTpl").Parse(importTpl),
	).Execute(w, f)

	if err != nil {
		return err
	}

	tpl := template.Must(
		template.New("messageTpl").Parse(messageTpl),
	)

	for _, msg := range f.Messages {
		if msg.Desc.IsMapEntry() {
			continue
		}

		err = tpl.Execute(w, struct {
			*protogen.Message
			*protogen.File
		}{msg, f})

		if err != nil {
			return err
		}
	}

	return nil
}
