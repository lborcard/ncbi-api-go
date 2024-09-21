package main

import (
	"context"
	"fmt"
	"github.com/lborcard/ncbi-api-go"
)

func main() {
	cfg := openapi.NewConfiguration()
	client := openapi.NewAPIClient(cfg)

	ctx := context.Background()
	GenomeApi := client.GenomeApi

	taxids := make([]string, 1)
	taxids[0] = "12342"

	_, resp, err := GenomeApi.DownloadAssemblyPackage(ctx, taxids).Filename("test2.zip").Execute()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status, resp.StatusCode)

}
