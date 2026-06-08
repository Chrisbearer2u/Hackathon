# Hackathon

# Question 1

Write a Go program that accepts a filename as a command-line argument (os.Args), reads the entire content of the file, and prints it to stdout. Handle the case where no argument is provided or the file doesn't exist with a meaningful error message.

# Question 2

Extend your file reader: now accept TWO command-line arguments (input file, output file). Read the input file, convert all text to uppercase using strings.ToUpper, and write the result to the output file using os.WriteFile. Make sure to use the correct file permission flags.

# Question 3

Write a function hexToDec(s string) (int64, error) that converts a hexadecimal string to its decimal integer equivalent. Test it with inputs like '1E' (should return 30), 'FF' (should return 255), and 'BADF00D'. Use only the standard library.
💡
# Question 4