# Project: kisaansathi – Farmer Empowerment App

## 🌟 Vision

An integrated platform for Indian farmers to connect, learn, access expert advice, track mandi prices, explore climate-smart farming, and get real-time support.

---

## ⚙️ Backend Architecture (projected)

### 🔹 Tech Stack

- **Go (Golang)** → Core APIs, Auth, User Management, Post Feed
- **Python** → Analytics (yield forecast, weather model, price prediction)
- **PostgreSQL** → Relational data (users, posts, schemes, prices)
- **Redis** → Caching mandi bhav, sessions
- **Firebase Cloud Messaging** → Push alerts

---

## 🗃️ Data Models (Sample)

### 🧑 User

- ID, Name, Role (Farmer, Advisor, Scientist)
- Phone, Language, Location (lat/lng, village, district)
- Soil Type (optional)

### 📝 Post

- ID, UserID, Caption, MediaURL
- CropTag, Comments\[], Likes\[]

### 📈 MandiPrice

- Crop, Region, Date, Price (₹/quintal)

### 📍 ServiceCenter

- Name, Type (Biotech, Vet), Location, Contact

### 📊 Scheme

- Title, Description, Eligibility, Tags, PDF URL

---

## 🧠 AI & Analytics (Python)

- 📉 Price Forecasts → LSTM/RNN models on mandi prices
- ☀️ Weather Recommendation → OpenWeather + soil type decision engine
- 📊 Crop Suitability → User's GPS + rainfall + historical yield match
