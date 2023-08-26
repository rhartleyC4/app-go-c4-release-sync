package main

import (
	"context"
	"fmt"
	"github.com/hooklift/gowsdl/soap"
	"github.com/snap-one/app-go-c4-release-sync/service"
)

//go:generate go run github.com/hooklift/gowsdl/cmd/gowsdl -p service -o release.go https://services.control4.com/Updates2x-internal/v2_0/updates.asmx?WSDL

func main() {
	client := soap.NewClient("https://services.control4.com/Updates2x-internal/v2_0/updates.asmx?WSDL")
	svc := service.NewUpdatesSoap(client)
	result, err := svc.GetAllVersionsContext(context.Background(), &service.GetAllVersions{
		CurrentVersion:         "3.4.0",
		IncludeEarlierVersions: false,
	})
	if err != nil {
		return
	}
	result2, err := svc.GetVersionsContext(context.Background(), &service.GetVersions{
		CurrentVersion: "3.3.0",
	})
	if err != nil {
		return
	}
	for _, v := range result2.GetVersionsResult.Astring {
		fmt.Println(fmt.Sprintf("%+v", v))
	}
	fmt.Println(fmt.Sprintf("%+v", result.GetAllVersionsResult))
}
