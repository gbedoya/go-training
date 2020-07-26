package main

import (
	"fmt"
)

func main() {

	fmt.Println("----------------")
	fmt.Println("[*] MemoryLeaks")
	fmt.Println("----------------")
	MemoryLeaks()

	fmt.Println("----------------")
	fmt.Println("[*] Variables")
	fmt.Println("----------------")
	//Variables()

	fmt.Println("----------------")
	fmt.Println("[*] Constants")
	fmt.Println("----------------")
	//Constants()

	fmt.Println("----------------")
	fmt.Println("[*] For")
	fmt.Println("----------------")
	//For()

	fmt.Println("----------------")
	fmt.Println("[*] ElseIf")
	fmt.Println("----------------")
	//ElseIf()

	fmt.Println("----------------")
	fmt.Println("[*] Switch")
	fmt.Println("----------------")
	//Switch()

	fmt.Println("----------------")
	fmt.Println("[*] Arrays")
	fmt.Println("----------------")
	//Arrays()

	fmt.Println("----------------")
	fmt.Println("[*] Slices")
	fmt.Println("----------------")
	//Slices()

	fmt.Println("----------------")
	fmt.Println("[*] Maps")
	fmt.Println("----------------")
	//Maps()

	fmt.Println("----------------")
	fmt.Println("[*] Range")
	fmt.Println("----------------")
	//Range()

	fmt.Println("----------------")
	fmt.Println("[*] Functions")
	fmt.Println("----------------")
	//Functions()

	fmt.Println("----------------")
	fmt.Println("[*] Closure")
	fmt.Println("----------------")
	//Closure()

	fmt.Println("----------------")
	fmt.Println("[*] Recursion")
	fmt.Println("----------------")
	//Recursive()

	fmt.Println("----------------")
	fmt.Println("[*] Pointers")
	fmt.Println("----------------")
	//Pointers()

	fmt.Println("----------------")
	fmt.Println("[*] Random")
	fmt.Println("----------------")
	//Random()

	fmt.Println("----------------")
	fmt.Println("[*] Structs")
	fmt.Println("----------------")
	//Struct()

	fmt.Println("----------------")
	fmt.Println("[*] Methods")
	fmt.Println("----------------")
	//Methods()

	fmt.Println("----------------")
	fmt.Println("[*] Interfaces")
	fmt.Println("----------------")
	//Interfaces()

	fmt.Println("----------------")
	fmt.Println("[*] Errors")
	fmt.Println("----------------")
	//Errors()

	fmt.Println("----------------")
	fmt.Println("[*] Go Routines")
	fmt.Println("----------------")
	//GoRoutines()

	fmt.Println("----------------")
	fmt.Println("[*] Channels")
	fmt.Println("----------------")
	//Channels()

	fmt.Println("----------------")
	fmt.Println("[*] Channels Buffering")
	fmt.Println("----------------")
	//ChannelsBuffering()

	fmt.Println("----------------")
	fmt.Println("[*] Channel Sync")
	fmt.Println("----------------")
	//ChannelSync()

	fmt.Println("----------------")
	fmt.Println("[*] Channel Directions")
	fmt.Println("----------------")
	//	ChannelsDirection()

	fmt.Println("----------------")
	fmt.Println("[*] Select")
	fmt.Println("----------------")
	//Select()

	fmt.Println("----------------")
	fmt.Println("[*] Non-Blocking Channels")
	fmt.Println("----------------")
	//NonBlockingChannels()

	fmt.Println("----------------")
	fmt.Println("[*] Closing Channels")
	fmt.Println("----------------")
	//CloseChannel()

	fmt.Println("----------------")
	fmt.Println("[*] Range Over Channels")
	fmt.Println("----------------")
	//RangeChannels()

	fmt.Println("----------------")
	fmt.Println("[*] Environment variables")
	fmt.Println("----------------")
	//EnvVars()

	fmt.Println("----------------")
	fmt.Println("[*] Defer")
	fmt.Println("----------------")
	//Defer()

	fmt.Println("----------------")
	fmt.Println("[*] Panic")
	fmt.Println("----------------")
	//Panic()

	fmt.Println("----------------")
	fmt.Println("[*] Sorting")
	fmt.Println("----------------")
	//Sorting()

	fmt.Println("----------------")
	fmt.Println("[*] Sorting by functions")
	fmt.Println("----------------")
	//SortByFunc()

	fmt.Println("----------------")
	fmt.Println("[*] Number Parsing")
	fmt.Println("----------------")
	//NumberParse()

	fmt.Println("----------------")
	fmt.Println("[*] URL Parsing")
	fmt.Println("----------------")
	//URLParse()

	fmt.Println("----------------")
	fmt.Println("[*] SHA1 Hashes")
	fmt.Println("----------------")
	//SHA1Hashes()

	fmt.Println("----------------")
	fmt.Println("[*] Base64 Encoding")
	fmt.Println("----------------")
	//Base64()

	fmt.Println("----------------")
	fmt.Println("[*] Rate Limiting")
	fmt.Println("----------------")
	//	RateLimit()

	fmt.Println("----------------")
	fmt.Println("[*] Reading Files")
	fmt.Println("----------------")
	//ReadFiles()

	fmt.Println("----------------")
	fmt.Println("[*] Writing Files")
	fmt.Println("----------------")
	//WriteFiles()

	fmt.Println("----------------")
	fmt.Println("[*] Line Filters")
	fmt.Println("----------------")
	//LineFilters -> line-filters.go

	fmt.Println("----------------")
	fmt.Println("[*] Signals")
	fmt.Println("----------------")
	//Signals()

	fmt.Println("----------------")
	fmt.Println("[*] Exit")
	fmt.Println("----------------")
	//Exit()

	fmt.Println("----------------")
	fmt.Println("[*] Regular Expressions")
	fmt.Println("----------------")
	//Regex()

	fmt.Println("----------------")
	fmt.Println("[*] Json")
	fmt.Println("----------------")
	//JSON()

	fmt.Println("----------------")
	fmt.Println("[*] Time")
	fmt.Println("----------------")
	//Time()

	fmt.Println("----------------")
	fmt.Println("[*] Epoch")
	fmt.Println("----------------")
	//Epoch()

	fmt.Println("----------------")
	fmt.Println("[*] Time Formatting")
	fmt.Println("----------------")
	//TimeFormatting()

	fmt.Println("----------------")
	fmt.Println("[*] String Functions")
	fmt.Println("----------------")
	//StringFunctions()

	fmt.Println("----------------")
	fmt.Println("[*] String Formatting")
	fmt.Println("----------------")
	//StringFormatting()

	fmt.Println("----------------")
	fmt.Println("[*] Collection Functions")
	fmt.Println("----------------")
	//CollectionFunctions()

	fmt.Println("----------------")
	fmt.Println("[*] Timers")
	fmt.Println("----------------")
	//Timer()

	fmt.Println("----------------")
	fmt.Println("[*] Tickers")
	fmt.Println("----------------")
	//Ticker()

	fmt.Println("----------------")
	fmt.Println("[*] Workers Pools")
	fmt.Println("----------------")
	//WorkerPool()

	fmt.Println("----------------")
	fmt.Println("[*] Atomic Counters")
	fmt.Println("----------------")
	//AtomicCounter()

	fmt.Println("----------------")
	fmt.Println("[*] Mutexes")
	fmt.Println("----------------")
	//Mutex()

	fmt.Println("----------------")
	fmt.Println("[*] Stateful Goroutines")
	fmt.Println("----------------")
	//	StatefulGoRoutines()

	// Command-Line Arguments
	// Spawning Processes
	// Exec'ing Processes
}
