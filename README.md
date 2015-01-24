#Go-Tasky
Tasky is a simple go tool that makes it easy to expose server side tasks with a RESTful api. Developed for the gophergala global hackathon. 

NOTE: this is preliminary ideas and nothing is set in stone. Consider it brainstorming via readme.  

##Basic Principles

Register a task with go-tasky that implements the worker interface and a set of routes will be created for that task. You can have any number of uniquely configured tasks for a worker. The go-tasky app would run on your server and expose api endpoints.  

##Workers
- copy a file to a new location  
- read the contents of a file  
- check the values of system settings 
- !TODO add more!  

##Use Case Example
Use this tool when debugging or checking values on your system.

##Routes
Workers:  
GET /tasky/v1/workers - returns a list of available worker endpoints   
GET /tasky/v1/workers/{worker_name} - returns a list of available tasks to run  
GET /tasky/v1/workers/{worker_name}/info - returns a description of the worker and it's usage   
POST /tasky/v1/workers/{worker_name} - Creates a new task to run with the worker and returns a unique task id  

Tasks:  
POST /tasky/v1/task - Create a new task. !!Should we model the task creation to be done here or from the worker?!!  
GET /tasky/v1/task/{task_id} - Fetch details of a single task.  
PUT /tasky/v1/task/{task_id} - Update the configuration of the task.  

POST /tasky/v1/task/{task_id}/actions - Modify the state of the task (cancel, pause, resume, run)  
GET /tasky/v1/task/{task_id}/status - returns the status of the task  
GET /tasky/v1/task/{task_id}/statistics - returns the statistics about the task, such as time to complete task  

RuleChains:  
For later, but used to chain multiple tasks together in an ordered fashion.  


## Worker Interface
The worker interface corresponds to an individual worker type
```go
type Worker interface {
    //provides a description and usage instructions for the worker
    Info()
    
    //implements the logic to create a new worker(config, etc)
    Create()
    
    //returns status of the worker
    Status()

    //controls to modify the state of the task(start, cancel, pause, resume)
    Update()

    //cleanup tasks when removing the task
    Delete()
}
```


## Task Interface
The task interface corresponds to details about a specific task running with a worker.  
```go
type Task interface {
    // logic to create a new task, called from the worker create endpoint
    Create()

    // returns the details of an individual task 
    Read(*Context)

    // modify the configuration of an individual task
    Update(*Context)

    // delete an individual task
    Delete(*Context)

}

```
