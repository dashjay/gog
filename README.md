# gog

Package gog provides some generic libraries.

- **gstd** provides some wrappers for standard library like sync.Map, sync.Pool, etc.
- **gmutex** provides some generics utils with mutex.
- **gstl** provides all kinds of containers stl.

[![codecov](https://codecov.io/gh/dashjay/gog/graph/badge.svg?token=QWD9F9EO1L)](https://codecov.io/gh/dashjay/gog)

generics library for golang


## Changelog

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