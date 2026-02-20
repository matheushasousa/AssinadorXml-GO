package handlers

import (
	"encoding/json"
	"event-publisher/internal/http/dto"
	"event-publisher/internal/rabbit"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func XmlRecebidoHandler(publisher *rabbit.Publisher) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var input dto.XmlRecebidoInputDTO
		json.NewDecoder(r.Body).Decode(&input)

		msg := map[string]interface{}{
			"eventId":       uuid.NewString(),
			"eventType":     "XML_RECEBIDO",
			"eventDateTime": time.Now().UTC(),
			"cnpj":          input.Cnpj,
			"chave":         input.Chave,
			"xml":           input.Xml,
		}

		publisher.Publish("xml.recebido", msg)

		w.WriteHeader(http.StatusOK)
	}
}
