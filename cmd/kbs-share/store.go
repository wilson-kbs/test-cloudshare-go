package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/tus/tusd/pkg/filelocker"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
	"github.com/tus/tusd/pkg/memorylocker"
	"github.com/tus/tusd/pkg/s3store"
)

func NewStore(storeType string) *tusd.StoreComposer {
	composer := tusd.NewStoreComposer()

	if storeType == "s3" {
		s3Config := aws.NewConfig()
		s3Config = s3Config.WithEndpoint("http://localhost:9000").WithS3ForcePathStyle(true)

		store := s3store.New("cloudshare", s3.New(session.Must(session.NewSession()), s3Config))
		store.ObjectPrefix = ""
		store.PreferredPartSize = 50 * 1024 * 1024
		store.DisableContentHashes = false

		locker := memorylocker.New()

		store.UseIn(composer)
		locker.UseIn(composer)
	} else {
		const dir = "./data/uploads"
		store := filestore.New(dir)
		locker := filelocker.New(dir)

		store.UseIn(composer)
		locker.UseIn(composer)
	}

	return composer
}
