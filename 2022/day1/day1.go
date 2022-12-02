package main
	
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    readFile, err := os.Open("day1/input.txt")
    check(err)

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    var numberOfCalories = 0
    var (
        firstElve int = 0
        secondElve    = 0
        thirdElve    = 0
    )

    for fileScanner.Scan() {
        if fileScanner.Text() == "" {
            if numberOfCalories > firstElve {
                thirdElve = secondElve
                secondElve = firstElve
                firstElve = numberOfCalories
            } else if numberOfCalories > secondElve{
                thirdElve = secondElve
                secondElve = numberOfCalories
            } else if numberOfCalories > thirdElve{
                thirdElve = numberOfCalories
            }
            numberOfCalories = 0
        } else {
            i, err := strconv.Atoi(fileScanner.Text())
            check(err)
            numberOfCalories += i
        }
    }

    fmt.Println("Top three Elves : ", firstElve, secondElve, thirdElve)
    fmt.Println("Total calories : ", firstElve + secondElve + thirdElve)

    readFile.Close()
}