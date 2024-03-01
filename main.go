package main

import (
	"fmt"
	"log"
	"nylas/nylas"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//Load the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	grantId := os.Getenv("GRANT_ID")

	//Get the thread with the ID 18df121c8f5b46df
	threadResponse := nylas.GetThread("18df121c8f5b46df", grantId)

	//varriable that will capture the ID of the Interview folder
	var threadfoldersIds []string
	threadfoldersIds = append(threadfoldersIds, threadResponse.Data.Folders...)

	//store the message IDs in an array
	messageIds := threadResponse.Data.MessageIDs

	//array that will have the subjects of the thread
	var subjects []string
	//array that will have the email objects that belongs to the thread
	var messages []nylas.MessageResponse
	//Array that will store the folders' names from the grant c0f2f708-567b-4a10-9b28-efad7d156283
	var folders []string

	//Capture all the message objects with each subject and store them in an array
	for i := 0; i < len(messageIds); i++ {
		message := nylas.GetEmail(messageIds[i], grantId)
		fmt.Println(message.Data.Body)
		messages = append(messages, message)
		subjects = append(subjects, message.Data.Subject)
	}

	//Get all the folders of the grant
	foldersResponse := nylas.GetFolders(grantId)
	for _, folder := range foldersResponse.Data {
		if folder.Name == "Interviews" { //capture the Interviews folder ID
			threadfoldersIds = append(threadfoldersIds, folder.Id)
		}
		folders = append(folders, folder.Name) //save all the folders names in an Array
	}

	//Create the body that will be send in the UpdateThread
	updateThreadBody := nylas.UpdateBody{
		Starred: false,
		Unread:  false,
		Folders: threadfoldersIds,
	}

	nylas.UpdateThread("18df121c8f5b46df", grantId, updateThreadBody)

}
