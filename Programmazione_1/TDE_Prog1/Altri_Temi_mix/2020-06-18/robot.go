package main

import "fmt"

type Robot struct {
	x, y int
	dir  string
}

func costruisci(n, m int, s string) (Robot, bool) {

	if n >= 0 && m >= 0 && (s == "su" || s == "giù" || s == "destra" || s == "sinistra") {

		return Robot{n, m, s}, true
	}

	return Robot{0, 0, ""}, false
}

func stato(r Robot) {

	fmt.Printf("{%d %d %s}\n", r.x, r.y, r.dir)
}

func avanza(robot *Robot) Robot {

	if robot.dir == "su" {

		robot.y -= 1
	}
	if robot.dir == "giù" {

		robot.y += 1
	}
	if robot.dir == "destra" {

		robot.x += 1
	}
	if robot.dir == "sinistra" {

		robot.x -= 1
	}

	return *robot
}

func ruota(robot *Robot, b bool) Robot {

	if b {

		robot.dir = "destra"
	} else {

		robot.dir = "sinistra"
	}

	return *robot
}

func griglia(slice []Robot, xmin, xmax, ymin, ymax int) {

	var c int

	for i := ymin; i < ymax+1; i++ {

		for j := xmin; j < xmax+1; j++ {

			for _, robot := range slice {

				if robot.x == i && robot.y == j {
					c++

					if robot.dir == "su" {

						fmt.Printf("%c", 11014)
					} else if robot.dir == "giù" {
						fmt.Printf("%c", 11015)

					} else if robot.dir == "sinistra" {

						fmt.Printf("%c", 11013)
					} else if robot.dir == "destra" {

						fmt.Printf("%c", 10145)
					}

				}

			}

			if c == 0 {
				fmt.Printf("%c", 9633)
			} else {
				c--
			}

		}
		fmt.Println()
	}

}

func vista(slice []Robot) [][]Robot {

	var previous, current Robot

	coppie := [][]Robot{}

	guardami := []Robot{}

	for i := 0; i < len(slice); i++ {

		previous = slice[i]

		for j := i + 1; j < len(slice); j++ {

			current = slice[j]

			if current.x == previous.x || current.y == previous.y {

				if (current.dir == "giù" && previous.dir == "su") || (current.dir == "su" && previous.dir == "giù") || (current.dir == "destra" && previous.dir == "sinistra") || (current.dir == "sinistra" && previous.dir == "destra") {

					guardami := append(guardami, previous, current)

					coppie = append(coppie, guardami)

					guardami = nil
				}
			}
		}
	}

	return coppie
}

func main() {

	var x, y, c int

	var dir string

	slice := []Robot{}

	for {

		fmt.Println("Inserisci un robot:")

		fmt.Scan(&x, &y, &dir)

		robot, flag := costruisci(x, y, dir)

		if flag == false {

			break
		}

		for _, i := range slice {

			if i.x == robot.x && i.y == robot.y {

				fmt.Println("occupato")
				c++
				break

			}
		}

		if c == 0 {

			slice = append(slice, robot)
		} else {

			c--
		}

	}

	fmt.Println("Robot inseriti")

	for i := len(slice) - 1; i >= 0; i-- {

		fmt.Println(slice[i])
	}

}
