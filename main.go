/*main.go
 *Tesla coding challenge
 *Author: Daniel D'Souza
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type expression struct {
	variables []string
	constants int // simplify storage by keeping track of the sum of constants
}

type equation struct {
	LHS string
	RHS expression
}

type equationSet struct {
	variables map[string]int // quickly search for duplicates
	equations []equation
}

func newEquationSet() equationSet {
	return equationSet{
		variables: map[string]int{},
		equations: make([]equation, 0),
	}
}

func (set *equationSet) addEquation(input string) {
	tokens := strings.Fields(input) // Fields handles extra whitespace
	eq := equation{}
	eq.RHS = expression{}

	for i, t := range tokens {
		if i == 0 {
			// store information on the right side 
			eq.LHS = t
			set.variables[t] = -1
			continue
		}

		if t == "=" || t == "+" {
			// discard these tokens becuase operators are known
			continue
		}

		v, err := strconv.Atoi(t)
		if err != nil {
			// token is a variable
			eq.RHS.variables = append(eq.RHS.variables, t)
			set.variables[t] = -1
		} else {
			// token is a constant
			eq.RHS.constants += v
		}
	}

	// add the equation to the list of known equations
	set.equations = append(set.equations, eq)
}

func (set *equationSet) solve() {
	// propogate values over the equation set until all variables's values are known
	for {
		totalUnknowns := 0

        // attempt to propogate values right for each equation
		for _, t := range set.equations {
			unknowns := 0
			variableValue := 0
			
			// attempt to calculate sum off RHS
			for _, name := range t.RHS.variables {
				v := set.variables[name]
				if v == -1 {
					unknowns++
				} else {
					variableValue += v
				}
			}

			if unknowns == 0 {
				// success, all values are known
				set.variables[t.LHS] = t.RHS.constants + variableValue
			}
			totalUnknowns += unknowns
		}

		if totalUnknowns == 0 {
			// success, all variables are known
			break
		}
	}
}

func main() {
	// Get the input filename from command line arguments
	inputFile := os.Args[1]

	// Grab a file object
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("could not open the file: ", inputFile)
		panic(err)
	}
	defer f.Close()

	equations := newEquationSet()

	// Read the equations
	for scanner := bufio.NewScanner(f); scanner.Scan(); {
		line := scanner.Text()
		equations.addEquation(line)
	}

	equations.solve()

	// Print out variables in alphabetical order
	var variableKeys []string
	for k := range equations.variables {
		variableKeys = append(variableKeys, k)
	}
	sort.Strings(variableKeys)

	for _, k := range variableKeys {
		fmt.Println(k, "=", equations.variables[k])
	}
}
