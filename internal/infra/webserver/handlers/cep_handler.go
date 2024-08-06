package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func SearchCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	viaCepChannel := make(chan string)
	brasilApiChannel := make(chan string)

	go makeRequest("ViaCEP", "https://viacep.com.br/ws/"+cep+"/json/", viaCepChannel)
	go makeRequest("BrasilAPI", "https://brasilapi.com.br/api/cep/v1/"+cep, brasilApiChannel)

	select {
	case msg := <-viaCepChannel:
		fmt.Println(msg)
	case msg := <-brasilApiChannel:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Erro de timeout")
	}
}

func makeRequest(apiName string, apiUrl string, resultChannel chan string) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resultChannel <- fmt.Sprintf("API vencedora: %s\nDados: %s", apiName, string(body))
}
