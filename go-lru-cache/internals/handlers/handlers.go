package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "go-lru-cache/internal/cache"
)

var lruCache = cache.NewLRUCache(100)

type SetRequest struct {
    Key        string `json:"key"`
    Value      string `json:"value"`
    Expiration int    `json:"expiration"` // expiration in seconds
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
    var req SetRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    lruCache.Set(req.Key, req.Value, time.Duration(req.Expiration)*time.Second)
    w.WriteHeader(http.StatusOK)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    if key == "" {
        http.Error(w, "key is required", http.StatusBadRequest)
        return
    }

    value, found := lruCache.Get(key)
    if !found {
        http.Error(w, "key not found", http.StatusNotFound)
        return
    }

    resp := map[string]interface{}{
        "key":   key,
        "value": value,
    }

    json.NewEncoder(w).Encode(resp)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    if key == "" {
        http.Error(w, "key is required", http.StatusBadRequest)
        return
    }

    lruCache.Delete(key)
    w.WriteHeader(http.StatusOK)
}
