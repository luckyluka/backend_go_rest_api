package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users/{user_id}/certificates", GetCertificatesEndpoint).Methods("GET")
	router.HandleFunc("/users/{user_id}/certificates/{user_id}", CreateCertificateEndpoint).Methods("POST")
	router.HandleFunc("/users/{user_id}/certificates/{cert_id}", UpdateCertificateEndpoint).Methods("PATCH")
	router.HandleFunc("/users/{user_id}/certificates/{cert_id}", DeleteCertificateEndpoint).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/certificates/{cert_id}/transfers", TransferCertificateEndpoint).Methods("POST")
	router.HandleFunc("/users/{user_id}/certificates/{cert_id}/transfers", AcceptCertificateEndpoint).Methods("PATCH")
	return router
}

func TestGetCertificatesEndpoint(t *testing.T) {
	fmt.Println("TEST_GET_CERTIFICATE")
	request, _ := http.NewRequest("GET", "/users/{user_id}/certificates", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestCreateCertificateEndpoint(t *testing.T) {
	fmt.Println("TEST_CREATE_CERTIFICATE")
	certificate := &Certificate{
		Title:     "Nice",
		CreatedAt: "Riga",
		OwnerId:   "3",
		Year:      "1987",
		Note:      "lovely",
	}
	jsonCertificate, _ := json.Marshal(certificate)
	request, _ := http.NewRequest("POST", "/users/{user_id}/certificates/{user_id}", bytes.NewBuffer(jsonCertificate))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestUpdateCertificateEndpoint(t *testing.T) {
	fmt.Println("TEST_UPDATE_CERTIFICATE")
	certificate := &Certificate{
		Title:     "Nice",
		CreatedAt: "Riga",
		OwnerId:   "3",
		Year:      "1987",
		Note:      "lovely",
	}
	jsonCertificate, _ := json.Marshal(certificate)
	request, _ := http.NewRequest("PATCH", "/users/{user_id}/certificates/{user_id}", bytes.NewBuffer(jsonCertificate))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestDeleteCertificatesEndpoint(t *testing.T) {
	fmt.Println("TEST_DELETE_CERTIFICATE")
	request, _ := http.NewRequest("DELETE", "/users/{user_id}/certificates/{user_id}", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestTransferCertificatesEndpoint(t *testing.T) {
	fmt.Println("TEST_TRANSFER_CERTIFICATE")
	request, _ := http.NewRequest("POST", "/users/{user_id}/certificates/{user_id}/transfers", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestAcceptCertificatesEndpoint(t *testing.T) {
	fmt.Println("TEST_ACCEPT_CERTIFICATE")
	certificate := &Certificate{

		Transfer: &Transfer{Status: "Finished"},
	}

	jsonCertificate, _ := json.Marshal(certificate)
	request, _ := http.NewRequest("PATCH", "/users/{user_id}/certificates/{user_id}", bytes.NewBuffer(jsonCertificate))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
