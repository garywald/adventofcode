package main
	
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getScore(firstElve, you string) int {
    var (
        shapeScore int = 0
        outComeRound   = 0
    )
    if you == "X" {
        shapeScore = 1
        if firstElve == "A" {
            outComeRound = 3
        } else if firstElve == "B" {
            outComeRound = 0
        } else {
            outComeRound = 6
        }
    } else if you == "Y" {
        shapeScore = 2
        if firstElve == "A" {
            outComeRound = 6
        } else if firstElve == "B" {
            outComeRound = 3
        } else {
            outComeRound = 0
        }
    } else {
        shapeScore = 3
        if firstElve == "A" {
            outComeRound = 0
        } else if firstElve == "B" {
            outComeRound = 6
        } else {
            outComeRound = 3
        }
    }
    return outComeRound + shapeScore
}

func main() {

    readFile, err := os.Open("day2/input.txt")
    check(err)

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    var total = 0

    for fileScanner.Scan() {
        manche := strings.Split(fileScanner.Text(), " ")
        fmt.Println(manche[0])
        total += getScore(manche[0],manche[1])
    }

    fmt.Println("Total : ", total)

    readFile.Close()
}