package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIpToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	var app application

	nextHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		val := request.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}
		ip, ok := val.(string)
		if !ok {
			t.Error("Not string")
		}
		t.Log(ip)
	})

	for _, test := range tests {
		handlerTotest := app.addIpToContext(nextHandler)

		req := httptest.NewRequest("Get", "http://testing", nil)
		if test.emptyAddr {
			req.RemoteAddr = ""
		}
		if len(test.headerName) > 0 {
			req.Header.Add(test.headerName, test.headerValue)
		}

		if len(test.addr) > 0 {
			req.RemoteAddr = test.addr
		}
		handlerTotest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	stringForContext := "whatever"

	var app application

	ctx := context.Background()
	ctx = context.WithValue(ctx, contextUserKey, stringForContext)
	ip := app.ipFromContext(ctx)

	if !strings.EqualFold(stringForContext, ip) {
		t.Errorf("Wrong value from context %s", ip)
	}
}
