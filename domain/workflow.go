package domain

//Workflow is a structure to store the service request details
type Request struct {
	ServiceFlow ServiceFlow
}
type ServiceFlow struct {
	Description string `json:"description"`
	FlowMode    string `json:"flowMode"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Enabled     bool   `json:"enabled"`
	Steps       Steps  `json:"steps"`
}

type Steps struct {
	Step []Step `json:"step"`
}

type Step struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

//Create a new work flow for a given service flow and return service flow details
func CreateWorkflow(serviceFlowRequest Request) Request {
	return newServiceFlow(serviceFlowRequest)
}


func newServiceFlow(serviceFlow Request) Request {
	return Request{ServiceFlow:serviceFlow.ServiceFlow}
}