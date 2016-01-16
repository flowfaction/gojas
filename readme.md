# Gojas: Golang JSON Assertions

## Values:
### Compatible, Fast, Hackable, Usable, Native, Minimal, Readable.

## Importing

To add to your project, import:
```
import "github.com/flowfaction/gojas"
```
There are no dependencies on 3rd party components, only native Go packages.


## Pull Requests

Your PRs are welcome! The goal is to evolve this package toward more testing scenarios and reliability. Accepting issue posts as well, thank you.

## Example

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

