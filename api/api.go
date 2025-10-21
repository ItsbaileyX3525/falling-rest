package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type SeasonalFacts struct {
	Fact string `json:"fact"`
}

type ScientificFacts struct {
	Fact string `json:"fact"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var seasonalFacts = []SeasonalFacts{
	{Fact: "Fall occurs between summer and winter. In the Northern Hemisphere, it typically starts around September 22-23 and ends around December 21, while in the Southern Hemisphere, it runs from March to June."},
	{Fact: "The season begins with the autumnal equinox, when day and night are roughly equal in length."},
	{Fact: "Temperatures gradually cool down as the Earth tilts away from the Sun."},
	{Fact: "Trees shed leaves due to chlorophyll breakdown, which reveals other pigments like carotenoids (yellow/orange) and anthocyanins (red/purple)."},
	{Fact: "Fall is traditionally associated with harvesting crops, such as apples, pumpkins, and corn."},
	{Fact: "Many animals prepare for winter by storing food or migrating to warmer regions."},
	{Fact: "Fall weather is often variable, ranging from warm sunny days to cold, windy, or rainy conditions."},
	{Fact: "Many cultures celebrate harvest festivals or holidays like Thanksgiving, Halloween, and Mid-Autumn Festival."},
	{Fact: "Fall is a common season for seasonal allergies, especially due to ragweed pollen and mold spores from fallen leaves."},
}

func Season() []byte {
	testJson := seasonalFacts[rand.Intn(len(seasonalFacts))]

	jsonStr, err := json.Marshal(testJson)

	if err != nil {
		fmt.Println("Error occured!")

		jsonErr := Response{Success: false, Message: "Something went wrong"}
		jsonStr, _ := json.Marshal(jsonErr)
		return jsonStr
	}

	return jsonStr
}

func GenRandom() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

}
