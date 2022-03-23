package topics

import (
	"accounting-service/core/environment"
	"accounting-service/events/handlers/company"
)

// Topics All topics to be consumed in the application will be registered here
type Topics struct {
	env            *environment.Environment
	companyHandler *company.EventHandler
	List           map[string]func([]byte, string) // Topics list
}

func New(env *environment.Environment, companyHandler *company.EventHandler) *Topics {
	var topics = map[string]func([]byte, string){
		"accounting.company-registered.request": companyHandler.HandleCompanyCreateEvent,
	}

	//println("Topics: ", len(topics))
	return &Topics{env: env, companyHandler: companyHandler, List: topics}
}
