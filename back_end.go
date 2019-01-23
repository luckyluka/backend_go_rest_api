package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Certificate struct {
	Id        string    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	CreatedAt string    `json:"createdAt,omitempty"`
	OwnerId   string    `json:"ownerId,omitempty"`
	Year      string    `json:"year,omitempty"`
	Note      string    `json:"note,omitempty"`
	Transfer  *Transfer `json:"transfer,omitempty"`
}

type Transfer struct {
	To     string `json:"to,omitempty"`
	Status string `json:"status,omitempty"`
}

type User struct {
	Id    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

var users []User
var certificates []Certificate

func checkUserId(w http.ResponseWriter, req *http.Request) error {
	params := mux.Vars(req)
	idExists := 0
	for i := range users {
		if users[i].Id == params["user_id"] {
			idExists = 1
		}
	}
	//fmt.Printf("idExists:%v\n", idExists)
	if idExists != 1 {
		json.NewEncoder(w).Encode("user_id non-existent")
		return errors.New("no such user")
	}
	return nil
}
func checkCertificateId(w http.ResponseWriter, req *http.Request) error {
	params := mux.Vars(req)
	idExists := 0
	for i := range certificates {
		if certificates[i].Id == params["cert_id"] {
			idExists = 1
		}
	}
	//fmt.Printf("idExists:%v\n", idExists)
	if idExists != 1 {
		json.NewEncoder(w).Encode("certificate_id non-existent")
		return errors.New("no such certificate")
	}
	return nil
}

func checkCertificateIdCreate(w http.ResponseWriter, req *http.Request) error {
	params := mux.Vars(req)
	idExists := 0
	for i := range certificates {
		if certificates[i].Id == params["cert_id"] {
			idExists = 1
		}
	}
	//fmt.Printf("idExists:%v\n", idExists)
	if idExists != 0 {
		json.NewEncoder(w).Encode("certificate_id already exists")
		return errors.New("no such certificate")
	}
	return nil
}

func checkResponse(w http.ResponseWriter, req *http.Request, certificate Certificate) (Certificate, error) {
	fmt.Printf("%v\n", req)
	error := json.NewDecoder(req.Body).Decode(&certificate)
	if error != nil {
		json.NewEncoder(w).Encode("non-valid request")
		return certificate, errors.New("not a valid json body")
	}
	return certificate, nil
}

func GetCertificatesEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	//CHECK IF USER_ID
	err := checkUserId(w, req)
	if err != nil {
		return
	}

	//RETRIEVE CERTIFICATES
	for i := range certificates {
		//fmt.Printf("i:%v\n", i)
		if certificates[i].OwnerId == params["user_id"] {
			json.NewEncoder(w).Encode(certificates[i])
			//return
		}
	}
	//json.NewEncoder(w).Encode(certificates)
}

func CreateCertificateEndpoint(w http.ResponseWriter, req *http.Request) {
	//fmt.Printf("CREATE CERTIFICATE\n")
	//CHECK IF USER_ID & CERTICATE_ID EXIST
	err := checkUserId(w, req)
	if err != nil {
		return
	}
	
	err = checkCertificateIdCreate(w, req)
	if err != nil {
		return
	}

	params := mux.Vars(req)
	var certificate Certificate
	reqOwner := req.Header.Get("OwnerId")
	certificate.Id = params["cert_id"]
	certificate.OwnerId = reqOwner
	fmt.Printf("%v\n", req)
	certificate, err = checkResponse(w, req, certificate)
	//if err != 1 {
	//	return
	//}
	//_ = json.NewDecoder(req.Body).Decode(&certificate)
	certificates = append(certificates, certificate)
	json.NewEncoder(w).Encode(certificate)

}

func UpdateCertificateEndpoint(w http.ResponseWriter, req *http.Request) {
	//CHECK IF USER_ID & CERTICATE_ID EXIST
	params := mux.Vars(req)
	err := checkUserId(w, req)
	if err != nil {
		return
	}
	err = checkCertificateId(w, req)
	if err != nil {
		return
	}

	var certificate Certificate
	reqOwner := req.Header.Get("OwnerId")

	certificate, err = checkResponse(w, req, certificate)

	if err != nil {
		return
	}
	certificate.Id = params["cert_id"]
	certificate.OwnerId = reqOwner

	for i := range certificates {
		if certificates[i].Id == params["cert_id"] && certificates[i].OwnerId == req.Header.Get("OwnerId") {
			if certificates[i].OwnerId == req.Header.Get("OwnerId") {
				certificates = append(certificates[:i], certificates[i+1:]...)
			}

			certificates = append(certificates, certificate)
			sort.Slice(certificates, func(i, j int) bool { return certificates[i].Id < certificates[j].Id })
			break

		}
	}
	json.NewEncoder(w).Encode(certificate)
}

func DeleteCertificateEndpoint(w http.ResponseWriter, req *http.Request) {
	//CHECK IF USER_ID & CERTICATE_ID EXIST
	err := checkUserId(w, req)
	if err != nil {
		return
	}
	err = checkCertificateId(w, req)
	if err != nil {
		return
	}

	params := mux.Vars(req)
	fmt.Printf("params:%v\n", params)
	for i := range certificates {
		if certificates[i].OwnerId == req.Header.Get("OwnerId") && certificates[i].Id == params["cert_id"] {
			json.NewEncoder(w).Encode(certificates[i])
			certificates = append(certificates[:i], certificates[i+1:]...)
			break
		}
	}
}

func TransferCertificateEndpoint(w http.ResponseWriter, req *http.Request) {
	//fmt.Printf("TRANSFER\n")
	//CHECK IF USER_ID & CERTICATE_ID EXIST
	err := checkUserId(w, req)
	if err != nil {
		return
	}
	err = checkCertificateId(w, req)
	if err != nil {
		return
	}

	params := mux.Vars(req)
	var certificate Certificate

	// COPY SELECTED CERTIFICATE
	for i := range certificates {
		if req.Header.Get("OwnerId") != params["user_id"] {
			fmt.Printf("user_id not equal to owner_id\n")
			return
		}

		if certificates[i].OwnerId == req.Header.Get("OwnerId") && certificates[i].Id == params["cert_id"] {
			certificate.Title = certificates[i].Title
			certificate.CreatedAt = certificates[i].CreatedAt
			certificate.Year = certificates[i].Year
			certificate.Note = certificates[i].Note
			certificate.OwnerId = certificates[i].OwnerId
			certificate.Id = certificates[i].Id
			certificates = append(certificates[:i], certificates[i+1:]...)
			break
		}
	}
	certificate, err = checkResponse(w, req, certificate)
	if err != nil {
		return
	}
	certificate.Transfer.Status = "pending"
	fmt.Printf("user_id:%v\n", params["user_id"])
	fmt.Printf("to:%v\n", certificate.Transfer.To)

	if params["user_id"] != certificate.Transfer.To {
		certificates = append(certificates, certificate)
		sort.Slice(certificates, func(i, j int) bool { return certificates[i].Id < certificates[j].Id })
	}
	json.NewEncoder(w).Encode(certificate)
}

func AcceptCertificateEndpoint(w http.ResponseWriter, req *http.Request) {
	//CHECK IF USER_ID & CERTICATE_ID EXIST
	err := checkUserId(w, req)
	if err != nil {
		return
	}
	err = checkCertificateId(w, req)
	if err != nil {
		return
	}

	params := mux.Vars(req)

	var certificate Certificate
	lastId := 0

	//RETRIEVE LAST CERTIFICATE ID FOR PROPER ID ORDER
	for i := range certificates {
		if certificates[i].OwnerId == req.Header.Get("OwnerId") {
			number, err := strconv.Atoi(certificates[i].Id)
			if err != nil {
				fmt.Printf("ERROR\n")
			}
			if lastId < number {
				lastId = number
			}
		}
		if params["user_id"] == certificate.OwnerId {
			return
		}
	}
	lastId = lastId + 1
	lastIdStr := strconv.Itoa(lastId)
	certificate.Id = lastIdStr

	certificate, err = checkResponse(w, req, certificate)

	if err != nil {
		return
	}
	// COPY SELECTED CERTIFICATE
	for i := range certificates {
		if certificates[i].OwnerId != req.Header.Get("OwnerId") && certificates[i].Id == params["cert_id"] {

			//certificate.OwnerId = params["user_id"]
			fmt.Printf("cert_ownid:%v\n", certificate.OwnerId)
			fmt.Printf("params[user_id]:%v\n", params["user_id"])
			fmt.Printf("certificate.Transfer.To:%v\n", certificate.Transfer.To)
			certificate.Title = certificates[i].Title
			certificate.CreatedAt = certificates[i].CreatedAt
			certificate.Year = certificates[i].Year
			certificate.Note = certificates[i].Note
			certificate.OwnerId = params["user_id"]

			if params["user_id"] != certificate.Transfer.To {
				sort.Slice(certificates, func(i, j int) bool { return certificates[i].Id < certificates[j].Id })
				certificates = append(certificates, certificate)
			}
			break
		}
	}
	// DELETE DUPLICATE
	for i := range certificates {
		if certificates[i].OwnerId != req.Header.Get("OwnerId") {
			fmt.Printf("%v\n", certificates[i])
			certificates = append(certificates[:i], certificates[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(certificate)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "VERISART")
	fmt.Println("Endpoint Hit: verisart")
}

func main() {
	router := mux.NewRouter()

	handler := cors.Default().Handler(router)

	//POPULATE "DATABASE" WITH SOME USERS AND CETIFICATES
	users = append(users, User{Id: "1", Email: "Vilko@gmail.com", Name: "Feldi"})
	users = append(users, User{Id: "2", Email: "Bejo@gmail.com", Name: "Zaere"})
	certificates = append(certificates, Certificate{Id: "1", Title: "SunScape", CreatedAt: "Florence", OwnerId: "1", Year: "1982", Note: "River"})
	certificates = append(certificates, Certificate{Id: "2", Title: "EveningGrass", CreatedAt: "Milan", OwnerId: "1", Year: "1972", Note: "Moon"})
	certificates = append(certificates, Certificate{Id: "1", Title: "Ecsta", CreatedAt: "Oxford", OwnerId: "2", Year: "2014", Note: "Shadow"})
	certificates = append(certificates, Certificate{Id: "2", Title: "Blisie", CreatedAt: "HongKong", OwnerId: "2", Year: "2018", Note: "Girls"})

	router.HandleFunc("/", homePage)
	router.HandleFunc("/users/{user_id}/certificates", GetCertificatesEndpoint).Methods("GET")
	router.HandleFunc("/users/{user_id}/certificates/{cert_id}", CreateCertificateEndpoint).Methods("POST")
	router.HandleFunc("/users/{user_id}/certificates/{cert_id}", UpdateCertificateEndpoint).Methods("PATCH")
	router.HandleFunc("/users/{user_id}/certificates/{cert_id}", DeleteCertificateEndpoint).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/certificates/{cert_id}/transfers", TransferCertificateEndpoint).Methods("POST")
	router.HandleFunc("/users/{user_id}/certificates/{cert_id}/transfers", AcceptCertificateEndpoint).Methods("PATCH")
	http.ListenAndServe(":8080", handler)
}
