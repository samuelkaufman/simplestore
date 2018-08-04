package simplestoreapp

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type SimpleStore struct {
	Mux *mux.Router
}

func (s *SimpleStore) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

func New() *SimpleStore {
	simpleStore := &SimpleStore{
		Mux: mux.NewRouter(),
	}
	simpleStore.Mux.HandleFunc("/messages", simpleStore.MessagesPost).Methods("GET")
	simpleStore.Mux.HandleFunc("/messages/", simpleStore.MessagesPost).Methods("GET")
	simpleStore.Mux.HandleFunc("/messages/{message}", simpleStore.MessagesGet).Methods("GET")
	simpleStore.Mux.HandleFunc("/messages/{message}/", simpleStore.MessagesGet).Methods("GET")
	return simpleStore
}
func (s *SimpleStore) MessagesPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("HELLO")
	w.Write([]byte(`{"hello":"world"}`))
	ctx := appengine.NewContext(r)

	k := datastore.NewKey(ctx, "Entity", "stringID", 0, nil)
	e := new(Entity)
	if err := datastore.Get(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	old := e.Value
	e.Value = r.URL.Path

	if _, err := datastore.Put(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *SimpleStore) MessagesGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["message"]
	log.Println("MessagesGet", message)
	w.Write([]byte(`{"hello":"` + message + `"}`))
}
