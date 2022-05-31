package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Contacts struct {
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("contacts.json")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Dont open file"))
		return
	}

	defer jsonFile.Close()

	var contacts Contacts

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &contacts)

	w.Write([]byte(byteValue))
}

func AddContact(w http.ResponseWriter, r *http.Request) {
	data := json.NewDecoder(r.Body)

	var dataContact Contact

	data.Decode(&dataContact)

	if dataContact.Name == "" || dataContact.Phone == "" || dataContact.Email == "" {
		w.Write([]byte("a data value is void"))
		return
	}

	newContact := new(Contact)

	var contacts Contacts

	byteValue, _ := ioutil.ReadFile("contacts.json")

	json.Unmarshal([]byte(byteValue), &contacts)

	newContact.Id = len(contacts.Contacts)
	newContact.Name = dataContact.Name
	newContact.Phone = dataContact.Phone
	newContact.Email = dataContact.Email

	contacts.Contacts = append(contacts.Contacts, *newContact)

	dataBytes, _ := json.Marshal(contacts)

	ioutil.WriteFile("contacts.json", dataBytes, 0644)
	w.Write([]byte("save sussefuly contact"))
}

func RemoveContact(w http.ResponseWriter, r *http.Request) {
	data := json.NewDecoder(r.Body)

	var dataContact Contact

	data.Decode(&dataContact)
}
