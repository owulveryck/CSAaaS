package server

import (
	"github.com/owulveryck/CSAaaS/subscription"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var Subscriptions = make(map[string]subscription.SubscriptionDetail, 0)

func (r *Routes) findRoute(p string, f os.FileInfo, err error) error {
	if path.Base(f.Name()) == "active" {
		// Get the method:
		dirname := path.Dir(p)
		sep := string(filepath.Separator)
		method := path.Base(dirname)
		routepath := filepath.Join(sep, path.Dir(dirname))
		elems := strings.Split(routepath, sep)
		routepath = filepath.Join(elems[2:]...)
		routepath = filepath.Join(sep, routepath)
		switch method {
		case "POST":
			route := Route{
				routepath,
				method,
				routepath,
				CsaCreate,
			}
			*r = append(*r, route)

		case "GET":
			route := Route{
				routepath,
				method,
				path.Join(routepath, "{id}"),
				CsaGet,
			}
			*r = append(*r, route)
		case "DELETE":
			route := Route{
				routepath,
				method,
				path.Join(routepath, "{id}"),
				CsaDelete,
			}
			*r = append(*r, route)
		default:
			log.Println("Warning: discarding unknown method ", method)
			return nil
		}
		var s subscription.SubscriptionDetail
		config := filepath.Join(dirname, "config.yaml")
		if _, err := os.Stat(config); err == nil {
			b, err := ioutil.ReadFile(config)
			err = yaml.Unmarshal(b, &s)
			if err != nil {
				log.Fatalf("Cannot unmarshal the content of %v (%v)", p, err)
			}
		}
		Subscriptions[filepath.Join(routepath, method)] = s
	}
	return nil
}
