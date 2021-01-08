package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main(){
	//for i := 0; i < 10; i++ {
	//	go getInstance()
	//}
	//fmt.Scanln()
	wrappedMain()
}

func wrappedMain() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	wi := &Wiki{Title: title}
	content, err := wi.loadContent()
	if err != nil {
		fmt.Println("Page is not found", err)
	}
	wi.Content = content

	t, _ := template.ParseFiles("../resources/wiki_view.html")
	t.Execute(w, wi)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	wi := &Wiki{Title: title}
	switch r.Method {
	case "GET":
		content, err := wi.loadContent()
		if err != nil {
			fmt.Println("Page is not found", err)
		}
		wi.Content = content

		t, _ := template.ParseFiles("../resources/wiki_edit.html")
		t.Execute(w, wi)
	case "POST":
		content := r.FormValue("content")

		wi := &Wiki{Title: title, Content: content}
		err := wi.save()
		if err != nil {
			fmt.Println("Cannot save file", err)
		}
		http.Redirect(w, r, "/view/" + title, http.StatusFound)
	}
}