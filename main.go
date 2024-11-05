package main
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)
const PORT = ":8080"
type outputData struct {
	PrintAscii string
}
var tmpl *template.Template
func init() {
	tmpl = template.Must(template.ParseGlob("view/*.html"))
}
// Method GET is used to request data from a specified resource and should have no other effect
// Method POST is used to send data to a server to create or update a resource
func main() {
	http.HandleFunc("/", home_page)          //home page
	http.HandleFunc("/ascii-art", ascii_art) //result page
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) //apply css
	fmt.Printf("Listening on port %v\n", PORT)
	fmt.Println("server started . . .")
	fmt.Println("ctrl + click: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(PORT, nil))
}
func home_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // error if wrong path
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil { //error if template not exists
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}
func ascii_art(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != "POST" {
		// If not, redirect the user to the index
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// Get the user input from the request form.
	input := strings.Split(r.FormValue("inputAscii"), "\r\n")
	if r.FormValue("inputAscii") == " " {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}
	//fmt.Println(r.FormValue("inputAscii"))
	// Create the ASCII art
	banner := r.FormValue("text_type")
	data := findAscii(input, banner)
	//tmpl.ExecuteTemplate(w, "ascii_art.html", nil)
	printAscii := outputData{
		PrintAscii: data,
	}
	tmpl.Execute(w, printAscii)
	//fmt.Print(findAscii(input, r.FormValue("text_type")))
}

func findAscii(splitText []string, inputFile string) string {
	var printAscii string
	content, _ := os.ReadFile("./fonts/" + strings.ToLower(inputFile) + ".txt")
	linesOFfile := strings.Split(string(content), "\n")
	for posNum := range splitText {
		if splitText[posNum] != " " {
			for i := 1; i < 9; i++ {
				for _, letter := range splitText[posNum] {
					if letter >= 32 && letter <= 126 {
						printAscii = printAscii + linesOFfile[((int(letter)-32)*9)+i]
					}
				}
				printAscii += "\n"
			}
		} else {
			printAscii += "\n"
		}
	}
	return printAscii
}
