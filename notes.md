# Notes

## Discipline

- Write a test

- Make the compiler pass

- Run the test, see that it fails and check the error message is meaningful

- Write enough code to make the test pass

- Refactor

## The TDD process and why the steps are important

- Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure

- Writing the smallest amount of code to make it pass so we know we have working software

- Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with.

## Others

- Strings in Go are immutable, meaning every concatenation, such as in our Repeat function, involves copying memory to accommodate the new string. This impacts performance, particularly during heavy string concatenation.

- The standard library provides the strings.BuilderstringsBuilder type which minimizes memory copying. It implements a WriteString method which we can use to concatenate strings

## pointers

- Go copies values when you pass them to functions/methods. To mutate state we will need it to take a pointer to the thing you want to change.

## nil

- Pointers can be nil

- When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.

- Useful for when you want to describe a value that could be missing
