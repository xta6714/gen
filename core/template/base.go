package template

const NotEditMark = `
// Code generated by github.com/xta6714/gen. DO NOT EDIT.
// Code generated by github.com/xta6714/gen. DO NOT EDIT.
// Code generated by github.com/xta6714/gen. DO NOT EDIT.
`

const Header = NotEditMark + `
package {{.Package}}

import(	
	{{range .ImportPkgPaths}}{{.}}` + "\n" + `{{end}}
)
`