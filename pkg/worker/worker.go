package worker

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	client "git.52retail.com/metro/coupon-client"
)

type Factory struct {
	// Workers  []*SendWorker    `json:"workers"`
	Workers  []*CustomerQueryWorker `json:"workers"`
	Provider Provider               `json:"provider"`
	ECNList  []string               `json:"ecnList"`
	Source   chan interface{}       `json:"-"`
	Signal   chan bool              `json:"-"`
	Status   int                    `json:"status"`
}

type Record struct {
	ID            uint
	TradeNo       string
	StoreKey      int
	CustKey       int
	CardholderKey int
	Cardnumber    string
	Status        string
}

func SaveRecord(record *Record) error {
	err := pgConn.Create(record).Error
	return err
}

func NewFactory() (factory *Factory) {
	factory = new(Factory)
	// factory.Workers = []*SendWorker{}
	factory.Workers = []*CustomerQueryWorker{}
	factory.Status = StatusWaiting
	factory.Signal = make(chan bool)
	factory.Source = make(chan interface{})

	// CSV only
	provider := NewCSVProvider("./source/prod.csv")
	provider.BaseProvider.Factory = factory
	factory.Provider = provider
	go provider.Supply()

	factory.ECNList = []string{
		"20201105IFC38",
		"20201105IFC68",
	}

	return
}

func (factory *Factory) Run() {
	factory.Status = StatusRunning
	defer func(factory *Factory) {
		factory.Status = StatusStop
	}(factory)

	count := 0
	for {
		if isEnd, ok := <-factory.Signal; ok {
			if isEnd {
				log.Infof("Worker end,count: %d", count)
				break
			}
			count++
			if count%10 == 0 {
				total := len(factory.Workers)
				log.Infof("[%d]: %d", total, count)
			}
		}
	}
}

func (factory *Factory) ChangeSpeed(isProd bool, count int) error {
	for _, worker := range factory.Workers {
		worker.Status = StatusStop
	}
	// factory.Workers = []*SendWorker{}
	factory.Workers = []*CustomerQueryWorker{}
	for i := 0; i < count; i++ {
		factory.AddWorker(isProd)
	}
	return nil
}

func (factory *Factory) AddWorker(isProd bool) error {
	// worker := NewSendWorker(isProd)
	worker := NewCustomerQueryWorker(isProd)
	factory.Workers = append(factory.Workers, worker)
	worker.Factory = factory
	go worker.DoJob()
	return nil
}

type SendWorker struct {
	Factory       *Factory  `json:"-"`
	ID            int       `json:"id"`
	StartTime     time.Time `json:"startTime"`
	Status        int       `json:"status"`
	FinishedCount int       `json:"finishedCount"`
	EvoEnv        string    `json:"-"`
	EvoAppID      string    `json:"-"`
}

func NewSendWorker(isProd bool) (worker *SendWorker) {
	worker = new(SendWorker)
	worker.StartTime = time.Now()
	if isProd {
		worker.EvoEnv = evoEnvProd
	} else {
		worker.EvoEnv = evoEnvTest
	}
	worker.EvoAppID = evoAppID
	worker.Status = StatusRunning
	worker.FinishedCount = 0
	return
}

func (worker *SendWorker) DoJob() {
	source := worker.Factory.Source
	ecnList := worker.Factory.ECNList
	for item := range source {
		if worker.Status == StatusStop {
			return
		}
		for _, ecn := range ecnList {
			var resp *client.CreateResponse
			slice := item.([]string)
			storeKey, _ := strconv.Atoi(slice[3])
			custKey, _ := strconv.Atoi(slice[4])
			chKey, _ := strconv.Atoi(slice[5])
			tradeNo := slice[0]
			record := &Record{
				TradeNo:       tradeNo,
				StoreKey:      storeKey,
				CustKey:       custKey,
				CardholderKey: chKey,
				Cardnumber:    slice[2],
			}
			// log.Infoln(tradeNo, ecn, storeKey)
			// continue
			err := client.Call(worker.EvoEnv, worker.EvoAppID,
				client.NewCreateRequest(tradeNo, ecn, "", "", storeKey, custKey, chKey, 0), &resp)
			if err == nil {
				switch resp.Code {
				case client.CodeOK:
					record.Status = "ok"
				case client.CodeCouponAlreadyActivated:
					record.Status = "already activated"
				case client.CodeCoupon:
					record.Status = "already redeemed"
				case client.CodeOutOfTime:
					record.Status = "out of time"
				default:
					record.Status = resp.Result
				}
			} else {
				record.Status = "call coupon client error"
			}
			SaveRecord(record)
		}
		worker.Factory.Signal <- false
	}
	// worker.Factory.Signal <- true
}

type CustomerQueryWorker struct {
	Factory       *Factory  `json:"-"`
	ID            int       `json:"id"`
	StartTime     time.Time `json:"startTime"`
	Status        int       `json:"status"`
	FinishedCount int       `json:"finishedCount"`
}

func NewCustomerQueryWorker(isProd bool) (worker *CustomerQueryWorker) {
	worker = new(CustomerQueryWorker)
	worker.StartTime = time.Now()
	worker.Status = StatusRunning
	worker.FinishedCount = 0
	return
}

func (worker *CustomerQueryWorker) DoJob() {
	output, err := os.OpenFile("./output/queryResult.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Warnln(err.Error())
	}
	w := csv.NewWriter(output)
	source := worker.Factory.Source
	for item := range source {
		if worker.Status == StatusStop {
			return
		}
		slice := item.([]string)
		this := slice
		uKey := slice[0]
		// lpNumber := customer.LPNumberByUKey(uKey)
		lpNumber := ""
		// mobile := customer.MobileByUKey(uKey)
		if err == nil {
			this = append(this, lpNumber)
		}
		w.Write(this)
		w.Flush()
		worker.FinishedCount++
		worker.Factory.Signal <- false
	}
	// worker.Factory.Signal <- true
}
