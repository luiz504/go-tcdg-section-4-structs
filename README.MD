
# [Section 4 | Go: The Complete Developer's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/)

This project was made to explore the Go structs

## 💡 Dev Notes: JS/TS → Go

* **Impressions:**:
  * Apparently, structs are kind like objects in JS.
  * Logging struct feels weird #console.log :), have to explore more fmt methods :/.
  * The , on the end of each line makes me miss prettier auto-formatting in VSCode :), let's check if there's a way to do it.
  * Mutation of struct field doesn't work like JS, lets dive into it. Pointers ahead captain!.
  * Humm, apparently, the receivers methods clone the (kind this value) to a different pointer, keeping the original intact. (Pass by value language)
  * Some magic syntax on pointers, kinda on my exploration, only change the receiver, argument adding the *, resolved but the class showed some boilerplate, maybe it's some old-school way to do it.
  * When logging a pointer, the output starts with the & symbol to indicate it's a pointer.
  * the (*x) operators only works if the argument has the "*" pointer indicator symbol.
  * the receiver ows the controls, if on the declaration, it says it's a pointer, even calling the method on a value, it works .

***Behavior***: When passed, the "Header" is copied, but it points to the same underlying data. You can mutate contents without an explicit pointer.

## Type |> Zero Value

When declaring some variables, and not assigning them a value, they are zero valued.
-- It's weird, not default undefined like JS, not nil either, let's check how they manage it.
-- Apparently, with structs, the zero value will follow the fields' types.

* string |> ""
* int |> 0
* float |> 0.0
* bool |> false

## Value Types

* int
* float
* string
* bool
* structs

***Behavior***: These are copied entirely. To mutate the original, you must use a pointer (*).

## Reference Types (Internal Pointer Types (Headers))

* slices
* maps
* channels
* pointers
* functions
