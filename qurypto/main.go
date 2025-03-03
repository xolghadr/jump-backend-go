package qurypto

import (
	"fmt"
	"io"
	"net/http"
)

func GetExchangeRate(source, destination string) (string, error) {
	if destination == "" {
		destination = "rls"
	}

	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:4001/rates?srcCurrency=%s&dstCurrency=%s", source, destination)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
