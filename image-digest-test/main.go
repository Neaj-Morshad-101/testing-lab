package main

import (
	"log"
	"time"

	"github.com/google/go-containerregistry/pkg/crane"
	"k8s.io/klog/v2"
)

func ImageWithDigest(image string) (string, error) {
	klog.Infoln("ImageWithDigest:", image)
	start := time.Now()

	// Drop the "@sha256:hash_string" part, if any (helper function assumed)
	// For demonstration, we're skipping auth and using simple image reference.
	digest, err := crane.Digest(image)
	if err != nil {
		klog.Errorln("Failed to get crane.Digest", image)
		return "", err
	}

	klog.Infof("ImageWithDigest execution time: %v\n", time.Since(start))
	return image + "@" + digest, nil
}

func main() {

	image := "mcr.microsoft.com/mssql/server:2022-CU12-ubuntu-22.04"
	result, err := ImageWithDigest(image)
	if err != nil {
		log.Fatalf("Failed to get image digest: %v", err)
	}
	klog.Infoln("Image with digest:", result)

	image = "neajmorshad/mssql-init-docker:upd"
	result, err = ImageWithDigest(image)
	if err != nil {
		log.Fatalf("Failed to get image digest: %v", err)
	}
	klog.Infoln("Image with digest:", result)
}
