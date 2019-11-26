package main

import (
	"log"

	"github.com/omerkaya1/abf-guard/cmd"
)

//go:generate mockgen -source=internal/domain/interfaces/bucket/bucket.go -destination=internal/domain/interfaces/bucket/bucket_mock.go -package=bucket
//go:generate mockgen -source=internal/domain/interfaces/bucket/manager.go -destination=internal/domain/interfaces/bucket/manager_mock.go -package=bucket
//go:generate mockgen -source=internal/domain/interfaces/bucket/store.go -destination=internal/domain/interfaces/bucket/store_mock.go -package=bucket
//go:generate mockgen -source=internal/domain/interfaces/db/storage.go -destination=internal/domain/interfaces/db/storage_mock.go -package=db

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
