package services

import (
	"baby-chain/blockchain/wallet"
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools/data"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

func DeleteUserIns(ctx *gin.Context) {
	receiveObj := new(models.ReceiveDeleteUserIns)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	var b = gpp.BC.MineBlock("DeleteUserIns", data.Data{
		"public_key":    receiveObj.PublicKey,
		"policy_ref_id": receiveObj.PolicyRefId,
	})
	sign, err := wallet.SignHash(receiveObj.PrivateKey, b.Hash)
	if err != nil {
		httpRes.Error(err)
		return
	}
	b.Header["signature"] = hex.EncodeToString(sign)
	err = gpp.CSAlgo.Exec(&gpp.BC, &gpp.SD, b)
	if err != nil {
		httpRes.Error(err)
		return
	}
	err = gpp.BC.AddBlock(b)
	if err != nil {
		httpRes.Error(err)
		return
	}

	httpRes.Text("ok")
}
