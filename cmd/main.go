package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	// application
	"github.com/sinakeshmiri/asset-notifier/internal/application/api"

	// adapters
	"github.com/sinakeshmiri/asset-notifier/internal/adapters/framework/left/http"
	"github.com/sinakeshmiri/asset-notifier/internal/adapters/framework/right/db"
	"github.com/sinakeshmiri/asset-notifier/internal/adapters/framework/right/notifier/telegram"
	"github.com/sinakeshmiri/asset-notifier/internal/adapters/framework/right/nsprovider/arvan"
	"github.com/sinakeshmiri/asset-notifier/internal/adapters/framework/right/nsprovider/cf"

	//ports
	"github.com/sinakeshmiri/asset-notifier/internal/ports"
)

func main() {
	var err error

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")
	cfKey := os.Getenv("CF_KEY")
	arvanKey := os.Getenv("ARVAN_KEY")
	telegramKey := os.Getenv("TG_KEY")
	telegramChats := os.Getenv("TG_IDS")
	telegramIdstr := strings.Split(telegramChats, ",")
	var telegramIDs []int64
	for _,id := range telegramIdstr {
		m, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalf("failed to parse telegram chat IDs: %v", err)
		}
		telegramIDs=append(telegramIDs, int64(m))
	}
	dbAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()
	cfAdaptor, err := cf.NewAdapter(cfKey)
	if err != nil {
		log.Fatalf("failed to initiate Cloudflare connection: %v", err)
	}
	arvanAdaptor, err := arvan.NewAdapter(arvanKey)
	if err != nil {
		log.Fatalf("failed to initiate Arvan connection: %v", err)
	}
	tgAdaptor, err := telegram.NewAdapter(telegramKey, telegramIDs)
	if err != nil {
		log.Fatalf("failed to initiate Telegram connection: %v", err)
	}

	var nsProviders []ports.NsproviderPort
	var notifProviders []ports.NotifierPort
	nsProviders = append(nsProviders, cfAdaptor)
	nsProviders = append(nsProviders, arvanAdaptor)
	notifProviders = append(notifProviders, tgAdaptor)

	// NOTE: The application's right side port for driven
	// adapters, in this case, a db adapter.
	// Therefore the type for the dbAdapter parameter
	// that is to be injected into the NewApplication will
	// be of type DbPort
	applicationAPI := api.NewApplication(dbAdapter, nsProviders, notifProviders)

	// NOTE: We use dependency injection to give the grpc
	// adapter access to the application, therefore
	// the location of the port is inverted. That is
	// the grpc adapter accesses the hexagon's driving port at the
	// application boundary via dependency injection,
	// therefore the type for the applicaitonAPI parameter
	// that is to be injected into the gRPC adapter will
	// be of type APIPort which is our hexagons left side
	// port for driving adapters
	httpAdapter := http.NewAdapter(applicationAPI)
	httpAdapter.Run()
}
