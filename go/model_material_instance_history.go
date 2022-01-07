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

type MaterialInstanceHistory struct {

	Id int32 `json:"id,omitempty"`

	MaterialInstanceId int32 `json:"materialInstanceId,omitempty"`

	SnapshotDate string `json:"snapshotDate,omitempty"`

	CreatedDate string `json:"createdDate,omitempty"`

	Unit string `json:"unit,omitempty"`

	MaterialQuantity float32 `json:"materialQuantity,omitempty"`

	MaterialStatus string `json:"materialStatus,omitempty"`

	OriginId int32 `json:"originId,omitempty"`
}