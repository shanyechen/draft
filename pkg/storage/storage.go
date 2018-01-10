package storage

import (
	"context"
)

// To regenerate the protocol buffer types for this package, run:
//		go generate

//go:generate protoc object.proto --go_out=.

// Deletor represents the delete APIs of the storage engine.
type Deletor interface {
	// DeleteBuilds deletes all draft builds for the application specified by appName.
	DeleteBuilds(ctx context.Context, appName string) ([]*Object, error)
	// DeleteBuild deletes the draft build given by buildID for the application specified by appName.
	DeleteBuild(ctx context.Context, appName, buildID string) (*Object, error)
}

// Creator represents the create APIs of the storage engine.
type Creator interface {
	// CreateBuild creates and stores a new build.
	CreateBuild(ctx context.Context, appName string, build *Object) error
}

// Getter represents the retrieval APIs of the storage engine.
type Getter interface {
	// GetBuilds retrieves all draft builds from storage.
	GetBuilds(ctx context.Context, appName string) ([]*Object, error)
	// GetBuild retrieves the draft build by id from storage.
	GetBuild(ctx context.Context, appName, buildID string) (*Object, error)
}

// Store represents a storage engine for application state stored by Draftd.
type Store interface {
	Creator
	Deletor
	Getter
}
