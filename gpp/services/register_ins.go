package services

import (
	"baby-chain/blockchain/wallet"
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

func RegisterIns(ctx *gin.Context) {
	receiveObj := new(models.ReceiveRegisterIns)
	sendObj := new(models.SendRegisterIns)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	receiveObj.Data["public_key"] = receiveObj.PublicKey
	b := gpp.BC.MineBlock("RegisterIns", receiveObj.Data)
	sign, err := wallet.SignHash(receiveObj.PrivateKey, b.Hash)
	if err != nil {
		httpRes.Error(err)
		return
	}
	b.Header["signature"] = hex.EncodeToString(sign)
	if err := gpp.CSAlgo.Exec(&gpp.BC, &gpp.SD, b); err != nil {
		httpRes.Error(err)
		return
	}
	err = gpp.BC.AddBlock(b)
	if err != nil {
		httpRes.Error(err)
		return
	}

	sendObj.PolicyRefId = hex.EncodeToString(b.Hash[:])
	httpRes.SendJson(sendObj)
}
