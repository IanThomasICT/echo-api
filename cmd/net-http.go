package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// type CatFact struct {
// 	Fact   string `json:"fact"`
// 	Length int64  `json:"length"`
// }

// type ResponseObj struct {
// 	Ok	bool `json:"ok"`
// 	Data interface{} `json:"data"`
// }

func main() {
	catFactHandler := func(w http.ResponseWriter, r *http.Request) {
		
		w.Header().Set("Content-Type", "application/json");
		var res ResponseObj 

		resp, err := http.Get("https://catfact.ninja/fa")
		if (err != nil) {
			log.Fatalln(err)
			json.NewEncoder(w).Encode(res);
		}

		defer resp.Body.Close()

		out, err := io.ReadAll(resp.Body)
		if (err != nil) {
			log.Fatalln("Failed to read the request body from CatFacts", err)
			json.NewEncoder(w).Encode(res);
		}
		
		var body CatFact
		if err := json.Unmarshal(out, &body); err != nil {
			log.Fatalln(err)
			json.NewEncoder(w).Encode(res);
		}
		
		res.Data = body;
		resObj, err := json.Marshal(res);
		if err != nil {
			log.Fatalln(err)
			json.NewEncoder(w).Encode(res);
		}
		w.Write(resObj);
	}

	http.HandleFunc("/cats", catFactHandler)

	server := &http.Server{
		Addr: ":8080",
	}

	log.Fatal(server.ListenAndServe())
}
