package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"BookingsWebsite/pkg/config"
	"BookingsWebsite/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// ---------------------------------------------------------------------------------------------------------

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// If we are in the development mode then usecache is false and then we need to load each and every template instead of using the cached template
	if app.UseCache {
		// If we are in the Production mode
		// Get the template from the App Config file
		tc = app.TemplateCache
	} else {
		// If we are in the development mode
		tc, _ = CreateTemplateCache()
	}

	// Create all the template Cache
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Println("Error creating template cache = ", err)
	// 	log.Fatal(err)
	// }

	// Get requested template cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not find template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		fmt.Println("Error executing template with the help of buffer = ", err)
	}
	// Render the template
	// If all goes well then write the template to our response writer
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to writer = ", err)
	}
}

// Create Template Cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all the file names which matches the pattern
	// Glob returns only the name home.page.tmpl (In this it only returns the name of the template)
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Println("Error while getting templates related to this pattern: ", err)
		return myCache, err
	}

	// We get all the templates in pages that match the pattern
	// Range over all the templates
	for _, page := range pages {
		// In name it saves the base(extension) part of the template
		name := filepath.Base(page)
		// In the below it creates a new template with the base value and the name of the template
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			fmt.Println("Error while creating template  = ", err)
			return myCache, err
		}

		// Get all the layout templates which matches the pattern
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			fmt.Println("Error while getting all the layout templates = ", err)
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				fmt.Println("Error while applying all layout template to a particular page template: ", err)
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

// -----------------------------------------------------------------------------------------------------------
// Renders Template using html templates
// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("Error executing template", err)
// 		return
// 	}
// }

// ---------------------------------------------------------------------------------------------------------------
// Creating Template Cache
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// Check to see if the template is already in the map or not
// 	_, inMap := tc[t]
// 	// If the template is not available in the map then isMap = false
// 	if !inMap {
// 		log.Println("Creating a new template and add it to the cache")
// 		// need to create a new template
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println("Error creating template cache = ", err)
// 		}
// 	} else {
// 		// We have template in the cache
// 		log.Println("Using Cache template")
// 	}
// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println("Error creating template cache = ", err)
// 	}
// }

// // Function to create a new template cache entry
// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	// Parse the above template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		log.Println("Error parsing template= ", err)
// 		return err
// 	}

// 	// add the template to create cache(map)
// 	tc[t] = tmpl
// 	return nil
// }
