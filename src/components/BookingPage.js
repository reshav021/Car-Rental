import React, { useState, useEffect } from "react"
import { useParams } from "react-router-dom"
import axios from "axios"

function BookingPage() {
  const { carId } = useParams()
  const [selectedTimeSlot, setSelectedTimeSlot] = useState("")
  const [timeSlots, setTimeSlots] = useState([])

  const handleTimeSlots = (data) => {
    setTimeSlots(data)
  }

  useEffect(() => {
    axios
      .get(`http://localhost:3000/booking/${carId}`)
      .then((response) => {})
      .catch((error) => {
        console.error("Error fetching data:", error)
      })
  }, [carId])

  useEffect(() => {
    axios
      .get("http://localhost:5000/api/cars")
      .then((response) => {
        const foundCar = response.data.find((car) => car._id === carId)
      
        if(foundCar) {
          handleTimeSlots(foundCar.timeSlots)
        }
      })
      .catch((error) => {
        console.error("Error fetching car data:", error)
      })
  }, [carId])

  const handleBooking = () => {
    if (selectedTimeSlot) {
      axios
        .post(`http://localhost:5000/api/cars/${carId}/book`, {
          timeSlot: selectedTimeSlot,
        })
        .then((response) => {
          alert("Booking successful!")
          window.location.reload();
        })
        .catch((error) => {
          console.error("Error booking car:", error)
        })
    } else {
      alert("Please select a time slot.")
    }
  }

  return (
    <div className="BookingPage">
      <h1>Book a Car</h1>
      <label>Select a Time Slot:</label>

      <select
        value={selectedTimeSlot}
        onChange={(e) => setSelectedTimeSlot(e.target.value)}
      >
        <option value="">Select a time slot</option>
        {timeSlots.map((timeSlot) =>
          timeSlot.Booked === false ? (
            <option key={timeSlot._id} value={timeSlot._id}>
              {timeSlot.StartTime} - {timeSlot.EndTime}
            </option>
          ) : (
            <option key={timeSlot._id} value={timeSlot._id} disabled>
              {timeSlot.StartTime} - {timeSlot.EndTime}
            </option>
          )
        )}
      </select>

      <button onClick={handleBooking}>Book Car</button>
    </div>
  )
}

export default BookingPage