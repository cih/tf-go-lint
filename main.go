package main

import (
	"fmt"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
	tensor, _ := tf.NewTensor([]int64{int64(1)})
	fmt.Printf("Tensor: %+v", tensor)
}
