package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var coats, trousers, shirts, caps, jackets []string
	var season string
	scanner := bufio.NewScanner(os.Stdin)
	scanComplete := false
	for !scanComplete {
		scanner.Scan()
		input := strings.Split(scanner.Text(), ":")
		switch input[0] {
		case "COAT":
			coats = strings.Fields(input[1])
		case "SHIRT":
			shirts = strings.Fields(input[1])
		case "PANTS":
			trousers = strings.Fields(input[1])
		case "CAP":
			caps = strings.Fields(input[1])
		case "JACKET":
			jackets = strings.Fields(input[1])
		default:
			scanComplete = true
			season = input[0]
		}
	}

	var options []string
	var newOption string
	switch season {
	case "SPRING":
		{
			caps = append(caps, "")
			coats = append(coats, "")
			//shirt and trousers, cap*, coat*
			for _, shirt := range shirts {
				for _, trouser := range trousers {
					for _, cap := range caps {
						for _, coat := range coats {
							newOption = makeOptionStatement(coat, shirt, trouser, cap, "")
							options = append(options, newOption)
						}
					}
				}
			}
		}
	case "SUMMER":
		{
			//shirt and trousers, cap
			for _, shirt := range shirts {
				for _, trouser := range trousers {
					for _, cap := range caps {
						newOption = makeOptionStatement("", shirt, trouser, cap, "")
						options = append(options, newOption)
					}
				}
			}
		}
	case "FALL":
		{
			//shirt and trousers, cap*, coat*
			caps = append(caps, "")
			coats = append(coats, "")
			for _, shirt := range shirts {
				for _, trouser := range trousers {
					for _, cap := range caps {
						for _, coat := range coats {
							if coat == "yellow" || coat == "orange" {
								coat = "" // I think this should be "break"
							}
							newOption = makeOptionStatement(coat, shirt, trouser, cap, "")
							options = append(options, newOption)
						}
					}
				}
			}
		}
	case "WINTER":
		{
			//shirt and trousers, coat or jacket
			for _, shirt := range shirts {
				for _, trouser := range trousers {
					for _, coat := range coats {
						newOption = makeOptionStatement(coat, shirt, trouser, "", "")
						options = append(options, newOption)
					}
					for _, jacket := range jackets {
						newOption = makeOptionStatement("", shirt, trouser, "", jacket)
						options = append(options, newOption)
					}
				}
			}
		}
	}
	for _, option := range options {
		fmt.Println(option)
	}
}

func makeOptionStatement(coat string, shirt string, trouser string, cap string, jacket string) string {
	var statement string
	if coat != "" {
		statement = fmt.Sprintf("COAT: %s ", coat)
	}
	statement = fmt.Sprintf("%sSHIRT: %s", statement, shirt)
	statement = fmt.Sprintf("%s PANTS: %s", statement, trouser)
	if cap != "" {
		statement = fmt.Sprintf("%s CAP: %s", statement, cap)
	}
	if jacket != "" {
		statement = fmt.Sprintf("%s JACKET: %s", statement, jacket)
	}
	return statement
}
