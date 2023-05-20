package services

import (
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools/data"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Status(ctx *gin.Context) {
	receiveObj := new(models.ReceiveStatus)
	sendObj := new(models.SendStatus)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	nodes, ok := gpp.SD.Data["Nodes"].(data.Data)
	if !ok {
		httpRes.Error(fmt.Errorf("no nodes"))
		return
	}
	node, ok := nodes[receiveObj.UserId].(data.Data)
	if !ok {
		httpRes.Error(fmt.Errorf("bad user id"))
		return
	}
	policies, ok := node["Policies"].(data.Array)
	if !ok {
		httpRes.Text("no policies")
		return
	}
	ins, ok := gpp.SD.Data["ins"].(data.Data)
	if !ok {
		httpRes.Error(fmt.Errorf("no ins"))
		return
	}
	insList := make(data.Array, 0)
	for _, policy := range policies {
		insData := ins[policy.(string)].(data.Data)
		insList = append(insList, insData)
	}

	sendObj.Array = insList
	sendObj.Info = node
	httpRes.SendJson(sendObj)
}
