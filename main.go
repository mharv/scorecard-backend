/*
 * scorecard
 *
 * inital pass at scorecard API
 *
 * API version: 1.0
 * Contact: mitchell.harvey@arup.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//sw "github.com/GIT_USER_ID/GIT_REPO_ID/go"
	//
	Config "github.com/mharv/scorecard-backend/db"
	sw "github.com/mharv/scorecard-backend/go"
)

var err error

func main() {

	Config.ConnectDatabase()

	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(router.Run(":8080"))
}
