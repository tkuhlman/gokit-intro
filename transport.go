package main

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
)

// These two decode/encode functions are just dealing with the io.Reader/Writer for the http transport. They would
// be more complicated if we were dealing with objects other than strings or a special encoding.

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	q, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return queryRequest{query: string(q)}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	queryResp, ok := response.(queryResponse)
	if !ok {
		return errors.New("failed to encode query response")
	}
	_, err := w.Write([]byte(queryResp.resp))
	return err
}
