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

// EpicMaterial - Material metadata
type EpicMaterial struct {
	Id int32 `json:"id,omitempty" gorm:"column:id"`

	Category string `json:"category,omitempty" gorm:"column:category"`

	Type string `json:"type,omitempty" gorm:"column:type"`

	Material string `json:"material,omitempty" gorm:"column:material"`

	FunctionalUnit string `json:"functionalUnit,omitempty" gorm:"column:functionalUnit"`

	EmbodiedEnergy float32 `json:"embodiedEnergy,omitempty" gorm:"column:embodiedEnergy"`

	EmbodiedWater float32 `json:"embodiedWater,omitempty" gorm:"column:embodiedWater"`

	EmbodiedGHGE float32 `json:"embodiedGHGE,omitempty" gorm:"column:embodiedGHGE"`

	MoreInformationUrl string `json:"moreInformationUrl,omitempty" gorm:"column:moreInformationUrl"`
}

func (b *EpicMaterial) TableName() string {
	return "epicMaterial"
}
