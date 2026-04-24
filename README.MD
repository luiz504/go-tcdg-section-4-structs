
# [Section 4 | Go: The Complete Developer's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/)

This project was made to explore the Go structs

## 💡 Dev Notes: JS/TS → Go

* **Impressions:**:
  * Apparently, structs are kind like objects in JS.
  * Logging struct feels weird #console.log :), have to explore more fmt methods :/.
  * The , on the end of each line makes me miss prettier auto-formatting in VSCode :), let's check if there's a way to do it.
  * Mutation of struct field doesn't work like JS, lets dive into it. Pointers ahead captain!.
  * Humm, apparently, the receivers methods clone the (kind this value) to a different pointer, keeping the original intact. (Pass by value language)

## Notes

When declaring some variables, and not assigning them a value, they are zero valued.
-- It's weird, not default undefined like JS, not nil either, let's check how they manage it.
-- Apparently, with structs, the zero value will follow the fields' types.

Type |> Zero Value
string |> ""
int |> 0
float |> 0.0
bool |> false
