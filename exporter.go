package exporter

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Tojson : Exports Socialmedia feed into a Json file
func Tojson(u SocialMedia, filename string) error {
	feedmap := make(map[string]string)

	for i := range u.Feed() {
		feedmap[strconv.Itoa(i+1)] = u.Feed()[i]
	}

	file, err := json.MarshalIndent(feedmap, "", " ")
	if err != nil {
		return errors.New("An error occured while parsing to Json: " + err.Error())
	}

	err = ioutil.WriteFile("./exportedfiles/"+filename+".json", file, 0644)
	if err != nil {
		return errors.New("An error occured while writing to Json file: " + err.Error())
	}

	fmt.Println("JSON file is ready")
	return nil
}

// Toxml : Exports Socialmedia feed into a XML file
func Toxml(u SocialMedia, filename string) error {
	file, err := xml.MarshalIndent(u.Feed(), "", " ")
	if err != nil {
		return errors.New("An error occured while parsing to XML: " + err.Error())
	}

	err = ioutil.WriteFile("./exportedfiles/"+filename+".xml", file, 0644)
	if err != nil {
		return errors.New("an error occured while writing to XML file: " + err.Error())
	}
	fmt.Println("XML file ready.")
	return nil
}

// Toyaml : Exports Socialmedia feed into a YAML file
func Toyaml(u SocialMedia, filename string) error {
	file, err := yaml.Marshal(u.Feed())
	if err != nil {
		return errors.New("An error occured while parsing to YAML: " + err.Error())
	}

	err = ioutil.WriteFile("./exportedfiles/"+filename+".yml", file, 0644)
	if err != nil {
		return errors.New("an error occured while writing to YAML file: " + err.Error())
	}
	fmt.Println("YAML file ready.")
	return nil
}

// Totext : Exports Socialmedia feed into a TEXT file
func Totext(u SocialMedia, filename string) error {
	f, err := os.OpenFile("./exportedfiles/"+filename+".txt", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		_, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occured writing to file: " + err.Error())
		}

	}
	fmt.Println("TEXT file is ready")
	return nil
}
