Statistics Calculator

This is a Go program that reads numerical data from a file and calculates basic statistical measures: average (mean), median, variance, and standard deviation. The results are printed to the console.
Usage

To use this program, you need to have Go installed on your system. You can run the program with the following command:

sh

go run main.go data.txt

Here, data.txt is the file containing the numerical data you want to analyze. Each number should be on a new line in the file.
How It Works
Reading Data

The readData function reads numerical data from a specified file and returns a slice of float64 values. It uses the bufio package to read the file line by line and the strconv package to convert each line to a float64.
Calculating Statistics

The calculateStatistics function takes a slice of float64 values and calculates the following:

    Average (mean): The sum of all the numbers divided by the count of numbers.
    Median: The middle value in the sorted list of numbers. If the list has an even number of values, the median is the average of the two middle values.
    Variance: The average of the squared differences from the mean.
    Standard Deviation: The square root of the variance, representing the dispersion of the dataset.

Main Function

The main function handles the following:

    Checks if the correct number of command-line arguments is provided.
    Reads the data from the specified file.
    Calculates the statistical measures.
    Prints the results to the console.

Example

Given a file data.txt with the following content:

1
2
3
4
5

Running the program:

sh

go run main.go data.txt

Will produce the following output:

yaml

Average: 3
Median: 3
Variance: 2
Standard Deviation: 1

Error Handling

The program handles errors in the following ways:

    If the file cannot be opened, an error message is printed and the program exits.
    If a line in the file cannot be parsed as a float, an error message is printed and the program exits.
    If the number of command-line arguments is incorrect, a usage message is printed and the program exits.

Requirements

    Go 1.18 or later

Installation

    Clone the repository or download the main.go file.
    Ensure you have Go installed on your system.
    Run the program with your data file as described in the Usage section.

License

This project is licensed under the MIT License.
Author

    eddy marvin

Feel free to customize this README file with your information and any additional details specific to your project.
