package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}

type Directory struct {
	name string
	files []*AOCFile
	directories map[string]*Directory
	parentDirectory *Directory
	totalSize int
}

type AOCFile struct {
	name string
	size int
}


func solveProblem(fileName string) int {
	fileSystem := buildFileSystem(fileName)
	result := 0
	
	//let's do maths
	computeDirectorySize(fileSystem, &result)

	currentFreeSpace := 70000000 - fileSystem.totalSize
	absoluteMinSpaceToBeDeleted := 30000000 - currentFreeSpace

	return traverseDirectories(fileSystem, absoluteMinSpaceToBeDeleted, fileSystem.totalSize)
}

func traverseDirectories(directory *Directory, threshold int, currentSmallestSize int) int{
	result := currentSmallestSize
	for k := range directory.directories {
		smallestSizeFromThatDirectory := traverseDirectories(directory.directories[k],threshold,result)
		if smallestSizeFromThatDirectory < currentSmallestSize && smallestSizeFromThatDirectory > threshold{
			result = smallestSizeFromThatDirectory
		}
	}

	if directory.totalSize < currentSmallestSize && directory.totalSize > threshold {
		result = directory.totalSize
	}
	
	return  result
}

func computeDirectorySize(directory *Directory,resultPointer *int){
	sizeFromDirectories := 0 
	for k := range directory.directories {
		computeDirectorySize(directory.directories[k],resultPointer)
		sizeFromDirectories += directory.directories[k].totalSize
	}
	sizeFromFiles := computeSizeComingFromFiles(directory.files)
	
	directory.totalSize = sizeFromFiles + sizeFromDirectories
	if (directory.totalSize <= 100000){
		*resultPointer += directory.totalSize
	}
}

func computeSizeComingFromFiles(files []*AOCFile) int {
	size := 0 
	for i := 0 ; i < len(files) ; i++ {
		size += files[i].size
	}
	return size
}

func buildFileSystem(fileName string) *Directory {
	readFile, err := os.Open(fileName)
	defer readFile.Close()
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	fileSystem := Directory{name:"/", directories: make(map[string]*Directory), files: make([]*AOCFile,0,10)}
	var currentDirectory *Directory

	for fileScanner.Scan() {
        line := fileScanner.Text()
		
		if line == "$ cd /"{
			currentDirectory = &fileSystem
			continue
		}

		if line == "$ ls"{
			continue
		}

		if strings.HasPrefix(line, "dir"){
			dirSplit := strings.Split(line," ")
			currentDirectory.directories[dirSplit[1]]= &Directory{name:dirSplit[1], parentDirectory: currentDirectory, directories: make(map[string]*Directory), files: make([]*AOCFile,0,10)}
			continue
		}

		if line == "$ cd .."{
			currentDirectory = currentDirectory.parentDirectory
			continue
		}

		if strings.HasPrefix(line,"$ cd"){
			cdSplit := strings.Split(line," ")
			currentDirectory = currentDirectory.directories[cdSplit[2]]
			continue
		}

		//just files now
		fileSplit := strings.Split(line," ")
		fileSize,err := strconv.Atoi(fileSplit[0])
		check(err)
		aocFile := AOCFile{name:fileSplit[1],size:fileSize}
		currentDirectory.files = append(currentDirectory.files,&aocFile)
	}

	return &fileSystem
}

