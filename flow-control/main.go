package main

func main() {
	// Three part loop
	/*for i := 0; i <= 10; i++ {
		fmt.Println("i is", i)
	}*/

	//while loop
	/*i := 10000
	for i > 100 {
		fmt.Println(i)
		i = i / 10
	}*/

	//Infinite loop
	// read input from user 5 times and write it to a log
	/*reader := bufio.NewReader(os.Stdin)

	ch := make(chan string)

	go myLogger.ListenFoLog(ch)

	fmt.Println("Enter something")
	for i := 0; i < 5; i++ {
		fmt.Printf("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}*/

	// Nested Loop
	/*for i := 1; i <= 10; i++ {
		fmt.Print("i is ", i)
		for j := 1; j <= 3; j++ {
			fmt.Print("   j: ", j)
		}
		fmt.Println()
	}*/

	// Do while loop
	for {
		// do some work

		// check condition and break
	}
}
