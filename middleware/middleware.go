package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/benwang2/ru_dining_api/models"
)

func FetchMenuFromURI(URI string) []models.MenuSection {
	sectionPatt, _ := regexp.Compile("<p style=\"margin: 3px 0;\"><b>-+\\s([A-z\\s]+)\\s-")
	itemPatt, _ := regexp.Compile("<div class=\"col-1.*>(.*)</div>\\n\\s+<div.*/div>\\n\\s*.*href=\"(.*)\"")

	resp, err := http.Get(URI)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	html := string(body)

	// var menu []models.MenuItem
	var sections []models.MenuSection
	var sectionIndices []int

	sectionsMatches := sectionPatt.FindAllStringSubmatch(html, -1)

	if sectionsMatches == nil {
		return nil
	}

	for i := range sectionsMatches {
		var section models.MenuSection
		section.Name = sectionsMatches[i][1]

		sections = append(sections, section)
		sectionIndices = append(sectionIndices, strings.Index(html, sectionsMatches[i][1]))
		// fmt.Printf("Got section %s\n", section.Name)
	}

	itemMatches := itemPatt.FindAllStringSubmatch(html, -1)
	for i := range itemMatches {
		var item models.MenuItem
		item.Name = itemMatches[i][1]
		item.Info = itemMatches[i][2]

		index := strings.Index(html, item.Name)

		for j := 0; j < len(sections); j++ {
			if index > sectionIndices[j] && (j == len(sections)-1 || index < sectionIndices[j+1]) {
				sections[j].Items = append(sections[j].Items, item)
				// fmt.Printf("Added %s to %s\n", item.Name, sections[j].Name)
			}
		}
	}

	return sections
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	location := strings.ToLower(params.Get("location"))
	meal := strings.ToLower(params.Get("meal"))
	date := params.Get("date")

	w.Header().Set("Content-Type", "application/json")
	if location == "" || (location != "busch" && location != "livingston" && location != "collegeave" && location != "cookdouglass") {
		w.WriteHeader(http.StatusBadRequest)
		var e models.ErrorResponse
		e.StatusCode = 400
		e.Status = "Invalid location name."
		json.NewEncoder(w).Encode(e)
		return
	}

	if meal == "" || (meal != "lunch" && meal != "dinner") {
		w.WriteHeader(http.StatusBadRequest)
		var e models.ErrorResponse
		e.StatusCode = 400
		e.Status = "Invalid meal name."
		json.NewEncoder(w).Encode(e)
		return
	}

	baseURI := "http://menuportal.dining.rutgers.edu/FoodPro/pickmenu.asp"

	var locationNum string
	var mealName string

	if location == "collegeave" {
		locationNum = models.CollegeAve
	} else if location == "livingston" {
		locationNum = models.Livingston
	} else if location == "busch" {
		locationNum = models.Busch
	} else if location == "cookdouglass" {
		locationNum = models.CookDouglass
	}

	if meal == "lunch" {
		mealName = "Lunch"
	} else if meal == "dinner" {
		mealName = "Dinner"
	}

	newURI := baseURI + "?locationNum=" + locationNum + "&mealName=" + mealName

	if date != "" {
		newURI = newURI + "&dtdate=" + date
	}
	w.WriteHeader(http.StatusOK)
	var resp models.MenuResponse
	resp.StatusCode = 200
	resp.Status = "OK"
	resp.URL = newURI

	var sections []models.MenuSection = FetchMenuFromURI(newURI)

	resp.Menu = sections
	json.NewEncoder(w).Encode(resp)
}
