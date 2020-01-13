// +build ignore

package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	err := buildTemplatesFile()
	if err != nil {
		log.Fatal(err)
	}
}

func buildTemplatesFile() error {
	f := &ast.File{
		Name: ast.NewIdent("main"),
	}

	f.Decls = []ast.Decl{
		&ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						ast.NewIdent("rawTemplates"),
					},
					Values: []ast.Expr{
						getTemplatesArray(),
					}},
			},
		},
	}

	fset := token.NewFileSet() // positions are relative to fset

	rawdata, err := os.Create("raw_templates.go")
	if err != nil {
		return err
	}

	err = printer.Fprint(rawdata, fset, f)
	if err != nil {
		return err
	}

	return nil
}

func getTemplatesArray() *ast.CompositeLit {
	templatesArray := &ast.CompositeLit{
		Type: &ast.MapType{
			Key:   ast.NewIdent("string"),
			Value: ast.NewIdent("string"),
		},
		Elts: []ast.Expr{},
	}

	for name, template := range getTemplates() {
		templatesArray.Elts = append(templatesArray.Elts, &ast.KeyValueExpr{
			Key: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(name),
			},
			Value: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(template),
			},
		})
	}

	return templatesArray
}

func getTemplates() map[string]string {
	templateFiles := []string{
		"templates/bin/console.tmpl",
		"templates/bin/setup.tmpl",
		"templates/bin/wp.tmpl",
		"templates/config/server.php.tmpl",
		"templates/docker-compose.yml.tmpl",
		"templates/setup/external.sh.tmpl",
		"templates/setup/internal.sh.tmpl",
		"templates/script/bootstrap.tmpl",
		"templates/script/console.tmpl",
		"templates/script/server.tmpl",
		"templates/script/setup.tmpl",
		"templates/script/update.tmpl",
	}
	mapping := map[string]string{}

	for _, filePath := range templateFiles {
		mapping[filePath] = mustReadFile(filePath)
	}

	return mapping
}

func mustReadFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
