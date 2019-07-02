package client


import (
	"github.com/parnurzeal/gorequest"
	"gtask/domain/gTask"
)

//GoogleTaskClient ...
type GoogleTaskClient struct {
	request *gorequest.SuperAgent
	baseURL string
	token string
}

//NewGoogleTaskClient ...
func NewGoogleTaskClient(token string) *GoogleTaskClient {
	googleTaskClient := new(GoogleTaskClient)
	googleTaskClient.baseURL = "https://www.googleapis.com/tasks/v1"
	googleTaskClient.token = token
	googleTaskClient.request = gorequest.New()
	return googleTaskClient
}

//Returns all the authenticated user's task lists ...
func (googleTaskClient GoogleTaskClient) getAllLists() (gorequest.Response, string, []error) {
	endpoint :=  googleTaskClient.baseURL + "/@me/lists" //build endpoint url
	//Build Request
	return googleTaskClient.request.
			Get(endpoint).
			Set("access-token", "my-access-token"). //TODO: Add function that goes and fetches the access token
			End() //Needed after each request
}
//Returns the authenticated user's specified task list ...
func (googleTaskClient GoogleTaskClient) getTaskList(tasklist string) (gorequest.Response, string, []error) {
	endpoint :=  googleTaskClient.baseURL + "/@me/lists/" + tasklist //build endpoint url
	//Build Request
	return googleTaskClient.request.
			Get(endpoint).
			Set("access-token", "my-access-token").//TODO
			End()
}
//Creates a new task list and adds it to the authenticated user's task lists ...
func (googleTaskClient GoogleTaskClient) createTaskList(taskList TaskList) (gorequest.Response, string, []error) {
	endpoint :=  googleTaskClient.baseURL + "/@me/lists" 
	return googleTaskClient.request.
			Post(endpoint).
			Set("access-token", "my-access-token").//TODO
			Set("Accept","application/json").
			Send(taskList). //GoRequest marshals struct into json format using Send()
			End()
}
//Updates the authenticated user's specified task list ...
func (googleTaskClient GoogleTaskClient) updateTaskList(taskList string, taskList TaskList) (gorequest.Response, string, []error) {
	endpoint :=  googleTaskClient.baseURL + "/@me/lists/" + tasklist
	return googleTaskClient.request.
			Put(endpoint).
			Set("access-token", "my-access-token").//TODO
			Set("Accept","application/json").
			Send(taskList).
			End()
}
//Deletes the authenticated user's specified task list ...
func (googleTaskClient GoogleTaskClient) deleteTaskList(taskList string, taskList TaskList) (gorequest.Response, string, []error) {
	endpoint :=  googleTaskClient.baseURL + "/@me/lists/" + tasklist
	return googleTaskClient.request.
			Delete(endpoint).
			Set("access-token", "my-access-token").//TODO
			End()
}
//Returns all tasks in the specified task list.
func (googleTaskClient GoogleTaskClient) getAllTasksInList(tasklist string) (gorequest.Response, string, []error) {
	endpoint :=  googleTaskClient.baseURL + "/@me/lists/" + tasklist + "tasks"
	return googleTaskClient.request.
			Get(endpoint).
			Set("access-token", "my-access-token").//TODO
			End()
}
//Returns the specified task.
func (googleTaskClient GoogleTaskClient) getTaskInList(tasklist string, task string) (gorequest.Response, string, []error) {
	endpoint :=  googleTaskClient.baseURL + "/@me/lists/" + tasklist + "/" + task
	return googleTaskClient.request.
			Get(endpoint).
			Set("access-token", "my-access-token").//TODO
			End()
}
//Creates a new task on the specified task list.
func (googleTaskClient GoogleTaskClient) createTaskInList(tasklist string, task Task) (gorequest.Response, string, []error) {
	endpoint :=  googleTaskClient.baseURL + "/@me/lists/" + tasklist + "tasks"
	return googleTaskClient.request.
			Put(endpoint).
			Set("access-token", "my-access-token").//TODO
			Set("Accept","application/json").
			Send(task).
			End()
}