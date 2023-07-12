// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0

// Package rest-wrapper.

// ESC-REST-Wrapper
//
// the purpose of this application is to wrap the ESC-Backbone behind a web2 API.
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//	Schemes: http
//	Host: localhost
//	BasePath: /
//	Version: 0.0.1
//	License: Apache-2.0 http://www.apache.org/licenses/
//	Contact: Lars Wegner<lars.wegner@bosch.com>
//
//	Consumes:
//	- application/json
//	Produces:
//	- application/json
//
// swagger:meta
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

func newRouter(
	config *Config,
	resourceSyncClient ResourceSyncClient,
) router {
	return router{
		addr:               config.HostAddress,
		mux:                mux.NewRouter().StrictSlash(true),
		resourceSyncClient: resourceSyncClient,
	}
}
func (t router) getStatus(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(r.RemoteAddr)
}

func (t router) createResource(w http.ResponseWriter, r *http.Request) {
	var resource RequestResource
	err := json.NewDecoder(r.Body).Decode(&resource)
	if err != nil {
		_, _ = fmt.Fprint(w, "invalid body")
		w.WriteHeader(400)
		return
	}
	resp, err := t.resourceSyncClient.CreateResource(r.Context(), types.Resource{
		OrigResId:    resource.OrigResId,
		TargetSystem: resource.TargetSystem,
		ResourceKey:  resource.ResourceKey,
		DataHash:     resource.DataHash,
	})
	if err != nil {
		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, err.Error())
	} else {
		response := ResponseResource{
			Height: resp.Height,
			TxHash: resp.TxHash,
			RawLog: resp.RawLog,
		}
		bz, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(500)
			_, _ = fmt.Fprintln(w, "Cannot marshal response: ", err.Error())
		}
		w.WriteHeader(201)
		_, _ = w.Write(bz)
	}
}

func (t router) updateResource(w http.ResponseWriter, r *http.Request) {
	var resource RequestResource
	err := json.NewDecoder(r.Body).Decode(&resource)
	if err != nil {
		_, _ = fmt.Fprint(w, "invalid body")
		w.WriteHeader(400)
		return
	}
	resp, err := t.resourceSyncClient.UpdateResource(r.Context(), types.Resource{
		OrigResId:    resource.OrigResId,
		TargetSystem: resource.TargetSystem,
		ResourceKey:  resource.ResourceKey,
		DataHash:     resource.DataHash,
	})
	if err != nil {
		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, err.Error())
	} else {
		response := ResponseResource{
			Height: resp.Height,
			TxHash: resp.TxHash,
			RawLog: resp.RawLog,
		}
		bz, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(500)
			_, _ = fmt.Fprintln(w, "Cannot marshal response: ", err.Error())
		}
		w.WriteHeader(201)
		_, _ = w.Write(bz)
	}
}

func (t router) deleteResourceWithObj(w http.ResponseWriter, r *http.Request) {
	var resource RequestResource
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
		response := ResponseResource{
			Height: resp.Height,
			TxHash: resp.TxHash,
			RawLog: resp.RawLog,
		}
		bz, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(500)
			_, _ = fmt.Fprintln(w, "Cannot marshal response: ", err.Error())
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
		response := ResponseResource{
			Height: resp.Height,
			TxHash: resp.TxHash,
			RawLog: resp.RawLog,
		}
		bz, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(500)
			_, _ = fmt.Fprintln(w, "Cannot marshal response: ", err.Error())
		}
		w.WriteHeader(201)
		_, _ = w.Write(bz)
	}
}

func (t router) handleRequests() {
	t.mux.HandleFunc("/", t.getStatus)
	t.mux.HandleFunc("/resource", t.createResource).Methods("POST")
	// swagger:operation POST /resource create-resource
	// ---
	// summary: Create a resource
	// description: Create a resource with specified parameters.
	// parameters:
	// - name: RequestResource
	//   in: body
	//   description: Resource to create
	//   schema:
	//     "$ref": "#/definitions/RequestResource"
	// responses:
	//   '200':
	//     description: Response on resource creation
	//     schema:
	//       type: object
	//       items:
	//         "$ref": "#/definitions/ResponseResource"
	//   '400':
	//     description: Transaction failed or bad request body
	//     schema:
	//       type: string
	//   '500':
	//     description: Failed to marshal the response to json
	//     schema:
	//       type: string
	t.mux.HandleFunc("/resource", t.updateResource).Methods("PATCH")
	// swagger:operation PATCH /resource update-resource
	// ---
	// summary: Update a resource
	// description: Update a resource with specified origResId.
	// parameters:
	// - name: RequestResource
	//   in: body
	//   description: Resource to create
	//   schema:
	//     "$ref": "#/definitions/RequestResource"
	// responses:
	//   '200':
	//     description: Response on resource update
	//     schema:
	//       type: object
	//       items:
	//         "$ref": "#/definitions/ResponseResource"
	//   '400':
	//     description: Transaction failed or bad request body
	//     schema:
	//       type: string
	//   '500':
	//     description: Failed to marshal the response to json
	//     schema:
	//       type: string
	t.mux.HandleFunc("/resource", t.deleteResourceWithObj).Methods("DELETE")
	// swagger:operation DELETE /resource delete-resource
	// ---
	// summary: Delete a resource
	// description: Deletes the resource with specified origResId.
	// produces:
	// - application/json
	// parameters:
	// - name: RequestResource
	//   in: body
	//   description: Resource to create
	//   schema:
	//     "$ref": "#/definitions/RequestResource"
	// responses:
	//   '200':
	//     description: Response on resource deletion
	//     schema:
	//       type: object
	//       items:
	//         "$ref": "#/definitions/ResponseResource"
	//   '400':
	//     description: Transaction failed or bad request body
	//     schema:
	//       type: string
	//   '500':
	//     description: Failed to marshal the response to json
	//     schema:
	//       type: string
	t.mux.HandleFunc(fmt.Sprintf("/resource/{%s}", origResIdParam), t.deleteResource).Methods("DELETE")
	// swagger:operation DELETE /resource/{origResId} delete-resource-query
	// ---
	// summary: Delete a resource
	// description: Deletes the resource with given origResId.
	// parameters:
	// - name: origResId
	//   in: path
	//   description: Deletes the resource with specified origResId.
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: Response on resource deletion
	//     schema:
	//       type: object
	//       items:
	//         "$ref": "#/definitions/ResponseResource"
	//   '400':
	//     description: Parameter missing or transaction failed
	//     schema:
	//       type: string
	//   '500':
	//     description: Failed to marshal the response to json
	//     schema:
	//       type: string

	log.Fatal(http.ListenAndServe(t.addr, t.mux))
}
