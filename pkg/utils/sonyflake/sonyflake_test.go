package sonyflake

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	InitSonyflake()
}

func TestGenerateSonyflake(t *testing.T) {
	for i := 0; i < 100; i++ {
		id := GenerateSonyflake()
		fmt.Println(id)
	}
}

func TestGenerateTime(t *testing.T) {
	// Second
	fmt.Printf("Timestamp (seconds): %v;\n", time.Now().Unix())
	fmt.Printf("Timestamp (convert nanoseconds to seconds): %v;\n", time.Now().UnixNano()/1e9)

	// Millisecond
	fmt.Printf("Timestamp (milliseconds): %v;\n", time.Now().UnixMilli())
	fmt.Printf("Timestamp (nanoseconds to milliseconds converted): %v;\n", time.Now().UnixNano()/1e6)

	// Microseconds
	fmt.Printf("Timestamp (microseconds): %v;\n", time.Now().UnixMicro())
	fmt.Printf("Timestamp (nanoseconds converted to microseconds): %v;\n", time.Now().UnixNano()/1e3)

	// Nanoseconds
	fmt.Printf("Timestamp (nanoseconds): %v;\n", time.Now().UnixNano())
}
