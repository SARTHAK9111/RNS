# **Realtime Notification System**

This project is a **real-time web application** built in **Go** with a **React frontend**. It allows users to **submit content** through a form and uses **Redis Pub/Sub** and **WebSocket** for real-time notifications. Submitted content is stored in a **MySQL database**.

The system supports both **user notifications** via user input and **clock-based events** to simulate real-time triggers.

---

## **Table of Contents**

- [Project Overview](#project-overview)
- [Features](#features)
- [Real-World Applications](#real-world-applications)
- [System Architecture](#system-architecture)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Database](#database)
- [Middleware](#middleware)
- [Test Output](#test-output)
- [Docker Deployment](#docker-deployment)
- [Contributing](#contributing)

---

## **Project Overview**

The **Realtime Notification System** integrates a **Go backend**, **Redis-based event-driven notifications**, and a **React frontend**. It allows:

- **user notifications** through user input via the **custom notification input panel** on the frontend.
- **Clock-based events** to simulate real-time notifications.
- **WebSocket connections** for instant updates to the frontend.

### **Frontend UI Example:**

![Frontend UI Example Placeholder](./project_documentation/assets/Frontend%20UI.png)  

---

## **Features**

- Submit text-based content via the **frontend input panel** or a POST request.
- Store submitted content in a MySQL database.
- **Redis Pub/Sub** for event-driven notifications.
- **WebSocket integration** for real-time updates.
- **Clock-based events** to simulate real-time triggers.
- Frontend built with **React and TypeScript**.
- Middleware for handling **CORS** between backend and frontend.

---

## **Real-World Applications**

This system drives **real-time engagement** across multiple industries. Below are some of its key use cases:

### **1. News and Media Platforms**  

- Deliver **breaking news alerts** and **content recommendations** instantly.  
- Boost user engagement and **ad revenue** with timely notifications.

### **2. Gaming Platforms and Monetization**  

- Notify players about **live events, rewards**, and **leader-board updates**.  
- Encourage **in-app purchases** and enhance **player retention**.

### **3. FinTech Platforms**  

- Provide real-time **stock alerts, transaction updates**, and **payment notifications**.  
- Help users manage finances and stay ahead with **market insights**.

### **4. E-Commerce Platforms**  

- Send updates on **order tracking, promotions**, and **stock availability**.  
- Improve **sales conversions** and enhance customer satisfaction.

### **5. Social Media and Messaging Apps**  

- Notify users about **mentions, messages**, and **new followers**.  
- Keep the user base engaged and connected.

### **6. IoT and Monitoring Systems**  

- Alert on **system anomalies, equipment failures**, and **temperature changes**.  
- Enable fast responses to prevent downtime and ensure safety.

### **7. Healthcare Systems**  

- Send **appointment reminders, lab results**, and **prescription updates**.  
- Reduce missed appointments and improve patient engagement.

### **8. Event and Calendar Management Systems**  

- Provide reminders for **meetings, events**, and **webinars**.  
- Improve participation and ensure **better time management**.

---

This **format**:

- Uses **bullet points** for clarity.
- Keeps each **use case brief** and highlights the **core impact**.
- Aligns with the rest of the README structure for a consistent look.

---

## **System Architecture**

```plaintext
Here is a system design diagram for the system
```

![SA Diagram](./project_documentation/assets/SA/System%20Architecture%20File.png)  

## **Installation**

### **Prerequisites**

- Go (v1.23 or later)
- Node.js and npm (for frontend)
- Redis Server
- MySQL Server

### **Clone the Repository**

```bash
git clone https://github.com/yourusername/Realtime-Notification-System.git
cd Realtime-Notification-System
```

---

### **Backend Setup**

1. **Install Go Dependencies:**

   ```bash
   cd backend_system
   go mod tidy
   ```

2. **Set Up MySQL Database:**

   ```bash
   mysql -u your_username -p realtime_notification_system < backend_system/server/mysql_schema.sql
   ```

3. **Run Redis Server:**

   ```bash
   redis-server
   ```

---

### **Frontend Setup**

1. **Navigate to Frontend Directory:**

   ```bash
   cd frontend_application
   ```

2. **Install Frontend Dependencies:**

   ```bash
   npm install
   ```

---

### **Set Up Environment Variables**

1. **Copy `.env.example` to `.env`:**

   ```bash
   cp .env.example .env
   ```

2. **Update Database Credentials in `.env`:**

   ```plaintext
   MYSQL_USER=your_username
   MYSQL_PASSWORD=your_password
   MYSQL_DB=realtime_notification_system
   MYSQL_HOST=127.0.0.1
   MYSQL_PORT=3306
   ```

---

### **Run the Backend**

```bash
cd backend_system/server
go run main.go
```

---

### **Run the Frontend**

```bash
cd frontend_application
npm start
```

---

## **Usage**

### **User Notification Submission**

Submit content via the frontend **custom notification input panel** or using Postman:

```bash
curl -X POST -d "content=Manual Notification" http://localhost:8080/submit
```

---

### **Clock-Based Event-Driven Notifications**

Every minute, a **clock-based event** publishes a notification via Redis to simulate real-time events.

---

## **API Endpoints**

- **POST /submit**: Submits content to the database and triggers a manual notification.
- **WebSocket /ws**: Listens for real-time notifications from Redis.

---

## **Database**

The **MySQL database** stores submitted content using the following schema:
Added a init.db (`backend_system/sql` folder) file for configuring the database

---

## **Middleware**

The **CORS middleware** allows the frontend to communicate with the backend. It is implemented in `backend_system/middleware_layer/cors.go` and applied to backend routes.

---

## **Test Output**

**User Notification Submission (Front End):**

   ![Front End Input Placeholder](./project_documentation/assets/Frontend_input.png)

**User Notification Submission (Postman):**

   ![Postman Input Placeholder](./project_documentation/assets/Postman_input.png)  

**Database Content View (Event Driven Notifications):**

   ![MySQL Output Placeholder](./project_documentation/assets/MySQL_event_driven_storage.png)  

**Database Content View (User Notifications):**

   ![MySQL Output Placeholder](./project_documentation/assets/MySQL_input_notifcation.png)  

---

## **Docker Deployment**

### **Docker Deployment Instructions**

Did following **changes and configurations** made for Docker deployment:

1. **Docker Compose File**:
   - Added `docker-compose.yml` to coordinate services: **backend**, **frontend**, **MySQL**, and **Redis**.
   - Specified health checks and dependencies for the backend to wait until MySQL and Redis are ready.

2. **Environment Variables**:
   - Extended the `.env.example` file to include Docker-specific variables.
   - Updated Redis and MySQL connection settings from `localhost` to their Docker container names.

3. **Redis and MySQL Client Changes**:
   - Changed Redis connection in the backend code to `redis:6379`.
   - Changed MySQL connection in the backend to `mysql:3306`.

4. **SQL Schema File for Database Setup**:
   - Added a **template SQL file** (`permission.sql.template`) schema file to add permissions are correctly, if the database does not grant them initially.

5. **Backend and Frontend Dockerfiles**:
   - Created separate Dockerfiles for **backend** and **frontend** services to containerize the applications efficiently.

---

#### **Running the Project with Docker**

1. **Build and Start Docker Containers**:

   ```bash
   docker-compose up --build
   ```

2. **Verify Services**:
   - Ensure MySQL, Redis, backend, and frontend services are running correctly with:

     ```bash
     docker ps
     ```

3. **Access Frontend**:
   - Open [http://localhost:3000](http://localhost:3000) in your browser.

4. **Backend API**:
   - Test the backend API via Postman or curl:

     ```bash
     curl -X POST -d "content=Hello Docker" http://localhost:8080/submit
     ```

---

## **Contributing**

Contributions are welcome! Fork the repository and submit a pull request.

---
