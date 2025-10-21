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

var scientificFacts = []ScientificFacts{
	{Fact: "All objects fall at the same rate in a vacuum. Without air resistance, a feather and a bowling ball fall at exactly the same rate — 9.81m/s^2 on Earth. Galileo demonstrated this principle, later confirmed by Apollo astronauts on the Moon."},
	{Fact: "Air resistance limits your speed (terminal velocity). When falling through air, drag counteracts gravity until forces balance. For a human in a belly-down position: ~55 m/s (≈200 km/h). For a skydiver head-down: ~90 m/s (≈320 km/h). A small insect might never reach lethal speed due to low mass and high drag."},
	{Fact: "Acceleration remains constant, speed does not. While falling freely, you accelerate at g = 9.81 m/s^2, meaning your velocity increases by ~9.81 m/s every second until drag balances gravity."},
	{Fact: "Microgravity (like in orbit) isn't zero gravity — it's continuous falling. Astronauts appear weightless because they're constantly falling toward Earth but moving forward fast enough to keep missing it — this is what an orbit actually is."},
	{Fact: "Your brain often can't tell the difference between free fall and weightlessness. Ignoring drag, the formula d = 1/2gt^2 tells you that if you fall for 2 seconds, you'll drop 4 times farther than in 1 second."},
	{Fact: "Your brain often can't tell the difference between free fall and weightlessness. The vestibular system senses acceleration changes — during a fall, it detects the rapid downward acceleration, triggering panic or adrenaline."},
	{Fact: "Your brain often can't tell the difference between free fall and weightlessness. That's why astronauts and people in zero-g planes experience the same “floating” feeling you get in the first instant of a fall."},
	{Fact: "Time feels slower during a fall — but it isn't. Under stress, your brain records more memories per second, creating the illusion that time is slowing down, even though external time is constant."},
	{Fact: "The 'stomach drop' feeling isn't your stomach moving. It's your internal organs lagging slightly behind the rest of your body as you accelerate downward — essentially inertia at work inside you."},
	{Fact: "Cats survive falls from higher places better than medium heights."},
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
