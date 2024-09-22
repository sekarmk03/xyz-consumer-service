package authorization

type AccessibleRoles map[string]map[string][]uint32

const (
	BasePath = "xyz-consumer-service"
	ConsumerSvc  = "ConsumerService"
	LimitSvc = "LimitService"
)

var roles = AccessibleRoles{
	"/" + BasePath + "." + ConsumerSvc + "/": {
		// "DeletePost":  {1, 2, 8},
	},
	"/" + BasePath + "." + LimitSvc + "/": {
		// "DeleteComment": {1, 2, 8},
	},
}

func GetAccessibleRoles() map[string][]uint32 {
	routes := make(map[string][]uint32)

	for service, methods := range roles {
		for method, methodRoles := range methods {
			route := service + method
			routes[route] = methodRoles
		}
	}

	return routes
}
