package server

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)

type Path struct {
	Method map[string]Method `yaml:",inline"`
}
type Method struct {
	Description string                `yaml:"description"`
	Consumes    []string              `yaml:"consumes"`
	Produces    []string              `yaml:"produces"`
	Security    []map[string][]string `yaml:"security"`
	Responses   map[string]Response   `yaml:"responses"`
	Parameters  []Parameter           `yaml:"parameters"`
	Tags        []string              `yaml:"tags"`
}
type Response struct {
	Description string            `yaml:"description"`
	Schema      map[string]string `yaml:"schema"`
}
type Parameter struct {
	Name        string            `yaml:"name"`
	In          string            `yaml:"in"`
	Description string            `yaml:"description"`
	Required    bool              `yaml:"required"`
	Schema      map[string]string `yaml:"schema"`
	Type        string            `yaml:"type"`
}
type Info struct {
	Version string `yaml:"version"`
	Title   string `yaml:"title"`
}
type Swagger struct {
	Swagger             string                 `yaml:"swagger"`
	Info                Info                   `yaml:"info"`
	Paths               map[string]Path        `yaml:"paths"`
	Definitions         map[string]Definition  `yaml:"definitions"`
	SecurityDefinitions map[string]SecurityDef `yaml:"securityDefinitions"`
}

type SecurityDef struct {
	Type string `yaml:"type"`
	In   string `yaml:"in"`
	Name string `yaml:"name"`
}
type Definition struct {
	Description string              `yaml:"description"`
	Properties  map[string]Property `yaml:"properties"`
}
type Property struct {
	Type        string `yaml:"type"`
	Format      string `yaml:"format,omitempty"`
	Description string `yaml:"description"`
}

func generateSwaggerFile() {
	var swagger Swagger
	swagger = Swagger{
		Swagger: "2.0",
		Info: Info{
			Version: "1.0.0",
			Title:   "Interface CSA",
		},
	}
	var paths map[string]Path
	var definitions map[string]Definition
	paths = make(map[string]Path, len(Subscriptions))
	definitions = make(map[string]Definition, len(Subscriptions))
	for p, s := range Subscriptions {
		m := path.Base(p)
		m = strings.ToLower(m)
		p = path.Dir(p)
		elems := strings.Split(p, string(filepath.Separator))
		tag := path.Base(elems[2])
		switch m {
		case "post":
			method := map[string]Method{
				m: Method{
					Description: s.Description,
					Consumes:    []string{"application/json"},
					Produces:    []string{"application/json"},
					Security: []map[string][]string{
						{
							"csa": []string{},
						},
					},
					Parameters: []Parameter{
						{
							Name:        "The payload",
							In:          "body",
							Description: "The payload",
							Required:    true,
							Schema: map[string]string{
								"$ref": fmt.Sprintf("#/definitions/%v", s.ID),
							},
						},
					},
					Tags: []string{tag},
					Responses: map[string]Response{
						"202": Response{
							Description: "L'identifiant de la demande",
							Schema: map[string]string{
								"$ref": "#/definitions/response",
							},
						},
						"400": Response{
							Description: "Request Malformed or mandatory parameter not present",
						},
						"500": Response{
							Description: "Unhandled error",
						},
						"502": Response{
							Description: "CSA backend unavailable",
						},
					},
				},
			}
			paths[p] = Path{
				Method: method,
			}
			properties := make(map[string]Property, len(s.Fields))
			for _, field := range s.Fields {
				if field.Needed {
					properties[field.Name] = Property{
						Type:        reflect.ValueOf(field.Value).Kind().String(),
						Description: field.DisplayName,
					}
				}
			}
			definitions[s.ID] = Definition{
				Description: s.Description,
				Properties:  properties,
			}
		case "delete":
			method := map[string]Method{
				m: Method{
					Description: s.Description,
					Consumes:    []string{"application/json"},
					Produces:    []string{"application/json"},
					Parameters: []Parameter{
						{
							Name:        "id",
							In:          "path",
							Description: "ID d'execution",
							Required:    true,
							Type:        "string",
						},
					},
					Tags: []string{tag},
					Responses: map[string]Response{
						"202": Response{
							Description: "L'identifiant de la demande",
							Schema: map[string]string{
								"$ref": "#/definitions/response",
							},
						},
						"400": Response{
							Description: "Request Malformed or mandatory parameter not present",
						},
						"500": Response{
							Description: "Unhandled error",
						},
						"502": Response{
							Description: "CSA backend unavailable",
						},
					},
				},
			}
			paths[fmt.Sprintf("%v/{id}", p)] = Path{
				Method: method,
			}
			properties := make(map[string]Property, len(s.Fields))
			for _, field := range s.Fields {
				if field.Needed {
					properties[field.Name] = Property{
						Type:        reflect.ValueOf(field.Value).Kind().String(),
						Description: field.DisplayName,
					}
				}
			}
			definitions[s.ID] = Definition{
				Description: s.Description,
				Properties:  properties,
			}
		case "get":
			method := map[string]Method{
				m: Method{
					Description: s.Description,
					Consumes:    []string{"application/json"},
					Produces:    []string{"application/json"},
					Parameters: []Parameter{
						{
							Name:        "id",
							In:          "path",
							Description: "ID d'execution",
							Required:    true,
							Type:        "string",
						},
					},
					Tags: []string{tag},
					Responses: map[string]Response{
						"202": Response{
							Description: "L'identifiant de la demande",
							Schema: map[string]string{
								"$ref": "#/definitions/response",
							},
						},
						"400": Response{
							Description: "Request Malformed or mandatory parameter not present",
						},
						"500": Response{
							Description: "Unhandled error",
						},
						"502": Response{
							Description: "CSA backend unavailable",
						},
					},
				},
			}
			paths[fmt.Sprintf("%v/{id}", p)] = Path{
				Method: method,
			}
			properties := make(map[string]Property, len(s.Fields))
			for _, field := range s.Fields {
				if field.Needed {
					properties[field.Name] = Property{
						Type:        reflect.ValueOf(field.Value).Kind().String(),
						Description: field.DisplayName,
					}
				}
			}
			definitions[s.ID] = Definition{
				Description: s.Description,
				Properties:  properties,
			}
		}
	}
	swagger.Paths = paths
	response := Definition{
		Description: "Reponse type",
		Properties: map[string]Property{
			"ID": Property{
				Type:        "string",
				Description: "L'identifiant interne de la demade",
			},
		},
	}
	definitions["response"] = response
	swagger.Definitions = definitions
	swagger.SecurityDefinitions = map[string]SecurityDef{
		"csa": SecurityDef{
			Type: "apiKey",
			In:   "header",
			Name: "X-Auth-Token",
		},
	}
	o, _ := yaml.Marshal(&swagger)
	err := ioutil.WriteFile("./dist/csa.yaml", o, 0644)
	if err != nil {
		log.Println("Cannot write the swagger file")
	}
}
