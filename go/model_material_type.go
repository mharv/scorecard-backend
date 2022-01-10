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

type MaterialType struct {
	Id int32 `json:"id,omitempty" gorm:"column:id"`

	Category string `json:"category,omitempty" gorm:"column:category"`

	SubCategory string `json:"subCategory,omitempty" gorm:"column:subCategory"`

	ItemType string `json:"itemType,omitempty" gorm:"column:itemType"`

	RawMaterial string `json:"rawMaterial,omitempty" gorm:"column:rawMaterial"`

	EndOfLifeAssessment string `json:"endOfLifeAssessment,omitempty" gorm:"column:endOfLifeAssessment"`

	ProductCertification string `json:"productCertification,omitempty" gorm:"column:productCertification"`
}
