package shared

import (
	"fmt"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

func PullProductsAndDnsRecords() map[string]*api.SimplyProduct {
	products := PullProducts()
	if products == nil {
		return nil
	}

	styles.WaitPrint("Getting DNS records from each product")
	var res = make(map[string]*api.SimplyProduct)
	for _, product := range products {
		res[product.Object] = product
		res[product.Object].DnsRecords = PullDnsRecordsForProduct(product, "  ")
	}

	styles.SuccessPrint("DNS records downloaded")
	return res
}

func PullProducts() []*api.SimplyProduct {
	styles.WaitPrint("Getting products from account")
	products, err := api.GetProducts()
	if err != nil {
		styles.FailPrint("Failed to get products")
		styles.FailPrint("Error: %v", err)
		return nil
	}

	styles.SuccessPrint("Products downloaded")
	return products
}

func PullProductNames() []string {
	products := PullProducts()
	var objNames = make([]string, 0)
	for _, product := range products {
		objNames = append(objNames, product.Object)
	}
	return objNames
}

func PullDnsRecords(productObject string, printPrefix string) []*api.SimplyDnsRecord {
	styles.WaitPrint("Getting dns records for '%s'", productObject)
	records, err := api.GetDnsRecords(productObject)
	if err != nil {
		styles.FailPrint("Failed to get DNS records")
		styles.FailPrint("Error: %v", err)
		return nil
	}

	styles.SuccessPrint("DNS records downloaded")
	return records
}

func PullDnsRecordsForProduct(product *api.SimplyProduct, printPrefix string) []*api.SimplyDnsRecord {
	styles.BlankPrint(fmt.Sprintf("%s%s", printPrefix, product.Name))
	records, err := api.GetDnsRecords(product.Object)
	if err != nil {
		styles.WarnPrint(fmt.Sprintf("%s%s >> %v", printPrefix, product.Name, err))
		return make([]*api.SimplyDnsRecord, 0)
	}
	return records
}
