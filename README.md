# Merge, with [tailwind-merge](https://github.com/dcastil/tailwind-merge)
[![Go Reference](https://pkg.go.dev/badge/github.com/will-lol/merge.svg)](https://pkg.go.dev/github.com/will-lol/merge)

Merge is a simple Go package to merge HTML attributes.
```go
merger := merge.New(map[string][]merge.MergeFunc{
    "class": {merge.ClassMergeFunc},
})
res := merger.Merge(map[string]any{
    "class": "px-4 mx-2 bg-red-500",
}, map[string]any{
    "class": "px-4 font-bold",
})
log.Println(res)
// map[class:px-4 font-bold mx-2 bg-red-500]
// 
```
## tailwind-merge
It also supports [tailwind-merge](https://github.com/dcastil/tailwind-merge) through [goja](https://github.com/dop251/goja).
```go
twMerge, err := merge.NewTailwindMerge()
if err != nil {
    log.Fatalln(err)
}
merger := merge.New(map[string][]merge.MergeFunc{
    "class": {twMerge.TailwindMergeFunc},
})
res := merger.Merge(map[string]any{
    "class": "px-2 py-1 bg-red hover:bg-dark-red",
}, map[string]any{
    "class": "p-3 bg-[#B91C1C]",
})
log.Println(res)
// "map[class:hover:bg-dark-red p-3 bg-[#B91C1C]]"
```
⚠️ [tailwind-merge](https://github.com/dcastil/tailwind-merge) support is rather slow due to the JavaScript runtime running in the background. Unfortunately this is the best solution I have found to use [tailwind-merge](https://github.com/dcastil/tailwind-merge) for now. Please refer to the [tailwind-merge](https://github.com/dcastil/tailwind-merge) documentation for details on usage. 

The slowest function is by far merge.NewTailwindMerge(). I suggest you run this in a separate goroutine! Subsequent calls to merger.Merge will usually complete in under 1ms.

## Custom merge logic
It allows you to define your own merge logic by defining a custom MergeFunc.
```go
// MergeId prioritises any ID starting with '!'
func MergeId(existing any, incoming any) (remaining any, committed any) {
	const exclaimationMark = rune(33) // 33 is the char code for !
	existingString := fmt.Sprint(existing)
	existingId := strings.TrimSpace(existingString)

	if rune(existingId[0]) == exclaimationMark {
		return nil, existing
	}

	return existing, incoming
}
```
Then using it in your merger.
```go
merger := merge.New(map[string][]merge.MergeFunc{
    "class": {merge.ClassMergeFunc},
    "id": {MergeId},
})
res := merger.Merge(map[string]any{
    "class": "px-4 mx-2 bg-red-500",
    "id": "!veryImportantId",
}, map[string]any{
    "class": "px-4 font-bold",
    "id": "lessImportantId",
})
log.Println(res)
// map[class:px-4 font-bold mx-2 bg-red-500 id:!veryImportantId]
```
