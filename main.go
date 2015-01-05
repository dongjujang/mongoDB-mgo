package main

import (
	"fmt"
  "log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
          Name string
          Phone string
}

func main() {
  
  session, err := mgo.Dial("127.0.0.1")//establishes a new session to the cluster identified by the given seed server
  if err != nil {
    panic(err)
  }
  defer session.Close()
      
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("test").C("people")//In the session,  make "people"collection into db"test"
  err =  c.Insert(&Person{"alole", "+55 5555 5555"}, //Insert data type of "Person"
                  &Person{"Cla", "+65 2938 2937"},
                  &Person{"baek", "+10 3425 3423"})
  
  result := Person{}
  err = c.Find(bson.M{"name": "Cla"}).One(&result)//find one data by using selector(bson.M) and store into result
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Phone:", result.Phone)
  c.Update(bson.M{"name": "Cla"}, bson.M{"name": "jang", "phone": "+10 8407 3212"})//update from first data to second data
  c.Update(bson.M{"name": "alole"}, bson.M{"name": "hot", "phone": "+09 3938 3233"})

  c.Remove(bson.M{"name": "baek"})//remove db by using selector
  //  c.RemoveAll(bson.M{})
}

  


