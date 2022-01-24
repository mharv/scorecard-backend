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

type MaterialInstance struct {
	Id int32 `json:"id,omitempty" gorm:"column:id"`

	MaterialTypeId int32 `json:"materialTypeId,omitempty" gorm:"column:materialTypeId"`

	MaterialStatus string `json:"materialStatus,omitempty" gorm:"column:materialStatus"`

	MaterialQuantity float32 `json:"materialQuantity,omitempty" gorm:"column:materialQuantity"`

	Unit string `json:"unit,omitempty" gorm:"column:unit"`

	CreatedDate string `json:"createdDate,omitempty" gorm:"column:createdDate"`

	SupplierLocation string `json:"supplierLocation,omitempty" gorm:"column:supplierLocation"`

	ManufacturerLocation string `json:"manufacturerLocation,omitempty" gorm:"column:manufacturerLocation"`

	SupplierName string `json:"supplierName,omitempty" gorm:"column:supplierName"`

	ManufacturerName string `json:"manufacturerName,omitempty" gorm:"column:manufacturerName"`

	OriginId int32 `json:"originId,omitempty" gorm:"column:originId"`

	StoreId int32 `json:"storeId,omitempty" gorm:"column:storeId"`

	RawMaterial string `json:"rawMaterial,omitempty" gorm:"column:rawMaterial"`

	EndOfLifeAssessment string `json:"endOfLifeAssessment,omitempty" gorm:"column:endOfLifeAssessment"`

	Category string `json:"category,omitempty" gorm:"column:category"`

	SubCategory string `json:"subCategory,omitempty" gorm:"column:subCategory"`

	ItemType string `json:"itemType,omitempty" gorm:"column:itemType"`

	ProductCertification string `json:"productCertification,omitempty" gorm:"column:productCertification"`

	RoadTransportKms float32 `json:"roadTransportKms,omitempty" gorm:"column:roadTransportKms"`

	RailTransportKms float32 `json:"railTransportKms,omitempty" gorm:"column:railTransportKms"`

	AirTransportKms float32 `json:"airTransportKms,omitempty" gorm:"column:airTransportKms"`

	OceanTransportKms float32 `json:"oceanTransportKms,omitempty" gorm:"column:oceanTransportKms"`
}

func (b *MaterialInstance) TableName() string {
	return "materialInstance"
}
