package main

import (
	"time"
)

//KIND ... (they are defaults values needed by googleapis)
const (
	KindTaskList string = "tasks#taskList"
	KindTask     string = "tasks#task"
)

//TaskList Resource ...
type TaskList struct {
	kind     string
	id       string
	etag     string
	title    string
	updated  string
	selfLink string
}

//NewTaskList ...
func NewTaskList() *TaskList {
	var TaskList = new(TaskList)
	TaskList.kind = KindTaskList
	return TaskList
}

//Kind ...
func (taskList TaskList) Kind() string {
	return taskList.kind
}

//ID ...
func (taskList TaskList) ID() string {
	return taskList.id
}

//SetID ...
func (taskList TaskList) SetID(ID string) {
	taskList.id = ID
}

//Etag ...
func (taskList TaskList) Etag() string {
	return taskList.etag
}

//SetEtag ...
func (taskList TaskList) SetEtag(etag string) {
	taskList.etag = etag
}

//Title ...
func (taskList TaskList) Title() string {
	return taskList.title
}

//SetTitle ...
func (taskList TaskList) SetTitle(title string) {
	taskList.title = title
}

//Updated ...
func (taskList TaskList) Updated() string {
	return taskList.updated
}

//SetUpdated ...
func (taskList TaskList) SetUpdated(updated string) {
	taskList.updated = updated
}

//SelfLink ...
func (taskList TaskList) SelfLink() string {
	return taskList.selfLink
}

//SetSelfLink ...
func (taskList TaskList) SetSelfLink(selfLink string) {
	taskList.selfLink = selfLink
}

//-------------------------------------------------------------------------------------------------------------------

//Task Resource ...
type Task struct {
	kind      string
	id        string
	etag      string
	title     string
	updated   time.Time
	selfLink  string
	parent    string
	position  string
	notes     string
	status    string
	due       time.Time
	completed time.Time
	deleted   bool
	hidden    bool
	links     []Link
}

//NewTask ...
func NewTask() *Task {
	var Task = new(Task)
	Task.kind = KindTask

	var Link = NewLink()
	Task.AddLink(*Link)

	return Task
}

//AddLink ...
func (task Task) AddLink(link Link) {
	task.links = append(task.links, link)
}

//Kind ...
func (task Task) Kind() string {
	return task.kind
}

//ID ...
func (task Task) ID() string {
	return task.id
}

//Etag ...
func (task Task) Etag() string {
	return task.etag
}

//Title ...
func (task Task) Title() string {
	return task.title
}

//Updated ...
func (task Task) Updated() time.Time {
	return task.updated
}

//SelfLink ...
func (task Task) SelfLink() string {
	return task.selfLink
}

//Parent ...
func (task Task) Parent() string {
	return task.parent
}

//Position ...
func (task Task) Position() string {
	return task.position
}

//Notes ...
func (task Task) Notes() string {
	return task.notes
}

//Status ...
func (task Task) Status() string {
	return task.status
}

//Due ...
func (task Task) Due() time.Time {
	return task.due
}

//Completed ...
func (task Task) Completed() time.Time {
	return task.completed
}

//Deleted ...
func (task Task) Deleted() bool {
	return task.deleted
}

//Hidden ...
func (task Task) Hidden() bool {
	return task.hidden
}

//Links ...
func (task Task) Links() []Link {
	return task.links
}

//SetID ...
func (task Task) SetID(ID string) {
	task.id = ID
}

//SetEtag ...
func (task Task) SetEtag(etag string) {
	task.etag = etag
}

//SetTitle ...
func (task Task) SetTitle(title string) {
	task.title = title
}

//SetUpdated ...
func (task Task) setUpdated(updated time.Time) {
	task.updated = updated
}

//SetSelfLink ...
func (task Task) SetSelfLink(selfLink string) {
	task.selfLink = selfLink
}

//SetParent ...
func (task Task) SetParent(parent string) {
	task.parent = parent
}

//SetPosition ...
func (task Task) SetPosition(position string) {
	task.position = position
}

//SetNotes ...
func (task Task) SetNotes(notes string) {
	task.notes = notes
}

//SetStatus ...
func (task Task) SetStatus(status string) {
	task.status = status
}

//SetDue ...
func (task Task) SetDue(due time.Time) {
	task.due = due
}

//SetCompleted ...
func (task Task) SetCompleted(completed time.Time) {
	task.completed = completed
}

//SetDeleted ...
func (task Task) SetDeleted(deleted bool) {
	task.deleted = deleted
}

//SetHidden ...
func (task Task) SetHidden(hidden bool) {
	task.hidden = hidden
}

//SetLinks ...
func (task Task) SetLinks(links []Link) {
	task.links = links
}

//Link Resource ...
type Link struct {
	_type       string
	description string
	_link       string
}

//NewLink ...
func NewLink() *Link {
	var Link = new(Link)
	return Link
}

//_type ...
func (link Link) _Type() string {
	return link._type
}

//Description ...
func (link Link) Description() string {
	return link.description
}

//Link ...
func (link Link) Link() string {
	return link._link
}

//_SetType ...
func (link Link) _SetType(_type string) {
	link._type = _type
}

//SetDescription ...
func (link Link) SetDescription(description string) {
	link.description = description
}

//SetLink ...
func (link Link) SetLink(_link string) {
	link._link = _link
}
