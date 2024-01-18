package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/Xacor/concurency-vulns/service"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()

	discountCodes := generateCodes()
	l.Infof("codes: %+v", discountCodes)

	discountWorker := service.NewDiscountWorker(discountCodes, l)
	l.Info("listening on :8080")
	http.HandleFunc(`/api/redeem-code`, discountWorker.RedeemCode)

	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		l.Panic(err)
	}
}

func generateCodes() map[string]bool {
	result := make(map[string]bool, 5)
	for i := 0; i < 5; i++ {
		result[randomString()] = false
	}

	return result
}

func randomString() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)

}
