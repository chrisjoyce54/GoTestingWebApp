package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type contextKey string

const contextUserKey contextKey = "user_ip"

func (app *application) ipFromContext(ctx context.Context) string {
	return ctx.Value(contextUserKey).(string)
}

func (app *application) addIpToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var ctx = context.Background()

		ip, err := getIp(request)
		if err != nil {
			ip, _, _ = net.SplitHostPort(request.RemoteAddr)
			if len(ip) == 0 {
				ip = "unknown"
			}
		}

		ctx = context.WithValue(request.Context(), contextUserKey, ip)

		fmt.Printf("ip 30: %s", ip)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func getIp(request *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	fmt.Printf("ip: %s", ip)
	if err != nil {
		return "unknown", err
	}

	userIp := net.ParseIP(ip)
	if userIp == nil {
		return "", fmt.Errorf("UserIp: %q is not ip:port", request.RemoteAddr)
	}

	forward := request.Header.Get("X-forwarded-for")
	if len(forward) > 0 {
		ip = forward
		fmt.Printf("forward: %s", forward)
	}

	if len(ip) == 0 {
		ip = "forward"
	}

	fmt.Printf("ip 57: %s", ip)
	return ip, nil
}
