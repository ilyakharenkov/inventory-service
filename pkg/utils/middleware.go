package utils

import (
	"log"
	"net/http"
	"runtime/debug"
)

func SafeHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v\n%s", err, debug.Stack())
				http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			}
		}()

		handler(w, r)
	}
}
