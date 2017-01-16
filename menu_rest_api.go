package main

import (
    "encoding/json"
    "net/http"
)

func main() {
     http.HandleFunc("/menu/beer", func(w http.ResponseWriter, r *http.Request) {
       

        data, err := queryBeers()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        json.NewEncoder(w).Encode(data)
    })

    http.HandleFunc("/menu/food", func(w http.ResponseWriter, r *http.Request) {
       

        data, err := queryFood()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        json.NewEncoder(w).Encode(data)
    })
    
    http.ListenAndServe(":8080", nil)
}
func queryBeers() (beer, error) {
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

func queryFood() (foods, error) {
    resp, err := http.Get("http://127.0.0.1:28017/menu/items/")
    if err != nil {
        return foods{}, err
    }

    defer resp.Body.Close()

    var d foods

    if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
        return foods{}, err
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

type foods struct {
    Food []struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Kind string `json:"kind"`
	About string `json:"about"`
	Quantity int `json:"quantity"`
	Ingredients []string `json:"ingredients"`
    } `json:"rows"`
}