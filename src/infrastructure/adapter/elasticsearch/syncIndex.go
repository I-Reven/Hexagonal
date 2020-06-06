package elasticsearch

import (
	"bytes"
	"fmt"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/juju/errors"
	"net/http"
	"os"
)

var (
	SyncIndexRequestBody = `{"mappings": {"%s" : {"discover" : ".*"}}}`
	SyncIndexRequestUrl  = "http://%s:%s/%s"
)

type SyncIndex struct {
	log    logger.Log
	client *http.Client
}

func (a SyncIndex) Send(keySpace string, table string) error {
	client := &http.Client{}
	json := []byte(fmt.Sprintf(SyncIndexRequestBody, table))
	url := fmt.Sprintf(SyncIndexRequestUrl, os.Getenv("ELASSANDRA_HOST"), os.Getenv("ELASTICSEARCH_PORT"), keySpace)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json))

	if err != nil {
		err = errors.NewNotSupported(err, "errors.can-not-create-index")
		a.log.Error(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)

	if err != nil {
		err = errors.NewNotSupported(err, "errors.can-not-create-index")
		a.log.Error(err)
		return err
	}

	return nil
}
