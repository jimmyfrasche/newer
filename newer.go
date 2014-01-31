//Command newer(1) takes two lists of files and returns true if any of the files
//in the first list are newer than any of the files in the second list.
//
//newer(1) is the essential logic of many build tools.
//It is useful when you have a few processes that could benefit
//from such logic, but are not quite ready to use a full build tool.
//
//All files in both lists must exist.
//
//EXAMPLES
//
//	newer a b && echo b is newer than a || echo a is newer than b
//
//	if newer a b
//	then
//		rebuild-b-from a
//	fi
//
//See if any of a, b, or c are newer than d
//	newer a b c d
//
//See if a is newer than any of b, c, or d
//	newer a -- b c d
//
//See if either a or b is newer than eitehr of c or d
//	newer a b -- c d
//
//EXIT STATUS
//
//If there is a usage error or a file cannot be stated, exit status is 2.
//
//If the second list contains an item newer than any item in the first,
//exit status is 0.
//
//Otherwise, the exit status is 1.
package main

import (
	"log"
	"os"
	"time"
)

func usage(msg string) {
	if msg != "" {
		log.Println("error:", msg)
	}
	log.Printf("Usage: %s left+ [-- right+]\n", os.Args[0])
	p := log.Println
	p("\twhere left and right are the names of files")
	p()
	p("If -- is not included the last value in left becomes the value of right.")
	p(os.Args[0], "compares the modtimes of left to the modtimes of right,")
	p("and returns true if any file in right is newer than any file in left.")
	os.Exit(2)
}

func max(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

func timestamps(ss []string) (time.Time, error) {
	var leader time.Time
	for _, s := range ss {
		fi, err := os.Stat(s)
		if err != nil {
			return time.Time{}, err
		}
		t := fi.ModTime()
		leader = max(t, leader)
	}
	return leader, nil
}

//Usage: %name files+ [-- files+]
func main() {
	log.SetFlags(0)

	var check, against []string
	args := os.Args[1:]
	switch len(args) {
	case 0, 1:
		usage("not enough arguments")
	default:
		partitionAt := len(args) - 1
		found := false
		for i, s := range args {
			if s == "--" {
				partitionAt = i
				if found {
					usage("arguments can contain only one instance of --")
				}
				found = true
			}
		}

		check, against = args[:partitionAt], args[partitionAt:]

		if against[0] == "--" {
			against = against[1:]
		}

		if len(check) == 0 {
			usage("no files to check")
		}
		if len(against) == 0 {
			usage("no files to check against")
		}
	}

	checkMod, err := timestamps(check)
	if err != nil {
		log.Fatalln(err)
	}

	againstMod, err := timestamps(against)
	if err != nil {
		log.Fatalln(err)
	}

	if againstMod.After(checkMod) {
		os.Exit(1)
	}
}
