package api

import (
	"fmt"
	"log"
	"sync"

	"github.com/sinakeshmiri/asset-notifier/internal/ports"

)

// Application implements the APIPort interface
type Application struct {
	db          ports.DbPort
	nsproviders []ports.NsproviderPort
	notifiers   []ports.NotifierPort
	an          Assetnotif
}

// NewApplication creates a new Application
func NewApplication(
	db ports.DbPort, nsproviders []ports.NsproviderPort, notifiers []ports.NotifierPort) *Application {
	return &Application{
		db:          db,
		nsproviders: nsproviders,
		notifiers:   notifiers,
	}
}

// CheckDNS will request all avaiable NS providers and will notify us if anything has changed
func (apia Application) CheckDNS() error {
	var wg sync.WaitGroup
	var  err  error
	var newRecs []ports.NSrecord
	var oldRecs []ports.NSrecord
	
	wg.Add(1)
	go func() {
		oldRecs, err = apia.db.GetRecords()
		wg.Done()	
	}()
	if err != nil {
		return err
	}
	for _,nsProv:= range(apia.nsproviders){
		recordsChan := make(chan ports.NSrecord)

		wg.Add(1)
		go func(c  chan ports.NSrecord){
			r,e:=nsProv.GetRecords()
			if err != nil {
				log.Println(e)
			}
			for _,rec:=range(r){
				c  <- rec
			}
			wg.Done()
		}(recordsChan)
		
		go func(c chan ports.NSrecord) {
			for {
				p := <-c
				newRecs = append(newRecs, p)
			}
		}(recordsChan)
		wg.Wait()
	}
	msgs:=func(a []ports.NSrecord, b []ports.NSrecord) []string {
		var res []string
		allrecs := append(a, b...)
	
		for _, rec := range allrecs {
			if constains(a, rec) && !constains(b, rec) {
				res = append(res, fmt.Sprintf("[ - ] %s : %s", rec.Name, rec.Value))
				e:=apia.db.DeleteRecord(rec)
				if err != nil {
					log.Println(e)
				}
			} else if !constains(a, rec) && constains(b, rec) {
				res = append(res, fmt.Sprintf("[ + ] %s : %s", rec.Name, rec.Value))
				e:=apia.db.AddRecord(rec)
				if err != nil {
					log.Println(e)
				}
			}
		}
	
		return res
	}(oldRecs,newRecs)
	for _,notifier:=range(apia.notifiers){
		notifier.SendNotif(msgs)
	}

	return nil
}



func constains(a []ports.NSrecord, x ports.NSrecord) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
