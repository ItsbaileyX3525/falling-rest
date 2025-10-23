package api

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

type Facts struct {
	Fact string `json:"fact"`
}

type Images struct {
	ImageUrl string `json:"imageUrl"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var seasonalFacts = []Facts{
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

var scientificFacts = []Facts{
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

var leavesImages = []Images{
	{ImageUrl: "/assets/images/leaves/leaves1.webp"},
	{ImageUrl: "/assets/images/leaves/leaves2.jpg"},
	{ImageUrl: "/assets/images/leaves/leaves3.jpg"},
	{ImageUrl: "/assets/images/leaves/leaves4.jpg"},
	{ImageUrl: "/assets/images/leaves/leaves5.jpg"},
	{ImageUrl: "/assets/images/leaves/leaves6.webp"},
	{ImageUrl: "/assets/images/leaves/leaves7.jpg"},
	{ImageUrl: "/assets/images/leaves/leaves8.webp"},
	{ImageUrl: "/assets/images/leaves/leaves9.jpg"},
	{ImageUrl: "/assets/images/leaves/leaves10.webp"},
}

var motionImages = []Images{
	{ImageUrl: "/assets/images/motion/cheese.jpg"},
	{ImageUrl: "/assets/images/motion/motion1.webp"},
	{ImageUrl: "/assets/images/motion/motion2.jpg"},
	{ImageUrl: "/assets/images/motion/motion3.jpg"},
	{ImageUrl: "/assets/images/motion/motion4.jpg"},
	{ImageUrl: "/assets/images/motion/motion5.jpg"},
	{ImageUrl: "/assets/images/motion/motion6.webp"},
	{ImageUrl: "/assets/images/motion/motion7.webp"},
}

func Science(params []string) []byte {
	genRandom()
	random := scientificFacts[rand.Intn(len(scientificFacts))]

	jsonStr, err := json.Marshal(random)

	if err != nil {
		fmt.Println("Something went wrong")

		jsonErr := Response{Success: false, Message: "Something went wrong"}
		jsonStr, _ := json.Marshal(jsonErr)
		return jsonStr
	}

	return jsonStr
}

func DecodeHash(params []string) []byte {
	var decoded []byte
	var err error

	decodeBase64 := func(hash string) []byte {
		decoded, err = base64.StdEncoding.DecodeString(hash)
		if err != nil {
			fmt.Println("Something went wrong")
			return []byte("false")
		}
		return decoded
	}

	decodeBase32 := func(hash string) []byte {
		decoded, err = base32.StdEncoding.DecodeString(hash)
		if err != nil {
			fmt.Println("Something went wrong")
			return []byte("false")
		}
		return decoded
	}

	if len(params) == 3 {

		var hash string
		var hashType string

		if !strings.Contains(params[2], "apiKey") {
			jsonErr := Response{Success: false, Message: "APIKey needs to be last appended item"}
			jsonStr, _ := json.Marshal(jsonErr)
			return jsonStr
		}

		if strings.Contains(params[0], "input") {
			hash = strings.Replace(params[0], "input=", "", -1)
			hashType = strings.Replace(params[1], "type=", "", -1)
		} else {
			hashType = strings.Replace(params[0], "type=", "", -1)
			hash = strings.Replace(params[1], "input=", "", -1)
		}

		if hashType == "base64" {
			decodeBase64(hash)
		} else if hashType == "base32" {
			decodeBase32(hash)
		}
	} else {
		jsonErr := Response{Success: false, Message: "Something went wrong"}
		jsonStr, _ := json.Marshal(jsonErr)
		return jsonStr
	}

	resp := map[string]string{"decoded": string(decoded)}
	jsonStr, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("Error occured!")

		jsonErr := Response{Success: false, Message: "Something went wrong"}
		jsonStr, _ := json.Marshal(jsonErr)
		return jsonStr
	}

	return jsonStr
}

func Season(params []string) []byte {
	genRandom()
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

func LeafImage(params []string) []byte {
	genRandom()
	random := leavesImages[rand.Intn(len(leavesImages))]

	jsonStr, err := json.Marshal(random)

	if err != nil {
		fmt.Println("Stinker happened")

		jsonErr := Response{Success: false, Message: "Something went wrong"}
		jsonStr, _ := json.Marshal(jsonErr)
		return jsonStr
	}

	return jsonStr
}

func MotionImage(params []string) []byte {
	genRandom()
	var random Images
	if slices.Contains(params, "noburger") {
		newRand := motionImages[1:]
		random = newRand[rand.Intn(len(newRand))]
	} else {
		random = motionImages[rand.Intn(len(motionImages))]
	}

	jsonStr, err := json.Marshal(random)

	if err != nil {
		fmt.Println("Stinker happened")

		jsonErr := Response{Success: false, Message: "Something went wrong"}
		jsonStr, _ := json.Marshal(jsonErr)
		return jsonStr
	}

	return jsonStr
}

func genRandom() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

}
