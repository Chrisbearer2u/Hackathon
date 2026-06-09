# Hackathon

# Question 1
Write a Go program that accepts a filename as a command-line argument (os.Args), reads the entire content of the file, and prints it to stdout. Handle the case where no argument is provided or the file doesn't exist with a meaningful error message.

# Question 2
Extend your file reader: now accept TWO command-line arguments (input file, output file). Read the input file, convert all text to uppercase using strings.ToUpper, and write the result to the output file using os.WriteFile. Make sure to use the correct file permission flags.

# Question 3
Write a function hexToDec(s string) (int64, error) that converts a hexadecimal string to its decimal integer equivalent. Test it with inputs like '1E' (should return 30), 'FF' (should return 255), and 'BADF00D'. Use only the standard library.

# Question 4
Write a function binToDec(s string) (int64, error) that converts a binary string to its decimal equivalent. Test it with '10' (should return 2), '1010' (should return 10), '11111111' (should return 255). Think about what base to pass to strconv.ParseInt.

# Question 5
Write a function that takes a string and returns a slice of all words/tokens split by whitespace. Then write a second version that preserves punctuation as separate tokens. Compare the outputs for: 'Hello, world! How are you?' Hint: look at strings.Fields vs strings.Split.

# Question 6
Write a function capitalize(s string) string that capitalizes only the first letter of a word (making the rest lowercase). Then write capitalizeN(words []string, n int) []string that capitalizes the last N words in the slice. This mimics the (cap, N) modifier in go-reloaded.

# Question 7
Write a function fixArticles(words []string) []string that scans a word slice and replaces 'a' with 'an' whenever the next word starts with a vowel (a, e, i, o, u) or the letter 'h'. Handle both uppercase and lowercase 'A'/'a'.

# Question 8
Write a function that takes a string and ensures punctuation marks (. , ! ? : ;) are directly attached to the preceding word with exactly one space after them. Example: 'Hello , world !' becomes 'Hello, world!'. Also handle groups like '...' and '!?' that should stay together.

# Question 9
Write a function that finds pairs of single quotes in a string and removes spaces between the quotes and the enclosed text. Example: ' awesome ' becomes 'awesome' and ' hello world ' becomes 'hello world'. Think about how to track odd/even quote occurrences.

# Question 10
Write a function that reads one of the ascii-art banner files (standard, shadow, thinkertoy) and parses it into a map[rune][8]string — mapping each ASCII character to its 8 lines of art. Remember: each character is 8 lines tall, separated by a newline.

# Question 11
Using the map you built in the previous question, write renderASCII(input string, banner map[rune][8]string) string that takes a string and returns the multi-line ASCII art representation. Each output line should be built by concatenating the corresponding art line for each character.

# Question 12
Modify your ASCII renderer to handle literal '\n' in the input string (e.g. 'Hello\nWorld'). When you encounter '\n', finish the current line group and start a new one. An empty string between two '\n\n' should produce a blank 8-line block.

# Question 13
Write a function using regexp to find all occurrences of modifier patterns like (hex), (bin), (up), (low), (cap), (up, 2), (low, 3), (cap, 5) in a string. Return each match along with its position and the optional number argument if present.

# Question 14
Refactor a function that does: read file → parse content → write output into idiomatic Go style. Each step should return an error. The caller should use if err != nil checks. Wrap errors with fmt.Errorf and %w for context. Demonstrate errors.Is() usage.

# Question 15
Write a _test.go file that tests your hexToDec and binToDec functions using table-driven tests. Create a struct with fields: input string, expected int64, wantErr bool. Use t.Run() for subtests and t.Errorf() to report failures. Run with go test ./...

# Question 16
Write a Go HTTP server that listens on port 8080. Register two routes: GET / serves an HTML page, and POST /convert accepts a JSON body {"text": "...", "type": "hex|bin"} and returns the converted result as JSON. Use only the standard library.

# Question 17
Design a TextProcessor struct that holds a slice of Transformer interfaces. Each Transformer has a single method: Transform(words []string) []string. Implement at least two concrete transformers: UpperCaseTransformer and HexTransformer. Chain them in the processor's Run() method.

# Question 18
Write two versions of a 'capitalize first letter' function: one using byte indexing (s[0]) and one using []rune(s). Test both with ASCII strings AND a string starting with a multi-byte unicode character like 'ñoño'. Explain why byte indexing can corrupt unicode strings.

# Question 19
Given a slice of 5 filenames, process each file concurrently using goroutines. Each goroutine should read the file and send its word count to a channel. The main goroutine collects all results and prints a summary. Use sync.WaitGroup OR a done channel to know when all are finished.

# Question 20
FINAL BOSS: Build a mini version of go-reloaded that handles ONLY these three transformations: (hex), (bin), and (up). Accept input/output filenames from os.Args. Read the input file, apply all three transformations in a single pass over the token slice, write the result. All edge cases handled.
