package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Car struct {
	Name      string      `json:"name"`
	TimeSlots []TimeSlot `json:"timeSlots"`
}

type TimeSlot struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Booked    bool   `json:"booked"`
}

var cars []Car

func main() {
	cars = []Car{
		{
			Name: "Toyota Camry",
			TimeSlots: []TimeSlot{
				{StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false},
				{StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false},
				{StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false},
			},
		},
		{
		  Name: "Honda Civic",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		  },
		},
		{
		  Name: "Ford Mustang",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		  },
		},
		{
		  Name: "Chevrolet Corvette",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Nissan Altima",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		  },
		},
		{
		  Name: "BMW 3 Series",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Mercedes-Benz E-Class",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Audi A4",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Tesla Model 3",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Hyundai Sonata",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		  },
		},
		{
		  Name: "Kia Optima",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Subaru Outback",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Mazda CX-5",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Jeep Wrangler",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Lexus RX 350",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		  },
		},
		{
		  Name: "GMC Yukon",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Cadillac Escalade",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Honda Creta",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "BMW M50",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
		{
		  Name: "Thar",
		  TimeSlots: []TimeSlot {
			{ StartTime: "9:00 AM", EndTime: "9:30 AM", Booked: false },
			{ StartTime: "10:00 AM", EndTime: "10:30 AM", Booked: false },
			{ StartTime: "11:00 AM", EndTime: "11:30 AM", Booked: false },
		},
		},
	}

	r := mux.NewRouter()

	corsOptions := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsHandler := handlers.CORS(corsOptions)
	r.Use(corsHandler)

	http.Handle("/", corsHandler(r))

	r.HandleFunc("/api/cars", GetCars).Methods("GET")
	r.HandleFunc("/api/cars/{id}/book", BookCar).Methods("POST")

	corsGetCars := handlers.CORS(corsOptions)(http.HandlerFunc(GetCars))
	corsBookCar := handlers.CORS(corsOptions)(http.HandlerFunc(BookCar))

	r.Handle("/api/cars", corsGetCars).Methods("GET")
	r.Handle("/api/cars/{id}/book", corsBookCar).Methods("POST")

	port := ":5000"
	fmt.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func BookCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	carName := vars["id"]

	var request struct {
		TimeSlot string `json:"timeSlot"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var car *Car
	for i := range cars {
		if cars[i].Name == carName {
			car = &cars[i]
			break
		}
	}

	if car == nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	var timeSlot *TimeSlot
	for i := range car.TimeSlots {
		if car.TimeSlots[i].StartTime == request.TimeSlot {
			timeSlot = &car.TimeSlots[i]
			break
		}
	}

	if timeSlot == nil {
		http.Error(w, "Time slot not found", http.StatusNotFound)
		return
	}

	if timeSlot.Booked {
		http.Error(w, "Time slot already booked", http.StatusConflict)
		return
	}

	timeSlot.Booked = true

	w.WriteHeader(http.StatusNoContent)
}
