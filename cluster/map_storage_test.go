package cluster_test

import (
	"testing"

	"github.com/thifnmi/mypaasdocker-cluster/cluster"
	storageTesting "github.com/thifnmi/mypaasdocker-cluster/storage/testing"
)

func TestMapStorageStorage(t *testing.T) {
	storageTesting.RunTestsForStorage(&cluster.MapStorage{}, t)
}
