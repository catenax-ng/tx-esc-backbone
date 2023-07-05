package main

import (
	"encoding/json"
	"fmt"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	origResIdParam = "origResId"
)

type router struct {
	addr               string
	mux                *mux.Router
	resourceSyncClient ResourceSyncClient
}

func newRouter(resourceSyncClient ResourceSyncClient) router {

	return router{
		addr:               ":8080",
		mux:                mux.NewRouter().StrictSlash(true),
		resourceSyncClient: resourceSyncClient,
	}
}
func (t router) getStatus(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(r.RemoteAddr)
}

func (t router) createResource(w http.ResponseWriter, r *http.Request) {
	var resource types.Resource
	err := json.NewDecoder(r.Body).Decode(&resource)
	if err != nil {
		_, _ = fmt.Fprint(w, "invalid body")
		w.WriteHeader(400)
		return
	}
	resp, err := t.resourceSyncClient.CreateResource(r.Context(), resource)
	if err != nil {
		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, err.Error())
	} else {
		bz, err := resp.Codec.Marshal(resp)
		if err != nil {
			w.WriteHeader(500)
			_, _ = fmt.Fprintln(w, "Cannot marshel response: ", err.Error())
		}
		w.WriteHeader(201)
		_, _ = w.Write(bz)
	}
}

func (t router) updateResource(w http.ResponseWriter, r *http.Request) {
	var resource types.Resource
	err := json.NewDecoder(r.Body).Decode(&resource)
	if err != nil {
		_, _ = fmt.Fprint(w, "invalid body")
		w.WriteHeader(400)
		return
	}
	resp, err := t.resourceSyncClient.UpdateResource(r.Context(), resource)
	if err != nil {
		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, err.Error())
	} else {
		bz, err := resp.Codec.Marshal(resp)
		if err != nil {
			w.WriteHeader(500)
			_, _ = fmt.Fprintln(w, "Cannot marshel response: ", err.Error())
		}
		w.WriteHeader(201)
		_, _ = w.Write(bz)
	}
}

func (t router) updateResourceWithObj(w http.ResponseWriter, r *http.Request) {
	var resource types.Resource
	err := json.NewDecoder(r.Body).Decode(&resource)
	if err != nil {
		_, _ = fmt.Fprint(w, "invalid body")
		w.WriteHeader(400)
		return
	}
	resp, err := t.resourceSyncClient.DeleteResource(r.Context(), resource.OrigResId)
	if err != nil {
		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, err.Error())
	} else {
		bz, err := resp.Codec.Marshal(resp)
		if err != nil {
			w.WriteHeader(500)
			_, _ = fmt.Fprintln(w, "Cannot marshel response: ", err.Error())
		}
		w.WriteHeader(201)
		_, _ = w.Write(bz)
	}
}

func (t router) deleteResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	origResId, ok := vars[origResIdParam]
	if !ok {
		_, _ = fmt.Fprintf(w, "%s missing", origResId)
		w.WriteHeader(400)
	}
	resp, err := t.resourceSyncClient.DeleteResource(r.Context(), origResId)
	if err != nil {
		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, err.Error())
	} else {
		bz, err := resp.Codec.Marshal(resp)
		if err != nil {
			w.WriteHeader(500)
			_, _ = fmt.Fprintln(w, "Cannot marshel response: ", err.Error())
		}
		w.WriteHeader(201)
		_, _ = w.Write(bz)
	}
}

func (t router) handleRequests() {
	t.mux.HandleFunc("/", t.getStatus)
	t.mux.HandleFunc("/resource", t.createResource).Methods("POST")
	t.mux.HandleFunc("/resource", t.updateResource).Methods("UPDATE")
	t.mux.HandleFunc("/resource", t.updateResourceWithObj).Methods("DELETE")
	t.mux.HandleFunc(fmt.Sprintf("/resource/{%s}", origResIdParam), t.deleteResource).Methods("DELETE")

	log.Fatal(http.ListenAndServe(t.addr, t.mux))
}
