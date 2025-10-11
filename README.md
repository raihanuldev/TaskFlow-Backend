# 🧠 TaskFlow Backend (GoLang)

**TaskFlow** is a team-based **Task Management System** built with pure **GoLang (`net/http`)**.  
It’s designed to help teams manage projects, assign tasks, and communicate effectively —  
with **real-time notifications**, **role-based access**, and **JWT-secured APIs**.

---

## 🚀 Features

- 👥 **User Authentication & Authorization**
  - Register / Login using JWT
  - Role-based access control (Admin, Project Manager, Engineer)
- ✅ **Task Management**
  - Create, update, delete, and assign tasks
  - Task priority and deadlines
- 💬 **Comments System**
  - Add comments under each task
  - Like / dislike system
- 🔔 **Real-time Notifications (WebSocket)**
  - Live updates for:
    - New tasks
    - Task status changes
    - New comments
    - Deadline reminders
- 📜 **Activity Logs**
  - Tracks who did what (e.g., created, updated, deleted tasks)
- 🔐 **JWT Security**
  - Token-based authentication middleware
- 🧱 **Clean Modular Structure**
  - Organized handlers, routes, models, and utils

---


---

## 🧰 Tech Stack

| Component | Technology |
|------------|-------------|
| **Language** | Go (net/http, encoding/json, gorilla/mux, gorilla/websocket) |
| **Database** | PostgreSQL / SQLite |
| **Auth** | JWT (github.com/golang-jwt/jwt/v5) |
| **Password Hashing** | bcrypt |
| **Real-time** | WebSocket |
| **Config** | `.env` file with `os.Getenv()` |
| **Logging** | Go log package |

---

## 🧱 Database Models

### 🧍‍♂️ UserModel
| Field | Type | Description |
|-------|------|-------------|
| ID | int | Primary key |
| Name | string | Full name |
| ImgURL | string | Profile image |
| Phone | string | Contact number |
| Email | string | Unique email |
| Note | string | Personal note |
| PreviousNotes | []string | Old notes list |
| Role | enum | Admin / ProjectManager / Engineer |

---

### 📋 TaskModel
| Field | Type | Description |
|-------|------|-------------|
| ID | int | Task ID |
| TaskName | string | Task title |
| Details | string | Task description |
| AuthorID | int | Creator (UserID) |
| CreatedAt | time.Time | Creation timestamp |
| Deadline | time.Time | Deadline date |
| Priority | string | Low / Medium / High |
| Comments | []Comment | Related comments |
| Status | string | To-Do / In Progress / Done |

---

### 💬 CommentModel
| Field | Type | Description |
|-------|------|-------------|
| ID | int | Comment ID |
| Comment | string | Text content |
| UserID | int | Comment author |
| LikedCount | int | Number of likes |
| DislikedCount | int | Number of dislikes |
| CreatedAt | time.Time | Timestamp |

---

### 🔔 NotificationModel
| Field | Type | Description |
|-------|------|-------------|
| ID | int | Notification ID |
| Type | string | TaskUpdate / Comment / Deadline / NewTask |
| Message | string | Notification content |
| ReceiverID | int | User ID who receives |
| CreatedAt | time.Time | Timestamp |
| IsRead | bool | Read/unread status |

---

## 🔐 Roles & Permissions

| Role | Permissions |
|------|--------------|
| **Admin** | Full control: manage users, tasks, comments, and notifications |
| **Project Manager** | Manage team tasks, assign engineers, comment, and view all logs |
| **Engineer** | Handle assigned tasks, comment, and receive notifications |

---

## 🌐 API Endpoints (Example)

| Method | Endpoint | Description |
|--------|-----------|-------------|
| POST | `/api/register` | Create new user |
| POST | `/api/login` | Login & get JWT |
| GET | `/api/users` | Get all users (Admin only) |
| POST | `/api/tasks` | Create new task |
| GET | `/api/tasks` | Fetch all tasks |
| PUT | `/api/tasks/{id}` | Update a task |
| DELETE | `/api/tasks/{id}` | Delete a task |
| POST | `/api/tasks/{id}/comments` | Add a comment to task |
| GET | `/api/notifications/ws` | Connect to real-time WebSocket |

---

## ⚡ Real-Time Notifications Flow

**Technology Used:** `gorilla/websocket`

**Use Case Example:**
1. Project Manager creates a new task →  
   All Engineers in that team instantly receive a notification.
2. Engineer updates task status →  
   Project Manager receives “Task Updated” notification.
3. New comment →  
   Task author gets notified instantly.

**Event Types:**
- `NEW_TASK`
- `TASK_UPDATED`
- `NEW_COMMENT`
- `DEADLINE_ALERT`

---
## 🧠 Future Improvements

- 📧 **Email Notifications**
  - Automatic email alerts for upcoming task deadlines
  - Notify users when a new task or comment is added

- 📊 **Task Analytics Dashboard**
  - Track completed vs pending tasks
  - Display productivity charts for teams or individuals

- 📎 **File Attachments**
  - Allow users to upload and attach files to tasks
  - Store in local or cloud (e.g., AWS S3 / Cloudinary)

- 🔍 **Advanced Pagination & Filtering**
  - Filter by status, priority, or deadline
  - Paginate task list for better performance

- 🐳 **Docker Deployment**
  - Dockerize the backend for smooth deployment
  - Include `Dockerfile` and `docker-compose.yml`

- 🧪 **Unit & Integration Tests**
  - Add tests for routes, handlers, and database logic
  - Use Go’s built-in `testing` package for coverage


## 🧠 Development Setup

```bash
# 1️⃣ Clone the repository
git clone https://github.com/<your-username>/taskflow-backend.git
cd taskflow-backend

# 2️⃣ Initialize Go module
go mod init github.com/raihanuldev/taskflow-backend
go mod tidy

# 3️⃣ Setup environment variables
cp .env.example .env

# 4️⃣ Run the server
go run main.go


