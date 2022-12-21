package arvan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"log"

	"github.com/sinakeshmiri/asset-notifier/internal/ports"	
)

///////////////

// List  domins
// curl example :
//	curl -X GET  https://napi.arvancloud.com/cdn/4.0/domains  -H "Authorization: Apikey its-in-your-ocean-eyes"

type Adapter struct {
	key string
}

// NewAdapter creates a new Adapter
func NewAdapter(key string) (*Adapter, error) {
	//
	return &Adapter{key: key}, nil
}

func (a Adapter) getDomains() ([]string, error) {
	arvan_api := "https://napi.arvancloud.com/cdn/4.0/"
	domins_ep := "domains"

	req, err := http.NewRequest(http.MethodGet, arvan_api+domins_ep, nil)
	if err != nil {
		return nil, err
	}
	req.Header["Authorization"] = []string{a.key}
	client := http.Client{
		Timeout: 300 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//return res, nil
	var resp dominsResp

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	var ret []string
	for _, dom := range resp.Data {
		ret = append(ret, dom.Domain)
	}
	return ret, nil
}

// List  Records
// curl example :
// 	curl -X GET  curl -X GET  https://napi.arvancloud.com/cdn/4.0/domains/example.com/dns-records -H "Authorization: Apikey its-in-your-ocean-eyes"


func askRecords(domin string, key string, recType string) ([] ports.NSrecord, error) {
	arvan_api := "https://napi.arvancloud.com/cdn/4.0/"
	recs_ep := fmt.Sprintf("domains/%s/dns-records", domin)

	req, err := http.NewRequest(http.MethodGet, arvan_api+recs_ep, nil)
	if err != nil {
		return nil, err
	}
	req.Header["Authorization"] = []string{key}
	q := req.URL.Query()
	q.Add("type", recType)
	req.URL.RawQuery = q.Encode()
	client := http.Client{
		Timeout: 300 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var records []ports.NSrecord

	if recType == "a" {
		var a aRecResp
		err := json.NewDecoder(res.Body).Decode(&a)
		if err != nil {
			return nil, err
		}
		for _, rec := range a.Data {
			for _, ip := range rec.Value {
				records = append(records, ports.NSrecord{
					Provider: "Arvan",
					Domin:    domin,
					Type:     recType,
					Name:     rec.Name,
					Value:    ip.IP,
				})
			}
		}

	} else if recType == "mx" {
		var mx mxRecResp
		err := json.NewDecoder(res.Body).Decode(&mx)
		if err != nil {
			return nil, err
		}
		for _, rec := range mx.Data {
			records = append(records, ports.NSrecord{
				Provider: "Arvan",
				Domin:    domin,
				Type:     recType,
				Name:     rec.Name,
				Value:    rec.Value.Host,
			})
		}

	} else if recType == "cname" {
		var cname cnameRecResp
		err := json.NewDecoder(res.Body).Decode(&cname)
		if err != nil {
			return nil, err
		}
		for _, rec := range cname.Data {
			records = append(records, ports.NSrecord{
				Provider: "Arvan",
				Domin:    domin,
				Type:     recType,
				Name:     rec.Name,
				Value:    rec.Value.Host,
			})
		}
	} else if recType == "ptr" {
		var ptr ptrRecResp
		err := json.NewDecoder(res.Body).Decode(&ptr)
		if err != nil {
			return nil, err
		}
		for _, rec := range ptr.Data {
			records = append(records, ports.NSrecord{
				Provider: "Arvan",
				Domin:    domin,
				Type:     recType,
				Name:     rec.Name,
				Value:    rec.Value.Domain,
			})
		}
	} else if recType == "ns" {
		var ns nsRecResp
		err := json.NewDecoder(res.Body).Decode(&ns)
		if err != nil {
			return nil, err
		}
		for _, rec := range ns.Data {
			records = append(records, ports.NSrecord{
				Provider: "Arvan",
				Domin:    domin,
				Type:     recType,
				Name:     rec.Name,
				Value:    rec.Value.Host,
			})
		}
	} else if recType == "txt" {
		var txt txtRecResp
		err := json.NewDecoder(res.Body).Decode(&txt)
		if err != nil {
			return nil, err
		}
		for _, rec := range txt.Data {
			records = append(records, ports.NSrecord{
				Provider: "Arvan",
				Domin:    domin,
				Type:     recType,
				Name:     rec.Name,
				Value:    rec.Value.Text,
			})
		}
	}
	return records, nil
}

// ////////////
func (a Adapter) GetRecords() ([]ports.NSrecord, error) {
	domins, err := a.getDomains()
	if err != nil {
		return nil, err
	}

	var recs []ports.NSrecord
	var wg sync.WaitGroup
	types := []string{"mx", "cname", "a", "ptr", "ns", "txt"}
	res := make(chan ports.NSrecord)
	for _, domin := range domins {
		wg.Add(1)
		go func(d string, c chan ports.NSrecord) {
			for _, recType := range types {
				wg.Add(1)
				go func(rt string) {
					t, err := askRecords(domin, a.key, recType)
					if err != nil {
						log.Println(err)
					}
					for _, r := range t {
						c <- r
					}
					wg.Done()
				}(recType)
			}
			wg.Done()
		}(domin, res)
		go func(c chan ports.NSrecord) {
			for {
				p := <-c
				recs = append(recs, p)
			}
		}(res)
		wg.Wait()
	}

	/*
		for _, domin := range domins {
			for _, recType := range types {
				t, err := askRecords(domin, a.key, recType)
				if err != nil {
					return nil, err
				}
				recs = append(recs, t...)
			}

		}*/

	return recs, nil
}
