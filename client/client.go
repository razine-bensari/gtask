package client


import (
	"github.com/parnurzeal/gorequest"
)

//GoogleTaskClient ...
type GoogleTaskClient struct {
	request *gorequest.SuperAgent
	baseURL string
	token string
}

//Init ...
func (googleTaskClient *GoogleTaskClient) Init(token string) *GoogleTaskClient{
	clientWrapper := new(GoogleTaskClient)
	clientWrapper.baseURL = "https://www.googleapis.com/tasks/v1"
	clientWrapper.token = token
	clientWrapper.request = gorequest.New()
	return clientWrapper
}

//Returns all the authenticated user's task lists ...
func (googleTaskClient *GoogleTaskClient) getAllLists(user string) (gorequest.Response, string, []error){
	endpoint :=  googleTaskClient.baseURL + "/" + user + "/lists" //build endpoint url
	//Build Request
	return googleTaskClient.request.
			Get(endpoint).
			Set("access-token", "my-access-token"). //TODO: Add function that goes and fetches the access token
			End() //Needed after each request
}
//Returns the authenticated user's specified task list ...
func (googleTaskClient *GoogleTaskClient) getTaskList(user string, tasklist string) (gorequest.Response, string, []error){
	endpoint :=  googleTaskClient.baseURL + "/" + user + "/lists/" + tasklist //build endpoint url
	//Build Request
	return googleTaskClient.request.
			Get(endpoint).
			Set("access-token", "my-access-token"). //TODO: Add function that goes and fetches the access token
			End() //Needed after each request
}
//Creates a new task list and adds it to the authenticated user's task lists ...
func (GoogleTaskClient *GoogleTaskClient) createTaskList(user string, tasklistPayload TaskList) (gorequest.Response, string, []error){
	endpoint :=  googleTaskClient.baseURL + "/" + user + "/lists" //build endpoint url
	//Build Request
	return googleTaskClient.request.
			Post(endpoint).
			Set("access-token", "my-access-token"). //TODO: Add function that goes and fetches the access token
			End() //Needed after each request
}