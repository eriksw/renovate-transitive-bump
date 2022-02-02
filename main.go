package main

import (
	"context"
	"io"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
)

func main() {
	ctx := context.Background()
	bucket, err := blob.OpenBucket(ctx, "s3://non-exitent-test-bucket")
	if err != nil {
		panic(err)
	}
	iter := bucket.List(nil)

	for {
		_, err := iter.Next(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}

	err = bucket.Close()
	if err != nil {
		panic(err)
	}
}
