package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	cg := coingecko.NewClient(httpClient)
	ids := []string{"cardano", "world-mobile-token", "muesliswap-milk", "sundaeswap", "hosky"}
	vc := []string{"eur"}
	sp, err := cg.SimplePrice(ids, vc)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(vc); j++ {
			fmt.Println(fmt.Sprintf("%s is worth %f %s", ids[i], (*sp)[ids[i]][vc[j]], vc[j]))

		}
	}
}
