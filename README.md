goset
======

goset provides a set implementation for Go.

### Usage

Removing duplicate items:
```go
    fruits := []string{"Apple", "Banana", "Apple", "Strawberry"}
	set := goset.From(fruits)
	deduped := set.Slice()
	sort.Strings(deduped)

	fmt.Println(deduped)
	// [Apple Banana Strawberry]
```

Iterating items:
```go
    set := goset.From([]int{1, 2, 3, 4, 5})
    set.Range(func(n int) bool {
        fmt.Println(n)
        // 1, 2, ...
        return false
    })
```
