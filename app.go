package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/readit-tw/readit-api/model"
	"github.com/readit-tw/readit-api/repository"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
	"gopkg.in/mgo.v2/bson"
)

func searchListResourcesHandler(rr repository.ResourceRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
    	term := vars["term"]
    	
    	log.Printf("term:" + term)
    	
		resources, err := rr.SearchByTerm(term)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		resourcesJson, err := json.Marshal(resources)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resourcesJson)
	}
}

func listResourcesHandler(rr repository.ResourceRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		resources, err := rr.GetAll()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		resourcesJson, err := json.Marshal(resources)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resourcesJson)
	}
}
func createResourceHandler(rr repository.ResourceRepository, tr repository.TagRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		
		log.Printf("create request body term:")
		w.Header().Set("Content-Type", "application/json")

		resource := &model.Resource{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&resource)
		
		log.Printf("request body term:")

		if err != nil {
			http.Error(w, "Bad JSON Request", http.StatusBadRequest)
			return
		}

		validationErrors := resource.Validate()
		

		if len(validationErrors) != 0 {
			returnedValidationErrors := make(map[string][]map[string][]string)
			returnedValidationErrors["errors"] = validationErrors
			validationErrorsJson, err := json.Marshal(returnedValidationErrors)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(validationErrorsJson)
			return
		}

		// Get the content type from source link
		linkResp, errorContentType := http.Get(resource.Link) 
	    if errorContentType == nil { 
	        contentType := linkResp.Header.Get("Content-Type") 
	        resource.Type = contentType // default, assign header content-type
	        if strings.Contains(contentType, "text/html") {
	        	bodyContents, bodyErr := ioutil.ReadAll(linkResp.Body)
	        	if bodyErr == nil {
	        		if strings.Contains(string(bodyContents), "application/x-shockwave-flash") {
	        			 resource.Type = "application/x-shockwave-flash"
	        		}
	        	}
	        } 
	    }
	    
	    // Inserts all tags and pull Ids for input tags
		tagList := []bson.ObjectId{}
		
		log.Printf("tag creation from resource :")
		for index := range resource.Tags {
	    	tag := &model.Tag{}
	    	tag.Name = resource.Tags[index]
	    	log.Printf("tag creation from resource each tag :" + tag.Name)
	    	tagFound, err := tr.Create(tag)
	    	if tagFound != nil && err == nil {
	    		tagList = append(tagList, tagFound.Id)
	    		log.Printf("append creation from resource :" + tagFound.Name)
	    	}
	    }
		
		resource.Tag = tagList
	
	log.Printf("before save resource")
		created, err := rr.Create(resource)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		resourceJson, err := json.Marshal(created)
		
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(resourceJson)
	}
}

func listTagsHandler(rr repository.TagRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		tags, err := rr.GetAll()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tagsJson, err := json.Marshal(tags)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(tagsJson)
	}
}

func searchListTagsHandler(rr repository.TagRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
    	term := vars["term"]
    	
    	log.Printf("term:" + term)
    	
		tags, err := rr.SearchByTerm(term)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tagsJson, err := json.Marshal(tags)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(tagsJson)
	}
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("readit")
	
	tagRepository := repository.NewMongoTagRepository(db)
	resourceRepository := repository.NewMongoResourceRepository(db)

	r := mux.NewRouter()

	// Resource
	r.HandleFunc("/resources", createResourceHandler(resourceRepository, tagRepository)).Methods("POST")
	r.HandleFunc("/resources", listResourcesHandler(resourceRepository)).Methods("GET")
	r.HandleFunc("/resources/{term}", searchListResourcesHandler(resourceRepository)).Methods("GET")
	
	// Tag
	r.HandleFunc("/tags", listTagsHandler(tagRepository)).Methods("GET")
	r.HandleFunc("/tags/{term}", searchListTagsHandler(tagRepository)).Methods("GET")
	
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("../readit-web"))))

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
