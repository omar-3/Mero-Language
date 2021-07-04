package main

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Fprintf(os.Stderr, "Must enter file name as a cli arg")
// 		os.Exit(1)
// 	}
// 	inFileName := os.Args[1]
// 	if _, err := os.Stat(inFileName); os.IsNotExist(err) {
// 		fmt.Fprintf(os.Stderr, "File with name ("+inFileName+") does not exist")
// 		os.Exit(1)
// 	}
// 	user, err := user.Current()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Hello %s! This is the Meroo programming language!\n",
// 		user.Username)
// 	repl.Start2(inFileName, os.Stdout)
// }
