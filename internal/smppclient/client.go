package smppclient

import (
	"log"
	"time"

	"github.com/fiorix/go-smpp/smpp"
	"github.com/fiorix/go-smpp/smpp/pdu"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutext"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
)

// Connect établit une connexion SMPP, envoie un message et écoute les DLR.
func Connect() error {
	tx := &smpp.Transceiver{
		Addr:   "localhost:2775",   // Remplace par l'adresse de ton SMSC
		User:   "smppuser",         // Nom d'utilisateur
		Passwd: "smpppass",         // Mot de passe
		Handler: func(p pdu.Body) {
			log.Printf("PDU reçu: %T => %+v\n", p, p)

			if p.Header().ID == pdu.DeliverSMID {
				sm := p.Fields()
				esmClass, ok := sm[pdufield.ESMClass]
				if ok && esmClass.Bytes()[0]&0x04 != 0 {
					// C’est un DLR
					log.Printf("DLR reçu : %s", sm[pdufield.ShortMessage].String())
					// Ici, ajoute le traitement ou mise à jour dans DB
				}
			}
		},
	}

	// Connexion
	client := tx.Bind()
	select {
	case status := <-client:
		if status.Error() != nil {
			return status.Error()
		}
		log.Println("Connexion SMPP établie")
	case <-time.After(5 * time.Second):
		return &TimeoutError{"Connexion SMPP expirée"}
	}

	// Envoi d'un message simple
	msg := &smpp.ShortMessage{
		Src:      "12345",                                 // Expéditeur
		Dst:      "2250701234567",                         // Destinataire
		Text:     pdutext.Latin1("Bonjour depuis Go SMPP"),
		Register: smpp.FinalDeliveryReceipt,               // Pour recevoir un DLR
	}

	resp, err := tx.Submit(msg)
	if err != nil {
		return err
	}
	log.Printf("Message envoyé, réponse: %+v", resp)

	time.Sleep(5 * time.Second) // On attend les éventuelles DLR

	tx.Close()
	return nil
}

// TimeoutError gère les erreurs de délai de connexion
type TimeoutError struct {
	Message string
}

func (e *TimeoutError) Error() string {
	return e.Message
}
