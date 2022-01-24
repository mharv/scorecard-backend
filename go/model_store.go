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

// Store - Store metadata for each individual Aesop store
type Store struct {
	Id int32 `json:"id,omitempty" gorm:"column:id"`

	StoreName string `json:"storeName,omitempty" gorm:"column:storeName"`

	Country string `json:"country,omitempty" gorm:"column:country"`

	ScorecardType string `json:"scorecardType,omitempty" gorm:"column:scorecardType"`

	LocationType string `json:"locationType,omitempty" gorm:"column:locationType"`

	RoicCapex string `json:"roicCapex,omitempty" gorm:"column:roicCapex"`

	Currency string `json:"currency,omitempty" gorm:"column:currency"`

	LeaseTermInYears float32 `json:"leaseTermInYears,omitempty" gorm:"column:leaseTermInYears"`

	ProjectManagerId int32 `json:"projectManagerId,omitempty" gorm:"column:projectManagerId"`

	ArchitectId int32 `json:"architectId,omitempty" gorm:"column:architectId"`

	ContractorId int32 `json:"contractorId,omitempty" gorm:"column:contractorId"`

	SignedCoc bool `json:"signedCoc,omitempty" gorm:"column:signedCoc"`

	SignedNda bool `json:"signedNda,omitempty" gorm:"column:signedNda"`

	TotalSqm float32 `json:"totalSqm,omitempty" gorm:"column:totalSqm"`

	RetailSqm float32 `json:"retailSqm,omitempty" gorm:"column:retailSqm"`

	BohSqm float32 `json:"bohSqm,omitempty" gorm:"column:bohSqm"`

	MarketingMediaType string `json:"marketingMediaType,omitempty" gorm:"column:marketingMediaType"`

	PosSqm float32 `json:"posSqm,omitempty" gorm:"column:posSqm"`

	MposSqm float32 `json:"mposSqm,omitempty" gorm:"column:mposSqm"`

	TapNumber int32 `json:"tapNumber,omitempty" gorm:"column:tapNumber"`

	TreatmentSpace bool `json:"treatmentSpace,omitempty" gorm:"column:treatmentSpace"`

	Bathroom bool `json:"bathroom,omitempty" gorm:"column:bathroom"`

	OtherSpace string `json:"otherSpace,omitempty" gorm:"column:otherSpace"`

	ReviewerId int32 `json:"reviewerId,omitempty" gorm:"column:reviewerId"`

	ReviewRequired int32 `json:"reviewRequired,omitempty" gorm:"column:reviewRequired"`

	Region string `json:"region,omitempty" gorm:"column:region"`

	SusInitGreenspace bool `json:"susInitGreenspace,omitempty" gorm:"column:susInitGreenspace"`

	SusInitRainwater bool `json:"susInitRainwater,omitempty" gorm:"column:susInitRainwater"`

	SusInitSolarPV bool `json:"susInitSolarPV,omitempty" gorm:"column:susInitSolarPV"`

	SusInitGreywater bool `json:"susInitGreywater,omitempty" gorm:"column:susInitGreywater"`

	SusInitBattery bool `json:"susInitBattery,omitempty" gorm:"column:susInitBattery"`

	SusInitVentilation bool `json:"susInitVentilation,omitempty" gorm:"column:susInitVentilation"`

	SusInitComposting bool `json:"susInitComposting,omitempty" gorm:"column:susInitComposting"`
}

func (b *Store) TableName() string {
	return "store"
}
