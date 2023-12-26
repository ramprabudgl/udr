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

// HTTPApplicationDataInfluenceDataSubsToNotifySubscriptionIdDelete -
func HTTPApplicationDataInfluenceDataSubsToNotifySubscriptionIdDelete(c *gin.Context) {
	auth_err := authorizationCheck(c)
	if auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err.Error()})
		return
	}

	// New HTTP request
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	// Handle request
	rsp := producer.HandleApplicationDataInfluenceDataSubsToNotifySubscriptionIdDelete(req)

	if rsp.Body != nil {
		// Serialize response body
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
	} else {
		c.Data(rsp.Status, "", nil)
	}
}

// HTTPApplicationDataInfluenceDataSubsToNotifySubscriptionIdGet -
func HTTPApplicationDataInfluenceDataSubsToNotifySubscriptionIdGet(c *gin.Context) {
	auth_err := authorizationCheck(c)
	if auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err.Error()})
		return
	}

	// New HTTP request
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	// Handle request
	rsp := producer.HandleApplicationDataInfluenceDataSubsToNotifySubscriptionIdGet(req)

	// Serialize response body
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

// HTTPApplicationDataInfluenceDataSubsToNotifySubscriptionIdPut -
func HTTPApplicationDataInfluenceDataSubsToNotifySubscriptionIdPut(c *gin.Context) {
	auth_err := authorizationCheck(c)
	if auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err.Error()})
		return
	}

	// Get HTTP request body
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

	// Deserialize request body
	var trafficInfluSub models.TrafficInfluSub
	err = openapi.Deserialize(&trafficInfluSub, requestBody, "application/json")
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

	// New HTTP request
	req := httpwrapper.NewRequest(c.Request, trafficInfluSub)
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	// Handle request
	rsp := producer.HandleApplicationDataInfluenceDataSubsToNotifySubscriptionIdPut(req)

	// Serialize response body
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
