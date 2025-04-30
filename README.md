# 📌 AI-Powered Task Management System

A real-time task management system featuring:

- JWT-based user authentication  
- Task creation, assignment, and tracking  
- AI-generated task suggestions (via OpenAI/Gemini)  
- Real-time updates using WebSockets  
- Fully deployed backend and frontend  

---

## 🚀 Live Demo

- **Backend: [Live URL](#)  
- **Video Demo (5 mins)**: [Watch here](#)

---

## 🛠️ Tech Stack

- **Language**: Golang (Gin/Fiber)  
- **Authentication**: JWT  
- **Database**: PostgreSQL
- **Real-time**: WebSockets & Goroutines  
- **AI Integration**: OpenAI 
- **Deployment**: Render 

---

## 🔍 Features

- ✅ User sign-up and login (JWT-based)  
- ✅ Task creation, assignment, and management  
- ✅ AI-generated task breakdowns  
- ✅ Real-time updates via WebSockets  
- ✅ Responsive UI using Tailwind CSS  
- ✅ Fully deployed and accessible online  

---

## 🤖 AI Utilization

- **Smart Suggestions**: Task breakdowns via OpenAI
- **Development Assistance**: Used Copilot & ChatGPT for:
  - Code generation & scaffolding  
  - Bug fixing and refactoring  
  - Real-time component integration  

---

## 📦 Bonus Implementations *(If applicable)*

- Docker

---

## 📄 Setup Instructions

### Prerequisites

- Go 1.24.2
- PostgreSQL
- OpenAI API key

### Environment Variables
Create a `.env` file in the root directory and add the following:

```bash
DATABASE_URL=postgres://<username>:<password>@localhost:5432/<database_name>?sslmode=disable
JWT_SECRET_KEY=<your_jwt_secret_key>
OPENAI_API_KEY=<your_openai_api_key>
```

### Backend Setup

#### Go

```bash
go run main.go
```
#### Docker Compose

```bash
docker-compose up
```