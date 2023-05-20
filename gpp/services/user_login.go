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

func UserLogin(ctx *gin.Context) {
	receiveObj := new(models.ReceiveLogin)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	nodes, ok := gpp.SD.Data["Nodes"].(data.Data)
	if !ok {
		httpRes.Error(fmt.Errorf("nodes not found"))
		return
	}
	_, ok = nodes[receiveObj.PublicKey].(data.Data)
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
