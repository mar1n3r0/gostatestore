### Install
* Go: ``` go get github.com/mar1n3r0/gostatestore ```

### Data Flow
**gostatestore** gives go a state store based on stateful routines:

* Creates read and write channel
* Creates a long running listener goroutine and short-lived read and write goroutines.
* Updates state map with any type you pass as interface.
* Reader uses the data type as a state key and stores the value.
* Writer retrieves the value by data type as state key.

### API Reference

* #### Store:
  * ` gostatestore.NewStore() `: Create read and write channels, start listener goroutine and wait for the other routines
  * ` godux.Reader(interface{}) `: Works with any type passed as a memory address ex. &User{Name: "test", Username: "tester"}
  * ` godux.Writer(interface{}) `: Works with any type passed as a memory address ex. &User{Name: "test", Username: "tester"}

### License
MIT License.
