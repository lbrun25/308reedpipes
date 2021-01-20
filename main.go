package main

import (
	"308reedpipes/reedpipes"
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("USAGE\n" +
		"\t./308reedpipes r0 r5 r10 r15 r20 n\n\n" +
		"DESCRIPTION\n" +
		"\tr0\tradius (in cm) of pipe at the 0cm abscissa\n" +
		"\tr5\tradius (in cm) of pipe at the 5cm abscissa\n" +
		"\tr10\tradius (in cm) of pipe at the 10cm abscissa\n" +
		"\tr15\tradius (in cm) of pipe at the 15cm abscissa\n" +
		"\tr20\tradius (in cm) of pipe at the 20cm abscissa\n" +
		"\tn\tnumber of points needed to display the radius")
}

func main() {
	if reedpipes.CheckHelp() {
		printHelp()
		os.Exit(0)
	}
	err := reedpipes.CheckArgs(); if err != nil {
		fmt.Println(err)
		printHelp()
		os.Exit(84)
	}
	err = reedpipes.Main(); if err != nil {
		fmt.Println(err)
		printHelp()
		os.Exit(84)
	}
}