package main

import (
	"github.com/Psycho-micheal/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	const filename ="Large.in"
	const n = 64
	file, err := os.Create(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer,p)
	writer.Flush()


	file,err = os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file),-1)
	count := 0
	for v := range p{
		fmt.Print(v)
		count++
		if count >=100{
			break
		}
	}
}
func  mergeDemo(){
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3,2,6,7,4)),
		pipeline.InMemSort(
			pipeline.ArraySource(1,9,7,0,4,5)))
	for   v := range p{
		fmt.Println(v)
	}
}
