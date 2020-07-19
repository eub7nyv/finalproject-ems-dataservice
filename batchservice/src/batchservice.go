package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const sleepTime = 5

type Config struct {
	Filepath    string
	Filepattern string
	Topic       string
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// initialize the logger
func init() {
	file, err := os.OpenFile("../logs/batchservicelogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// read config file and return all configuration
func readconfigandreturnbytevalue() ([]byte, error) {
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

func searchandprocess(Filepath string, Filepattern string, Topic string) error {
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
			errp := processfile(file, Topic)
			if errp != nil {
				fmt.Println(errp)
				ErrorLogger.Println(errp)
				return errp
			}
			errb := backupfile(file)
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

func processfile(file string, topic string) error {
	InfoLogger.Println("starting the processfile method")
	InfoLogger.Println("Processing file: " + file + " Topic: " + topic)

	dataFile, err := os.Open(file)
	defer dataFile.Close()
	if err != nil {
		fmt.Println(err)
		ErrorLogger.Println(err)
		return err
	}
	//process above open file by looping thru all records one by one and calling producer service for each record along with topic
	fileScanner := bufio.NewScanner(dataFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	for i, eachline := range fileTextLines {
		fmt.Println("record #: " + strconv.Itoa(i) + " : " + eachline)
		InfoLogger.Println("record #: " + strconv.Itoa(i) + " : " + eachline)
		// ??? Call producer end point and pass on each record/line along with Topic
	}

	InfoLogger.Println("Completed Processing file: " + file)
	InfoLogger.Println("exiting the processfile method")
	return err
}

func backupfile(file string) error {
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

func main() {
	InfoLogger.Println("starting the Main method")
	// read config file and unmarshal values
	byteValue, err := readconfigandreturnbytevalue()
	if err != nil {
		fmt.Println(err)
		ErrorLogger.Println(err)
		return
	}

	var config []Config
	json.Unmarshal(byteValue, &config)

	fmt.Printf("\nFull list of config values: %+v", config)
	InfoLogger.Printf("Full list of config values: %+v", config)
	fmt.Printf("\nLenght of config: %+v", len(config))
	InfoLogger.Printf("Lenght of config: %+v", len(config))

	for _, prop := range config {
		fmt.Printf("\nconfig: %+v", prop.Filepath+prop.Filepattern+" "+prop.Topic)
		InfoLogger.Printf("config: %+v", prop.Filepath+prop.Filepattern+" "+prop.Topic)
		if prop.Filepath == "" || prop.Filepattern == "" || prop.Topic == "" {
			fmt.Printf("\nInvalid Config file .. missing values")
			ErrorLogger.Printf("Invalid Config file .. missing values")
			return
		}
	}

	InfoLogger.Println("End loading the config.json values ..")

	//if config len > 0 and config file has valid values then proceed otherwise exit
	if len(config) > 0 {
		// Add infinite loop and few secs of sleep to keep scanning continously and to process new files as they become available
		for {
			fmt.Printf("\nProceed with scanning for files")
			InfoLogger.Printf("Proceed with scanning for files")
			for _, prop := range config {
				// search for file(s) in folder and process them
				err := searchandprocess(prop.Filepath, prop.Filepattern, prop.Topic)
				if err != nil {
					fmt.Println(err)
					ErrorLogger.Println(err)
					return
				}
			}
			fmt.Printf("\nSleeping for 5 secs")
			InfoLogger.Printf("Sleeping for " + strconv.Itoa(sleepTime) + " secs")
			time.Sleep(sleepTime * time.Second)
		}
	}

	InfoLogger.Println("exiting the Main method successfully")
}
