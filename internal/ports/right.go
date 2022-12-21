package ports

type NSrecord struct {
	Provider  string
	Domin string
	Type  string
	Name  string
	Value string
}

// DbPort is the port for a db adapter
type DbPort interface {
	CloseDbConnection()
	GetRecords() ([]NSrecord ,error)
	AddRecord(NSrecord)error
	DeleteRecord(NSrecord)error
}


type NsproviderPort interface {
	GetRecords() ([]NSrecord ,error)
}

type NotifierPort interface {
	SendNotif([]string) (error)
}