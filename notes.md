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

## maps

- A map value is a pointer to a runtime.hmap structure.

- An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)

- Go has a built-in function delete that works on maps. It takes two arguments and returns nothing. The first argument is the map and the second is the key to be removed.

## sync

- Sync provides basic synchronization primitives that help manage councurrent programming.

- Mutex allows us to add locks to our data. Which help us avoid conflict in mutation of same variable.

- WaitGroup is a means of waiting for goroutines to finish jobs.

### go vet

- use go vet to check your build scripts. It alerts you with any subtle bugs in the code.

## context

- The context package provides functions to derive new Context values from existing ones. These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.
