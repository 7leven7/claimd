package services

import (
	"fmt"
	"github.com/7leven7/claimd/app/configs"
	"github.com/7leven7/claimd/app/db"
	"github.com/7leven7/claimd/app/models"
	"github.com/gocolly/colly"
	"net/http"
	"strings"
	"time"
)

// ProfileData is the struct for the scraped data
type ProfileData struct {
	Profile   string `json:"profile"`
	Username  string `json:"username"`
	Followers string `json:"follower"`
	Following string `json:"following"`
	Likes     string `json:"likes"`
	PostCount string `json:"postcount"`
}

// Scrape calls the Handler function to scrape the platforms
func Scrape() error {
	platforms, err := GetAllPlatforms()
	if err != nil {
		return fmt.Errorf("Error retrieving platforms: %v", err)
	}
	for _, platform := range platforms {
		err = Handler(platform)
		if err != nil {
			return err
		}
	}
	return nil
}

// Handler handles the scraping of the platforms
func Handler(platform models.Platform) error {

	fmt.Println(platform.Name)
	data := ProfileData{}

	modelName := platform.Name
	profile := platform.Profile
	username := platform.Name
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"))

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
	})
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.SetRequestTimeout(120 * time.Second)

	if modelName == "instagram" {
		c.OnHTML("html", func(e *colly.HTMLElement) {

			data = *InstagramSelectors(e)

		})
	} else if modelName == "tiktok" {
		c.OnHTML("h2.tiktok-7k173h-H2CountInfos.e1457k4r1", func(e *colly.HTMLElement) {
			data = *TiktokSelectors(e)
		})
		c.OnHTML("div.tiktok-yvmafn-DivVideoFeedV2.ecyq5ls0", func(e *colly.HTMLElement) {
			data.PostCount = TiktokPostCounter(e).PostCount
		})
	}

	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.Visit(profile)

	followers := data.Followers
	following := data.Following
	likes := data.Likes
	posts := data.PostCount

	if modelName == "instagram" {
		platform := models.Instagram{
			Profile:   profile,
			Username:  username,
			Followers: followers,
			Following: following,
			Likes:     likes,
			PostCount: posts,
		}
		err := InsertPlatform(modelName, platform.Profile, platform.Username, platform.Followers, platform.Following, platform.Likes, platform.PostCount)
		if err != nil {
			return err
		}
	} else if modelName == "tiktok" {
		platform := models.Tiktok{
			Profile:   profile,
			Username:  username,
			Followers: followers,
			Following: following,
			Likes:     likes,
			PostCount: posts,
		}
		err := InsertPlatform(modelName, platform.Profile, platform.Username, platform.Followers, platform.Following, platform.Likes, platform.PostCount)
		if err != nil {
			return err
		}
	}
	return nil
}

// InstagramSelectors scrapes the instagram profile for followers, following and likes
func InstagramSelectors(e *colly.HTMLElement) (data *ProfileData) {
	meta := strings.Split(e.ChildAttr("meta[name=\"description\"]", "content"), "\n")[0]
	parts := strings.Split(meta, ", ")

	data = &ProfileData{
		Followers: strings.Fields(parts[0])[0],
		Following: strings.Fields(parts[1])[0],
		PostCount: strings.Fields(parts[2])[0],
	}

	return data
}

// TiktokSelectors scrapes the tiktok profile for followers, following and likes
func TiktokSelectors(e *colly.HTMLElement) (data *ProfileData) {
	data = &ProfileData{}
	e.ForEach("strong", func(_ int, e *colly.HTMLElement) {
		if e.Attr("title") == "Followers" {
			data.Followers = e.Text
		}
		if e.Attr("title") == "Following" {
			data.Following = e.Text
		}
		if e.Attr("title") == "Likes" {
			data.Likes = e.Text
		}
	})
	return data
}

// TiktokPostCounter counts the number of posts on the tiktok profile
func TiktokPostCounter(e *colly.HTMLElement) (data *ProfileData) {
	data = &ProfileData{}
	count := 0
	e.ForEach("div.tiktok-x6y88p-DivItemContainerV2.e19c29qe7", func(_ int, e *colly.HTMLElement) {
		count++
	})
	data.PostCount = fmt.Sprintf("%d", count)

	return data
}

// GetAllPlatforms returns all platforms from the database
func GetAllPlatforms() ([]models.Platform, error) {
	var platforms []models.Platform

	err := configs.PGConnection()
	if err != nil {
		return nil, err
	}

	err = db.DB.Find(&platforms).Error
	if err != nil {
		return nil, err
	}
	return platforms, nil
}

// InsertPlatform inserts a new platform to the database
func InsertPlatform(platformName, profile string, username string, followers string, following string, likes string, posts string) error {

	err := configs.PGConnection()
	if err != nil {
		return err
	}
	switch platformName {
	case "instagram":
		platform := models.Instagram{
			Profile:   profile,
			Username:  username,
			Followers: followers,
			Following: following,
			Likes:     likes,
			PostCount: posts,
		}
		err := db.DB.Where("profile = ?", &platform.Profile).FirstOrCreate(&platform).Error
		if err != nil {
			return err
		}
		err = db.DB.Save(&platform).Error
		if err != nil {
			return err
		}
		return nil
	case "tiktok":
		platform := &models.Tiktok{
			Profile:   profile,
			Username:  username,
			Followers: followers,
			Following: following,
			Likes:     likes,
			PostCount: posts,
		}
		err := db.DB.Where("profile = ?", platform.Profile).FirstOrCreate(&platform).Error
		if err != nil {
			return err
		}
		err = db.DB.Save(&platform).Error
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unsupported platform: %s", platformName)
	}
}
