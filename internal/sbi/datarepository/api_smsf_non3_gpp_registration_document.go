/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/udr/internal/logger"
	"github.com/free5gc/udr/internal/sbi/producer"
	"github.com/free5gc/util/httpwrapper"
)

// HTTPCreateSmsfContextNon3gpp - Create the SMSF context data of a UE via non-3GPP access
func HTTPCreateSmsfContextNon3gpp(c *gin.Context) {
	auth_err := authorizationCheck(c)
	if auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err.Error()})
		return
	}

	var smsfRegistration models.SmsfRegistration

	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.DataRepoLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&smsfRegistration, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.DataRepoLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := httpwrapper.NewRequest(c.Request, smsfRegistration)
	req.Params["ueId"] = c.Params.ByName("ueId")

	rsp := producer.HandleCreateSmsfContextNon3gpp(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// HTTPDeleteSmsfContextNon3gpp - To remove the SMSF context data of a UE via non-3GPP access
func HTTPDeleteSmsfContextNon3gpp(c *gin.Context) {
	auth_err := authorizationCheck(c)
	if auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err.Error()})
		return
	}

	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	rsp := producer.HandleDeleteSmsfContextNon3gpp(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// HTTPQuerySmsfContextNon3gpp - Retrieves the SMSF context data of a UE using non-3gpp access
func HTTPQuerySmsfContextNon3gpp(c *gin.Context) {
	auth_err := authorizationCheck(c)
	if auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err.Error()})
		return
	}

	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	rsp := producer.HandleQuerySmsfContextNon3gpp(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
