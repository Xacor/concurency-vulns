package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/Xacor/concurency-vulns/service"
	"github.com/sirupsen/logrus"
)

func main() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	discountCodes := generateCodes()
	logrus.Infof("codes: %v", discountCodes)

	discountWorker := service.NewDiscountWorker(discountCodes)

	http.HandleFunc(`/api/redeem-code`, discountWorker.RedeemCode)

	logrus.Info("listening on :8080")
	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		logrus.Panic(err)
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
