//usr/bin/go run $0 $@ ; exit
package main

import (
	"fmt"
	"os"

	"github.com/namedotcom/go/namecom"
)

func main() {
	username := os.Getenv("NAMECOM_USERNAME")
	apiToken := os.Getenv("NAMECOM_API_TOKEN")
	domainName := os.Getenv("NAMECOM_DOMAIN_NAME")
	host := os.Getenv("NAMECOM_DOMAIN_HOST")

	if len(username) == 0 || len(apiToken) == 0 || len(domainName) == 0 || len(host) == 0 {
		fmt.Println("NAMECOM environment variables missing")
		os.Exit(1)
	}

	nc := namecom.New(username, apiToken)

	listRecordsRequest := namecom.ListRecordsRequest{
		DomainName: domainName,
		PerPage:    100,
		Page:       1,
	}
	response, err := nc.ListRecords(&listRecordsRequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ourRecord namecom.Record
	for _, record := range response.Records {
		if record.Type == "A" && record.Host == host {
			ourRecord = *record
			break
		}
	}

	ip := "95.19.243.87"
	if ourRecord.Answer == ip {
		fmt.Println("No update required")
		os.Exit(0)
	}

	ourRecord.Answer = ip
	_, err = nc.UpdateRecord(&ourRecord)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
