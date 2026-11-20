package handler

import (
	"net/http"
	"strconv"

	"school-management-backend/internal/middleware"
	"school-management-backend/internal/models"
	"school-management-backend/internal/service"
)

type AttendanceHandler struct {
	attendanceService *service.AttendanceService
}

func NewAttendanceHandler(attendanceService *service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{attendanceService: attendanceService}
}

func (h *AttendanceHandler) GetClassAttendance(w http.ResponseWriter, r *http.Request) {
	classIDStr := r.URL.Query().Get("class_id")
	if classIDStr == "" {
		WriteError(w, http.StatusBadRequest, "class_id query parameter is required")
		return
	}

	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "invalid class_id")
		return
	}

	date := r.URL.Query().Get("date")
	records, err := h.attendanceService.GetClassAttendance(r.Context(), classID, date)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, records)
}

func (h *AttendanceHandler) SaveAttendance(w http.ResponseWriter, r *http.Request) {
	var req models.SaveAttendanceRequest
	if err := DecodeJSON(r, &req); err != nil {
		WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.attendanceService.SaveAttendance(r.Context(), req); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"message": "attendance saved successfully"})
}

func (h *AttendanceHandler) GetMyAttendance(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	records, err := h.attendanceService.GetStudentAttendance(r.Context(), claims.UserID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, records)
}

func (h *AttendanceHandler) GetStudentAttendance(w http.ResponseWriter, r *http.Request) {
	studentIDStr := r.URL.Query().Get("student_id")
	if studentIDStr == "" {
		WriteError(w, http.StatusBadRequest, "student_id is required")
		return
	}

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "invalid student_id")
		return
	}

	records, err := h.attendanceService.GetStudentAttendance(r.Context(), studentID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, records)
}
