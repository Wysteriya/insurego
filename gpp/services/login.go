package services

import (
	"baby-chain/blockchain/wallet"
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools"
	"baby-chain/tools/data"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	receiveObj := new(models.ReceiveLogin)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	master, ok := gpp.SD.Data["master"].(data.Data)
	if !ok {
		httpRes.Error(fmt.Errorf("master not found"))
		return
	}
	_, ok = master[receiveObj.PublicKey].(data.Data)
	if !ok {
		httpRes.Error(fmt.Errorf("user not registered"))
		return
	}

	sign := tools.HashB([]byte("signature"))
	signedHash, err := wallet.SignHash(receiveObj.PrivateKey, sign)
	if err != nil {
		httpRes.Error(err)
		return
	}
	if ok := wallet.VerifySignature(receiveObj.PublicKey, sign, signedHash); !ok {
		httpRes.Error(fmt.Errorf("invalid public-private key pair"))
		return
	}

	httpRes.Text("successfully logged in")
}
