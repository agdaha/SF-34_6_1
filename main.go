package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	inFile := flag.String("in", "", "input file")
	outFile := flag.String("out", "output.txt", "output file")
	flag.Parse()

	if *inFile == "" {
		fmt.Println("Укажите параметры:")
		flag.PrintDefaults()
		return
	}

	inList, shouldReturn := readFile(inFile)
	if shouldReturn {
		return
	}

	ouList, _ := Resolver(inList)

	shouldReturn1 := writeFile(outFile, ouList)
	if shouldReturn1 {
		return
	}
}

func writeFile(outFile *string, ouList []string) bool {
	f2, err := os.Create(*outFile)
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer f2.Close()

	w := bufio.NewWriter(f2)
	defer w.Flush()

	for _, s := range ouList {
		_, err = w.WriteString(s + "\n")
	}
	return false
}

func readFile(inFile *string) ([]string, bool) {
	f, err := os.Open(*inFile)
	if err != nil {
		fmt.Printf("не удалось открыть файл %v: %v", inFile, err)
		return nil, true
	}
	defer f.Close()

	inList := []string{}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			inList = append(inList, line)
			break
		}
		if err != nil {
			fmt.Printf("ошибка чтения из файла: %s", err)
			break
		}
		inList = append(inList, line)
	}
	return inList, false
}
