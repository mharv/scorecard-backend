/*
 * scorecard
 *
 * inital pass at scorecard API
 *
 * API version: 1.0
 * Contact: mitchell.harvey@arup.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	Config "github.com/mharv/scorecard-backend/db"
)

// DeleteCommentsCommentId -
func DeleteCommentsCommentId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteMaterialTypesMaterialTypeId -
func DeleteMaterialTypesMaterialTypeId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteStoresStoreId -
func DeleteStoresStoreId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteUsersUserId -
func DeleteUsersUserId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetCommentsMaterialInstanceId - Your GET endpoint
func GetCommentsMaterialInstanceId(c *gin.Context) {
	id := c.Params.ByName("materialInstanceId")
	var allMaterialComments []Comment

	if result := Config.DB.Where("materialInstanceId = ?", id).Find(&allMaterialComments); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, allMaterialComments)
	}
}

// GetEpicMaterials - Your GET endpoint
func GetEpicMaterials(c *gin.Context) {
	var epicMaterials []EpicMaterial

	if result := Config.DB.Find(&epicMaterials); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, epicMaterials)
	}
}

// GetEpicMaterialsEpicMaterialId - Your GET endpoint
func GetEpicMaterialsEpicMaterialId(c *gin.Context) {
	id := c.Params.ByName("materialTypeId")
	var epicMaterial EpicMaterial

	if result := Config.DB.Where("id = ?", id).First(&epicMaterial); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, epicMaterial)
	}
}

// GetMaterialInstanceHistoryMaterialInstanceId - Your GET endpoint
func GetMaterialInstanceHistoryMaterialInstanceId(c *gin.Context) {
	id := c.Params.ByName("materialInstanceId")
	var materialInstance MaterialInstance
	var materialInstanceHistory []MaterialInstanceHistory

	type MaterialInstanceHistoryResponse struct {
		instance MaterialInstance
		history  []MaterialInstanceHistory
	}

	if result := Config.DB.Where("id = ?", id).Find(&materialInstance); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		if result := Config.DB.Where("materialInstanceId = ?", id).Find(&materialInstanceHistory); result.Error != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, MaterialInstanceHistoryResponse{materialInstance, materialInstanceHistory})
		}
	}
}

// GetMaterialInstancesStoreId - Your GET endpoint
func GetMaterialInstancesStoreId(c *gin.Context) {
	id := c.Params.ByName("storeId")
	var materialInstances []MaterialInstance

	if result := Config.DB.Where("storeId = ?", id).Find(&materialInstances); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, materialInstances)
	}
}

// GetMaterialTypes - Your GET endpoint
func GetMaterialTypes(c *gin.Context) {
	var materialTypes []MaterialType

	if result := Config.DB.Find(&materialTypes); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, materialTypes)
	}
}

// GetMaterialTypesMaterialTypeId - Your GET endpoint
func GetMaterialTypesMaterialTypeId(c *gin.Context) {
	id := c.Params.ByName("materialTypeId")
	var materialType MaterialType

	if result := Config.DB.Where("id = ?", id).First(&materialType); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, materialType)
	}
}

// GetStores - Your GET endpoint
func GetStores(c *gin.Context) {
	var stores []Store

	if result := Config.DB.Find(&stores); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, stores)
	}
}

// GetStoresStoreId - Your GET endpoint
func GetStoresStoreId(c *gin.Context) {
	id := c.Params.ByName("storeId")
	var store Store

	if result := Config.DB.Where("id = ?", id).First(&store); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, store)
	}
}

// GetUsers - Your GET endpoint
func GetUsers(c *gin.Context) {
	var users []User

	if result := Config.DB.Find(&users); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// GetUsersUserId - Get User Info by User ID
func GetUsersUserId(c *gin.Context) {
	id := c.Params.ByName("userId")
	var user User

	if result := Config.DB.Where("id = ?", id).First(&user); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// PostComment -
func PostComment(c *gin.Context) {
	var comment Comment
	c.BindJSON(&comment)

	comment.CommentDate = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	if result := Config.DB.Create(&comment); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, comment)
	}
}

// PostMaterialInstance -
func PostMaterialInstance(c *gin.Context) {
	var materialInstance MaterialInstance
	c.BindJSON(&materialInstance)

	if result := Config.DB.Create(&materialInstance); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, materialInstance)
	}
}

// PostMaterialType -
func PostMaterialType(c *gin.Context) {
	var materialType MaterialType
	c.BindJSON(&materialType)

	if result := Config.DB.Create(&materialType); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, materialType)
	}
}

// PostStore -
func PostStore(c *gin.Context) {
	var store Store
	c.BindJSON(&store)

	if result := Config.DB.Create(&store); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, store)
	}
}

// PostUser - Create New User
func PostUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	user.CreatedDate = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	if result := Config.DB.Create(&user); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// PutMaterialInstancesMaterialInstanceId -
func PutMaterialInstancesMaterialInstanceId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PutMaterialTypesMaterialTypeId -
func PutMaterialTypesMaterialTypeId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PutStoresStoreId -
func PutStoresStoreId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PutUsersUserId -
func PutUsersUserId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
