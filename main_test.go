package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func resetFiles() {
	files = map[int]string{}
	nextID = 1
}

func uploadRequest(t *testing.T, method, url, filename string) *http.Request {
	t.Helper()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		t.Fatal(err)
	}
	part.Write([]byte("fake data"))
	writer.Close()

	req := httptest.NewRequest(method, url, &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}

func TestPostFile(t *testing.T){
	resetFiles()
	req := uploadRequest(t, http.MethodPost, "/files/1", "report.xlsx")
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	uploadHandler(rec,req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "report.xlsx") {
	t.Errorf("expected report.xlsx in response, got %q", rec.Body.String())
	}
}

func TestPostUnsupportedFileType(t *testing.T){
	resetFiles()
	req := uploadRequest(t, http.MethodPost, "/files/1", "report.pdf")
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	uploadHandler(rec,req)

	if rec.Code != http.StatusUnsupportedMediaType {
		t.Errorf("expected 415, got %d", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "report.pdf") {
		t.Errorf("did not expect report.pdf in response, got %q", rec.Body.String())	
	}


}

func TestDeleteFileReturns(t *testing.T) {
	resetFiles()
	files[1] = "report.xlsx"

	req := httptest.NewRequest(http.MethodDelete, "/files/1", nil)
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	deleteHandler(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", rec.Code)
	}
	if _, exists := files[1]; exists {
		t.Errorf("expected file 1 to be deleted, but it still exists")
	}
}

func TestGetWithOneFile(t *testing.T){
	resetFiles()
	files[1] = "report.xlsx"

	req := httptest.NewRequest(http.MethodGet, "/files/1", nil)
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	getHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
	if files[1] != "report.xlsx" {
		t.Errorf("expected file 1 to be report.xlsx, got %q", files[1])
	}
}

func TestGetAllWithFilesReturns(t *testing.T) {
	resetFiles()
	files[1] = "report.xlsx"
	files[2] = "report_2.xlsx"

	req := httptest.NewRequest(http.MethodGet, "/files", nil)
	rec := httptest.NewRecorder()

	getAllHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
	if files[1] != "report.xlsx" {
		t.Errorf("expected file 1 to be report.xlsx, got %q", files[1])
	}
	if files[2] != "report_2.xlsx" {
		t.Errorf("expected file 2 to be report_2.xlsx, got %q", files[2])
	}
}

func TestReplaceOneFile(t *testing.T) {
	resetFiles()
	files[1] = "report.xlsx"

	req := uploadRequest(t, http.MethodPut, "/files/1", "report_2.xlsx")
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	replaceHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
	if files[1] != "report_2.xlsx" {
		t.Errorf("expected file 1 to be report_2.xlsx, got %q", files[1])
	}
}