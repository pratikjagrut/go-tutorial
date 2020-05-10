//Parsing numbers from strings is a basic but common task in many programs;
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//The most common numeric conversions are Atoi (string to int).
	i, _ := strconv.Atoi("786")
	fmt.Println(i)

	_, e := strconv.Atoi("big")
	fmt.Println(e)

	//ParseInt, the 0 means infer the base from the string.
	//64 requires that the result fit in 64 bits.
	j, _ := strconv.ParseInt("738", 0, 64)
	fmt.Println(j)

	//ParseInt will recognize hex-formatted numbers.
	k, _ := strconv.ParseInt("0x3f5", 0, 64)
	fmt.Println(k)

	u, _ := strconv.ParseUint("648", 0, 64)
	fmt.Println(u)

	//ParseFloat, this 5 tells how many bits of precision to parse.
	f, _ := strconv.ParseFloat("4.7494", 64)
	fmt.Println(f)

}
