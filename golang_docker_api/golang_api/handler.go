package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sort"
	"time"
	"strings"

)

// Just entring the dns name will be redirected to upload page
func Index(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/data/index.html", http.StatusFound)

}

// This will help for whether the server is up and running successfully
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome To the health check!")
}


// Function helps to get all the data data from database
func getalldata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	session, err := mgo.Dial(mongodbHostName+":"+mongodbPort)
	if err != nil {
	        panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)

	collection := session.DB(mongodbDatabaseName).C(mongodbCollectionName)

	err = collection.Find(nil).All(&allResult)
	if err != nil {
	        log.Fatal(err)
	}
	jsonResult, _ := json.Marshal(allResult)
	w.Write(jsonResult)
	io.WriteString(w, ",\n")
}

//Below function helps to find the respective data using id
func getdatabyid(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	session, err := mgo.Dial(mongodbHostName+":"+mongodbPort)
	if err != nil {
	        panic(err)
	}
	defer session.Close()
	collection := session.DB(mongodbDatabaseName).C(mongodbCollectionName)
	err = collection.Find(bson.M{"id": id}).One(&idResult)
	if err != nil {
		searchNotFoundText := "search Id not found"
		jsonResult, _ := json.Marshal(searchNotFoundText)
		w.Write(jsonResult)
		io.WriteString(w, "\n")
	} else {
		jsonResult, _ := json.Marshal(idResult)
		w.Write(jsonResult)
		io.WriteString(w, ",\n")
	}
}

// Below Function helps to get the data data using title regex search
func getdatabytitle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	searchtitle := vars["searchtitle"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	session, err := mgo.Dial(mongodbHostName+":"+mongodbPort)
	if err != nil {
	        panic(err)
	}
	defer session.Close()
	collection := session.DB(mongodbDatabaseName).C(mongodbCollectionName)

	err = collection.Find(bson.M{"title" : bson.M{"$regex":searchtitle}}).All(&regexResult)
	if err != nil {
		searchNotFoundText := "search Id not found"
		jsonResult, _ := json.Marshal(searchNotFoundText)
		w.Write(jsonResult)
		io.WriteString(w, "\n")
	} else {
		jsonResult, _ := json.Marshal(regexResult)
		w.Write(jsonResult)
		io.WriteString(w, ",\n")
	}
}



type AscendingSorter []data_created_at_time
func (a AscendingSorter) Len() int           { return len(a) }
func (a AscendingSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AscendingSorter) Less(i, j int) bool { return a[i].CreatedAt.Before(a[j].CreatedAt) }

type DescendingSorter []data_created_at_time
func (a DescendingSorter) Len() int           { return len(a) }
func (a DescendingSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DescendingSorter) Less(i, j int) bool { return a[i].CreatedAt.After(a[j].CreatedAt) }

// Below Function helps to get the data data by either ascending or descending order
func getdatasorted(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	orderby := vars["orderby"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	session, err := mgo.Dial(mongodbHostName+":"+mongodbPort)
	if err != nil {
	        panic(err)
	}
	defer session.Close()

	if strings.ToUpper(orderby) =="ASC"{
		collection := session.DB(mongodbDatabaseName).C(mongodbCollectionName)
		err = collection.Find(nil).All(&allResult)
		if err != nil {
			log.Fatal(err)
		}
		results_data_time := make([]data_created_at_time,len(allResult)) 
		layout := "Mon Jan 2 2006 15:04:05 GMT-0700 (MST)"
		for i := 0; i < len(allResult); i++ {
			results_data_time[i].Id = allResult[i].Id
			results_data_time[i].Picture=allResult[i].Picture
			results_data_time[i].Title=allResult[i].Title
			results_data_time[i].CreatedAt,err=time.Parse(layout, allResult[i].CreatedAt)
			results_data_time[i].dataCreator.CreatorId=allResult[i].dataCreator.CreatorId
			results_data_time[i].dataCreator.CreatorName=allResult[i].dataCreator.CreatorName
		}

		sort.Sort(AscendingSorter(results_data_time))
		jsonResult, _ := json.Marshal(results_data_time)
		w.Write(jsonResult)
		io.WriteString(w, ",\n")
	}else if strings.ToUpper(orderby) =="DESC"{
		collection := session.DB(mongodbDatabaseName).C(mongodbCollectionName)
		err = collection.Find(nil).All(&allResult)
		if err != nil {
			log.Fatal(err)
		}
		results_data_time := make([]data_created_at_time,len(allResult)) 
		layout := "Mon Jan 2 2006 15:04:05 GMT-0700 (MST)"
		for i := 0; i < len(allResult); i++ {
			results_data_time[i].Id = allResult[i].Id
			results_data_time[i].Picture=allResult[i].Picture
			results_data_time[i].Title=allResult[i].Title
			results_data_time[i].CreatedAt,err=time.Parse(layout, allResult[i].CreatedAt)
			results_data_time[i].dataCreator.CreatorId=allResult[i].dataCreator.CreatorId
			results_data_time[i].dataCreator.CreatorName=allResult[i].dataCreator.CreatorName
		}

		sort.Sort(DescendingSorter(results_data_time))
		jsonResult, _ := json.Marshal(results_data_time)
		w.Write(jsonResult)
		io.WriteString(w, ",\n")
	}else{
		searchNotFoundText := "search Id not found"
		jsonResult, _ := json.Marshal(searchNotFoundText)
		w.Write(jsonResult)
		io.WriteString(w, "\n")		
	}

}
