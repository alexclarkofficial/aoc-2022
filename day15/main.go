package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getSensors(filename string) ([][4]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sensors := [][4]int{}

	sensorX := 0
	sensorY := 0
	beaconX := 0
	beaconY := 0
	for sc.Scan() {
		line := sc.Text()
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		sensors = append(sensors, [4]int{sensorX, sensorY, beaconX, beaconY})
	}

	return sensors, nil
}

func noBeaconCount(sensors [][4]int, line int) int {
	notPossible := map[int]bool{}
	for _, sensor := range sensors {
		distance := absVal(sensor[0]-sensor[2]) + absVal(sensor[1]-sensor[3])
		distanceToLine := absVal(sensor[1] - line)
		surplus := distance - distanceToLine
		if surplus > 0 {
			notPossible[sensor[0]] = true
			for i := 1; i <= surplus; i++ {
				left := sensor[0] - i
				right := sensor[0] + i
				notPossible[left] = true
				notPossible[right] = true
			}
		}
		if sensor[3] == line {
			notPossible[sensor[2]] = false
		}
	}

	count := 0
	for _, isNotPossible := range notPossible {
		if isNotPossible {
			count++
		}
	}

	return count
}

func possibleBeaconLocation(sensors [][4]int, boundary int) int {
	for _, sensor := range sensors {
		permimeter := absVal(sensor[0]-sensor[2]) + absVal(sensor[1]-sensor[3]) + 1
		// Top Right
		for i := permimeter; i >= 0; i-- {
			inverse := permimeter - i
			candidate := [2]int{sensor[0] + i, sensor[1] + inverse}
			if candidate[0] <= 0 || candidate[0] >= boundary ||
				candidate[1] <= 0 || candidate[1] >= boundary {
				continue
			}
			if outOfRange(sensors, candidate) {
				return candidate[0]*4000000 + candidate[1]
			}
		}
		// Top Left
		for i := permimeter; i >= 0; i-- {
			inverse := permimeter - i
			candidate := [2]int{sensor[0] + i, sensor[1] - inverse}
			if candidate[0] <= 0 || candidate[0] >= boundary ||
				candidate[1] <= 0 || candidate[1] >= boundary {
				continue
			}
			if outOfRange(sensors, candidate) {
				return candidate[0]*4000000 + candidate[1]
			}
		}
		// Bottom Right
		for i := permimeter; i >= 0; i-- {
			inverse := permimeter - i
			candidate := [2]int{sensor[0] - i, sensor[1] + inverse}
			if candidate[0] <= 0 || candidate[0] >= boundary ||
				candidate[1] <= 0 || candidate[1] >= boundary {
				continue
			}
			if outOfRange(sensors, candidate) {
				return candidate[0]*4000000 + candidate[1]
			}
		}
		// Bottom Left
		for i := permimeter; i >= 0; i-- {
			inverse := permimeter - i
			candidate := [2]int{sensor[0] + i, sensor[1] - inverse}
			if candidate[0] <= 0 || candidate[0] >= boundary ||
				candidate[1] <= 0 || candidate[1] >= boundary {
				continue
			}
			if outOfRange(sensors, candidate) {
				return candidate[0]*4000000 + candidate[1]
			}
		}
	}

	return 0
}

func outOfRange(sensors [][4]int, candidate [2]int) bool {
	for _, otherSensor := range sensors {
		otherSensorDistance := absVal(otherSensor[0]-otherSensor[2]) + absVal(otherSensor[1]-otherSensor[3])
		distanceFromOther := absVal(candidate[0]-otherSensor[0]) + absVal(candidate[1]-otherSensor[1])
		if distanceFromOther <= otherSensorDistance {
			return false
		}
	}
	return true
}

func main() {
	filename := os.Args[1]

	sensors, err := getSensors(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	partOneValue := noBeaconCount(sensors, 2000000)
	fmt.Printf("Part one score: %d\n", partOneValue)
	partTwoValue := possibleBeaconLocation(sensors, 4000000)
	fmt.Printf("Part two score: %d\n", partTwoValue)
}

func absVal(integer int) int {
	if integer < 0 {
		return -integer
	} else {
		return integer
	}
}
