/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package processor

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/udr/internal/logger"
)

func (p *Processor) GetIdentityDataProcedure(c *gin.Context, collName string, ueId string) {
	logger.DataRepoLog.Debugf("Handle GetIdentityDataProcedure: %+v", ueId)
	filter := bson.M{
		"$or": []bson.M{
			{"gpsi": ueId},
			{"ueId": ueId},
		},
	}

	data, pd := p.GetDataFromDB(collName, filter)
	if pd != nil {
		logger.DataRepoLog.Errorf("GetIdentityDataProcedure err: %+v", pd)
		c.JSON(int(pd.Status), pd)
		return
	}

	res := models.IdentityData{}
	gpsi, ok := data["gpsi"]
	if ok {
		res.GpsiList = append(res.GpsiList, gpsi.(string))
	}

	ueID, ok := data["ueId"]
	if ok {
		res.SupiList = append(res.SupiList, ueID.(string))
	}
	c.JSON(http.StatusOK, res)
}
