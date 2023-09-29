/**

Program: redlineBodyFile
Author: Duane Dunston
Date: 2023-08-29

Redline File to Body File.

This will read the files-api.urn files from the RedLine audit and allow creating a body file
on a specific file or directory path. Use the path in the <FilePathName> XML tag.

./redlineBodyFile -f <RedLineAuditFile.xml> -d "path to search"

Example:

./redlineBodyFile -f redline_audit.xml -d "c:\\documents"

Usage of redlineBodyFile:

	-d string
	      The directory to scan (no trailing slash) or full file path.
	-f string
	      The RedLine Audit file.
*/
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// XML Structure.
type FileItem struct {
	XMLName     xml.Name `xml:"FileItem"`
	FullPath    string   `xml:"FullPath"`
	FileName    string   `xml:"FileName"`
	SizeInBytes int      `xml:"SizeInBytes"`
	Modified    string   `xml:"Modified"`
	Accessed    string   `xml:"Accessed"`
	Changed     string   `xml:"Changed"`
	Username    string   `xml:"Username"`
	SecurityID  string   `xml:"SecurityID"`
	Group       string   `xml:"Group"`
	GroupID     string   `xml:"GroupID"`
	Permissions string   `xml:"Permissions"`
	Attributes  string   `xml:"FileAttributes"`
}

// FileItem is the tag that separates each XML object.
type ItemList struct {
	XMLName  xml.Name   `xml:"itemList"`
	FileItem []FileItem `xml:"FileItem"`
}

/** Function converts time to the Unix Timestamp */
func convertTime(theTime string) int64 {
	unixTime, err := time.Parse(time.RFC3339, theTime)
	if err != nil {

		log.Fatal(err)

	}
	return unixTime.Unix()
}

func main() {

	var dir string
	var filename string

	// Set a commandline parameter for the directory using flags.
	flag.StringVar(&dir, "d", "", "The directory to scan (no trailing slash) or full file path.")
	flag.StringVar(&filename, "f", "", "The RedLine Audit file.")
	flag.Parse()

	// Check if a directory was specified.
	if dir == "" || filename == "" {

		flag.Usage()
		os.Exit(1)

	}

	// Open the file.
	file, err := os.Open(filename)
	if err != nil {

		log.Fatalf("Error opening the file: %v", err)

	}
	defer file.Close()

	// Create an XML decoder.
	decoder := xml.NewDecoder(file)

	// Loop through the file.
	for {

		// Read the file as a stream instead of loading it into memory.
		// Here it is reading each XML FileItem.
		tok, err := decoder.Token()
		if err != nil {

			break

		}

		// Check if the token is a StartElement
		switch checkXML := tok.(type) {

		case xml.StartElement:

			// Check if the tag name is FileItem.
			if checkXML.Name.Local == "FileItem" {

				var item FileItem

				// Decode the XML element.
				if err := decoder.DecodeElement(&item, &checkXML); err != nil {

					log.Fatalf("DecodeElement error: %v", err)

				}

				// Search only for the directory or file path.
				if !strings.HasPrefix(item.FullPath, dir) {

					// Skip this item if there is no match.
					continue

				}

				// Convet time.
				modified := convertTime(item.Modified)
				accessed := convertTime(item.Accessed)
				changed := convertTime(item.Changed)

				// Format for the results.
				result := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%d|%d|%d|%d|%d|",

					// MD5 is not available from the RedLine audit file.
					"0",
					item.FullPath,

					// The inode is not available from the RedLine audit file.
					"0",
					"0"+item.Permissions,
					item.SecurityID,
					item.GroupID,
					item.SizeInBytes,
					accessed,
					modified,
					changed,
					changed,
				)

				fmt.Println(strings.TrimSuffix(result, "|"))
			}
		}
	}
}
