// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package http

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	dshelp "github.com/ipfs/go-ipfs-ds-help"
	dag "github.com/ipfs/go-merkledag"
	"github.com/multiformats/go-multihash"
	"github.com/pkg/errors"

	corecrdt "github.com/sourcenetwork/defradb/core/crdt"
	"github.com/sourcenetwork/defradb/events"
)

const (
	contentTypeJSON           = "application/json"
	contentTypeGraphQL        = "application/graphql"
	contentTypeFormURLEncoded = "application/x-www-form-urlencoded"
)

func rootHandler(rw http.ResponseWriter, req *http.Request) {
	sendJSON(
		req.Context(),
		rw,
		simpleDataResponse(
			"response", "Welcome to the DefraDB HTTP API. Use /graphql to send queries to the database."+
				" Read the documentation at https://docs.source.network/.",
		),
		http.StatusOK,
	)
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	sendJSON(
		req.Context(),
		rw,
		simpleDataResponse("response", "pong", "test"),
		http.StatusOK,
	)
}

func dumpHandler(rw http.ResponseWriter, req *http.Request) {
	db, err := dbFromContext(req.Context())
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusInternalServerError)
		return
	}

	err = db.PrintDump(req.Context())
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusInternalServerError)
		return
	}

	sendJSON(
		req.Context(),
		rw,
		simpleDataResponse("response", "ok"),
		http.StatusOK,
	)
}

type gqlRequest struct {
	Query string `json:"query"`
}

func execGQLHandler(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("query")
	if query == "" {
		// extract the media type from the content-type header
		contentType, _, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
		// mime.ParseMediaType will return an error (mime: no media type)
		// if there is no media type set (i.e. application/json).
		// This however is not a failing condition as not setting the content-type header
		// should still make for a valid request and hit our default switch case.
		if err != nil && err.Error() != "mime: no media type" {
			handleErr(req.Context(), rw, err, http.StatusBadRequest)
			return
		}

		switch contentType {
		case contentTypeJSON:
			gqlReq := gqlRequest{}

			err := getJSON(req, &gqlReq)
			if err != nil {
				handleErr(req.Context(), rw, err, http.StatusBadRequest)
				return
			}

			query = gqlReq.Query

		case contentTypeFormURLEncoded:
			handleErr(
				req.Context(),
				rw,
				ErrFormNotSupported,
				http.StatusBadRequest,
			)
			return

		case contentTypeGraphQL:
			fallthrough

		default:
			if req.Body == nil {
				handleErr(req.Context(), rw, ErrBodyEmpty, http.StatusBadRequest)
				return
			}
			body, err := io.ReadAll(req.Body)
			if err != nil {
				handleErr(req.Context(), rw, errors.WithStack(err), http.StatusBadRequest)
				return
			}
			query = string(body)
		}
	}

	// if at this point query is still empty, return an error
	if query == "" {
		handleErr(req.Context(), rw, ErrMissingGQLQuery, http.StatusBadRequest)
		return
	}

	db, err := dbFromContext(req.Context())
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusInternalServerError)
		return
	}
	result := db.ExecQuery(req.Context(), query)

	if result.Pub != nil {
		subscriptionHandler(result.Pub, rw, req)
		return
	}

	sendJSON(req.Context(), rw, result.GQL, http.StatusOK)
}

func loadSchemaHandler(rw http.ResponseWriter, req *http.Request) {
	sdl, err := io.ReadAll(req.Body)
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusBadRequest)
		return
	}

	db, err := dbFromContext(req.Context())
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusInternalServerError)
		return
	}

	err = db.AddSchema(req.Context(), string(sdl))
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusBadRequest)
		return
	}

	sendJSON(
		req.Context(),
		rw,
		simpleDataResponse("result", "success"),
		http.StatusBadRequest,
	)
}

func getBlockHandler(rw http.ResponseWriter, req *http.Request) {
	cidStr := chi.URLParam(req, "cid")

	// try to parse CID
	cID, err := cid.Decode(cidStr)
	if err != nil {
		// If we can't try to parse DSKeyToCID
		// return error if we still can't
		key := ds.NewKey(cidStr)
		var hash multihash.Multihash
		hash, err = dshelp.DsKeyToMultihash(key)
		if err != nil {
			handleErr(req.Context(), rw, err, http.StatusBadRequest)
			return
		}
		cID = cid.NewCidV1(cid.Raw, hash)
	}

	db, err := dbFromContext(req.Context())
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusInternalServerError)
		return
	}

	block, err := db.Blockstore().Get(req.Context(), cID)
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusBadRequest)
		return
	}

	nd, err := dag.DecodeProtobuf(block.RawData())
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusBadRequest)
		return
	}

	buf, err := nd.MarshalJSON()
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusInternalServerError)
		return
	}

	reg := corecrdt.LWWRegister{}
	delta, err := reg.DeltaDecode(nd)
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusInternalServerError)
		return
	}

	data, err := delta.Marshal()
	if err != nil {
		handleErr(req.Context(), rw, err, http.StatusInternalServerError)
		return
	}

	sendJSON(
		req.Context(),
		rw,
		simpleDataResponse(
			"block", string(buf),
			"delta", string(data),
			"val", delta.Value(),
		),
		http.StatusOK,
	)
}

func peerIDHandler(rw http.ResponseWriter, req *http.Request) {
	peerID, ok := req.Context().Value(ctxPeerID{}).(string)
	if !ok || peerID == "" {
		handleErr(req.Context(), rw, ErrPeerIdUnavailable, http.StatusNotFound)
		return
	}

	sendJSON(
		req.Context(),
		rw,
		simpleDataResponse(
			"peerID", peerID,
		),
		http.StatusOK,
	)
}

func subscriptionHandler(pub *events.Publisher[events.Update], rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)
	if !ok {
		handleErr(req.Context(), rw, ErrStreamingUnsupported, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")

	for {
		select {
		case <-req.Context().Done():
			pub.Unsubscribe()
			return
		case s, open := <-pub.Stream():
			if !open {
				return
			}
			b, err := json.Marshal(s)
			if err != nil {
				handleErr(req.Context(), rw, err, http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(rw, "data: %s\n\n", b)
			flusher.Flush()
		}
	}
}