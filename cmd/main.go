package main

import (
	exporter "parsego_to"
	"parsego_to/facebook"
	"parsego_to/linkedin"
	"parsego_to/twitter"
)

func main() {
	/*
		The exporter package contains 4 functions:
		-Tojson
		-Toxml
		-Toyaml
		-Totext

		The functions accepts two arguments:
		-type either [facebook, twitter or Linkedin]
		-then the file name as a string [no need to add extensions]
	*/
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnked := new(linkedin.Linkedin)

	// exports to json
	exporter.Tojson(fb, "fbdata")
	exporter.Tojson(twtr, "twtrdata")
	exporter.Tojson(lnked, "lnkeddata")

	// exports to xml
	exporter.Toxml(fb, "fbdata")
	exporter.Toxml(twtr, "twtrdata")
	exporter.Toxml(lnked, "lnkeddata")

	// exports to yaml
	exporter.Toyaml(fb, "fbdata")
	exporter.Toyaml(twtr, "twtrdata")
	exporter.Toyaml(lnked, "lnkeddata")

	// exports to text
	exporter.Totext(fb, "fbdata")
	exporter.Totext(twtr, "twtrdata")
	exporter.Totext(lnked, "lnkeddata")

}
