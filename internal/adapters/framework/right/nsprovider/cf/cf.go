package cf

import (
	"context"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/sinakeshmiri/asset-notifier/internal/ports"	
)

type Adapter struct {
	api *cloudflare.API
}

// NewAdapter creates a new Adapter
func NewAdapter(key string) (*Adapter, error) {
	
	//
	api, err := cloudflare.NewWithAPIToken(key)

	if err != nil {
		log.Fatalf("failed to  connect to the CloudFlare: %v", err)
		return nil, err
	}

	return &Adapter{api: api}, nil
}

func (cfa Adapter) GetRecords() ([]ports.NSrecord, error) {
	var recs []ports.NSrecord
	// Fetch all zones available to this user.
	zones, err := cfa.api.ListZones(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, z := range zones {
		cfRecs, err := cfa.api.DNSRecords(context.Background(), z.ID, cloudflare.DNSRecord{})
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		for _, r := range cfRecs {
			recs = append(recs, ports.NSrecord{
				Provider: "CloudFlare",
				Type:     z.Name,
				Name:     r.Name,
				Value:    r.Content,
			})
		}
	}
	return recs, nil
}
