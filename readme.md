# Gojas: Golang JSON Assertions. v0.1.0

## Values

### minimal
### compatible
### performant
### hackable
### simplifying
### low dep

## Pull Requests

Your PRs are welcome! The goal is to evolve this package toward more testing scenarios and reliability. Accepting issue posts as well, thank you.

## Basic Use

For most routine testing, just use the assertions in the package. For example

```
func TestSomething(t *testing.T) {

//...

    jsonDocumentAsString := LoadTheJsonFromSomewhere()

    // The path string represents a hierarchy within the json document.

    passed := gojas.AssertNumberAtPath(t, jsonDocumentAsString, "/user/properties/age", 42.0)

    // If the assertion does not pass, the 't' pointer will be used to set the test to Fail, with an error string.

    // There are several other assertions available.
}
```

