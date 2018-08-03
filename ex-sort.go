package main

import (
	"github.com/Psycho-micheal/pipeline"
	"os"
	"bufio"
	"fmt"
)

func main() {
	p :=  createPipeline("large.in",800000000,100)
	writetoFile(p,"large.out")
	printFile("large.out")
}

func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	p := pipeline.ReaderSource(file,-1)
	count := 0

	for v := range  p{
		fmt.Println(v)
		count++
		if count >= 100{
			break
		}
	}
}

func writetoFile(p <- chan int, filename string){
	file,err := os.Create(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer,p)
}

func createPipeline(
	filename string,
	fileSize, chunkCount int) <-chan int{
		chunkSize := fileSize / chunkCount
		pipeline.Init()

		sortResults := [] <- chan int{}
		for i := 0; i< chunkCount; i++{
			file,err := os.Open(filename)
			if err != nil{
				panic(err)
			}

			file.Seek(int64(i * chunkSize),0)

			source := pipeline.ReaderSource(
				bufio.NewReader(file),chunkSize)

			sortResults = append(sortResults,
			    pipeline.InMemSort(source))
		}
		return pipeline.MergeN(sortResults...)
}

