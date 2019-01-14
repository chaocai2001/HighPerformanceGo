# High Performance Go Programming
## Avoid of Lock
https://github.com/easierway/concurrent_map
https://github.com/chaocai2001/HighPerformanceGo/blob/master/rlock_test.go

## Avoid of string key
https://github.com/chaocai2001/HighPerformanceGo/blob/master/map_string_key_vs_int_key_test.go
https://medium.com/@ConnorPeet/go-maps-are-not-o-1-91c1e61110bf

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
https://github.com/chaocai2001/HighPerformanceGo/blob/master/reference_vs_value_passing_test.go

## Avoid of auto-growing
https://github.com/chaocai2001/HighPerformanceGo/blob/master/slice_auto_grow_test.go
