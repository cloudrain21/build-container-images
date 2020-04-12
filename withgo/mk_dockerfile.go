package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var (
	templateDir = "./template/"
	dockerfilePrefix = "Dockerfile"
	dockerfileDir = "./dockerfile/"
	prefix = "\"$"
	postfix = "\""
)

func fatalf(err error) {
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}
}

func readConfig(filename string) *viper.Viper {
	v := viper.New()

	v.SetConfigName(filename)
	v.SetConfigType("yaml")
	v.AddConfigPath("./")

	err := v.ReadInConfig()
	fatalf(err)

	return v
}

// Construct map for section name and each sections' item list
func convertToConfigMap(v *viper.Viper) (map[string][]string) {
	keys := v.AllKeys()
	configMap := make(map[string][]string)
	var sectionName string
	var itemName string

	for _, key := range keys {
		s := strings.Split(key, ".")
		sectionName = s[0]
		itemName = s[1]
		configMap[sectionName] = append(configMap[sectionName], itemName)
	}
	return configMap
}

func readTemplate(tfile string) (string, error) {
	var buf []byte
	buf, err := ioutil.ReadFile(tfile)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func validItemConfig(sectionName string, itemList []string) error {
	exist := false
	for _, item := range itemList {
		if item == "template" {
			exist = true
			break
		}
	}
	if !exist {
		return errors.New(fmt.Sprintf("template item doesn't exist in [%s] section", sectionName))
	}
	return nil
}

func makeDockerFile(v *viper.Viper) {

	configMap := convertToConfigMap(v)

	for sectionName, itemList := range configMap {
		fatalf( validItemConfig(sectionName, itemList) )

		vkey := sectionName + "." + "template"
		templateFileName := fmt.Sprintf("%v", v.Get(vkey))
		content, err := readTemplate(templateDir + templateFileName)
		fatalf(err)

		for _, item := range itemList {
			if item == "template" {
				continue
			}

			templateStr := prefix + item + postfix
			valueStr := fmt.Sprintf("%v", v.Get(sectionName + "." + item))

			content = strings.ReplaceAll(content, templateStr, valueStr)
		}

		outputFileName := dockerfileDir + dockerfilePrefix + "." + sectionName
		fatalf( ioutil.WriteFile(outputFileName, []byte(content), 0644) )
		fmt.Printf("Done : %s\n", outputFileName)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage : go run %s.go config_file_name\n", path.Base(os.Args[0]))
		os.Exit(0)
	}

	configFile := os.Args[1]

	v := readConfig(configFile)
	makeDockerFile(v)
}
