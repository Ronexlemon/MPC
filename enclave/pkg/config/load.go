package config

import "os"

func LoadEnv()(string,string){
	nodeName := os.Getenv("NODE_NAME")
	nodeIndexStr := os.Getenv("NODE_INDEX")
	return nodeName,nodeIndexStr

}