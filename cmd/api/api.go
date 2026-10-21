package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Vijay6922/Go-DBconnection/sqlMethods"
	"log"
	"net/http"
	"strconv"
)

func API(db *sql.DB) {
	mux := http.NewServeMux()

	//get all the details of rows
	mux.HandleFunc("GET /info", func(w http.ResponseWriter, r *http.Request) {
		details, err := sqlMethods.ReturnAllDetails(db)
		if err != nil {
			http.Error(w, "failed to show all details", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(details); err != nil {
			log.Printf("failed to write details list response: %v", err)
		}
	})

	//insert values of given data 
	mux.HandleFunc("POST /info", func(w http.ResponseWriter, r *http.Request) {

		var details sqlMethods.Details
		err := json.NewDecoder(r.Body).Decode(&details)
		if err != nil {
			http.Error(w, "Invalid JSON payload format", http.StatusBadRequest)
			return
		}

		status, err := sqlMethods.InsertData(db, details)
		if err != nil {
			http.Error(w, "Database insertion failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // HTTP 201 Created
		fmt.Fprintf(w, `{"status": %d, "message": "details created successfully"}`, status)
	})

	//delete rows of given id 
	mux.HandleFunc("DELETE /info/{id}", func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid details ID", http.StatusBadRequest)
			return
		}

		rowsAffected, err := sqlMethods.DeleteDetails(db, id)
		if err != nil {
			log.Printf("failed to delete details %d: %v", id, err)
			http.Error(w, "failed to delete details", http.StatusInternalServerError)
			return
		}
		if rowsAffected == 0 {
			http.Error(w, "details not found", http.StatusNotFound)
			return
		}

		sqlMethods.VerifyRows(int(rowsAffected),id,w,db)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{
			"deleted": true,
			"id":      id,
		}); err != nil {
			log.Printf("failed to write delete response for details %d: %v", id, err)
		}
	})

	//find rows of given id 
	mux.HandleFunc("GET /info/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid details ID", http.StatusBadRequest)
			return
		}

		details, err := sqlMethods.FindRows(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "details not found", http.StatusNotFound)
				return
			}
			http.Error(w, "failed to find details", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(details); err != nil {
			log.Printf("failed to write details response: %v", err)
		}
	})

	//update phone number of given id
	mux.HandleFunc("PATCH /info/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid details ID", http.StatusBadRequest)
			return
		}

		var request struct {
			PhoneNumber string `json:"phonenumber"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid JSON payload format", http.StatusBadRequest)
			return
		}
		if request.PhoneNumber == "" {
			http.Error(w, "phonenumber is required", http.StatusBadRequest)
			return
		}

		rowsAffected, err := sqlMethods.UpdatePhone(db, id, request.PhoneNumber)
		if err != nil {
			log.Printf("failed to update phone number for details %d: %v", id, err)
			http.Error(w, "failed to update phone number", http.StatusInternalServerError)
			return
		}

		sqlMethods.VerifyRows(int(rowsAffected),id,w,db)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{
			"updated":     true,
			"id":          id,
			"phonenumber": request.PhoneNumber,
		}); err != nil {
			log.Printf("failed to write phone update response for details %d: %v", id, err)
		}
	})

	log.Println("API server listening at http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err)
	}
}
