package seeders

import (
	"github.com/7leven7/claimd/app/configs"
	"github.com/7leven7/claimd/app/db"
	"github.com/7leven7/claimd/app/models"
)

// PlatformsSeeder seeds the platforms table
func PlatformsSeeder() {
	platforms := []models.Platform{
		{
			Name:    "instagram",
			Profile: "https://www.instagram.com/utopiantravel/",
		},
		{
			Name:    "instagram",
			Profile: "https://www.instagram.com/milan_milanovic1997/",
		},
		{
			Name:    "instagram",
			Profile: "https://www.instagram.com/prodaja_telefona/",
		},
		{
			Name:    "instagram",
			Profile: "https://www.instagram.com/lazybagprodaja/",
		},
		{
			Name:    "instagram",
			Profile: "https://www.instagram.com/prodaja_tehnike/",
		},
		{
			Name:    "tiktok",
			Profile: "https://www.tiktok.com/@florablue_",
		},
		{
			Name:    "tiktok",
			Profile: "https://www.tiktok.com/@miroljub.teokratija",
		},
		{
			Name:    "tiktok",
			Profile: "https://www.tiktok.com/@factslobby",
		},
		{
			Name:    "tiktok",
			Profile: "https://www.tiktok.com/@swiss_beautiful",
		},
		{
			Name:    "tiktok",
			Profile: "https://www.tiktok.com/@krazymations",
		},
	}
	err := configs.PGConnection()
	if err != nil {
		return
	}
	for _, platform := range platforms {
		db.DB.Create(&platform)
	}
}
