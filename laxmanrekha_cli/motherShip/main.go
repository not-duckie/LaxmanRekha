package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type malware struct {
	Name    string
	Md5sum  string
	Sha1sum string
	Count   int
}

var samples = make(map[string]*malware)

func addSample(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if i, ok := samples[r.PostForm["sha256sum"][0]]; ok {
		i.Count++
		w.Write([]byte("Alreay have this sample.\n"))
		w.Write([]byte("Thank you for your time\n"))
		return
	}

	fmt.Println(r.PostForm)
	samples[r.PostForm["sha256sum"][0]] = &malware{
		Name:    r.PostForm["name"][0],
		Md5sum:  r.PostForm["md5sum"][0],
		Sha1sum: r.PostForm["sha1sum"][0],
		Count:   1,
	}

	w.Write([]byte("Sample Added.\n"))
	w.Write([]byte("Thank you"))

}

func returnSamples(w http.ResponseWriter, r *http.Request) {
	jsonString, err := json.Marshal(samples)
	if err != nil {
		w.Write([]byte("Something Went Wrong"))
	}
	w.Write([]byte(jsonString))
}

func home(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("malwareSamples.html")
	temp.Execute(w, samples)
}

func main() {

	http.HandleFunc("/addSample", addSample)
	http.HandleFunc("/", home)
	http.HandleFunc("/api/samples", returnSamples)
	log.Println("Starting listening on Port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
