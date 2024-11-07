# gog

Package gog provides some generic libraries.

⚠️WARNING: this repo is still under development, we found that the stdlib iter.Seq has not good performance in some simple cases.

- **gstd** provides some wrappers for standard library like sync.Map, sync.Pool, etc.
- **gmutex** provides some generics utils with mutex.
- **gstl** provides all kinds of containers stl.
- **gslice** provides some utils for slices.

[![codecov](https://codecov.io/gh/dashjay/gog/graph/badge.svg?token=QWD9F9EO1L)](https://codecov.io/gh/dashjay/gog)

generics library for golang


## Interfaces

### Slices
- [gslice.All](https://pkg.go.dev/github.com/dashjay/gog/gslice#All)
- [gslice.Any](https://pkg.go.dev/github.com/dashjay/gog/gslice#Any)
- [gslice.Avg](https://pkg.go.dev/github.com/dashjay/gog/gslice#Avg)
- [gslice.AvgN](https://pkg.go.dev/github.com/dashjay/gog/gslice#AvgN)
- [gslice.AvgBy](https://pkg.go.dev/github.com/dashjay/gog/gslice#AvgBy)
- [gslice.Contains](https://pkg.go.dev/github.com/dashjay/gog/gslice#Contains)
- [gslice.ContainsBy](https://pkg.go.dev/github.com/dashjay/gog/gslice#ContainsBy)
- [gslice.ContainsAny](https://pkg.go.dev/github.com/dashjay/gog/gslice#ContainsAny)
- [gslice.ContainsAll](https://pkg.go.dev/github.com/dashjay/gog/gslice#ContainsAll)
- [gslice.Count](https://pkg.go.dev/github.com/dashjay/gog/gslice#Count)
- [gslice.Find](https://pkg.go.dev/github.com/dashjay/gog/gslice#Find)
- [gslice.FindO](https://pkg.go.dev/github.com/dashjay/gog/gslice#FindO)
- [gslice.ForEach](https://pkg.go.dev/github.com/dashjay/gog/gslice#ForEach)
- [gslice.ForEachIdx](https://pkg.go.dev/github.com/dashjay/gog/gslice#ForEachIdx)
- [gslice.HeadO](https://pkg.go.dev/github.com/dashjay/gog/gslice#HeadO)
- [gslice.Head](https://pkg.go.dev/github.com/dashjay/gog/gslice#Head)
- [gslice.Join](https://pkg.go.dev/github.com/dashjay/gog/gslice#Join)
- [gslice.Min](https://pkg.go.dev/github.com/dashjay/gog/gslice#Min)
- [gslice.MinN](https://pkg.go.dev/github.com/dashjay/gog/gslice#MinN)
- [gslice.MinBy](https://pkg.go.dev/github.com/dashjay/gog/gslice#MinBy)
- [gslice.Max](https://pkg.go.dev/github.com/dashjay/gog/gslice#Max)
- [gslice.MaxN](https://pkg.go.dev/github.com/dashjay/gog/gslice#MaxN)
- [gslice.MaxBy](https://pkg.go.dev/github.com/dashjay/gog/gslice#MaxBy)
- [gslice.Map](https://pkg.go.dev/github.com/dashjay/gog/gslice#Map)
- [gslice.Clone](https://pkg.go.dev/github.com/dashjay/gog/gslice#Clone)
- [gslice.CloneBy](https://pkg.go.dev/github.com/dashjay/gog/gslice#CloneBy)
- [gslice.Concat](https://pkg.go.dev/github.com/dashjay/gog/gslice#Concat)
- [gslice.Subset](https://pkg.go.dev/github.com/dashjay/gog/gslice#Subset)
- [gslice.SubsetInplace](https://pkg.go.dev/github.com/dashjay/gog/gslice#SubsetInplace)
- [gslice.Replace](https://pkg.go.dev/github.com/dashjay/gog/gslice#Replace)
- [gslice.ReplaceAll](https://pkg.go.dev/github.com/dashjay/gog/gslice#ReplaceAll)
- [gslice.ReverseClone](https://pkg.go.dev/github.com/dashjay/gog/gslice#ReverseClone)
- [gslice.Reverse](https://pkg.go.dev/github.com/dashjay/gog/gslice#Reverse)
- [gslice.Repeat](https://pkg.go.dev/github.com/dashjay/gog/gslice#Repeat)
- [gslice.RepeatBy](https://pkg.go.dev/github.com/dashjay/gog/gslice#RepeatBy)
- [gslice.Shuffle](https://pkg.go.dev/github.com/dashjay/gog/gslice#Shuffle)
- [gslice.ShuffleInPlace](https://pkg.go.dev/github.com/dashjay/gog/gslice#ShuffleInPlace)
- [gslice.Chunk](https://pkg.go.dev/github.com/dashjay/gog/gslice#Chunk)
- [gslice.ChunkInPlace](https://pkg.go.dev/github.com/dashjay/gog/gslice#ChunkInPlace)
## Changelog

v0.1.4:
- Implement more slice tools like subset, replace, reverse, etc.

v0.1.3:
- Implement slice tools ()
  - gslice.*

v0.1.2:
- Implement stl containers
  - [x] stack
  - [x] list

v0.1.1:
- Implement locked value provides interface like rust
  - [x] gmutex.RWLockedValue
  - [x] gmutex.LockedValue

v0.1.0:
- Initial repo
- Implement wrappers for following types
  - [x] sync.Map
  - [x] sync.Pool