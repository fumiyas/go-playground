package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func main() {
	var password_csv_b strings.Builder
	password_writer := csv.NewWriter(&password_csv_b)
	stdout_writer := csv.NewWriter(os.Stdout)

	data := [][]string{
		{"username", "password"},
		{"alice", `A1 secret`, `A2 secret`},
		{"bob", `B1 secret`, `B2,secret`},
		{"carol", `"C1 secret`, `"C2,secret"`},
	}

	for _, record := range data {
		password_csv_b.Reset()
		password_writer.Write(record[1:])
		password_writer.Flush()
		password_csv := password_csv_b.String()
		stdout_writer.Write([]string{record[0], password_csv[:len(password_csv)-1]})
	}

	stdout_writer.Flush()
}
