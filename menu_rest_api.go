package main

import (
    "encoding/json"
    "net/http"
    "strings"
)

func main() {
     http.HandleFunc("/menu/", func(w http.ResponseWriter, r *http.Request) {
        item := strings.SplitN(r.URL.Path, "/", 3)[2]

        data, err := queryMenu(item)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        json.NewEncoder(w).Encode(data)
    })
    http.ListenAndServe(":8080", nil)
}
func queryMenu(item string) (beer, error) {
    resp, err := http.Get("http://127.0.0.1:28017/menu/beers/")
    if err != nil {
        return beer{}, err
    }

    defer resp.Body.Close()

    var d beer

    if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
        return beer{}, err
    }

    return d, nil
}

type beer struct {
    Beers []struct {
		Name string `json:"name"`
		Price int `json:"price"`
		Kind string `json:"kind"`
    } `json:"rows"`
}
