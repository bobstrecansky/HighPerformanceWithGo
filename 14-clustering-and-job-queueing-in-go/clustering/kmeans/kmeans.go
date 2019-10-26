package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
)

func createDataset(datasetSize int) clusters.Observations {
	var dataset clusters.Observations
	for i := 1; i < datasetSize; i++ {
		dataset = append(dataset, clusters.Coordinates{
			rand.Float64(),
			rand.Float64(),
		})
	}
	return dataset
}

func printCluster(clusters clusters.Clusters) {
	for i, c := range clusters {
		fmt.Printf("\nCluster %d center points: x: %.2f y: %.2f\n", i, c.Center[0], c.Center[1])
		fmt.Printf("\nDatapoints assigned to this cluster: : %+v\n\n", c.Observations)
	}
}

func main() {

	var clusterSize = 3
	var datasetSize = 30
	var thresholdSize = 0.01
	rand.Seed(time.Now().UnixNano())
	dataset := createDataset(datasetSize)
	fmt.Println("Dataset: ", dataset)

	km, err := kmeans.NewWithOptions(thresholdSize, kmeans.SimplePlotter{})
	if err != nil {
		log.Printf("Your K-Means configuration struct was not initialized properly")
	}

	clusters, err := km.Partition(dataset, clusterSize)
	if err != nil {
		log.Printf("There was an error in creating your K-Means relation")
	}

	printCluster(clusters)
}
