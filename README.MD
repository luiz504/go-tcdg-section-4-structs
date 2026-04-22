
# [Section 4 | Go: The Complete Developer's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/)

This project was made to explore the Go structs

## 💡 Dev Notes: JS/TS → Go

* **First Impressions:**:
  * Apparently, structs are kind like objects in JS.
  * Logging struct feels weird #console.log :), have to explore more fmt methods :/.

## Notes

When declaring some variables, and not assigning them a value, they are zero valued.
-- It's weird, not default undefined like JS, not nil either, let's check how they manage it.
-- Apparently, with structs, the zero value will follow the fields' types.

Type |> Zero Value
string |> ""
int |> 0
float |> 0.0
bool |> false
