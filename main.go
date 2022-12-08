package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/aireet/protoc-gen-go-proto-json/gen"
)

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range plugin.Files {

			if len(f.Messages) <= 0 {
				continue
			}

			gen.Gen(plugin, f)
		}
		return nil
	})
}
