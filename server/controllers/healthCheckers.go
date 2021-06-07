package controllers

import (
	"encoding/json"
	"fmt"
	"gin-first/server/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func HealthResponceAddsvc() gin.HandlerFunc {

	tags := []string{
		"primary",
	}

	// How to get information from lib and put in response
	checks := api.HealthChecks{
		&api.HealthCheck{
			Node:        "gin-first",
			CheckID:     "service:gin-first",
			Name:        "Service 'gin-first' check",
			Status:      "passing",
			Notes:       "",
			Output:      "",
			ServiceID:   "gin-first",
			ServiceName: "gin-first",
			ServiceTags: tags,
			Namespace:   "default",
		},
		&api.HealthCheck{
			Node:        "gin-first",
			CheckID:     "serfHealth",
			Name:        "Serf Health Status",
			Status:      "passing",
			Notes:       "",
			Output:      "",
			ServiceID:   "",
			ServiceName: "",
			ServiceTags: tags,
			Namespace:   "default",
		},
	}

	serviceEntry := api.ServiceEntry{
		Node: &api.Node{
			ID:         "40e4a748-2192-161a-0510-9bf59fe950b5",
			Node:       "gin-first",
			Address:    "127.0.0.1",
			Datacenter: "dc1",
			TaggedAddresses: map[string]string{
				"lan": "127.0.0.1",
				"wan": "127.0.0.1",
			},
			Meta: map[string]string{
				"instance_type": "t2.medium",
			},
		},
		Service: &api.AgentService{
			ID:      "redis",
			Service: "redis",
			Tags:    tags,
			Address: "127.0.0.1",
			TaggedAddresses: map[string]api.ServiceAddress{
				"lan": {
					Address: "127.0.0.1",
					Port:    4444,
				},
				"wan": {
					Address: "198.18.1.2",
					Port:    80,
				},
			},
			Meta: map[string]string{
				"redis_version": "4.0",
			},
			Port: 4444,
			Weights: api.AgentWeights{
				Passing: 10,
				Warning: 1,
			},
			Namespace: "default",
		},
		Checks: checks,
	}

	res := []api.ServiceEntry{
		serviceEntry,
	}

	return func(c *gin.Context) {

		t, err := helpers.TlsTcpConn()

		if err != nil {
			fmt.Println("Tls conn error ---", err)
		}

		defer t.Close()

		fmt.Println("HealthResponceAddsvc entry ---", t)

		fmt.Println("HealthResponceAddsvc clother func ---", c, t)
		// resJson, err := json.Marshal(res)
		// if err != nil {
		// 	c.JSON(401, err)
		// }

		// c.JSON(http.StatusOK, resJson)

		resJson, err := json.Marshal(res)

		fmt.Println("HealthResponceAddsvc closer func 2---", resJson)

		if err != nil {
			log.Fatal("resJson", err)
		}

		n, err := t.Write(resJson)
		// d := []byte("hello ----- this is kiss\n")
		// fmt.Println("Test data d ----", d)
		// n, err := t.Write(d)

		fmt.Println("HealthResponceAddsvc clother func 3---", n)

		if err != nil {
			fmt.Println("HealthResponceAddsvc Write error ---", n, err)
			return
		}

		buf := make([]byte, 1000000)

		fmt.Println("HealthResponceAddsvc clother func 3.5---", n)

		n, err = t.Read(buf)

		fmt.Println("HealthResponceAddsvc clother func 4---", n)

		if err != nil {
			fmt.Println(n, err)
			return
		}

		println(string(buf[:n]))

		fmt.Println("HealthResponceAddsvc clother func 5---", n)

		c.JSON(http.StatusOK, string(buf[:n]))
		// c.JSON(http.StatusOK, map[string]string{
		// 	"hello": "Hello go--------",
		// })

	}

}

func HealthResponceStringsvc() gin.HandlerFunc {

	tags := []string{
		"primary",
	}

	// How to get information from lib and put in response
	checks := api.HealthChecks{
		&api.HealthCheck{
			Node:        "foobar",
			CheckID:     "service:redis",
			Name:        "Service 'redis' check",
			Status:      "passing",
			Notes:       "",
			Output:      "",
			ServiceID:   "redis",
			ServiceName: "redis",
			ServiceTags: tags,
			Namespace:   "default",
		},
		&api.HealthCheck{
			Node:        "foobar",
			CheckID:     "serfHealth",
			Name:        "Serf Health Status",
			Status:      "passing",
			Notes:       "",
			Output:      "",
			ServiceID:   "",
			ServiceName: "",
			ServiceTags: tags,
			Namespace:   "default",
		},
	}

	serviceEntry := api.ServiceEntry{
		Node: &api.Node{
			ID:         "40e4a748-2192-161a-0510-9bf59fe950b5",
			Node:       "foobar",
			Address:    "10.1.10.12",
			Datacenter: "dc1",
			TaggedAddresses: map[string]string{
				"lan": "10.1.10.12",
				"wan": "10.1.10.12",
			},
			Meta: map[string]string{
				"instance_type": "t2.medium",
			},
		},
		Service: &api.AgentService{
			ID:      "redis",
			Service: "redis",
			Tags:    tags,
			Address: "10.1.10.12",
			TaggedAddresses: map[string]api.ServiceAddress{
				"lan": {
					Address: "10.1.10.12",
					Port:    8000,
				},
				"wan": {
					Address: "198.18.1.2",
					Port:    80,
				},
			},
			Meta: map[string]string{
				"redis_version": "4.0",
			},
			Port: 8000,
			Weights: api.AgentWeights{
				Passing: 10,
				Warning: 1,
			},
			Namespace: "default",
		},
		Checks: checks,
	}

	res := []api.ServiceEntry{
		serviceEntry,
	}

	return func(c *gin.Context) {

		fmt.Println(res)
		// resJson, err := json.Marshal(res)
		// if err != nil {
		// 	c.JSON(401, err)
		// }

		// c.JSON(http.StatusOK, resJson)
		c.JSON(http.StatusOK, res)

	}
}

// func HealthResponceRedis() gin.HandlerFunc {

// 	tags := []string{
// 		"primary",
// 	}

// 	// How to get information from lib and put in response
// 	checks := api.HealthChecks{
// 		&api.HealthCheck{
// 			Node:        "foobar",
// 			CheckID:     "service:redis",
// 			Name:        "Service 'redis' check",
// 			Status:      "passing",
// 			Notes:       "",
// 			Output:      "",
// 			ServiceID:   "redis",
// 			ServiceName: "redis",
// 			ServiceTags: tags,
// 			Namespace:   "default",
// 		},
// 		&api.HealthCheck{
// 			Node:        "foobar",
// 			CheckID:     "serfHealth",
// 			Name:        "Serf Health Status",
// 			Status:      "passing",
// 			Notes:       "",
// 			Output:      "",
// 			ServiceID:   "",
// 			ServiceName: "",
// 			ServiceTags: tags,
// 			Namespace:   "default",
// 		},
// 	}

// 	serviceEntry := api.ServiceEntry{
// 		Node: &api.Node{
// 			ID:         "40e4a748-2192-161a-0510-9bf59fe950b5",
// 			Node:       "foobar",
// 			Address:    "10.1.10.12",
// 			Datacenter: "dc1",
// 			TaggedAddresses: map[string]string{
// 				"lan": "10.1.10.12",
// 				"wan": "10.1.10.12",
// 			},
// 			Meta: map[string]string{
// 				"instance_type": "t2.medium",
// 			},
// 		},
// 		Service: &api.AgentService{
// 			ID:      "redis",
// 			Service: "redis",
// 			Tags:    tags,
// 			Address: "10.1.10.12",
// 			TaggedAddresses: map[string]api.ServiceAddress{
// 				"lan": {
// 					Address: "10.1.10.12",
// 					Port:    8000,
// 				},
// 				"wan": {
// 					Address: "198.18.1.2",
// 					Port:    80,
// 				},
// 			},
// 			Meta: map[string]string{
// 				"redis_version": "4.0",
// 			},
// 			Port: 8000,
// 			Weights: api.AgentWeights{
// 				Passing: 10,
// 				Warning: 1,
// 			},
// 			Namespace: "default",
// 		},
// 		Checks: checks,
// 		// "Checks": [
// 		//   {
// 		// 	"Node": "foobar",
// 		// 	"CheckID": "service:redis",
// 		// 	"Name": "Service 'redis' check",
// 		// 	"Status": "passing",
// 		// 	"Notes": "",
// 		// 	"Output": "",
// 		// 	"ServiceID": "redis",
// 		// 	"ServiceName": "redis",
// 		// 	"ServiceTags": ["primary"],
// 		// 	"Namespace": "default"
// 		//   },
// 		//   {
// 		// 	"Node": "foobar",
// 		// 	"CheckID": "serfHealth",
// 		// 	"Name": "Serf Health Status",
// 		// 	"Status": "passing",
// 		// 	"Notes": "",
// 		// 	"Output": "",
// 		// 	"ServiceID": "",
// 		// 	"ServiceName": "",
// 		// 	"ServiceTags": [],
// 		// 	"Namespace": "default"
// 		//   }
// 		// ]
// 	}

// 	res := []api.ServiceEntry{
// 		serviceEntry,
// 	}

// 	return func(c *gin.Context) {

// 		fmt.Println(res)
// 		// resJson, err := json.Marshal(res)
// 		// if err != nil {
// 		// 	c.JSON(401, err)
// 		// }

// 		// c.JSON(http.StatusOK, resJson)
// 		c.JSON(http.StatusOK, res)

// 	}
// }
