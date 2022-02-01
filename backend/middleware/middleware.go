package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"backend/models"
	"github.com/gorilla/mux"
	//"backend/Searching"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://brent:roxxy@cluster0.sft8a.mongodb.net/test?retryWrites=true&w=majority"

const dbName = "test"

const collName = "todolist"

var collection *mongo.Collection

func init(){
	clientOptions :=options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(),clientOptions)

	if err != nil{
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(),nil)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

//GetAllTask get all the tasks route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	json.NewEncoder(w).Encode(payload)
}

//// KmpSearch
//func KmpSearch(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	payload := kmpSearch(pattern)
//	json.NewEncoder(w).Encode(payload)
//}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)
	// fmt.Println(task, r.Body)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}
func TaskComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// UndoTask undo the complete task route
func UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteTask delete one task route
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	// json.NewEncoder(w).Encode("Task not found")

}

// DeleteAllTask delete all tasks route
func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
	// json.NewEncoder(w).Encode("Task not found")

}

// get all task from the DB and return it
func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
	printAll()
}

// task complete method, update task's status to true
func taskComplete(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// task undo method, update task's status to false
func undoTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

//func bmSearch(substr string){
//	//get all of the tasks
//	//strings := bson.M{"task":task}
//
//	filter := bson.M{"Task": true}
//	tasks , err := collection.Distinct(context.Background(),"Task",filter)
//	fmt.Println(tasks)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	table := Searching.CalculateSlideTable(substr)
//	t := fmt.Sprintf("%v",tasks)
//	//to string conversion
//	strTask :=  make([]string{},len(tasks))
//	for i,v := range tasks{
//		strTask[i] = v
//	}
//
//	for i,v := range strTask {
//		if pos := Searching.IndexWithTable(&table, strTask[i], substr); pos > -1 {
//			fmt.Println("Found in position: ", pos)
//			print("found substr in task: ", t)
//
//		}
//	}

// need to  range loop over docs so i can get each task
func printAll() []bson.M{
	cursor,err := collection.Find(context.Background(),bson.M{})
	var docs []bson.M
	for cursor.Next(context.Background()) {
		var document bson.M
		err = cursor.Decode(&document)
		if err != nil {
			log.Println(err)
		}
		docs = append(docs, document)
		fmt.Println(docs[i]["task"])
	}
	return docs


}

//func kmpSearch(pattern string) bson.M{
//	cursor,err := collection.Find(context.Background(),bson.M{})
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	var data  = bson.M{}
//	if err = cursor.All(context.Background(),&data);err != nil{
//		log.Fatal(err)
//	}
//	Searching.Knp(bson.M{"task":"hello"},pattern string)
//
//
//}





//}

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}
