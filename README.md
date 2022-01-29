[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse: `muid` **M**ultiversal *(Unique)* **Id**
**URL** [multiverse-os.org](https://multiverse-os.org)

Multiversal Unique ID, or `muid` (or `mid`), is an ultra lightweight id library,
similar to `snowflake` or `xid`. Producing ids using a variety of techniques to
build 8 to 12 byte identification that is time sortable and avoid collision. 

The library is built so that developers can customize their generated ids with
enough resolution to reproduce MongoDB `bsonid` using a methodology which
improves upon the existing go `bsonid` library. 

The library is designed to easily be customizable, use less memory than
alternatives, less over-head, while acheiving the same results or even better
using some unique techniques to shrink the byte size of the resulting ids. 

The resulting Ids can easily be converted to Hex, Bytes, Base32, URL-Safe and 
other encoding types.

Additionally some 2-way hash functions to hide the source id system while still 
providing developers a way to access all the unique features provided by `muid`.

While the original inspiration comes from `bsonid` and `xid`, these libraries 
both use more resources, produce bigger ids, and their code does make use of the
standard libraries making them hard for some developers to understand and more
importantly hard for some developers to moidfy, customize, or contribute to. 
For example, `xid` has a lot of code pulled from the standard library 
`encoding/base32` to customize it to support downcased `base32`, and we simply
use the standard library, and downcase it.

`muid` like almost all Multiverse libraries uses no external dependencies, are
written in pure Go, using only the standard libraries. Making code review,
easier and resulting code footprint smaller. 

#### Features
`muid` utilizes a unique solution for a compressed (2 byte instead of 4 byte ) 
version of time, but the decision to use the 2 byte or 4 byte version is left
to the developers using the library. 

Ability to add a checksum, both a 2 byte and 4 byte version. 

```
  2..4 Bytes  + 2..3 Bytes  + 2 Bytes + 0..N Bytes + 2..4 Bytes
  (Timestamp)  (MachineID)     (PID)     (Random)    (Checksum)
```

The resulting `muid` is minimum 8 bytes, default is 10 bytes, and can be as 
large as 32 bytes. Providing additional functionality over existing libraries, 
with far less resource usage, overhead and can approach 0% probability of 
collision. 

Machine ID is 2 bytes of the `crc32` checksum pulled from the system, otherwise
it falls back on random bytes which is perfectly fine for almost all use cases. 

The resulting Base32/Hex string (default, but this is customizable) is 14 bytes. 

xid is 12 bytes, and snowflake is >12 but both use more resources and have less
fucntionality. 

**A basic deterministic ID system based on a given seed is also planned.**

#### Commands
Included with the project are example commands, but also a generic command which
can produce ids so that it can easily be used with scripting languages. 

### Development
Development is specifically for the Multiverse project but all of our projects
are open to the public, we accept pull requests, and requests for features that
go beyond the scope of our needs as long as the features could reasonably be
used by other developers and do not increase the code footprint too
significiantly. 

The goal is to create libraries for use within the Multiverse project but also
for use by the wider Go developer community which is why we make effort to avoid
any external dependencies.



