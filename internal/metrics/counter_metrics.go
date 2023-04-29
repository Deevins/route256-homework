package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var UserCreateCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "create_user_requests",
	Help: "count of create_user_requests",
})

var UserDeleteCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "delete_user_requests",
	Help: "count of delete_user_requests",
})
