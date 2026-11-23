package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Role constants
const (
	RoleAdmin   = "admin"
	RoleTeacher = "teacher"
	RoleStudent = "student"
)

// User represents a user entity
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	Phone        string    `json:"phone,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt,
	}
}

// Auth DTOs
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone,omitempty"`
}

// JWT Claims
type JWTClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// Class represents a course/class
type Class struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Section     string    `json:"section"`
	TeacherID   *int      `json:"teacher_id,omitempty"`
	TeacherName *string   `json:"teacher_name,omitempty"`
	Room        string    `json:"room,omitempty"`
	Schedule    string    `json:"schedule,omitempty"`
	StudentCount int      `json:"student_count,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateClassRequest struct {
	Name      string `json:"name"`
	Section   string `json:"section"`
	TeacherID *int   `json:"teacher_id,omitempty"`
	Room      string `json:"room,omitempty"`
	Schedule  string `json:"schedule,omitempty"`
}

type ClassDetail struct {
	Class
	Students []UserResponse `json:"students"`
}

type EnrollRequest struct {
	StudentID int `json:"student_id"`
}

// Attendance
type AttendanceRecord struct {
	ID          int       `json:"id"`
	ClassID     int       `json:"class_id"`
	StudentID   int       `json:"student_id"`
	StudentName string    `json:"student_name,omitempty"`
	Date        string    `json:"date"`
	Status      string    `json:"status"` // present, absent, late, excused
	Remarks     string    `json:"remarks,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type AttendanceItem struct {
	StudentID int    `json:"student_id"`
	Status    string `json:"status"`
	Remarks   string `json:"remarks,omitempty"`
}

type SaveAttendanceRequest struct {
	ClassID int              `json:"class_id"`
	Date    string           `json:"date"`
	Records []AttendanceItem `json:"records"`
}

// Grade
type Grade struct {
	ID             int       `json:"id"`
	ClassID        int       `json:"class_id"`
	ClassName      string    `json:"class_name,omitempty"`
	StudentID      int       `json:"student_id"`
	StudentName    string    `json:"student_name,omitempty"`
	AssignmentName string    `json:"assignment_name"`
	Score          float64   `json:"score"`
	MaxScore       float64   `json:"max_score"`
	GradedAt       time.Time `json:"graded_at"`
}

type CreateGradeRequest struct {
	ClassID        int     `json:"class_id"`
	StudentID      int     `json:"student_id"`
	AssignmentName string  `json:"assignment_name"`
	Score          float64 `json:"score"`
	MaxScore       float64 `json:"max_score"`
}

// Announcement
type Announcement struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Category   string    `json:"category"`
	AuthorID   *int      `json:"author_id,omitempty"`
	AuthorName *string   `json:"author_name,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateAnnouncementRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

// Dashboard Stats
type DashboardStats struct {
	TotalStudents   int               `json:"total_students"`
	TotalTeachers   int               `json:"total_teachers"`
	TotalClasses    int               `json:"total_classes"`
	AttendanceRate  float64           `json:"attendance_rate"`
	RecentGrades    []Grade           `json:"recent_grades"`
	Announcements   []Announcement    `json:"announcements"`
}
