# High Performance Go Programming
## Avoid of Lock
https://github.com/easierway/concurrent_map

## Avoid of string key

## Avoid of unnecessary access
```Go
func matchMraidImage(cond *Condition, indexElem *indexElem, country, size string, sRankMraid *mvutil.SRankCreativeIdx) {

	for _, mraid := range indexElem.MapCreative[country].MraidImage[size] {

        ....
}

// When "contry" is decided, it is not necessary to access the map many times, the result can be stored and reused

func matchMraidImage(cond *Condition, creatives *map[string]mvutil.SCreative,  size string, sRankMraid *mvutil.SRankCreativeIdx) {

	for _, mraid := range creatives.MraidImage[size] {

        ....

}

```

## Avoid of memory allocation

```Go
for i:=0;i<100;i++{
  x:=i*2
  ...
}

// Avoid of memory allocation by reusing memory

var x int
for i:=0;i<100;i++{
  x=i*2
  ...
}
```
## Reference vs Value
There's no reference-passing in Go at all.

When using Map, Slice and Channel, you would feel like to passing by reference. Actually, it is not true. Map, Slice, Channel is the struct type which includes the point (memory address) to the backend storage. So, the value of the struct will be copied when passing, but they both point to the same backend storage.


## Avoid of auto-growing
