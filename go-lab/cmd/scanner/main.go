// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var lines []string
	i, maxLines := 0,2

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && i <= maxLines {
		lines = append(lines, scanner.Text())
		i++
	}
	fmt.Printf("Halo, saya %s. Saya tinggal di %s. Alamat email saya adalah %s\n",lines[0],lines[1],lines[2])
}
