package main

import (
	"log"

	"github.com/omerkaya1/abf-guard/cmd"
)

//go:generate mockgen -source=internal/domain/bucket.go -destination=internal/domain/bucket_mock.go -package=domain
//go:generate mockgen -source=internal/domain/manager.go -destination=internal/domain/manager_mock.go -package=domain
//go:generate mockgen -source=internal/domain/storage.go -destination=internal/domain/storage_mock.go -package=domain
//go:generate mockgen -source=internal/db/storage.go -destination=internal/db/storage_mock.go -package=db

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
