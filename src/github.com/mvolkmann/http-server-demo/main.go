package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type any interface{}

// Can"t just omit the unused parameters.
func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func getCoupon(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	couponURL :=
		"http://www.extremecouponingcanada.com/wp-content/uploads/2011/09/purina-one.jpg"
	sendText(w, couponURL)
}

func hello(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", params.ByName("name"))
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Sorry, I could not find that.")
}

func postActivityLevel(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	output := []string{
		"is fairly active.",
		"Examples: enjoys walks around the park or yard, " +
			"games of fetch, and other types of interactive play.",
	}
	sendJSON(w, output)
}

func postBreed(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	type pet struct {
		Species string
		Breed   string
	}
	output := pet{Species: "dog", Breed: "Whippet"}
	sendJSON(w, output)
}

func postBodyCondition(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	type input struct {
		BreedID       int32
		Age           int8
		Gender        string
		BodyCondition int8
	}
	decoder := json.NewDecoder(r.Body)
	var body input
	err := decoder.Decode(&body)
	if err != nil {
		msg := "invalid request body in POST to body-condition"
		http.Error(w, msg, http.StatusBadRequest)
	} else {
		//fmt.Printf("body = %+v\n", body)

		output := []string{
			"Ribs easily palpable with minimal fat covering.",
			"Waist easily noted viewed from above.",
			"Abdominal tuck evident.",
		}
		sendJSON(w, output)
	}
}

func postFoodRecommendations(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	type foodRec struct {
		ID          int32
		Name        string
		Description string
		Stars       int8
		ImageURL    string
	}
	output := []foodRec{
		foodRec{
			ID:          3,
			Name:        "Beyond Grain Free Beef & Egg Recipe",
			Description: "It all starts with the ingredients. Naturally. It’s about knowing exactly what your pet is eating – real, recognizable ingredients you know and trust, plus vitamins and minerals. We thoughtfully select natural ingredients that work together to deliver the right nutrition to help your pet live a long healthy life.",
			Stars:       4,
			ImageURL:    "https://www.beyondpetfood.com/sites/g/files/auxxlc511/files/styles/kraken_product_carousel_regular/public/BYD_DOGDry_17800_17406_GFBEEF_3LB_FRONT_New.png?itok=YuHZmdvS",
		},
		foodRec{
			ID:          7,
			Name:        "Bright Mind Adult 7+ Large Breed Formula",
			Description: "From less interaction with you, to lower engagement in daily activities, there are many signs your dog may be aging. Purina Pro Plan Bright Mind was created out of proprietary research that shows enhanced botanical oils provide an efficient fuel source for the brain in dogs age 7 and older--helping naturally nourish their minds to help them think more like they did when they were younger.",
			Stars:       5,
			ImageURL:    "https://www.proplan.com/media/4398/dog_dry_bag_brmnd_a7_lb_chknrc_f_16.png?mode=crop&width=323&anchor=middlecenter&scale=both",
		},
	}
	sendJSON(w, output)
}

func postWeightRange(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	type weightRange struct {
		MinWeight int8
		MaxWeight int8
	}
	output := weightRange{MinWeight: 11, MaxWeight: 17}
	sendJSON(w, output)
}

func sendJSON(w http.ResponseWriter, obj any) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		msg := fmt.Sprintf("error marshaling %+v to JSON", obj)
		http.Error(w, msg, http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func sendText(w http.ResponseWriter, text string) {
	w.Write([]byte(text))
}

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/hello/:name", hello)

	router.GET("/coupon/:foodId", getCoupon)
	router.POST("/activity-level", postActivityLevel)
	router.POST("/body-condition", postBodyCondition)
	router.POST("/breed", postBreed)
	router.POST("/food-recommendations", postFoodRecommendations)
	router.POST("/weight-range", postWeightRange)

	router.ServeFiles("/public/*filepath", http.Dir("public"))
	router.NotFound = http.HandlerFunc(notFound)

	log.Fatal(http.ListenAndServe(":8080", router))
}
