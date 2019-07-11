package client

import (
	"gorequest"
	gtask "gtask/domain"
)

//GoogleTaskClient ...
type GoogleTaskClient struct {
	request *gorequest.SuperAgent
	baseURL string
	token   string
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
func (googleTaskClient GoogleTaskClient) GetAllLists() (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/users/@me/lists" //build endpoint url
	//Build Request
	return googleTaskClient.request.
		Get(endpoint).
		Set("access-token", googleTaskClient.token). //TODO: Add function that goes and fetches the access token
		End()                                        //Needed after each request
}

//Returns all the authenticated user's task lists ...
func (googleTaskClient GoogleTaskClient) GetAllListsWithQuery(query string) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/users/@me/lists" //build endpoint url
	//Build Request
	return googleTaskClient.request.
		Get(endpoint).
		Set("access-token", googleTaskClient.token). //TODO: Add function that goes and fetches the access token
		Query(query).
		End() //Needed after each request
}

//Returns the authenticated user's specified task list ...
func (googleTaskClient GoogleTaskClient) GetTaskList(taskList gtask.TaskList) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/users/@me/lists/" + taskList.ID() //build endpoint url
	//Build Request
	return googleTaskClient.request.
		Get(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		End()
}

//Creates a new task list and adds it to the authenticated user's task lists ...
func (googleTaskClient GoogleTaskClient) CreateTaskList(taskList gtask.TaskList) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/users/@me/lists"
	return googleTaskClient.request.
		Post(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		Set("Accept", "application/json").
		Send(taskList). //GoRequest marshals struct into json format using Send()
		End()
}

//Updates the authenticated user's specified task list ...
func (googleTaskClient GoogleTaskClient) UpdateTaskList(taskList gtask.TaskList) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/users/@me/lists/" + taskList.ID()
	return googleTaskClient.request.
		Put(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		Set("Accept", "application/json").
		Send(taskList).
		End()
}

//Deletes the authenticated user's specified task list ...
func (googleTaskClient GoogleTaskClient) DeleteTaskList(taskList gtask.TaskList) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/users/@me/lists/" + taskList.ID()
	return googleTaskClient.request.
		Delete(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		End()
}

//Returns all tasks in the specified task list.
func (googleTaskClient GoogleTaskClient) GetAllTasksInList(taskList gtask.TaskList) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/lists/" + taskList.ID() + "tasks"
	return googleTaskClient.request.
		Get(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		End()
}

//Returns all tasks in the specified task list.
func (googleTaskClient GoogleTaskClient) GetAllTasksInListWithQuery(taskList gtask.TaskList, query string) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/lists/" + taskList.ID() + "tasks"
	return googleTaskClient.request.
		Get(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		Query(query).
		End()
}

//Returns the specified task.
func (googleTaskClient GoogleTaskClient) GetTaskInList(taskList gtask.TaskList, task gtask.Task) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/lists/" + taskList.ID() + "/tasks/" + task.ID()
	return googleTaskClient.request.
		Get(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		End()
}

//Creates a new task on the specified task list.
func (googleTaskClient GoogleTaskClient) CreateTaskInList(taskList gtask.TaskList, task gtask.Task) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/lists/" + taskList.ID() + "tasks"
	return googleTaskClient.request.
		Post(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		Set("Accept", "application/json").
		Send(task).
		End()
}

//Updates the specified task.
func (googleTaskClient GoogleTaskClient) UpdateTaskInList(taskList gtask.TaskList, task gtask.Task) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/lists/" + taskList.ID() + "/tasks/" + task.ID()
	return googleTaskClient.request.
		Put(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		Set("Accept", "application/json").
		Send(task).
		End()
}

//Deletes the specified task from the task list.
func (googleTaskClient GoogleTaskClient) DeleteTaskInList(taskList gtask.TaskList, task gtask.Task) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/lists/" + taskList.ID() + "/tasks/" + task.ID()
	return googleTaskClient.request.
		Delete(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		End()
}

//Clears all completed tasks from the specified task list.
//The affected tasks will be marked as 'hidden' and no longer be returned by default when retrieving all tasks for a task list.
func (googleTaskClient GoogleTaskClient) ClearAllTasksInTaskList(taskList gtask.TaskList) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/lists/" + taskList.ID() + "/clear"
	return googleTaskClient.request.
		Post(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		End()
}

//Moves the specified task to another position in the task list.
//This can include putting it as a child task under a new parent and/or move it to a different position among its sibling tasks.
func (googleTaskClient GoogleTaskClient) MoveTaskInTaskList(taskList gtask.TaskList, task gtask.Task) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/lists/" + taskList.ID() + "/tasks/" + task.ID() + "move"
	return googleTaskClient.request.
		Post(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		End()
}

//Moves the specified task to another position in the task list.
//This can include putting it as a child task under a new parent and/or move it to a different position among its sibling tasks.
func (googleTaskClient GoogleTaskClient) MoveTaskInTaskListWithQuery(taskList gtask.TaskList, task gtask.Task, query string) (gorequest.Response, string, []error) {
	endpoint := googleTaskClient.baseURL + "/lists/" + taskList.ID() + "/tasks/" + task.ID() + "move"
	return googleTaskClient.request.
		Post(endpoint).
		Set("access-token", googleTaskClient.token). //TODO
		Query(query).
		End()
}
