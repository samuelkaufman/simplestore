package simplestoreapp

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Entity struct {
	Value string
}

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
	simpleStore.Mux.HandleFunc("/messages", simpleStore.MessagesPost).Methods("POST")
	simpleStore.Mux.HandleFunc("/messages/", simpleStore.MessagesPost).Methods("POST")
	simpleStore.Mux.HandleFunc("/messages/{message}", simpleStore.MessagesGet).Methods("GET")
	simpleStore.Mux.HandleFunc("/messages/{message}/", simpleStore.MessagesGet).Methods("GET")
	return simpleStore
}

type Resp struct {
	Digest string `json:"digest"`
}

func (s *SimpleStore) MessagesPost(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	if len(b) == 0 {
		http.Error(w, err.Error(), 500)
		return
	}
	sum := sha256.Sum256(b)
	resp := &Resp{
		Digest: fmt.Sprintf("%x", sum),
	}
	log.Printf("sha256 sum of [%s]\n%s", b, resp.Digest)
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	ctx := appengine.NewContext(r)
	k := datastore.NewKey(ctx, "SimpleStoreData", resp.Digest, 0, nil)
	e := &Entity{
		Value: string(b),
	}
	if _, err := datastore.Put(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(jsonBytes)
}

func (s *SimpleStore) MessagesGet(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	vars := mux.Vars(r)
	message := vars["message"]
	log.Println("MessagesGet", message)
	e := &Entity{}
	k := datastore.NewKey(ctx, "SimpleStoreData", message, 0, nil)
	if err := datastore.Get(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	w.Write([]byte(e.Value))
}
