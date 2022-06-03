package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	decode := flag.Bool("d", false, "Decode input")
	width := flag.Uint("w", 0, "Wrap output length")
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	input = bytes.TrimSpace(input)

	if *decode {
		b, err := hex.DecodeString(string(input))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Stdout.Write(b)
		return
	}
	out := hex.EncodeToString(input)
	if *width > 0 {
		out = wrap(out, *width)
	}
	fmt.Println(out)
}

func wrap(str string, w uint) string {
	var buf strings.Builder
	n := 0
	for _, r := range str {
		if n > int(w) {
			buf.WriteRune('\n')
			n = 0
		}
		buf.WriteRune(r)
		n++
	}
	return buf.String()
}
