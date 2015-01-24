#Go-Tasky
Tasky is a simple go framework that makes it easy to expose server side tasks with a RESTful api. Developed for the gophergala global hackathon.  

NOTE: this is preliminary ideas and nothing is set in stone. Consider it brainstorming via readme.  

##Basic Principles

Register a task with go-tasky that implements the worker interface and a set of routes will be created for that task. 

##Routes

GET /tasky/v1/workers - returns a list of available worker endpoints
GET /tasky/v1/workers/{worker_name} - returns a description of the worker and it's usage  
  
GET /tasky/v1/{worker_name} - returns a list of available tasks to run  
POST /tasky/v1/{worker_name} - Creates a new task to run with the worker  

PUT /tasky/v1/{worker_name} - Modify the state of the task (cancel, pause, resume)  
GET /tasky/v1/{worker_name}/status - returns the status of the worker  

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
