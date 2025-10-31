
// প্রয়োজনীয় package থাকতে হবে
package main  // এই ফাইলটি main package এর, অর্থাৎ এটি executable program তৈরি করবে

// main go file এর package এর মধ্যে main function থাকা আবশ্যক
// main() function থেকে প্রোগ্রাম শুরু হয়

// fmt মানে হলো format
// fmt package এ অনেক built-in function আছে যেমন: প্রিন্ট করা, ফরম্যাটিং ইত্যাদি
import "fmt"

// func মানে function
// main() হলো প্রোগ্রামের entry point, যেখান থেকে কোড এক্সিকিউট হবে
func main() {
    // fmt.Println() function দিয়ে console/terminal এ লেখা প্রিন্ট করা হয়
    // এটি লেখা শেষ হলে নতুন লাইনে চলে যায়
    fmt.Println("Hello Bangladesh") // এখানে "Hello Bangladesh" প্রিন্ট হবে
}
