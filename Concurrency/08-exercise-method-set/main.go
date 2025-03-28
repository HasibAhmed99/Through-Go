package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("nothing to say")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	saySomething(&p1)

	// does not work because it will only take pointer
	// saySomething(p1)

	p1.speak()
}

/* The reason you have to call `h.speak()` explicitly inside the function `saySomething(h human)` even though the `human` interface already declares `speak()` is because:

1. **Interfaces in Go Only Define Behavior, Not Implementation**
   - The `human` interface defines that any type satisfying it **must have** a `speak()` method, but it does not **automatically execute** that method.
   - Interfaces only tell Go: *"Any type that has a `speak()` method can be treated as a `human`."*

2. **The Interface Variable (`h`) Holds a Value of the Underlying Type**
   - When you pass `p1` (as `&p1`) to `saySomething()`, the function receives `h`, which is of type `human`.
   - However, Go does **not** call `speak()` just because `h` exists. It just **stores the reference** to a value that has a `speak()` method.
   - To actually execute `speak()`, you must explicitly call `h.speak()`.

3. **Go’s Static Type System and Explicit Calls**
   - Unlike some object-oriented languages where methods might get called automatically, Go follows an explicit approach:
     - The interface acts as a **contract** (defining what methods exist).
     - But you must explicitly call those methods on an interface variable when you need them.

### **Example to Demonstrate This**
Consider this code:
```go
type human interface {
    speak()
}

type person struct {
    name string
}

func (p *person) speak() {
    fmt.Println("Hello, my name is", p.name)
}

func saySomething(h human) {
    // h holds a value that satisfies human, but speak() must be explicitly called
    h.speak()
}

func main() {
    p1 := person{name: "Alice"}

    // saySomething(p1)  // ERROR: person does not implement human (only *person does)
    saySomething(&p1)   // WORKS: *person implements human
}
```
🔹 The **interface does not execute methods automatically**; you need to explicitly call `h.speak()` inside `saySomething()`.

*/
