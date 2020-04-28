package main

import (
    "fmt"
    "demography"
    "os"
)

func printHelp() {
    fmt.Println("USAGE")
    fmt.Println("\t./207demography code [...]")
    fmt.Println("")
    fmt.Println("DESCRIPTION")
    fmt.Println("\tcode\t\tcountry code")
}

func main() {
    if demography.CheckHelp() {
        printHelp()
        os.Exit(0)
    }
    if !demography.CheckArgs() {
        printHelp()
        os.Exit(84)
    }
    demography.Main()
    os.Exit(0)
}