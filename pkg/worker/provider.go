package worker

import (
	"encoding/csv"
	"io"
	"os"
)

func NewCSVProvider(sourceFileDir string) (provider CSVProvider) {
	provider = CSVProvider{}
	provider.Status = StatusWaiting
	provider.Type = ProviderByCSV
	provider.SourceFileDir = sourceFileDir
	return
}

func (provider CSVProvider) GetFactory() *Factory {
	return provider.BaseProvider.Factory
}
func (provider CSVProvider) SetFactory(factory *Factory) {
	provider.BaseProvider.Factory = factory
	return
}

func (provider CSVProvider) GetStatus() int {
	return provider.Status
}

func (provider CSVProvider) GetType() int {
	return provider.Type
}

func (provider CSVProvider) Supply() {
	provider.Status = StatusRunning

	log.Debugln(provider.Factory)
	to := provider.Factory.Source
	// log.Infoln("[SUPPLYING START...]")
	file, err := os.OpenFile(provider.SourceFileDir, os.O_RDONLY, 0644)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	r := csv.NewReader(file)
	for line, err := r.Read(); err != io.EOF; line, err = r.Read() {
		to <- line
	}

	// log.Infoln("[SUPPLYING END!]")
	provider.Status = StatusStop
	close(provider.Factory.Source)
	return
}

func (provider CSVProvider) GetDetail() (result map[string]interface{}) {
	result = make(map[string]interface{})
	result["sourceFileDir"] = provider.SourceFileDir
	result["status"] = provider.Status
	return
}
