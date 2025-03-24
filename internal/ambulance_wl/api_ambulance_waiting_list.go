/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xbartosp2@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ambulance_wl

import (
	"github.com/gin-gonic/gin"
)

type AmbulanceWaitingListAPI interface {


    // CreateWaitingListEntry Post /api/waiting-list/:ambulanceId/entries
    // Saves new entry into waiting list 
     CreateWaitingListEntry(c *gin.Context)

    // DeleteWaitingListEntry Delete /api/waiting-list/:ambulanceId/entries/:entryId
    // Deletes specific entry 
     DeleteWaitingListEntry(c *gin.Context)

    // GetWaitingListEntries Get /api/waiting-list/:ambulanceId/entries
    // Provides the ambulance waiting list 
     GetWaitingListEntries(c *gin.Context)

    // GetWaitingListEntry Get /api/waiting-list/:ambulanceId/entries/:entryId
    // Provides details about waiting list entry 
     GetWaitingListEntry(c *gin.Context)

    // UpdateWaitingListEntry Put /api/waiting-list/:ambulanceId/entries/:entryId
    // Updates specific entry 
     UpdateWaitingListEntry(c *gin.Context)

}