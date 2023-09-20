import React, { useState, useEffect } from "react";
import axios from "axios";

function HomePage() {
  const [cars, setCars] = useState([]);

  useEffect(() => {
    axios
      .get("http://localhost:5000/api/cars")
      .then((response) => {
        setCars(response.data);
      })
      .catch((error) => {
        console.error("Error fetching car data:", error);
      });
  }, []);

  return (
    <div className="container">
      <h1>Available Cars</h1>

      <div className="car-list">
        {cars.map((car) => (
          <div className="car-item" key={car._id}>
            <h3>{car.name}</h3>
            {car.timeSlots.length > 0 && (
              <a href={`/booking/${car._id}`} className="book-link">
                Book
              </a>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}

export default HomePage;
