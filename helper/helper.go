package main

import (
	"encoding/json"
	"flag"
	"github.com/owulveryck/CSAaaS/subscription"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func subscription2yaml() error {
	var s subscription.SubscriptionDetail

	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&s); err != nil {
		return err
	}
	o, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	os.Stdout.Write(o)
	return nil
}

func yaml2json() error {
	var s subscription.SubscriptionDetail

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Println("error1:", err)
		return err
	}
	err = yaml.Unmarshal(b, &s)
	if err != nil {
		log.Println("error2:", err)
		return err
	}
	o, err := json.MarshalIndent(s, " ", "  ")
	if err != nil {
		log.Println("error2:", err)
		return err
	}

	os.Stdout.Write(o)
	return nil

}
func main() {
	action := flag.String("action", "yaml", "depending of the parameter (yaml of json) generates the file from stdin")
	flag.Parse()
	switch *action {
	case "json":
		log.Println("calling yaml2json")
		err := yaml2json()
		if err != nil {
			log.Fatal(err)
		}
	case "yaml":
		log.Println("calling subscription2yaml")
		err := subscription2yaml()
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unknown action")
	}
}
