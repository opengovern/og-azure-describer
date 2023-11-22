package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"regexp"
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

type ResourceType struct {
	ResourceName         string
	ResourceLabel        string
	ServiceName          string
	ListDescriber        string
	GetDescriber         string
	TerraformName        []string
	TerraformNameString  string `json:"-"`
	TerraformServiceName string
	FastDiscovery        bool
	SteampipeTable       string
	Model                string
}

func main() {
	rt := "../../../kaytu-deploy/kaytu/inventory-data/azure-resource-types.json"
	b, err := os.ReadFile(rt)
	if err != nil {
		panic(err)
	}
	var resourceTypes []ResourceType
	err = json.Unmarshal(b, &resourceTypes)
	if err != nil {
		panic(err)
	}

	flag.CommandLine.Init("gen", flag.ExitOnError)
	flag.Parse()

	tpl := template.New("types")
	_, err = tpl.Parse(`
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

func (r *{{ .Name }}) UnmarshalJSON(b []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(b, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for k, v := range rawMsg {
		switch k {
		case "description":
			wrapper := {{ .SourceType }}Describer.JSONAllFieldsMarshaller{
				Value: r.Description,
			}
			if err := json.Unmarshal(v, &wrapper); err != nil {
				return fmt.Errorf("unmarshalling type %T: %v", r, err)
			}
			var ok bool
			r.Description, ok = wrapper.Value.({{ .SourceType }}.{{ .Name }}Description)
			if !ok {
				return fmt.Errorf("unmarshalling type %T: %v", r, fmt.Errorf("expected type %T, got %T", r.Description, wrapper.Value))
			}
		case "metadata":
			if err := json.Unmarshal(v, &r.Metadata); err != nil {
				return fmt.Errorf("unmarshalling type %T: %v", r, err)
			}
		case "resource_job_id":		
			if err := json.Unmarshal(v, &r.ResourceJobID); err != nil {
				return fmt.Errorf("unmarshalling type %T: %v", r, err)
			}
		case "source_job_id":
			if err := json.Unmarshal(v, &r.SourceJobID); err != nil {
				return fmt.Errorf("unmarshalling type %T: %v", r, err)
			}
		case "resource_type":
			if err := json.Unmarshal(v, &r.ResourceType); err != nil {
				return fmt.Errorf("unmarshalling type %T: %v", r, err)
			}
		case "source_type":
			if err := json.Unmarshal(v, &r.SourceType); err != nil {
				return fmt.Errorf("unmarshalling type %T: %v", r, err)
			}
		case "id":
			if err := json.Unmarshal(v, &r.ID); err != nil {
				return fmt.Errorf("unmarshalling type %T: %v", r, err)
			}
		case "arn":
			if err := json.Unmarshal(v, &r.ARN); err != nil {
				return fmt.Errorf("unmarshalling type %T: %v", r, err)
			}
		case "source_id":
			if err := json.Unmarshal(v, &r.SourceID); err != nil {
				return fmt.Errorf("unmarshalling type %T: %v", r, err)
			}
		default:
		}
	}
	return nil
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
	Total essdk.SearchTotal       ` + "`json:\"total\"`" + `
	Hits  []{{ .Name }}Hit ` + "`json:\"hits\"`" + `
}

type {{ .Name }}SearchResponse struct {
	PitID string          ` + "`json:\"pit_id\"`" + `
	Hits  {{ .Name }}Hits ` + "`json:\"hits\"`" + `
}

type {{ .Name }}Paginator struct {
	paginator *essdk.BaseESPaginator
}

func (k Client) New{{ .Name }}Paginator(filters []essdk.BoolFilter, limit *int64) ({{ .Name }}Paginator, error) {
	paginator, err := essdk.NewPaginator(k.ES(), "{{ .Index }}", filters, limit)
	if err != nil {
		return {{ .Name }}Paginator{}, err
	}

	p := {{ .Name }}Paginator{
		paginator: paginator,
	}

	return p, nil
}

func (p {{ .Name }}Paginator) HasNext() bool {
	return !p.paginator.Done()
}

func (p {{ .Name }}Paginator) NextPage(ctx context.Context) ([]{{ .Name }}, error) {
	var response {{ .Name }}SearchResponse
	err := p.paginator.Search(ctx, &response)
	if err != nil {
		return nil, err
	}

	var values []{{ .Name }}
	for _, hit := range response.Hits.Hits {
		values = append(values, hit.Source)
	}

	hits := int64(len(response.Hits.Hits))
	if hits > 0 {
		p.paginator.UpdateState(hits, response.Hits.Hits[hits-1].Sort, response.PitID)
	} else {
		p.paginator.UpdateState(hits, nil, "")
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
	cfg := essdk.GetConfig(d.Connection)
	ke, err := essdk.NewClientCached(cfg, d.ConnectionCache, ctx)
	if err != nil {
		plugin.Logger(ctx).Error("List{{ .Name }} NewClientCached", "error", err)
		return nil, err
	}
	k := Client{Client: ke}

	sc, err := steampipesdk.NewSelfClientCached(ctx, d.ConnectionCache)
	if err != nil {
		plugin.Logger(ctx).Error("List{{ .Name }} NewSelfClientCached", "error", err)
		return nil, err
	}
	accountId, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.KaytuConfigKeyAccountID)
	if err != nil {
		plugin.Logger(ctx).Error("List{{ .Name }} GetConfigTableValueOrNil for KaytuConfigKeyAccountID", "error", err)
		return nil, err
	}
	encodedResourceCollectionFilters, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.KaytuConfigKeyResourceCollectionFilters)
	if err != nil {
		plugin.Logger(ctx).Error("List{{ .Name }} GetConfigTableValueOrNil for KaytuConfigKeyResourceCollectionFilters", "error", err)
		return nil, err
	}
	clientType, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.KaytuConfigKeyClientType)
	if err != nil {
		plugin.Logger(ctx).Error("List{{ .Name }} GetConfigTableValueOrNil for KaytuConfigKeyClientType", "error", err)
		return nil, err
	}

	paginator, err := k.New{{ .Name }}Paginator(essdk.BuildFilter(ctx, d.QueryContext, list{{ .Name }}Filters, "{{ .SourceType }}", accountId, encodedResourceCollectionFilters, clientType), d.QueryContext.Limit)
	if err != nil {
		plugin.Logger(ctx).Error("List{{ .Name }} New{{ .Name }}Paginator", "error", err)
		return nil, err
	}

	for paginator.HasNext() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("List{{ .Name }} paginator.NextPage", "error", err)
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
	cfg := essdk.GetConfig(d.Connection)
	ke, err := essdk.NewClientCached(cfg, d.ConnectionCache, ctx)
	if err != nil {
		return nil, err
	}
	k := Client{Client: ke}

	sc, err := steampipesdk.NewSelfClientCached(ctx, d.ConnectionCache)
	if err != nil {
		return nil, err
	}
	accountId, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.KaytuConfigKeyAccountID)
	if err != nil {
		return nil, err
	}
	encodedResourceCollectionFilters, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.KaytuConfigKeyResourceCollectionFilters)
	if err != nil {
		return nil, err
	}
	clientType, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.KaytuConfigKeyClientType)
	if err != nil {
		return nil, err
	}

	limit := int64(1)
	paginator, err := k.New{{ .Name }}Paginator(essdk.BuildFilter(ctx, d.QueryContext, get{{ .Name }}Filters, "{{ .SourceType }}", accountId, encodedResourceCollectionFilters, clientType), &limit)
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
	fmt.Fprintf(&buf, "package kaytu")

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
			for _, resourceType := range resourceTypes {
				if resourceType.Model == s.Name {
					var stopWordsRe = regexp.MustCompile(`\W+`)
					index := stopWordsRe.ReplaceAllString(resourceType.ResourceName, "_")
					index = strings.ToLower(index)
					s.Index = index

					plugin := "steampipe-plugin-azure/azure"
					if strings.HasPrefix(resourceType.SteampipeTable, "azuread") {
						plugin = "steampipe-plugin-azuread/azuread"
					}
					fileName := "../../" + plugin + "/table_" + resourceType.SteampipeTable + ".go"
					tableFileSet := token.NewFileSet()
					tableNode, err := parser.ParseFile(tableFileSet, fileName, nil, parser.ParseComments)
					if err != nil {
						panic(err)
					}

					ast.Inspect(tableNode, func(tnode ast.Node) bool {
						if c, ok := tnode.(*ast.CompositeLit); ok {

							var columnName, transformer string
							for _, arg := range c.Elts {
								if kv, ok := arg.(*ast.KeyValueExpr); ok {
									if i, ok := kv.Key.(*ast.Ident); ok {
										if i.Name == "Name" {
											if bl, ok := kv.Value.(*ast.BasicLit); ok {
												columnName = strings.Trim(bl.Value, "\"")
											}
										} else if i.Name == "Transform" {
											if cl, ok := kv.Value.(*ast.CallExpr); ok {
												transformer = extractTransformer(cl)
											}
										}
									}
								}
							}

							if columnName != "" && transformer != "" {
								if strings.HasPrefix(transformer, "Description") || strings.HasPrefix(transformer, "Metadata") {
									transformer = strings.ToLower(transformer[:1]) + transformer[1:]
								}
								s.GetFilters[columnName] = transformer
								s.ListFilters[columnName] = transformer
							}
							return true
						}
						return true
					})
				}
			}

			if decl.Doc != nil {
				for _, c := range decl.Doc.List {
					if strings.HasPrefix(c.Text, "//index:") {
						//s.Index = strings.TrimSpace(strings.TrimPrefix(c.Text, "//index:"))
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
				s.GetFilters["kaytu_account_id"] = "metadata.SourceID"
				s.ListFilters["kaytu_account_id"] = "metadata.SourceID"
			}

			if s.Index != "" {
				sources = append(sources, s)
			} else {
				fmt.Println("failed to find the index:", s.Name)
			}
		}
		return false
	})

	if len(sources) > 0 {
		fmt.Fprintln(&buf, `
		import (
			"context"
			"encoding/json"
			"fmt"
			essdk "github.com/kaytu-io/kaytu-util/pkg/kaytu-es-sdk"
			steampipesdk "github.com/kaytu-io/kaytu-util/pkg/steampipe"
			"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
			`+*sourceType+`Describer "github.com/kaytu-io/kaytu-`+*sourceType+`-describer/`+*sourceType+`/describer"
			`+*sourceType+` "github.com/kaytu-io/kaytu-`+*sourceType+`-describer/`+*sourceType+`/model"
		)

		type Client struct {
			essdk.Client
		}

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

func extractTransformer(cl *ast.CallExpr) string {
	if sl, ok := cl.Fun.(*ast.SelectorExpr); ok {
		if sl.Sel.Name == "Transform" {
			return ""
		}
		if call, ok := sl.X.(*ast.CallExpr); ok {
			return extractTransformer(call)
		}
		if sl.Sel.Name == "FromField" {
			for _, arg := range cl.Args {
				if bl, ok := arg.(*ast.BasicLit); ok {
					return strings.Trim(bl.Value, "\"")
				}
			}
		}
	}
	return ""
}
