package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const ratePerKm float64 = 200 // Fare par kilometer in Naira

// location struct
type location struct {
	Lat  float64 // Location Longitude
	Long float64 // location Latitude
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance calculates the distance in meters between two locations using haversin function.
func (loc1 location) distance(loc2 location) float64 {
	var la1, lo1, la2, lo2, r float64
	lat1 := loc1.Lat
	lon1 := loc1.Long
	lat2 := loc2.Lat
	lon2 := loc2.Long
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return math.Round(2 * r * math.Asin(math.Sqrt(h)))
}

// Fare Returns the fare between loc1 and loc2 locations.
func (loc1 location) fare(loc2 location) float64 {
	dist := loc1.distance(loc2) // Distance between locations
	return (dist / 1000) * ratePerKm
}

func main() {

	// Locations Map
	var locations = map[string]location{
		"Alakahia":  location{4.8856561, 6.9210004},
		"Choba":     location{4.8671919, 6.9040659},
		"Rumuosi":   location{4.8807082, 6.9358458},
		"Rumuokoro": location{4.8737726, 6.9754668},
		"Mgbuoba":   location{4.8444011, 6.9336596},
		"Aluu":      location{4.9371896, 6.9365787},
		"Rumuola":   location{4.8271002, 7.0076752},
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Enter Your Pick up Location: ")
	pickUp, _ := reader.ReadString('\n') // Read pickup location string from stdin
	pickUp = strings.Trim(pickUp, "\n")
	pickUp = strings.ToLower(pickUp)
	pickUp = strings.Title(pickUp)

	fmt.Print("Please Enter Your Drop off Location: ")
	dropOff, _ := reader.ReadString('\n') // read drop off loaction from stdin
	dropOff = strings.Trim(dropOff, "\n")
	dropOff = strings.Title(dropOff)

	pickUpLoc, aval1 := locations[pickUp]
	dropOffLoc, aval2 := locations[dropOff]

	// Check if location one and two is in the locations map.
	if !aval1 && !aval2 {
		fmt.Print("Your pickup or dropoff locations are not within the areas Cabby operates at the moment\nWe Currently operate in the following cities:\n")
		for city := range locations {
			println(city)
		}

		println("Bye, later")

		return
	}

	fare := pickUpLoc.fare(dropOffLoc) // Calculate fare
	fmt.Printf("You fare is %v Naira\n", fare)

	time.Sleep(10 * time.Second) // Simulate travel for 10 secs.

	fmt.Print("Please Enter Amount To Pay: ")
	amt, _ := reader.ReadString('\n') // read fare from stdin
	amt = strings.Trim(amt, "\n")

	paidFare, err := strconv.ParseFloat(amt, 64) // convert fare string to base 64 float

	// Error converting fare from string to base 64
	if err != nil {
		fmt.Println("You enetered wrong monetery format")
		return
	}

	// Ascertain that the amount paid covers the fare.
	for numTrials, paid := 1, false; !paid; {
		if paidFare == fare {
			fmt.Println("Thank you for the payment")
			paid = true
			break
		} else if paidFare > fare {
			fmt.Printf("Your change is: %.2f\n", math.Round((paidFare-fare)*100)/100)
			paid = true
			break
		} else if paidFare < fare {
			if numTrials == 5 {
				fmt.Println("You will be reported to the police.")
			}
			fmt.Print("The amount you entered cannot cover you fare, please renter amount: ")
			amt, _ = reader.ReadString('\n') // read fare from stdin
			amt = strings.Trim(amt, "\n")

			paidFare, err = strconv.ParseFloat(amt, 64) // convert fare string to base 64 float

			// Error converting fare from string to base 64
			if err != nil {
				fmt.Println("You enetered wrong monetery format")
				return
			}

			numTrials++
		}

	}

	fmt.Print("Tip will be appreciated sir: ")
	tip, _ := reader.ReadString('\n') // Read tip from stdin
	tip = strings.Trim(tip, "\n")

	tipAmt, err := strconv.ParseFloat(tip, 64) // convert tip to base 64 float.

	// Error converting fare from string to base 64
	if err != nil {
		fmt.Println("You enetered wrong monetery format")
		return
	}

	// Greet the customer based on the tip amount.
	if tipAmt <= 0 {
		fmt.Println("You are a stingy fool")
	} else if tipAmt > 0 && tipAmt <= fare {
		fmt.Println("Thank you")
	} else {
		fmt.Println("gracias mucho")
	}

}
