package main

import (
	"os"
	"testing"
)

func TestLoadConfigFile(t *testing.T) {

	config, err := loadConfigFile()

	if err != nil {
		t.Errorf("Issues with Loading Config file ..")
	}

	for _, prop := range config {
		if prop.Filepath == "" || prop.Filepattern == "" || prop.Appname == "" || prop.Ignorehdr == "" || (prop.Ignorehdr != "Y" && prop.Ignorehdr != "N") || prop.Fileformat == "" || prop.Fileformat != "delimitedtext" {
			t.Errorf("Issues with missing values in Config file ..")
		}
	}

}

func TestSearchAndProcess(t *testing.T) {

	_, erros := os.Stat("../incoming_files/")
	if os.IsNotExist(erros) {
		t.Errorf("Invalid folder/dir ..")
	}

	err := searchAndProcess("../incoming_files/", "ncpdp*.csv", "ncpdp", "Y")
	if err != nil {
		t.Errorf("Issues with Searching and Processing files due to invalid folder or file name..")
	}
}

func TestHttpPost(t *testing.T) {
	if producerEndPoint == "" {
		t.Errorf("Missing producerEndPoint URL ..")
	}
	err := httpPost("abc", "xyz")
	if err != nil {
		t.Errorf("Issues with posting the record to the producer end point..")
	}
}
