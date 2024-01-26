package main

import "google.golang.org/protobuf/compiler/protogen"

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + ".gin.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	return g
}
