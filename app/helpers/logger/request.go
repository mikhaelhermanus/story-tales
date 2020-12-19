package logger

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

// RecordThridParty ...
func RecordThridParty(ctx context.Context, req *http.Request, start time.Time, status int, response []byte) context.Context {
	var body string
	t := time.Since(start)

	if req.Method != http.MethodGet {
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			body = ""
		} else {
			body = string(reqBody)
			req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		}
	}

	v, ok := ctx.Value(logKey).(*Data)
	if ok {
		third := ThirdParty{}
		third.URL = req.Host + req.URL.Path
		third.RequestBody = body
		third.Response = string(response)
		third.StatusCode = status
		third.RequestHeader = DumpRequest(req)
		third.ExecTime = t.Seconds()
		third.RequestBody = body

		v.ThirdParty = append(v.ThirdParty, third)

		ctx = context.WithValue(ctx, logKey, v)

		return ctx
	}

	return ctx
}

// RecordThridPartyFailed ...
func RecordThridPartyFailed(ctx context.Context, req *http.Request, start time.Time, status int, messages string) context.Context {
	var (
		body = ""
		url  = req.Host + req.URL.Path
	)
	t := time.Since(start)

	if req != nil {
		if req.Method != http.MethodGet {
			b, _ := ioutil.ReadAll(req.Body)

			body = string(b)
		}
	} else {
		url = ""
	}

	v, ok := ctx.Value(logKey).(*Data)
	if ok {
		third := ThirdParty{}
		third.URL = url
		third.RequestBody = body
		third.Response = messages
		third.StatusCode = status
		third.RequestHeader = DumpRequest(req)
		third.ExecTime = t.Seconds()

		v.ThirdParty = append(v.ThirdParty, third)

		ctx = context.WithValue(ctx, logKey, v)

		return ctx
	}

	return ctx
}
