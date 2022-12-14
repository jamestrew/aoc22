package day13

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jamestrew/aoc22/utils"
)

// packets out of order
// pairs of packets signified by blank line
// packet is always a list, one per line

// (left,right)
// if both ints -> low,high
// if both list -> compare ints -> small len, high len - compare values
// if one int   -> convert int to list & retry comparison

type Packet struct {
	vals  []int
	depth int
}

type PacketPair struct {
	left, right Packet
}

func (p *PacketPair) inOrder() bool {
	if p.eitherEmpty() {
		return p.emptyInOrder()
	}

	ret := true
	length := utils.Min(len(p.left.vals), len(p.right.vals))
	for i := 0; i < length; i++ {
		if p.left.vals[i] > p.right.vals[i] {
			ret = false
			break
		}
		if p.left.vals[i] < p.right.vals[i] {
			return true
		}
	}
	if ret && len(p.left.vals) > len(p.right.vals) {
		ret = false
	}

	return ret
}

func (p *PacketPair) eitherEmpty() bool {
	return len(p.left.vals) == 0 || len(p.right.vals) == 0
}

func (p *PacketPair) emptyInOrder() bool {
	if len(p.left.vals) != 0 && len(p.right.vals) != 0 {
		return true
	}
	if len(p.left.vals) == 0 && len(p.right.vals) != 0 {
		return true
	}
	if len(p.left.vals) != 0 && len(p.right.vals) == 0 {
		return false
	}
	return p.left.depth < p.right.depth
}

func part1(input string) int {
	ret := 0
	// packets := parseJson(input)
	// // fmt.Println(packets)
	// for _, pair := range packets {
	// 	checkOrder(pair.left, pair.right)
	// }
	return ret
}

func checkOrder(left, right []interface{}) {
	// length := utils.Min(len(left), len(right))

	// for i := 0; i < length; i++ {
	// 	b := left[i]
	// 	switch b := b.(int) {
	// 	case int:
	// 		fmt.Println("int", b)
	// 	case interface{}:
	// 		fmt.Println("blah", b)
	// 	}
	// }
}

func compareInts(left, right int) bool {
	return false
}

func compareLists(left, right int) bool {
	return false
}

func asymmetricType(left, right interface{}) {

}

func part2(input string) int {
	ret := 0
	return ret
}

type foo struct {
	f json.RawMessage
}

func parseInput(input string) []PacketPair {
	ret := []PacketPair{}
	input = strings.TrimSpace(input)

	for _, pair := range strings.Split(input, "\n\n") {
		packets := strings.Split(pair, "\n")
		p := PacketPair{
			left:  flattenPacketArray(packets[0]),
			right: flattenPacketArray(packets[1]),
		}
		ret = append(ret, p)
	}
	return ret
}

func flattenPacketArray(packet string) Packet {
	depth := getDepth(packet)
	packet = strings.ReplaceAll(packet, "[", "")
	packet = strings.ReplaceAll(packet, "]", "")
	data := strings.Split(packet, ",")
	return Packet{utils.MapStrInt(data), depth}
}

func getDepth(packet string) int {
	return strings.Count(packet, "[")
}

func Answers() {
	fmt.Println(part1(utils.GetInput(13))) // 5316 -> too low, 5478 -> too low
	fmt.Println(part2(utils.GetInput(13)))
}
