package services

import (
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools/data"
	"github.com/gin-gonic/gin"
)

func InsHavs(ctx *gin.Context) {
	receiveObj := new(models.ReceiveINSHavs)
	sendObj := new(models.SendINSHavs)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	ins := gpp.SD.Data["ins"].(data.Data)
	buyers := ins[receiveObj.PolicyRefId].(data.Data)["Buyer"].(data.Array)
	sendObj.PubKeys = buyers

	httpRes.SendJson(sendObj)
}
