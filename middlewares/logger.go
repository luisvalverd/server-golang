package middleware

import (
	"fmt"
	"net/http"
)

type LoggerRoute struct {
	method, path, protocol string
}

func printLog(log *LoggerRoute) {
	fmt.Printf("%s%s%s%s%s%s\n", "INFO:  ", log.protocol, " - ", log.path, " - ", log.method)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := new(LoggerRoute)
		log.method = r.Method
		log.protocol = r.Proto
		log.path = r.URL.Path

		//	printLog(log)

		fmt.Printf("%s%s%s%s%s%s\n", "INFO:  ", log.protocol, " - ", log.path, " - ", log.method)
		next.ServeHTTP(w, r)
	})
}
