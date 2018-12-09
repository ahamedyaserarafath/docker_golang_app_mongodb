package main
import (
        "time"
)


// Json stuture in DB
type data struct {
        Id string `bson:"id"`
        Picture string `bson:"picture"`
        Title string `bson:"title"`
        CreatedAt string `bson:"createdAt"`

        dataCreator struct{
                 CreatorId string `bson:"id"`
                 CreatorName string `bson:"name"`
        } `bson:"creator"`

}


// Json stuture in for sorting
type data_created_at_time struct {
        Id string `bson:"id"`
        Picture string `bson:"picture"`
        Title string `bson:"title"`
        CreatedAt time.Time `bson:"createdAt"`
        dataCreator struct{
                 CreatorId string `bson:"id"`
                 CreatorName string `bson:"name"`
        } `bson:"creator"`

}
var allResult []data
var idResult data
var regexResult []data

var sortResult []data
