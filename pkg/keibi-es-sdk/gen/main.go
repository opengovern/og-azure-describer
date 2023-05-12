package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"strings"
)

var (
	file       = flag.String("file", "", "Location of the model file")
	sourceType = flag.String("type", "", "Type of resource clients (e.g. aws, azure). Should match the model import path")
	output     = flag.String("output", "", "Location of the output file")
)

type SourceType struct {
	Name        string
	Index       string
	ListFilters map[string]string
	GetFilters  map[string]string
	SourceType  string
}

func main() {
	flag.CommandLine.Init("gen", flag.ExitOnError)
	flag.Parse()

	tpl := template.New("types")
	_, err := tpl.Parse(`
// ==========================  START: {{ .Name }} =============================

type {{ .Name }} struct {
	Description   {{ .SourceType }}.{{ .Name }}Description 	` + "`json:\"description\"`" + `
	Metadata      {{ .SourceType }}.Metadata 					` + "`json:\"metadata\"`" + `
	ResourceJobID int ` + "`json:\"resource_job_id\"`" + `
	SourceJobID   int ` + "`json:\"source_job_id\"`" + `
	ResourceType  string ` + "`json:\"resource_type\"`" + `
	SourceType    string ` + "`json:\"source_type\"`" + `
	ID            string ` + "`json:\"id\"`" + `
	ARN            string ` + "`json:\"arn\"`" + `
	SourceID      string ` + "`json:\"source_id\"`" + `
}

type {{ .Name }}Hit struct {
	ID      string            ` + "`json:\"_id\"`" + `
	Score   float64           ` + "`json:\"_score\"`" + `
	Index   string            ` + "`json:\"_index\"`" + `
	Type    string            ` + "`json:\"_type\"`" + `
	Version int64             ` + "`json:\"_version,omitempty\"`" + `
	Source  {{ .Name }}       ` + "`json:\"_source\"`" + `
	Sort    []interface{}     ` + "`json:\"sort\"`" + `
}

type {{ .Name }}Hits struct {
	Total SearchTotal       ` + "`json:\"total\"`" + `
	Hits  []{{ .Name }}Hit ` + "`json:\"hits\"`" + `
}

type {{ .Name }}SearchResponse struct {
	PitID string          ` + "`json:\"pit_id\"`" + `
	Hits  {{ .Name }}Hits ` + "`json:\"hits\"`" + `
}

type {{ .Name }}Paginator struct {
	paginator *baseESPaginator
}

func (k Client) New{{ .Name }}Paginator(filters []BoolFilter, limit *int64) ({{ .Name }}Paginator, error) {
	paginator, err := newPaginator(k.es, "{{ .Index }}", filters, limit)
	if err != nil {
		return {{ .Name }}Paginator{}, err
	}

	p := {{ .Name }}Paginator{
		paginator: paginator,
	}

	return p, nil
}

func (p {{ .Name }}Paginator) HasNext() bool {
	return !p.paginator.done
}

func (p {{ .Name }}Paginator) NextPage(ctx context.Context) ([]{{ .Name }}, error) {
	var response {{ .Name }}SearchResponse
	err := p.paginator.search(ctx, &response)
	if err != nil {
		return nil, err
	}

	var values []{{ .Name }}
	for _, hit := range response.Hits.Hits {
		values = append(values, hit.Source)
	}

	hits := int64(len(response.Hits.Hits))
	if hits > 0 {
		p.paginator.updateState(hits, response.Hits.Hits[hits-1].Sort, response.PitID)
	} else {
		p.paginator.updateState(hits, nil, "")
	}

	return values, nil
}

var list{{ .Name }}Filters = map[string]string{
	{{ range $key, $value := .ListFilters }}"{{ $key }}": "{{ $value }}",
	{{ end }}
}

func List{{ .Name }}(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("List{{ .Name }}")

	// create service
	cfg := GetConfig(d.Connection)
	k, err := NewClientCached(cfg, d.ConnectionManager.Cache, ctx)
	if err != nil {
		return nil, err
	}

	paginator, err := k.New{{ .Name }}Paginator(buildFilter(d.KeyColumnQuals, list{{ .Name }}Filters, "{{ .SourceType }}", *cfg.AccountID), d.QueryContext.Limit)
	if err != nil {
		return nil, err
	}

	for paginator.HasNext() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page {
			d.StreamListItem(ctx, v)
		}
	}

	return nil, nil
}


var get{{ .Name }}Filters = map[string]string{
	{{ range $key, $value := .GetFilters }}"{{ $key }}": "{{ $value }}",
	{{ end }}
}

func Get{{ .Name }}(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("Get{{ .Name }}")

	// create service
	cfg := GetConfig(d.Connection)
	k, err := NewClientCached(cfg, d.ConnectionManager.Cache, ctx)
	if err != nil {
		return nil, err
	}

	limit := int64(1)
	paginator, err := k.New{{ .Name }}Paginator(buildFilter(d.KeyColumnQuals, get{{ .Name }}Filters, "{{ .SourceType }}", *cfg.AccountID), &limit)
	if err != nil {
		return nil, err
	}

	for paginator.HasNext() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page {
			return v, nil
		}
	}

	return nil, nil
}

// ==========================  END: {{ .Name }} =============================

`)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, *file, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(&buf, "// Code is generated by go generate. DO NOT EDIT.")
	fmt.Fprintf(&buf, "package keibi")

	// fmt.Fprintln(&buf, "import (")
	// for _, pkg := range []string{
	// 	"context",
	// 	"github.com/turbot/steampipe-plugin-sdk/plugin",
	// 	"github.com/kaytu-io/kaytu-aws-describer/aws/model",
	// 	"github.com/kaytu-io/kaytu-azure-describer/azure/model",
	// } {
	// 	fmt.Fprintf(&buf, "\"%s\"\n", pkg)
	// }
	// fmt.Fprintln(&buf, ")")

	var sources []SourceType

	ast.Inspect(node, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			return true
		}

		for _, spec := range decl.Specs {
			t := spec.(*ast.TypeSpec)

			if !strings.HasSuffix(t.Name.String(), "Description") {
				fmt.Println("Ignoring type " + t.Name.String())
				continue
			}

			s := SourceType{
				Name:        strings.TrimSuffix(t.Name.String(), "Description"),
				SourceType:  *sourceType,
				GetFilters:  map[string]string{},
				ListFilters: map[string]string{},
			}

			if decl.Doc != nil {
				for _, c := range decl.Doc.List {
					if strings.HasPrefix(c.Text, "//index:") {
						s.Index = strings.TrimSpace(strings.TrimPrefix(c.Text, "//index:"))
					} else if strings.HasPrefix(c.Text, "//getfilter:") {
						f := strings.TrimSpace(strings.TrimPrefix(c.Text, "//getfilter:"))
						fparts := strings.Split(f, "=")
						s.GetFilters[fparts[0]] = fparts[1]
					} else if strings.HasPrefix(c.Text, "//listfilter:") {
						f := strings.TrimSpace(strings.TrimPrefix(c.Text, "//listfilter:"))
						fparts := strings.Split(f, "=")
						s.ListFilters[fparts[0]] = fparts[1]
					}
				}
				s.GetFilters["keibi_account_id"] = "metadata.SourceID"
				s.ListFilters["keibi_account_id"] = "metadata.SourceID"
			}

			if s.Index != "" {
				sources = append(sources, s)
			}
		}

		return false
	})

	if len(sources) > 0 {
		fmt.Fprintln(&buf, `
		import (
			"context"
			"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
			`+*sourceType+` "github.com/kaytu-io/kaytu-`+*sourceType+`-describer/`+*sourceType+`/model"
		)
		`)
	}

	for _, source := range sources {
		err := tpl.Execute(&buf, source)
		if err != nil {
			log.Fatal(err)
		}
	}

	source, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	_, err = out.Write(source)
	if err != nil {
		log.Fatal(err)
	}
}
