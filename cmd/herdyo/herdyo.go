package main

import (
	"github.com/gin-contrib/sessions/memstore"
	"github.com/istream-philion/cowyo/server"

	"net/http"
	"strings"
)

func main() {
	store := memstore.NewStore([]byte("secret"))

	first := server.Site{
		PathToData:    "site1",
		Debounce:      500,
		SessionStore:  store,
		AllowInsecure: true,
		Fileuploads:   true,
		MaxUploadSize: 2,
	}.Router()

	second := server.Site{
		PathToData:    "site2",
		Debounce:      500,
		SessionStore:  store,
		AllowInsecure: true,
		Fileuploads:   true,
		MaxUploadSize: 2,
	}.Router()
	panic(http.ListenAndServe("localhost:8000", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Host, "first") {
			first.ServeHTTP(rw, r)
		} else if strings.HasPrefix(r.Host, "second") {
			second.ServeHTTP(rw, r)
		} else {
			http.NotFound(rw, r)
		}
	})))
}
