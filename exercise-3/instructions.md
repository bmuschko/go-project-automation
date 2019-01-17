# Exercise 3

In this exercise, you will use Go's `testing` package to implement a simple test. Furthermore, you will generate code coverage metric for the whole project.

1. Create a new Go file named `file_test.go` in the package `utils`.
2. Write two different test cases for the function `CreateDir`. The first test case should verify that the directory is created if it didn't exist before. The second test case should verify that not recreate the directory if it already exists. Create the new directory in a temporary directory which is deleted after the test cases finish.
3. Generate the code coverage for the `utils` package.
4. Create the HTML report for the coverage metrics in the `utils` package.
5. Generate code coverage metric for the whole project. Merge metrics if necessary to generate a single, unified coverage report.