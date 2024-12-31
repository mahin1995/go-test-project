package page

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/mahin19/students-api/internal/storage"
	typesutils "github.com/mahin19/students-api/internal/typesUtils"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Parse the HTML template file
	tmplPath := "templates/" + tmpl
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

func HomeHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := storage.GetAllStudent()
		if err != nil {
			data := typesutils.PageData{
				Title:   "Error",
				Message: "Something went wrong",
				Error:   "Database error: unable to fetch students.",
			}
			renderTemplate(w, "index.html", data)
			return
		}

		data := typesutils.PageData{
			Title:    "Student List",
			Message:  "Here is the list of students:",
			Error:    "",
			Students: students,
		}
		renderTemplate(w, "index.html", data)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	data := typesutils.PageData{
		Title:   "Form Example",
		Message: "Please fill in the form below:",
	}
	renderTemplate(w, "form.html", data)
}

// Handler for processing the form submission
func SubmitHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form data", http.StatusBadRequest)
			return
		}

		// Retrieve form values
		name := r.FormValue("name")
		email := r.FormValue("email")
		age := r.FormValue("age")
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			fmt.Fprintf(w, "Form Submitted unsuccessfull!\n")
			fmt.Fprintf(w, "Name: %s\n", name)
			fmt.Fprintf(w, "Email: %s\n", email)
		}
		storage.CreateStudent(name, email, ageInt)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
