## JSONTypeWiz

---
Convert your JSON API Response to valid TypeScript types.

### Milestones

1. Convert json with only primitives to typescript type

Input:
```json lines
{
  "key": "value",
  "key2": true,
  "key3": 100
}
```
Output:
```typescript
type JSONTypeWiz ={
    "key": string,
    "key2":boolean
    "key3": number
}
```

Potential struct:

````go
package main
type Wiz struct {
    key string
    parent *Wiz 
    objType string 
    children *Wiz 
}

func getTypeString(wiz *Wiz)string {
	return ""
}
````