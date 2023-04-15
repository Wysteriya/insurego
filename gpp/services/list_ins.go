package services

import (
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools/data"
	"github.com/gin-gonic/gin"
)

func INSList(ctx *gin.Context) {
	sendObj := new(models.SendInsList)
	httpRes := gpp.NewHttpResponse(ctx)

	ins, ok := gpp.SD.Data["ins"].(data.Data)
	if !ok {
		httpRes.SendJson(sendObj)
		return
	}
	sendObj.Data = ins

	httpRes.SendJson(sendObj)
}
