# Go Beginner Labs — 50 Programs
### From zero to confident with Go fundamentals

---

## How to Use This Document

Each lab is one small program. Work through them in order — they build on each other.

**For every lab:**
1. Create a new file: `lab-XX/main.go`
2. Write `package main` and `func main()` at the top
3. Build the program described
4. Run it: `go run lab-XX/main.go`

**Rules:**
- Do not look up solutions. Look up Go documentation, not answers.
- If you're stuck for more than 20 minutes, re-read the hint. Still stuck? Ask.
- The goal is working code, not perfect code. Get it running first, clean it up second.

---

## Tier 1 — Output, Variables, Basic Math
*Labs 1–10. Get comfortable with Go syntax, variables, and simple calculations.*

---

### Lab 1 — Hello, World!

Your first Go program. Every programmer starts here.

**You will practice:** `fmt.Println`, `package main`, `func main()`

**Build this:**
- Print the text `Hello, World!` to the terminal

**Example run:**
```
Hello, World!
```

**Hint:** Every Go program needs `package main` at the top and a `func main()`. Use `fmt.Println("your text")` to print.

---

### Lab 2 — Personal Greeting

Programs become interesting when they respond to input.

**You will practice:** `fmt.Scan`, variables, string printing

**Build this:**
- Ask the user: `Enter your name:`
- Read one word from the user
- Print: `Hello, Ali!` (replacing Ali with whatever was entered)

**Example run:**
```
Enter your name: Ali
Hello, Ali!
```

**Hint:** `fmt.Scan(&name)` reads one word into the variable `name`. Declare `name` as a `string` first.

---

### Lab 3 — Add Two Numbers

**You will practice:** integer variables, `fmt.Scan`, arithmetic, `fmt.Printf`

**Build this:**
- Ask the user to enter two numbers (separately)
- Print their sum in the format: `5 + 3 = 8`

**Example run:**
```
Enter first number: 5
Enter second number: 3
5 + 3 = 8
```

**Hint:** `fmt.Printf("%d + %d = %d\n", a, b, a+b)` formats integers into a string. `%d` means "integer here".

---

### Lab 4 — Simple Calculator

**You will practice:** switch statement, operators, reading multiple types

**Build this:**
- Ask the user for two numbers and an operator (`+`, `-`, `*`, `/`)
- Print the result
- For division, print a special message if the second number is 0

**Example run:**
```
Enter first number: 10
Enter operator (+, -, *, /): /
Enter second number: 2
10 / 2 = 5
```
```
Enter first number: 5
Enter operator: /
Enter second number: 0
Error: cannot divide by zero
```

---

### Lab 5 — Temperature Converter

**You will practice:** `float64`, arithmetic with decimals, `fmt.Printf` with `%.2f`

**Build this:**
- Ask the user for a temperature in Celsius
- Convert it to Fahrenheit: `F = (C × 9/5) + 32`
- Print the result rounded to 2 decimal places

**Example run:**
```
Enter temperature in Celsius: 100
100.00°C = 212.00°F
```
```
Enter temperature in Celsius: 37
37.00°C = 98.60°F
```

**Hint:** Use `float64` for the variable. `%.2f` in `fmt.Printf` prints 2 decimal places.

---

### Lab 6 — Seconds Converter

**You will practice:** integer division (`/`), modulo operator (`%`), multiple variables

**Build this:**
- Ask for a number of seconds
- Convert to hours, minutes, and remaining seconds
- Print in the format: `1h 1m 1s`

**Example run:**
```
Enter seconds: 3661
1h 1m 1s
```
```
Enter seconds: 7322
2h 2m 2s
```

**Hint:** Hours = seconds / 3600. Remaining after hours = seconds % 3600. Minutes = remaining / 60. Seconds left = remaining % 60.

---

### Lab 7 — Rectangle Calculator

**You will practice:** multiple inputs, float arithmetic, `fmt.Printf`

**Build this:**
- Ask the user for the width and height of a rectangle (can be decimals)
- Print the area and perimeter

**Example run:**
```
Enter width: 5.5
Enter height: 3.0
Area: 16.50
Perimeter: 17.00
```

---

### Lab 8 — Circle Calculator

**You will practice:** `math` package, `math.Pi`, `math.Pow`

**Build this:**
- Ask the user for the radius of a circle
- Print the area (`π × r²`) and circumference (`2 × π × r`)
- Round to 2 decimal places

**Example run:**
```
Enter radius: 7
Area: 153.94
Circumference: 43.98
```

**Hint:** Import the `math` package. Use `math.Pi` for π and `math.Pow(r, 2)` for r².

---

### Lab 9 — Simple Interest Calculator

**You will practice:** multiple float inputs, formula application

**Build this:**
- Ask for Principal (PKR), annual interest rate (%), and time (years)
- Calculate simple interest: `SI = (P × R × T) / 100`
- Print the interest earned and total amount

**Example run:**
```
Principal (PKR): 100000
Annual rate (%): 8.5
Time (years): 3
Interest earned: PKR 25500.00
Total amount: PKR 125500.00
```

---

### Lab 10 — BMI Calculator

**You will practice:** float arithmetic, comparing with constants

**Build this:**
- Ask for weight in kilograms and height in meters
- Calculate BMI: `weight / (height × height)`
- Print the BMI and category:
  - Below 18.5 → Underweight
  - 18.5–24.9 → Normal weight
  - 25–29.9 → Overweight
  - 30 and above → Obese

**Example run:**
```
Weight (kg): 70
Height (m): 1.75
BMI: 22.86 — Normal weight
```

---

## Tier 2 — Control Flow
*Labs 11–20. if/else, switch, and making decisions in code.*

---

### Lab 11 — Even or Odd

**You will practice:** `if/else`, modulo operator

**Build this:**
- Ask for a number
- Print whether it is even or odd

**Example run:**
```
Enter a number: 7
7 is odd
```
```
Enter a number: 12
12 is even
```

**Hint:** A number is even if `number % 2 == 0`.

---

### Lab 12 — Number Classifier

**You will practice:** `if / else if / else` chains

**Build this:**
- Ask for a number
- Print whether it is positive, negative, or zero

**Example run:**
```
Enter a number: -5
-5 is negative
```
```
Enter a number: 0
0 is zero
```

---

### Lab 13 — Largest of Three

**You will practice:** nested if statements, logical operators

**Build this:**
- Ask for three numbers
- Print which one is the largest
- Handle ties: if two or more are equal and largest, say so

**Example run:**
```
Enter three numbers: 12 45 30
45 is the largest
```
```
Enter three numbers: 7 7 3
7 is the largest (tied)
```

---

### Lab 14 — Leap Year Checker

**You will practice:** logical operators (`&&`, `||`), complex conditions

**Build this:**
- Ask for a year
- A year is a leap year if: divisible by 4, AND (not divisible by 100, OR divisible by 400)
- Print whether it is a leap year

**Example run:**
```
Enter year: 2000
2000 is a leap year
```
```
Enter year: 1900
1900 is not a leap year
```
```
Enter year: 2024
2024 is a leap year
```

---

### Lab 15 — Letter Grade

**You will practice:** `if / else if` chains, ranges

**Build this:**
- Ask for a score (0–100)
- Print the letter grade:
  - 90–100 → A
  - 80–89 → B
  - 70–79 → C
  - 60–69 → D
  - Below 60 → F
- If the score is outside 0–100, print an error

**Example run:**
```
Enter score: 85
Grade: B
```
```
Enter score: 110
Invalid score. Must be between 0 and 100.
```

---

### Lab 16 — Day of the Week

**You will practice:** `switch` statement

**Build this:**
- Ask for a number (1–7)
- Print the corresponding day (1 = Monday, 7 = Sunday)
- If the number is outside 1–7, print an error

**Example run:**
```
Enter day number (1-7): 5
Friday
```
```
Enter day number: 8
Invalid. Please enter a number between 1 and 7.
```

**Hint:** Use `switch number { case 1: ... case 2: ... default: ... }`. No `break` needed in Go — each case stops automatically.

---

### Lab 17 — Simple Login

**You will practice:** string comparison, `if/else`

**Build this:**
- Hardcode a username and password in your program (e.g., `"admin"` and `"go123"`)
- Ask the user to enter a username and password
- Print `Login successful!` or `Invalid credentials`

**Example run:**
```
Username: admin
Password: go123
Login successful!
```
```
Username: admin
Password: wrong
Invalid credentials
```

---

### Lab 18 — FizzBuzz

Classic interview warm-up question.

**You will practice:** loops, modulo, multiple conditions

**Build this:**
- Print numbers from 1 to 30
- For multiples of 3 print `Fizz` instead of the number
- For multiples of 5 print `Buzz`
- For multiples of both 3 and 5 print `FizzBuzz`

**Example run:**
```
1
2
Fizz
4
Buzz
Fizz
7
...
FizzBuzz
```

---

### Lab 19 — Vowel or Consonant

**You will practice:** `switch` with multiple values per case, `strings.ToLower`

**Build this:**
- Ask the user to enter a single letter
- Print whether it is a vowel or consonant
- Handle uppercase and lowercase the same way
- Print an error if the input is not a letter

**Example run:**
```
Enter a letter: A
A is a vowel
```
```
Enter a letter: b
b is a consonant
```

**Hint:** `switch letter { case 'a', 'e', 'i', 'o', 'u': ... }` matches multiple values in one case.

---

### Lab 20 — Triangle Type

**You will practice:** multiple conditions, logical operators

**Build this:**
- Ask for the lengths of three sides of a triangle (as integers)
- First check if the three sides can form a valid triangle: each side must be less than the sum of the other two
- If valid, print the type:
  - All sides equal → Equilateral
  - Two sides equal → Isosceles
  - No sides equal → Scalene

**Example run:**
```
Enter three sides: 5 5 5
Equilateral triangle
```
```
Enter three sides: 3 3 5
Isosceles triangle
```
```
Enter three sides: 1 2 10
Not a valid triangle
```

---

## Tier 3 — Loops
*Labs 21–30. for loops, while-style loops, nested loops, and iteration patterns.*

---

### Lab 21 — Count to N

**You will practice:** `for` loop, loop variable

**Build this:**
- Ask for a number N
- Print all numbers from 1 to N, one per line

**Example run:**
```
Enter N: 5
1
2
3
4
5
```

**Hint:** `for i := 1; i <= n; i++ { ... }` is Go's for loop. There is no `while` keyword in Go — the `for` loop covers all cases.

---

### Lab 22 — Sum from 1 to N

**You will practice:** accumulator pattern, `for` loop

**Build this:**
- Ask for a number N
- Print the sum of all numbers from 1 to N

**Example run:**
```
Enter N: 10
Sum from 1 to 10 = 55
```

---

### Lab 23 — Multiplication Table

**You will practice:** `for` loop, formatted output

**Build this:**
- Ask for a number
- Print its multiplication table from 1 to 10

**Example run:**
```
Enter number: 7
7 x  1 =  7
7 x  2 = 14
7 x  3 = 21
...
7 x 10 = 70
```

**Hint:** `fmt.Printf("%d x %2d = %3d\n", n, i, n*i)` right-aligns numbers using width specifiers.

---

### Lab 24 — Factorial

**You will practice:** accumulator pattern, `for` loop

**Build this:**
- Ask for a number N
- Print N! (factorial)
- Print an error if N is negative

**Example run:**
```
Enter N: 5
5! = 120
```
```
Enter N: 0
0! = 1
```
```
Enter N: -3
Factorial is not defined for negative numbers
```

**Hint:** Factorial of 0 is 1 (this is a special case you must handle). For N > 0: multiply all numbers from 1 to N.

---

### Lab 25 — Fibonacci Sequence

**You will practice:** multiple variables, `for` loop, sequence generation

**Build this:**
- Ask for how many Fibonacci numbers to print
- Print the sequence starting from 0

**Example run:**
```
How many Fibonacci numbers? 8
0 1 1 2 3 5 8 13
```

**Hint:** The Fibonacci sequence: each number is the sum of the two before it. Start with `a = 0, b = 1`. In each step, print `a`, then `a, b = b, a+b`.

---

### Lab 26 — Countdown

**You will practice:** loop going backwards, `for` with decrement

**Build this:**
- Ask for a number N
- Count down from N to 1 and print each number
- Print `Go!` at the end

**Example run:**
```
Enter N: 5
5
4
3
2
1
Go!
```

---

### Lab 27 — Star Triangle

**You will practice:** nested `for` loops, `strings.Repeat`

**Build this:**
- Ask for the number of rows
- Print a right-aligned triangle of stars

**Example run (5 rows):**
```
    *
   **
  ***
 ****
*****
```

**Hint:** For each row `i` (starting at 1), print `(rows - i)` spaces then `i` stars. Use `strings.Repeat("*", i)` and `strings.Repeat(" ", rows-i)`.

---

### Lab 28 — Reverse a Number

**You will practice:** modulo operator, integer division, `for` loop

**Build this:**
- Ask for a positive integer
- Print the digits in reverse order

**Example run:**
```
Enter a number: 12345
Reversed: 54321
```
```
Enter a number: 100
Reversed: 1
```

**Hint:** `number % 10` gives the last digit. `number / 10` removes the last digit. Repeat until the number is 0.

---

### Lab 29 — Sum of Digits

**You will practice:** modulo operator, accumulator in a loop

**Build this:**
- Ask for a positive integer
- Print the sum of its individual digits

**Example run:**
```
Enter a number: 1234
Sum of digits: 10
```
```
Enter a number: 9999
Sum of digits: 36
```

---

### Lab 30 — Number Guessing Game

**You will practice:** `for` loop with no condition (`for {}`), `break`, comparison

**Build this:**
- Hardcode a secret number (e.g., `42`)
- Keep asking the user to guess until they get it right
- After each wrong guess, print `Too high` or `Too low`
- Print how many attempts it took

**Example run:**
```
Guess the number: 50
Too high
Guess the number: 30
Too low
Guess the number: 42
Correct! You got it in 3 attempts.
```

**Hint:** `for { ... }` loops forever. Use `break` to exit the loop when the guess is correct.

---

## Tier 4 — Functions and Strings
*Labs 31–40. Writing reusable functions and working with text.*

---

### Lab 31 — Is Prime?

**You will practice:** writing a function, returning a `bool`, `for` loop

**Build this:**
- Write a function `isPrime(n int) bool` that returns true if n is prime
- In `main`, ask for a number and call the function

**Example run:**
```
Enter a number: 17
17 is prime
```
```
Enter a number: 15
15 is not prime
```

**Hint:** A number is prime if it has no divisors other than 1 and itself. Check all numbers from 2 up to `n-1` (or more efficiently, up to `√n`). Return false as soon as you find a divisor.

---

### Lab 32 — Power Function

**You will practice:** function with two parameters, `for` loop, returning a value

**Build this:**
- Write a function `power(base, exp int) int` that computes `base^exp`
- Do not use the `math` package — implement the multiplication loop yourself
- In `main`, ask for base and exponent and print the result

**Example run:**
```
Enter base: 2
Enter exponent: 10
2^10 = 1024
```
```
Enter base: 5
Enter exponent: 0
5^0 = 1
```

**Hint:** Any number to the power 0 is 1. For other cases, multiply `base` by itself `exp` times.

---

### Lab 33 — Reverse a String

**You will practice:** strings and runes, `range` over a string, `string()` conversion

**Build this:**
- Ask the user to enter a word or sentence
- Print it reversed

**Example run:**
```
Enter text: hello
olleh
```
```
Enter text: Go is fun
nuf si oG
```

**Hint:** In Go, strings are sequences of bytes. To handle non-ASCII characters correctly, convert to `[]rune` first: `runes := []rune(text)`. Then iterate backwards.

---

### Lab 34 — Count Vowels

**You will practice:** `range` over a string, character comparison, `strings.ToLower`

**Build this:**
- Ask for a sentence
- Count and print the number of vowels (a, e, i, o, u) — case insensitive

**Example run:**
```
Enter sentence: Hello World
Vowels: 3
```

---

### Lab 35 — Palindrome Checker

**You will practice:** functions, string manipulation, comparison

**Build this:**
- Write a function `isPalindrome(s string) bool`
- A palindrome reads the same forwards and backwards (ignore spaces and case)
- Ask for a word and print whether it is a palindrome

**Example run:**
```
Enter a word: racecar
racecar is a palindrome
```
```
Enter a word: hello
hello is not a palindrome
```

**Hint:** Use your reverse string logic from Lab 33. Compare the original (lowercased) with the reversed version.

---

### Lab 36 — Word Counter

**You will practice:** `strings.Fields`, `len()`, slices of strings

**Build this:**
- Ask for a sentence
- Print the number of words

**Example run:**
```
Enter a sentence: the quick brown fox jumps
Word count: 5
```

**Hint:** `strings.Fields(sentence)` splits a string by whitespace and returns a slice of words. `len(slice)` gives the length.

---

### Lab 37 — Title Case

**You will practice:** `strings.Fields`, `strings.ToUpper`, `strings.Join`, slices

**Build this:**
- Ask for a sentence in lowercase
- Print it in title case (first letter of each word capitalized)

**Example run:**
```
Enter sentence: the quick brown fox
The Quick Brown Fox
```

**Hint:** Split the sentence into words with `strings.Fields`. For each word, capitalize the first letter: `strings.ToUpper(string(word[0])) + word[1:]`. Join them back with `strings.Join(words, " ")`.

---

### Lab 38 — Caesar Cipher

**You will practice:** rune arithmetic, modulo for wrapping, building strings

**Build this:**
- Ask for a message and a shift number (1–25)
- Encrypt the message by shifting each letter forward by `shift` positions
- Leave non-letter characters (spaces, punctuation) unchanged
- Print the encrypted message

**Example run:**
```
Enter message: Hello World
Enter shift: 3
Encrypted: Khoor Zruog
```

**Hint:** To shift a letter: `encrypted = 'a' + (letter - 'a' + shift) % 26`. Handle uppercase and lowercase separately. Non-letters stay unchanged.

---

### Lab 39 — Remove Spaces

**You will practice:** `strings.ReplaceAll`, or manually building a string with `strings.Builder`

**Build this:**
- Ask for a sentence
- Print it with all spaces removed
- Also print how many spaces were removed

**Example run:**
```
Enter sentence: go is great
goIsgreat
Removed 2 spaces
```

Wait — actually print without spaces: `goIsgreat` → just remove spaces, no camelCase needed:

**Corrected example:**
```
Enter sentence: go is great
goIsgreat
```

Actually just remove spaces:
```
Enter sentence: hello world how are you
helloworldhowareyou
Spaces removed: 4
```

**Hint:** `strings.ReplaceAll(s, " ", "")` removes all spaces in one call.

---

### Lab 40 — Count Character Occurrences

**You will practice:** `range` over a string, character comparison, counter variable

**Build this:**
- Ask for a sentence and a character to search for
- Print how many times that character appears (case insensitive)

**Example run:**
```
Enter sentence: Mississippi
Enter character: s
's' appears 4 times in "Mississippi"
```

---

## Tier 5 — Slices, Maps, and Structs
*Labs 41–50. Working with collections and your own data types.*

---

### Lab 41 — Min and Max

**You will practice:** slices, `for range`, comparison

**Build this:**
- Define a slice of integers in your code (you can hardcode 7–8 numbers)
- Find and print the minimum and maximum values

**Example run:**
```
Numbers: [3 1 4 1 5 9 2 6]
Minimum: 1
Maximum: 9
```

**Hint:** Start by assuming the first element is both min and max. Then loop through the rest and update as you find smaller/larger values.

---

### Lab 42 — Slice Average

**You will practice:** slices, accumulator, float division

**Build this:**
- Define a slice of `float64` numbers
- Calculate and print the sum, count, and average

**Example run:**
```
Numbers: [85.0 92.5 78.0 90.0 88.5]
Sum: 434.00
Count: 5
Average: 86.80
```

**Hint:** When dividing, make sure both values are `float64`. If you have `sum` as `float64` and `count` as `int`, use `float64(count)` in the division.

---

### Lab 43 — Remove Duplicates

**You will practice:** slices, maps as a "seen" tracker, appending to a slice

**Build this:**
- Start with a slice that has duplicate numbers: `[1, 2, 3, 2, 4, 3, 5, 1]`
- Print a new slice with duplicates removed, preserving the original order

**Example run:**
```
Original: [1 2 3 2 4 3 5 1]
Unique:   [1 2 3 4 5]
```

**Hint:** Use a `map[int]bool` to track which numbers you've already seen. For each number, if it's not in the map, add it to the result and mark it in the map.

---

### Lab 44 — Word Frequency

**You will practice:** `map[string]int`, `strings.Fields`, `range` over a map

**Build this:**
- Ask for a sentence
- Count how many times each word appears
- Print each word and its count

**Example run:**
```
Enter sentence: the cat sat on the mat the cat
the: 3
cat: 2
sat: 1
on: 1
mat: 1
```

**Hint:** `wordCount := make(map[string]int)`. For each word: `wordCount[word]++`. If the key doesn't exist yet, Go automatically gives it a zero value — so `++` on a missing key starts from 0 and becomes 1.

---

### Lab 45 — Phone Book

**You will practice:** `map[string]string`, CRUD operations on a map

**Build this:**
- Create a phone book (map of name → phone number)
- Pre-fill it with 3 entries
- Allow the user to:
  - Look up a number by name
  - Add a new entry
  - Delete an entry
- Use a simple menu loop

**Example run:**
```
Phone Book
1. Look up
2. Add
3. Delete
4. List all
5. Quit
> 1
Enter name: Ali
Ali: 0300-1234567

> 4
Ali: 0300-1234567
Sara: 0321-9876543
Umar: 0333-1111222
```

**Hint:** To check if a key exists in a map: `number, exists := phonebook[name]`. If `exists` is false, the name is not in the book.

---

### Lab 46 — Student Report

**You will practice:** structs, struct fields, slice of structs

**Build this:**
- Define a `Student` struct with fields: `Name string`, `Score float64`
- Create a slice of 5 students (hardcoded)
- Print each student's name, score, and letter grade
- Print the class average at the end

**Example run:**
```
Name         Score   Grade
Ali          88.5    B
Sara         95.0    A
Umar         72.0    C
Fatima       61.5    D
Hassan       45.0    F

Class average: 72.40
```

---

### Lab 47 — Rectangle Methods

**You will practice:** structs, methods on structs, multiple methods

**Build this:**
- Define a `Rectangle` struct with `Width` and `Height float64`
- Write an `Area()` method that returns `Width × Height`
- Write a `Perimeter()` method that returns `2 × (Width + Height)`
- Write a `IsSquare()` method that returns true if Width == Height
- In `main`, create two rectangles and call all three methods

**Example run:**
```
Rectangle 1: 5.0 × 3.0
Area: 15.00
Perimeter: 16.00
Is square: false

Rectangle 2: 4.0 × 4.0
Area: 16.00
Perimeter: 16.00
Is square: true
```

**Hint:** A method is a function with a receiver: `func (r Rectangle) Area() float64 { return r.Width * r.Height }`. Call it like: `rect.Area()`.

---

### Lab 48 — Bank Account

**You will practice:** structs, methods that modify state, pointer receivers

**Build this:**
- Define a `BankAccount` struct with `Owner string` and `Balance float64`
- Write a `Deposit(amount float64)` method — adds to balance
- Write a `Withdraw(amount float64) error` method — subtracts from balance; return an error if funds are insufficient
- Write a `PrintStatement()` method — prints owner name and balance
- In `main`, simulate a few transactions

**Example run:**
```
Account: Ali
Balance: PKR 0.00

Deposited PKR 5000.00
Balance: PKR 5000.00

Withdrew PKR 2000.00
Balance: PKR 3000.00

Error: insufficient funds (tried to withdraw PKR 5000.00, balance is PKR 3000.00)
Balance: PKR 3000.00
```

**Hint:** Methods that modify the struct must use a pointer receiver: `func (a *BankAccount) Deposit(amount float64)`. Without the `*`, the method works on a copy and changes are lost.

---

### Lab 49 — In-Memory To-Do List

**You will practice:** struct with a `bool` field, slice of structs, menu loop, multiple methods

**Build this:**
- Define a `Todo` struct with `ID int`, `Text string`, `Done bool`
- Build a simple command-line app with a menu:
  1. Add a task
  2. List all tasks (show ✓ for done, ○ for pending)
  3. Mark a task as done (by ID)
  4. Delete a task (by ID)
  5. Quit

**Example run:**
```
1. Add  2. List  3. Done  4. Delete  5. Quit
> 1
Task: Buy groceries
Added ✓

> 1
Task: Call the vet
Added ✓

> 2
1. ○ Buy groceries
2. ○ Call the vet

> 3
Mark task ID as done: 1

> 2
1. ✓ Buy groceries
2. ○ Call the vet
```

---

### Lab 50 — Simple Stack

**You will practice:** structs, slice as underlying storage, methods, error handling

**Build this:**
- Define a `Stack` struct that holds a `[]int` internally
- Write a `Push(value int)` method — adds to the top
- Write a `Pop() (int, error)` method — removes and returns the top; error if empty
- Write a `Peek() (int, error)` method — returns top without removing; error if empty
- Write a `IsEmpty() bool` method
- Write a `Size() int` method
- In `main`, demonstrate all operations

**Example run:**
```
Stack is empty: true

Pushed: 10
Pushed: 20
Pushed: 30
Size: 3

Peek: 30
Size: 3

Pop: 30
Pop: 20
Size: 1

Pop: 10
Stack is empty: true

Pop from empty stack: error: stack is empty
```

**Hint:** The top of the stack is the last element of the slice. `Push` appends to the slice. `Pop` reads the last element, then re-slices to remove it: `s.items = s.items[:len(s.items)-1]`.

---

## Quick Reference

Packages used across these labs:

| Package | Import | Used for |
|---|---|---|
| `fmt` | `"fmt"` | Printing, scanning input |
| `math` | `"math"` | `math.Pi`, `math.Pow`, `math.Sqrt` |
| `strings` | `"strings"` | `Fields`, `ToLower`, `ToUpper`, `ReplaceAll`, `Repeat`, `Join` |
| `errors` | `"errors"` | `errors.New("message")` to create an error |
| `strconv` | `"strconv"` | `strconv.Atoi` (string → int) |

---

## Reading Input Tips

```go
// Read one word (stops at space)
var name string
fmt.Scan(&name)

// Read one integer
var n int
fmt.Scan(&n)

// Read one float
var f float64
fmt.Scan(&f)

// Read a full line (including spaces)
var line string
fmt.Scanln(&line)               // reads one line
// OR use bufio for reliable line reading:
reader := bufio.NewReader(os.Stdin)
line, _ := reader.ReadString('\n')
line = strings.TrimSpace(line)  // remove trailing newline
```

---

## What to Do When Stuck

1. **Read the error message carefully.** Go errors are specific. `undefined: strings` = you forgot `import "strings"`.
2. **Search [pkg.go.dev](https://pkg.go.dev)** for the package you're using. The docs have small examples.
3. **Write a smaller test.** If you're stuck on one piece of logic, write a 5-line program that only tests that piece.
4. **After 20 minutes on the same problem, ask.** Come with the error message and what you've tried.
