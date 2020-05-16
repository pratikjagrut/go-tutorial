# Map

Map maps the key to value. The zero value of a map is nil. 
A nil map has no keys, nor can keys be added.

**Create map using make function**

The make function returns a map of the given type, initialized and ready for use.

```go
m := make(map[int]string)
```

**Map literals**

```go
n := map[string]int{"foo": 0, "bar": 1}
```

**Mutating Maps**

- Insert or update an element in map m:
    `m[key] = element`

- Retrieve an element:
    `element = m[key]`

- Delete an element:
    `delete(m, key)`

- Test that a key is present with a two-value assignment:
    `element, ok := m[key]`
    - If `ok -> true` key is present in the map, if not, `ok -> false`.

***You can refer main.go file for examples***

To run:
```
go run main.go
```