package service

import (
	"bytes"
	"net/http"
	"time"

	"github.com/Xacor/concurency-vulns/models"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

type DiscountWorker struct {
	codes  map[string]bool
	logger *logrus.Logger
}

func NewDiscountWorker(codes map[string]bool, logger *logrus.Logger) *DiscountWorker {
	return &DiscountWorker{
		codes:  codes,
		logger: logger,
	}
}

func (d *DiscountWorker) RedeemCode(w http.ResponseWriter, r *http.Request) {
	// чтение тела запроса
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		d.logger.Error("failed to read body", err)
		return
	}

	// десериализация json в структуру
	var data models.RedeemRequest
	err := jsoniter.Unmarshal(buf.Bytes(), &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		d.logger.Error("failed to unmarshal request", err)
		return
	}

	// проверка валидности кода и подготовка тела ответа
	resp := models.RedeemResponse{Status: models.Fail}
	if ok := d.checkCode(data.DiscountCode); ok {
		resp.Status = models.Success
	}

	// сериализация ответа в json
	bytes, err := jsoniter.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		d.logger.Error("failed to marshal response", err)
		return
	}

	// отправка ответа
	w.Write(bytes)

	time.Sleep(time.Second * 5)

	// отметка, что промокод был использован
	d.codes[data.DiscountCode] = true
}

func (d *DiscountWorker) checkCode(code string) bool {
	// если код на скидку отсутвует или уже использован
	if redeemed, ok := d.codes[code]; !ok || redeemed {
		return false
	}

	return true
}
