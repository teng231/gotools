package common

import (
	"log"
	"testing"
)

// https://apia.urbox.dev/api/internal/voucher/validate-receive?client_id=701ee755-ed45-4591-b4c5-e6cf208ff767&timestamp=1656398486&signature=2762bab9cc81efbc90c845c398a5f89b597e18315ca2fa7666094ffdc8ae4bac&receive_code=Pk_6iJdHpZ-8CGdsdq-B9-0iiGiBR27KwTGFD_-LpTxX0hTR42zZdOxR2oP-XiHu

func TestSig(t *testing.T) {
	ts := 1656398486
	ins := map[string]interface{}{
		"client_id":                            "secret",
		"701ee755-ed45-4591-b4c5-e6cf208ff767": "upgOsFXumJcUUkli",
	}
	for clientId, secret := range ins {
		// log.Print(secret, clientId, ts)
		out := Hash(secret, clientId, ts)
		log.Print("-->", out)
	}
}
