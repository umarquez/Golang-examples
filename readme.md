# Go (_Golang_) code examples and notes.

## blank_identifier
### http_get
Makes a **http** get request using the _blank identifier_ to ignore exceptions.

### pollution
Simple code that declares a variable that is never used, its asign the var to a _blank identifier_ to avoid `declared and unused` error.

## by_ref_and_by_val
A program that calls two functions,   one recive arguments **_by val_** and the other recive arguments **_by ref_**. It uses **pointers** to pass **_by ref_**.

## Concurrency & Parallelism
### Atomicity
Provides the capability to access a var only one process at a time.

### concurrency
> Is the composition of independently executing processes (_dealing with a lot of things at once_).

An example code that implements concurrency using the  **go** keyword for creating _goroutine_ and a `WaitGroup` just for keeping the main thread alive until all other threads ends.

### parallelism
> Is the simulraneously execution of (possibly related) computations (_doing a lot of things at once_).

An example code that shows how to implements parallelism using `runtime.GOMAXPROCS()` function inside the _init()_ function.

### rc_mutex
Example of a rance condition code and mutex for preventing unwanted behavio r.

## constants
Program that declares and use some constants, it also use the _`iota`_ value.
> **iota:** a small amount of something

## control_flow
Code example that implements the uses for the `for` statement.
- C/Java `for` style
- Like `while` statement sustitution
- Like `do` for infinite loops
- With the `continue` keyword

## Convertion and Assertion
**Convertion**: Transform a value of certain type to another (casting or parsing in other languajes).
**Assertion**: Treats a value of one type as if it were another, without converting it.

### assertion_ex
Code that shows how to use assertion (only works with interfaces).

### convertion_ex
Code example that shows how to convert a typed value to another using ***strconv*** library.

## data_structures
Basic data structuresand how to deal with.

### array_and_slice
Example with an _**Array**_ and a _**Slice**_ and the differences between these types.

### slice_example
How to declare a slice and append elements to it.

### maps
Example code of maps declaration and manipulation.

### struct
This code shows how to create a custom type based on **struct** type, how to use it in a var declaration and how to change its field's values.

## fmt_verbs
Uses some format verbs to display a number.
- Deciamal
- Binary
- Char
- Octal
- Hex (_3 formats_)

## for_and_format_examples
Uses a for to generate a numerical secuences and display every generated number in three formats.
- Decimal `%d`
- Binary `%b`
- Hex `%#X`
- Char `%q`

## for_and_module_op
Code example that implements a for to generate a counter and a conditional (_if_) to clasificate generated numbers in **odd** and **even** using a module operation `x % 2`.

## functions
Working with functions.

### anon_function
A simple code that shows how to declare and execute a anonymous function.

### callback
Example of a function call that needs a callback function to do certain things.

### returns
Example of three types of function's _returns_
- **helloWorld**: A simple (and _traditionsl_) single value return
- **greet**: A named var return
- **names**: A multiple values return

### defer
Code that uses the word `defer` to make a function call be executed at the end of the current function.

### variadic_args
Example of a function call that passes variadic arguments using  the variadic technique.

### variadic_params
Example of a function that can recive a unlimited number of arguments and stores in a collection, it's called variadic function.

## inner_functions
Program with inner functions and functions that returns functions as result. Something like _functions inception_.

## Interfaces
Some interfaces examples.

### empty_interface
How to use empty interface to create a function that could manage any type of data/object.

### simple_interface
Code that shows how to declare an interface and work with it.

### sort_exercise
Requirements:

```
Use https://godoc.org/sort to sortthe following:

(1)
type people []string
studyGroup := people{"Zeno", "John", "Al", "Jenny"}

(2)
s := []string{"Zeno", "John", "Al", "Jenny"}

(3)
n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
```

## JSON
Working with JSON

### json_decoder
How to convert a streammed JSON text to an object.

**Note:** Use this for reading JSON through the web.

### json_encoder
How to convert an object to a JSON string to write it on a stream.

**Note:** Use this for sending Objects JSON through the web.

### json_to_struct_unmarshal
Having a string with JSON content, this code will assign its values to an instance of a type that has the same **exported** fields using `json.unmarshal()` method.

### struct_to_json_marshal
Having a custom class object(_struct_), this code will convert it to a **JSON** formated string using `json.Marshal()` method.

## memmory_address
A wey to get var's memmory address, good introduction to pointers.

## nested_loops
Example of nested loops, prints values of both vars.

## packages_and_imports
How to use the `ìmport` statement for multiple libs.

## pointers
How to declares and use pointers in _**GO**_.

## runes
runes generator example `rune = char`, _**rune**_ type is an alias of int32.

## switch_case_default
Example code using switch statements.

## switch_on_types
An example of how to implement a type based switch function.

## types
A program that defines twoo custom types using **struct**, one inside the other; both of them has a method with the same name that the code calls for each type, it also contains a function that receive a pointer to a custom type.

## variables
Example about the both ways to declare a variable.

## vars_scope
How does scopes works on _**Go**_:

### main
Main code, only imports a lib and calls a functions

### outter
Has two files in the same package, one has the var declaration and the other has the function that prints text messages.
