package main

import (
	"fmt"

	"./turing"
)

func main() {

	var incrementer = turing.NewMachine([]turing.Rule{
		{State: "q0", Symbol: '1', Write: '1', Motion: turing.Right, Next: "q0"},
		{State: "q0", Symbol: 'B', Write: '1', Motion: turing.Stay, Next: "qf"},
	})
	input := turing.NewTape('B', 0, []turing.Symbol{'1', '1', '1'})
	cnt, output := incrementer.Run(input)
	fmt.Println("Turing machine halts after", cnt, "operations")
	fmt.Println("Resulting tape:", output)

	var sort = turing.NewMachine([]turing.Rule{
		// Moving right, first b→B;s1
		{State: "s0", Symbol: 'a', Write: 'a', Motion: turing.Right, Next: "s0"},
		{State: "s0", Symbol: 'b', Write: 'B', Motion: turing.Right, Next: "s1"},
		{State: "s0", Symbol: ' ', Write: ' ', Motion: turing.Left, Next: "se"},
		// Conintue right to end of tape → s2
		{State: "s1", Symbol: 'a', Write: 'a', Motion: turing.Right, Next: "s1"},
		{State: "s1", Symbol: 'b', Write: 'b', Motion: turing.Right, Next: "s1"},
		{State: "s1", Symbol: ' ', Write: ' ', Motion: turing.Left, Next: "s2"},
		// Continue left over b.  a→b;s3, B→b;se
		{State: "s2", Symbol: 'a', Write: 'b', Motion: turing.Left, Next: "s3"},
		{State: "s2", Symbol: 'b', Write: 'b', Motion: turing.Left, Next: "s2"},
		{State: "s2", Symbol: 'B', Write: 'b', Motion: turing.Left, Next: "se"},
		// Continue left until B→a;s0
		{State: "s3", Symbol: 'a', Write: 'a', Motion: turing.Left, Next: "s3"},
		{State: "s3", Symbol: 'b', Write: 'b', Motion: turing.Left, Next: "s3"},
		{State: "s3", Symbol: 'B', Write: 'a', Motion: turing.Right, Next: "s0"},
		// Move to tape start → halt
		{State: "se", Symbol: 'a', Write: 'a', Motion: turing.Left, Next: "se"},
		{State: "se", Symbol: ' ', Write: ' ', Motion: turing.Right, Next: "see"},
	})
	input = turing.NewTape(' ', 0, []turing.Symbol("ababhabk"))
	cnt, output = sort.Run(input)
	fmt.Println("Turing machine halts after", cnt, "operations")
	fmt.Println("Resulting tape:", output)
}
