package main

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/kyokomi/emoji"
)

func main() {
	end := httpEndpoint{}
	end.Endpoint = "https://github.com"

	response, err := doHTTPEndpoint(end)
	if err != nil {
		fmt.Println(err)
	}

	if response.TLS != nil && len(response.TLS.PeerCertificates) != 0 {
		value := response.TLS.PeerCertificates[0].NotAfter
		days := daysToExpiry(value)

		expiryStatus := confirmCertExpiry(end.Endpoint, days)

		if expiryStatus != "" {
			fmt.Println(expiryStatus)
		} else {
			fmt.Println("Certificate valid for more than 4 months")
		}
	} else {
		fmt.Println("Site does not use HTTPS certificates.")
	}
}

type httpEndpoint struct {
	Name     string
	Endpoint string
}

func doHTTPEndpoint(endpoint httpEndpoint) (*http.Response, error) {
	body := ""
	return http.Post(endpoint.Endpoint, "application/json", bytes.NewBufferString(body))
}

func daysToExpiry(expiryDate time.Time) float64 {

	duration := expiryDate.Sub(time.Now())
	return math.Floor(duration.Hours() / 24)
}

func confirmCertExpiry(endpoint string, expiryDays float64) string {

	if expiryDays <= 40 {
		return emoji.Sprintf(":rotating_light: SSL certificate for %s expires withing 40 days!", endpoint)
	}
	if expiryDays <= 60 {
		return emoji.Sprintf(":rotating_light: SSL certificate for %s expires withing 60 days!", endpoint)
	}
	if expiryDays <= 80 {
		return emoji.Sprintf(":rotating_light: SSL certificate for %s expires withing 80 days!", endpoint)
	}
	if expiryDays <= 100 {
		return emoji.Sprintf(":rotating_light: SSL certificate for %s expires withing 100 days!", endpoint)
	}
	if expiryDays <= 120 {
		return emoji.Sprintf(":rotating_light: SSL certificate for %s expires withing 120 days!", endpoint)
	}
	return ""
}
