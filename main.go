package main

import (
	cmdproto "github.com/domgolonka/ent2proto/cmd"
)


func main() {
	cmdproto.Execute()
	//// Get GOPATH var.
	//goPath := os.Getenv("GOPATH")
	//if goPath == "" {
	//	goPath = "~/go"
	//}
	//
	//// Define flags.
	//var src, dst string
	//flag.StringVar(&src, "src", filepath.Join(goPath, "src/github.com/osquery/osquery/specs"), "source path")
	//flag.StringVar(&dst, "dst", filepath.Join(goPath, "src/github.com/mickep76/osquery-protobuf"), "destination path")
	//flag.Parse()
	//
	//// Expand home dir.
	//src, _ = homedir.Expand(src)
	//dst, _ = homedir.Expand(dst)
	//
	//// Compile template's.
	////t := template.Must(template.ParseGlob(filepath.Join("*.tmpl")))
	//
	//
	////// Parse database schemas and template protobuf.
	////s := internal.Service{Tables: []*internal.Table{}}
	////if err := s.Parse(src, dst, t, 0); err != nil {
	////	logrus.Fatal(err)
	////}
	//opts := []Option{}
	//cfg := gen.Config{
	//	Target: dst,
	//}
	//idtype    := idType(field.TypeInt)
	//if cfg.Target != "" {
	//	pkgPath, err := pkg.PkgPath(pkg.DefaultConfig, cfg.Target)
	//	if err != nil {
	//		logrus.Fatalln(err)
	//	}
	//	cfg.Package = pkgPath
	//}
	//cfg.IDType = &field.TypeInfo{Type: field.Type(idtype)}
	//
	//err := Generate(src,&cfg , opts...)
	//if (err !=nil) {
	//	logrus.Error(err)
	//}

}