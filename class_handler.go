package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"school-management-backend/internal/middleware"
	"school-management-backend/internal/models"
	"school-management-backend/internal/service"
)

type ClassHandler struct {
	classService *service.ClassService
}

func NewClassHandler(classService *service.ClassService) *ClassHandler {
	return &ClassHandler{classService: classService}
}

func (h *ClassHandler) ListClasses(w http.ResponseWriter, r *http.Request) {
	claims, _ := middleware.GetUserFromContext(r.Context())

	var teacherID *int
	var studentID *int

	// Optional query params or role scoping
	if r.URL.Query().Get("my_classes") == "true" && claims != nil {
		if claims.Role == models.RoleTeacher {
			teacherID = &claims.UserID
		} else if claims.Role == models.RoleStudent {
			studentID = &claims.UserID
		}
	} else if tidStr := r.URL.Query().Get("teacher_id"); tidStr != "" {
		if tid, err := strconv.Atoi(tidStr); err == nil {
			teacherID = &tid
		}
	}

	classes, err := h.classService.ListClasses(r.Context(), teacherID, studentID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, classes)
}

func (h *ClassHandler) GetClass(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "invalid class id")
		return
	}

	classDetail, err := h.classService.GetClassDetails(r.Context(), id)
	if err != nil {
		WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, classDetail)
}

func (h *ClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var req models.CreateClassRequest
	if err := DecodeJSON(r, &req); err != nil {
		WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	class, err := h.classService.CreateClass(r.Context(), req)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteJSON(w, http.StatusCreated, class)
}

func (h *ClassHandler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "invalid class id")
		return
	}

	if err := h.classService.DeleteClass(r.Context(), id); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"message": "class deleted successfully"})
}

func (h *ClassHandler) EnrollStudent(w http.ResponseWriter, r *http.Request) {
	classIDStr := chi.URLParam(r, "id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "invalid class id")
		return
	}

	var req models.EnrollRequest
	if err := DecodeJSON(r, &req); err != nil {
		WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.classService.EnrollStudent(r.Context(), classID, req.StudentID); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"message": "student enrolled successfully"})
}

func (h *ClassHandler) RemoveStudent(w http.ResponseWriter, r *http.Request) {
	classIDStr := chi.URLParam(r, "id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "invalid class id")
		return
	}

	studentIDStr := chi.URLParam(r, "studentId")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "invalid student id")
		return
	}

	if err := h.classService.RemoveStudent(r.Context(), classID, studentID); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"message": "student removed from class"})
}
