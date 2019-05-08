package regex

import (
	"fmt"
	"testing"

	"github.com/fatih/stopwatch"
)

func TestMatchingRegex(t *testing.T) {
	result := MatchRegex("aaa4C")
	if result == false {
		t.Error("Regex should have matched")
	}
	println("Regex matches")
}

func TestNonMatchingRegex(t *testing.T) {
	result := MatchRegex("aaaC")
	if result == true {
		t.Error("Regex should not match")
	}
	println("Regex do not match")
}

func TestRegexManual(t *testing.T) {

	fmt.Println("Payload \t\t\t\t\t\t\t\t\t\t   Lenght\t\t\tDuration ")
	fmt.Println("============================================================================")
	for i := 15; i < 50; i++ {
		var payload = RandomString(i) + "!"
        // time start
        stopwatch := stopwatch.New()
		stopwatch.Start(0)
        
        MatchRegex(payload)

        //time stops
        stopwatch.Stop()
        
		ElapsedTime := stopwatch.ElapsedTime()
        fmt.Printf("%-50s %d chars \t \t %s \n", payload, i, ElapsedTime)
        
        //resetting the watcher
		stopwatch.Reset()
	}
}
//TestSingleLongString generates a 10000000 character lengh string and tests the regex method
func TestSingleLongString(t *testing.T) {
	//defining a payload size
	var payloadSize = 10000000

	fmt.Println("Payload Size      Duration ")
	fmt.Println("===================================")
	var payload = RandomString(payloadSize) + "4C"
    // time start
    stopwatch := stopwatch.New()
	stopwatch.Start(0)
        
    MatchRegex(payload)
    //time stops
    stopwatch.Stop()
        
	ElapsedTime := stopwatch.ElapsedTime()
    fmt.Printf("%d chars \t  %s \n", payloadSize, ElapsedTime)
		
    //resetting the watcher
	stopwatch.Reset()
}
/* 
BenchmarkMatchRegexExecution runs a benchmark agains MatchRegex function (which uses MatchString)
MatchRegex function is invoked with a 10000 character length payload 
*/ 
func BenchmarkMatchRegexExecution(b * testing.B){
	// generating a random string of 10000 character length
	var randomString = RandomString(100000)

	//benchmarking a 10000 char string agains the regex
	for i := 0; i < b.N; i++ {
		MatchRegex(randomString)
	}
}

/* 
BenchmarkMatchCompiledRegexExecution runs a benchmark agains MatchCompiledRegex function (which uses Compile)
MatchRegex function is invoked with a 10000 character length payload 
*/ 
func BenchmarkMatchCompiledRegexExecution(b * testing.B){
	// generating a random string of 10000 character length
	var randomString = RandomString(100000)

	//benchmarking a 10000 char string agains the regex
	for i := 0; i < b.N; i++ {
		MatchCompiledRegex(randomString)
	}
}

func BenchmarkRegex(b *testing.B){
	var payloadMaxSize = 50
	for i := 15; i < payloadMaxSize; i++ {
		benchmarkRegexFunction (i, b)
	}
}

func benchmarkRegexFunction(i int, b *testing.B){
	for index := 0; index < b.N; index++ {
        MatchRegex(RandomString(i))
    }
}



func BenchmarkCompiledRegex(b *testing.B) {
	for i := 14; i < 25; i++ {
		MatchCompiledRegex(RandomString(i) + "!")
	}
}

