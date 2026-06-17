package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string  `jsoon:"userID"`
	PackageSlug       string  `jsoon:"packageslugs"`
	TotalPriceinCents float64 `jsoon:"totalpriceincents"`
}
