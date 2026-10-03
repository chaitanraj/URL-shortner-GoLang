package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	// "net/http"
	"time"
)

type URL struct{
	ID string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL string `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

/*
d9736711 ---> {
			ID:"d9736711",
			OriginalURL: "https://github.com/chaitanraj",
			ShortURL: "d9736711",
			CreationDate: time.Now()
}
*/

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string{
	hasher := md5.New();
	hasher.Write([]byte(OriginalURL));
	fmt.Println("hasher: ",hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data: ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("EncodeToString: ", hash)
	fmt.Println("final String: ", hash[:8])
	return "https://github.com/chaitanraj"
}

func storeURL(originalURL string){
	shortURL := generateShortURL(originalURL)
	id := shortURL //use ths short url as id for simplicity

	urlDB[id]=URL{
		ID : id,
		OriginalURL: originalURL,
		ShortURL: shortURL,
		CreationDate: time.Now(),
	}

	return shortURL
}

func main(){
	fmt.Println("Url shortner listening");
	// http.ListenAndServe(":9999",nil);
	OriginalURL := "https://github.com/chaitanraj"
	generateShortURL(OriginalURL)

}