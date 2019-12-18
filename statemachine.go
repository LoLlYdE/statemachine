package main

import (
	"fmt"
	"github.com/lollyde/fileLib"
	"os"
	"strings"
)

type instruction struct {
	nextState int /* -1 stay on current */
	do string /* R L F / Right Left Flip*/
}

type state struct {
	instructions []instruction
}

type stateMachine struct {
	pointer int
	states []state
	data []bool
}

func main() {
	argsWithProg := os.Args

	fmt.Println("args:")
	fmt.Println(argsWithProg)
	if len(os.Args) < 1 {
		fmt.Println("Please provide a machinefile via command line argument")
		return
	}
	var machine,err = fileLib.ReadFileIntoSlice(os.Args[1])
	if err == nil {
		for _,item := range machine {
			fmt.Println(item)
		}
		initializeMachine(machine)
	} else {
		fmt.Println(err)
	}
}

func initializeMachine(src []string){
	var machine stateMachine
	machine.pointer = 0
	machine.states = make([]state, 0, 10)


	for _, line := range src{
		parsedLine := preParseLine(line)

		switch parsedLine[0] {
		case "DATA":
			machine.data = parseData(parsedLine[1])
		}
	}
}

func parseData(data string) []bool {
	boolData := make([]bool, len(data))
	for i, item := range data {
		boolData[i] = item == '1'
	}

	return boolData
}

func preParseLine(line string) []string {
	/* remove comments from line */
	i := strings.Index(line, "//")
	if i > 0 {
		line = line[:i]
	}

	toReturn := strings.Split(strings.TrimSpace(line), " ")

	if len(toReturn) == 1 && len(toReturn[0]) == 0 {
		return make([]string, 0)
	}

	return toReturn
}
