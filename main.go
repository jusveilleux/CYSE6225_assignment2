package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	files  = map[int]string{}
	nextID = 1
	mu     sync.Mutex 
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".xlsx" {
		http.Error(w, "Only .xlsx files are allowed", http.StatusUnsupportedMediaType)
		return
	}

	mu.Lock()
	id := nextID
	files[id] = header.Filename
	nextID++
	mu.Unlock()

	fmt.Println("Upload received:", header.Filename)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s received", header.Filename)}


func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := files[id]; !exists {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	delete(files, id)
	w.WriteHeader(http.StatusNoContent)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	name, exists := files[id]
	if !exists {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %d: %s", id, name)
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if len(files) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "No files uploaded yet")
		return
	}

	w.WriteHeader(http.StatusOK)
	for id, name := range files {
		fmt.Fprintf(w, "File %d: %s\n", id, name)
	}
}

func replaceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	oldName, exists := files[id]
	if !exists {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".xlsx" {
		http.Error(w, "Only .xlsx files are allowed", http.StatusUnsupportedMediaType)
		return
	}

	files[id] = header.Filename

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %d replaced: %s -> %s", id, oldName, header.Filename)
}

func main() {
	http.HandleFunc("POST /upload", uploadHandler)
	http.HandleFunc("DELETE /files/{id}", deleteHandler)
	http.HandleFunc("PUT /files/{id}", replaceHandler)
	http.HandleFunc("GET /files/{id}", getHandler)
	http.HandleFunc("GET /files", getAllHandler)
	http.ListenAndServe(":8000", nil)
}