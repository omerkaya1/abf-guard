package main

import (
	"log"

	"github.com/omerkaya1/abf-guard/cmd"
)

//go:generate mockgen -source=internal/domain/bucket.go -destination=internal/bucket/bucket_mock.go -package=bucket
//go:generate mockgen -source=internal/domain/manager.go -destination=internal/bucket/manager_mock.go -package=bucket
//go:generate mockgen -source=internal/domain/storage.go -destination=internal/bucket/storage_mock.go -package=bucket
//go:generate mockgen -source=internal/db/storage.go -destination=internal/db/storage_mock.go -package=db

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
