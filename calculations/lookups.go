package Lookups

var Sourcing = map[string]float32{
	"Globally":   33,
	"Regionally": 67,
	"Locally":    100,
}

var ManufacturingMaterials = map[string]float32{
	"Virgin grade material":            40,
	"Part recycled, part virgin grade": 50,
	"Recycled material":                75,
	"Reused material":                  100,
	"Retained material":                100,
}

var EndOfLifeSustainability = map[string]float32{
	"Not recyclable":                     0,
	"Part recyclable":                    40,
	"Finish to remain after decomission": 90,
	"Fully recyclable":                   90,
	"Recycling program":                  90,
	"Able to be reused/repurposed":       100,
	"Take Back Scheme":                   100,
}

var CircularityAssessment = map[string]string{
	"Not recyclable":                     "Not recyclable",
	"Part recyclable":                    "Part recyclable",
	"Finish to remain after decomission": "Finish to remain",
	"Fully recyclable":                   "Recyclable",
	"Recycling program":                  "Recyclable",
	"Able to be reused/repurposed":       "Reusable",
	"Take Back Scheme":                   "Reusable",
}

var Certifications = map[string]float32{
	"Aluminum Stewardship Initiative": 100,
	"Cradle to Cradle":                100,
	"Declare":                         100,
	"Forest Stewardship Council (FSC) certification": 100,
	"Global Green Tag": 60,
	"Green Tick":       80,
	"Not Certified":    0,
	"Other":            60,
	"Programme for the Endorsement of Forest Certification (PEFC)": 100,
	"Smart Certified": 80,
	"Retained":        100,
}

var CategoryLookup = map[string]string{
	"Flooring":         "Core structure & finish",
	"Wall":             "Core structure & finish",
	"Ceiling":          "Core structure & finish",
	"Facade":           "Core structure & finish",
	"POS":              "Module units",
	"Product shelving": "Module units",
	"BOH":              "Module units",
	"Basin":            "Module units",
	"Furniture":        "Furniture & fittings",
	"Lighting":         "Furniture & fittings",
	"Fixtures":         "Furniture & fittings",
}

var SubCategoryContribution = map[string]float32{
	"Flooring":         0.125,
	"Wall":             0.125,
	"Ceiling":          0.125,
	"Facade":           0.125,
	"POS":              0.1,
	"Product shelving": 0.1,
	"BOH":              0.1,
	"Basin":            0.1,
	"Furniture":        0.05,
	"Lighting":         0.03,
	"Fixtures":         0.02,
}

var CategoryContribution = map[string]float32{
	"Core structure & finish": 0.5,
	"Module units":            0.4,
	"Furniture & fittings":    0.1,
}

var ScoreWeights = map[string]float32{
	"Manufacturer Location":  0.15,
	"Raw material":           0.30,
	"End of life assessment": 0.45,
	"Product Certification":  0.1,
}

var ManufacturingMaterialsLCA = map[string]float32{
	"Virgin grade material":            1,
	"Part recycled, part virgin grade": 1,
	"Recycled material":                0.5,
	"Reused material":                  0,
	"Retained material":                0,
}

var CarbonFactors = map[string]float32{
	"RigidTruck":       0.0004814,
	"ArticulatedTruck": 0.0001645,
	"Rail":             0.0000688,
	"Sea":              0.0000326,
	"Air":              0.0095791,
}

var TravelFactors = map[string]float32{
	"Locally": (50 * CarbonFactors["RigidTruck"]),
	"RegionallyRail": (15 * CarbonFactors["RigidTruck"]) +
		(150 * CarbonFactors["ArticulatedTruck"]) +
		(1000 * CarbonFactors["Rail"]),
	"RegionallyAir": (15 * CarbonFactors["RigidTruck"]) +
		(150 * CarbonFactors["ArticulatedTruck"]) +
		(1000 * CarbonFactors["Air"]),
	"GloballySea": (15 * CarbonFactors["RigidTruck"]) +
		(150 * CarbonFactors["ArticulatedTruck"]) +
		(10000 * CarbonFactors["Sea"]),
	"GloballyAir": (15 * CarbonFactors["RigidTruck"]) +
		(150 * CarbonFactors["ArticulatedTruck"]) +
		(10000 * CarbonFactors["Air"]),
}

var RetailEmissionsMultiplier = map[string]float32{
	"Low":    1.2,
	"Medium": 1,
	"High":   0.8,
}

var RetailStoreEmissions = map[string]float32{
	"Low":    200,
	"Medium": 400,
	"High":   600,
}
var RetailCounterEmissions = map[string]float32{
	"Low":    100,
	"Medium": 200,
	"High":   300,
}
