package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const sleepTime = 5
const producerEndPoint = "http://mckessonproducer:8081/mckesson/produce"

//const producerEndPoint = "http://httpbin.org/post"

type Config struct {
	Filepath    string
	Filepattern string
	Appname     string
	Ignorehdr   string
	Fileformat  string
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type Message struct {
	Appname string `jason:"appName"`
	Message string `json:"incomingMessage"`
}

// initialize the logger
func init() {
	file, err := os.OpenFile("../logs/batchservicelogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func httpPost(appname string, eachline string) error {
	InfoLogger.Println("starting the httpPost method")

	//jsonData := map[string]string{"appName": appname, "message": eachline}
	message := Message{appname, eachline}
	//jsonValue, _ := json.Marshal(jsonData)
	jsonReq, err := json.Marshal(message)
	response, err := http.Post(producerEndPoint, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		ErrorLogger.Println(err)
		return err
	} else {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		InfoLogger.Println("producerEndPoint: " + producerEndPoint + " request: " + string(jsonReq))
		InfoLogger.Println("response: " + string(data))
		fmt.Println("response: " + string(data))

	}
	InfoLogger.Println("exiting the httpPost method")
	return err
}

func backupFile(file string) error {
	InfoLogger.Println("starting the backupfile method")
	InfoLogger.Println("Backing up file: " + file)
	currentTime := time.Now()
	appendStr := ".processed_" + currentTime.Format("2006-01-02_15.04.05.000000000")
	err := os.Rename(file, file+appendStr)
	if err != nil {
		fmt.Println(err)
		ErrorLogger.Println(err)
		return err
	}
	InfoLogger.Println("Backed up file name: " + file + appendStr)
	InfoLogger.Println("exiting the backupfile method")
	return err
}

func processFile(file string, appname string, ignorehdr string) error {
	InfoLogger.Println("starting the processfile method")
	InfoLogger.Println("Processing file: " + file + " appname: " + appname)

	dataFile, err := os.Open(file)
	defer dataFile.Close()
	if err != nil {
		fmt.Println(err)
		ErrorLogger.Println(err)
		return err
	}
	//process above open file by looping thru all records one by one and calling producer service for each record along with appName
	fileScanner := bufio.NewScanner(dataFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	for i, eachline := range fileTextLines {
		fmt.Println("record #: " + strconv.Itoa(i) + " : " + eachline)
		InfoLogger.Println("record #: " + strconv.Itoa(i) + " : " + eachline)
		if !(ignorehdr == "Y" && i == 0) { // if header record exists, then don't send first record
			errp := httpPost(appname, eachline)
			if errp != nil {
				fmt.Println(errp)
				ErrorLogger.Println(errp)
				return errp
			}
		}
	}

	InfoLogger.Println("Completed Processing file: " + file)
	InfoLogger.Println("exiting the processfile method")
	return err
}

func searchAndProcess(Filepath string, Filepattern string, Appname string, Ignorehdr string) error {
	InfoLogger.Println("starting the searchandprocess method for " + Filepath + Filepattern)
	matches, err := filepath.Glob(Filepath + Filepattern)
	if err != nil {
		fmt.Println(err)
		ErrorLogger.Println(err)
		return err
	}
	if len(matches) != 0 {
		fmt.Println("\nFound : ", matches)
		InfoLogger.Println("Found : ", matches)
		for _, file := range matches {
			fmt.Println("\nFound : ", file)
			// open each file and loop thru all records and call producer end point .. Once process, move the file too backup and add datetimestamp to the file
			errp := processFile(file, Appname, Ignorehdr)
			if errp != nil {
				fmt.Println(errp)
				ErrorLogger.Println(errp)
				return errp
			}
			errb := backupFile(file)
			if errb != nil {
				fmt.Println(errb)
				ErrorLogger.Println(errb)
				return errb
			}
		}
	} else {
		InfoLogger.Println("No files found for " + Filepath + Filepattern)
	}
	InfoLogger.Println("exiting the searchandprocess method successfully")
	return err
}

// read config file and return all configuration
func readConfigAndReturnByteValue() ([]byte, error) {
	InfoLogger.Println("starting the readconfigandreturnbytevalue method")
	var byte_value []byte
	jsonFile, err := os.Open("../config/batchservicefileconfig.json")
	defer jsonFile.Close()
	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		ErrorLogger.Println(err)
		return byte_value, err
	}
	// read our opened File as a byte array.
	file_byte_value, _ := ioutil.ReadAll(jsonFile)
	InfoLogger.Println("exiting the readconfigandreturnbytevalue method successfully before return statement")
	return file_byte_value, err
}

func loadConfigFile() ([]Config, error) {
	InfoLogger.Println("starting the loadconfigfile method")
	// read config file and unmarshal values
	var config []Config
	byteValue, err := readConfigAndReturnByteValue()
	if err != nil {
		fmt.Println(err)
		ErrorLogger.Println(err)
		return config, err
	}

	json.Unmarshal(byteValue, &config)

	fmt.Printf("\nFull list of config values: %+v", config)
	InfoLogger.Printf("Full list of config values: %+v", config)
	fmt.Printf("\nLenght of config: %+v", len(config))
	InfoLogger.Printf("Lenght of config: %+v", len(config))

	for _, prop := range config {
		fmt.Printf("\nconfig: %+v", prop.Filepath+prop.Filepattern+" "+prop.Appname)
		InfoLogger.Printf("config: %+v", prop.Filepath+prop.Filepattern+" "+prop.Appname)
		if prop.Filepath == "" || prop.Filepattern == "" || prop.Appname == "" || prop.Ignorehdr == "" || (prop.Ignorehdr != "Y" && prop.Ignorehdr != "N") || prop.Fileformat == "" || prop.Fileformat != "delimitedtext" {
			fmt.Printf("\nInvalid Config file .. missing values")
			ErrorLogger.Printf("Invalid Config file .. missing values")
			return config, errors.New("Invalid Config file .. missing values")
		}
	}

	InfoLogger.Println("End loading the config.json values ..")
	InfoLogger.Println("exiting the loadconfigfile method")
	return config, err
}

func main() {

	InfoLogger.Println("starting the Main method")
	// read config file and unmarshal values
	config, err := loadConfigFile()
	if err != nil {
		fmt.Println(err)
		ErrorLogger.Println(err)
		return
	}
	//if config len > 0 and config file has valid values then proceed otherwise exit
	if len(config) > 0 {
		// Add infinite loop and few secs of sleep to keep scanning continously and to process new files as they become available
		for {
			fmt.Printf("\nProceed with scanning for files")
			InfoLogger.Printf("Proceed with scanning for files")
			for _, prop := range config {
				// search for file(s) in folder and process them
				err := searchAndProcess(prop.Filepath, prop.Filepattern, prop.Appname, prop.Ignorehdr)
				if err != nil {
					fmt.Println(err)
					ErrorLogger.Println(err)
					return
				}
			}
			fmt.Printf("\nSleeping for " + strconv.Itoa(sleepTime) + " secs")
			InfoLogger.Printf("Sleeping for " + strconv.Itoa(sleepTime) + " secs")
			time.Sleep(sleepTime * time.Second)
		}
	} else {
		InfoLogger.Println("No config entries found for files to scan and proocess.")
	}
	InfoLogger.Println("exiting the Main method..")
}
