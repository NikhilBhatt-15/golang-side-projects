package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func async(ch chan string) {
	for i:=0;i<3;i++{
		time.Sleep(500 * time.Millisecond)
		ch <- "sending data from marrs"
	}
	close(ch)
}

func main() {
	dup()

}

func dup(){
	count := make(map[string]int)
	files := os.Args[1:]
	if len(files) ==0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan(){
			f,err := os.Open(input.Text())
			if err!=nil{
				fmt.Fprintf(os.Stderr,"Error in opening file %s\t %v\n",input.Text(),err)
			}
			countlines(f,count)
			f.Close()
		}
	}else{
		for _,file := range files{
			f,err := os.Open(file)
			if err!=nil{
				fmt.Fprintf(os.Stderr,"dup: %v\n",err)
			}
			countlines(f,count)
			f.Close()
		}
	}

	for line,num := range count{
		if num>1{
			fmt.Printf("%d %s\n",num,line)
		}
	}

}

func countlines(f *os.File,count map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		
		(count)[input.Text()]++
	}
}