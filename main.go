// All in one file because It's just a small project.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"github.com/reujab/wallpaper"
	"github.com/skratchdot/open-golang/open"
)

type RedditAPIResponse struct {
	Data struct {
		Children []struct {
			Data struct {
				Title     string `json:"title"`
				Author    string `json:"author"`
				URL       string `json:"url"`
				Permalink string `json:"permalink"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

var currentWallpaperURL string
var httpClient = &http.Client{Timeout: 10 * time.Second}

func getRequest(url string) (io.ReadCloser, error) {
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"User-Agent": {"walp"},
		"Accept":     {"*/*"},
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

func fetchWallpapers() (RedditAPIResponse, error) {
	var redditResponse RedditAPIResponse
	body, err := getRequest("https://www.reddit.com/r/wallpaper/top.json?t=month&limit=100")
	if err != nil {
		return redditResponse, err
	}
	defer body.Close()
	err = json.NewDecoder(body).Decode(&redditResponse)
	if err != nil {
		return redditResponse, err
	}
	return redditResponse, nil
}

func saveWallpaper(id int, url string) (path string, err error) {
	// fetching the wallpaper
	resp, err := getRequest(url)
	if err != nil {
		return "", err
	}

	// saving the wallpaper to the cache directory
	cache, err := os.UserCacheDir()
	filename := fmt.Sprintf("wallpaper%d.%s", id, strings.Split(url, ".")[3])
	path = filepath.Join(cache, filename)
	println(path)
	out, err := os.Create(path)
	if err != nil {
		return path, err
	}
	defer out.Close()
	_, err = io.Copy(out, resp)
	if err != nil {
		return path, err
	}
	return path, nil
}

func GetAndSetWallpaper(desc *systray.MenuItem) {
	systray.SetTitle("Loading...")
	desc.SetTitle("Loading...")
	desc.Disable()

	// Fetching the top 100 posts this month from r/wallpaper
	redditResponse, err := fetchWallpapers()
	if err != nil {
		systray.SetTitle("Error")
		desc.SetTitle("Error - Could not fetch wallpapers.")
		return
	}

	// Randomly selecting a wallpaper from the top 20 wallpapers
	randomWallpaper := redditResponse.Data.Children[rand.Intn(99)].Data

	/*
		Downloading the wallpaper because the wallpaper library saves all the wallpapers
		we setFromURL in the same name, which apperantly doesn't really work as intended
		with macOS when changing the wallpaper multiple times.
	*/
	path, err := saveWallpaper(rand.Intn(100), randomWallpaper.URL)
	if err != nil {
		systray.SetTitle("Error")
		desc.SetTitle("Error - Could not save the wallpaper.")
		return
	}

	err = wallpaper.SetFromFile(path)
	if err == nil {
		systray.SetTitle("Walp")

		// Giving credit to the author of the wallpaper and linking to the post
		desc.SetTitle(randomWallpaper.Title)
		desc.SetTooltip(fmt.Sprintf("%s by %s", randomWallpaper.Title, randomWallpaper.Author))
		desc.Enable()
		currentWallpaperURL = "https://www.reddit.com" + randomWallpaper.Permalink

		// Deleting the wallpaper from the cache directory after setting it to not clutter the cache directory
		time.Sleep(3 * time.Second)
		os.Remove(path)
	}

}

func main() {
	rand.Seed(time.Now().UnixMilli())
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Walp")
	desc := systray.AddMenuItem("r/wallpaper", "r/wallpaper")
	desc.Disable()
	systray.AddSeparator()
	changeWall := systray.AddMenuItem("Get New Wallpaper", "Click to fetch a new wallpaper from r/wallpaper")
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	for {
		select {
		case <-changeWall.ClickedCh:
			GetAndSetWallpaper(desc)
			break
		case <-mQuitOrig.ClickedCh:
			systray.Quit()
			break
		case <-desc.ClickedCh:
			if currentWallpaperURL != "" {
				_ = open.Run(currentWallpaperURL)
			}
		}

	}
}

func onExit() {
	return
}
