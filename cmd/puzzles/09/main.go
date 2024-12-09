package main

import (
	"advent-of-code-2024/pkg/util"
	"fmt"
	"slices"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	denseDiskMap, diskMapSize := readDenseDiskMap("09/input")
	diskMap := expandDiskMap1(denseDiskMap, diskMapSize)
	compactDiskMap1(diskMap)
	return calculateChecksum(diskMap)
}

func part2() int {
	denseDiskMap, diskMapSize := readDenseDiskMap("09/input")
	convertedDiskMap := convertDiskMap(denseDiskMap)
	convertedDiskMap = compactDiskMap2(convertedDiskMap)
	diskMap := expandDiskMap2(convertedDiskMap, diskMapSize)
	return calculateChecksum(diskMap)
}

func expandDiskMap1(denseDiskMap []int, diskMapSize int) []int {
	diskMap := make([]int, diskMapSize)
	slot, fileID := 0, 0
	for n, blockSize := range denseDiskMap {
		if n%2 == 0 {
			for i := 1; i <= blockSize; i++ {
				diskMap[slot] = fileID
				slot++
			}
			fileID++
		} else {
			for i := 1; i <= blockSize; i++ {
				diskMap[slot] = -1
				slot++
			}
		}
	}
	return diskMap
}

func compactDiskMap1(diskMap []int) {
	for {
		left := 0
		right := len(diskMap) - 1

		for {
			if diskMap[left] == -1 {
				break
			}
			left++
		}
		for {
			if diskMap[right] > -1 {
				break
			}
			right--
		}
		if left < right {
			temp := diskMap[right]
			diskMap[right] = diskMap[left]
			diskMap[left] = temp
		} else {
			break
		}
	}
}

func convertDiskMap(denseDiskMap []int) [][]int {
	diskMap := make([][]int, len(denseDiskMap))
	fileID := 0
	for d, blockSize := range denseDiskMap {
		if d%2 == 0 {
			diskMap[d] = []int{fileID, blockSize}
			fileID++
		} else {
			diskMap[d] = []int{-1, blockSize}
		}
	}
	return diskMap
}

func compactDiskMap2(diskMap [][]int) [][]int {
	lastIndex := len(diskMap) - 1
	lastFileID := diskMap[lastIndex][0]

	for fileID := lastFileID; fileID > 0; fileID-- {
		lastIndex = len(diskMap) - 1

		for k := lastIndex; k > 0; k-- {
			if diskMap[k][0] == fileID {
				for l := 0; l < k; l++ {
					if diskMap[l][0] == -1 && diskMap[l][1] >= diskMap[k][1] {
						if diskMap[l][1] == diskMap[k][1] {
							diskMap[l][0] = fileID
							diskMap[k][0] = -1
						} else {
							diskMap[l][1] -= diskMap[k][1]
							diskMap[k][0] = -1
							// insert new file element
							diskMap = slices.Insert(diskMap, l, []int{fileID, diskMap[k][1]})
						}
						break
					}
				}
				break
			}
		}
	}
	return diskMap
}

func expandDiskMap2(denseDiskMap [][]int, diskMapSize int) []int {
	diskMap := make([]int, diskMapSize)
	slot := 0
	for _, blocks := range denseDiskMap {
		for i := 1; i <= blocks[1]; i++ {
			diskMap[slot] = blocks[0]
			slot++
		}
	}
	return diskMap
}

func calculateChecksum(diskMap []int) int {
	checksum := 0
	for n, fileID := range diskMap {
		if fileID > -1 {
			checksum += n * fileID
		}
	}
	return checksum
}

func readDenseDiskMap(day string) ([]int, int) {
	lines := util.ReadFile(day)
	line := lines[0]

	denseDiskMap := make([]int, len(line))
	diskMapSize := 0
	for n, r := range line {
		denseDiskMap[n] = util.MustParseInt(string(r))
		diskMapSize += denseDiskMap[n]
	}
	return denseDiskMap, diskMapSize
}
