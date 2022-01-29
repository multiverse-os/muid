[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse: `muid` **M**ultiversal *(Unique)* **Id**
**URL** [multiverse-os.org](https://multiverse-os.org)

Multiversal Unique ID, or `muid` (or `mid`), is an ultra lightweight id library,
similar to `snowflake` or `xid`. Producing ids using a variety of techniques to
build 8 to 12 byte identification that is time sortable and avoid collision. 

The default currenty is:
  * 4 Bytes for Timestamp (which functions as a nonce, and make it sortable)
  * 3 Bytes for Machine ID
  * 2 Bytes for PID
  * 3 Bytes of Random
  * 1 Byte for simple checksum (Add first 11 bytes and divide by 11)

*This project is still alpha phase, and the default may shift slightly but
developers taking advantage of the customability of our library will be able to
easily guarantee their id format remains static across versions. And this
library properly uses standard libraries making it incredibly easy to understand
and modify as needed.*

~3.67µs to generate a command using the `muid` command provided as both an
example and a tool for scripting languages. 

The library is built so that developers can customize their generated ids with
enough resolution to reproduce MongoDB `bsonid` using a methodology which
improves upon the existing go `bsonid` library. 

#### Example
Examples are included in the project in the `cmd` folder, in addition, a simple
command is included that simply outputs default keys, and soon will support
various customizations using flags so that it can be easily used from any
scripting language or other programming languages. 

The simplest possible example using the library from your go application:

```
package main

import (
  "fmt"

  id "github.com/multiverse-os/muid"
)

func main() {
  fmt.Println("id:", id.Generate().Base32().String())
  fmt.Println("id:", id.Generate().Base32().Prefix("mv-").String())
}
```

##### Why Use MUID?

It uses less memory than the alternatives, for example, there are no global 
variables or even constants.

Overall there is less over-head, as there are no external dependencies and 
the library just works after importing it into a project.

The codebase is smaller, easier to understand, modify, and customize, while 
not just simply ensuring feature parity with similar libraries, `muid` has
many new features. 

While the original inspiration comes from `bsonid` and `xid`, these libraries 
both use more resources, produce bigger ids, and their code does make use of the
standard libraries making them hard for some developers to understand and more
importantly hard for some developers to moidfy, customize, or contribute to. 
For example, `xid` has a lot of code pulled from the standard library 
`encoding/base32` to customize it to support downcased `base32`, and we simply
use the standard library. 

A simple example of this, we were able to replace ~30% of the `xid` codebase 
with simple use of the standard go library `encoding/base32`. `xid` implements
a custom `base32` encoding system to acheive no-padding and downcased hex
encoding. In contrast, we replaced over 100 lines with:
           
```        
encoder := base32.NewEncoding("0123456789abcdefghijklmnopqrstuv").WithPadding(base32.NoPadding)
base32Id := encoder.EncodeToString(id)
fmt.Println("custom encoder:", base32Id)
```


#### Features
`muid` provides a wide variety of features, but with no dependencies beyond Go
standard libraries. 

  * **Deterministic Keys**
  * **Checksums**
  * **Time Sorting**
  * **Ultra Low Collision Probability**
  * **Key Hashing**
  * **Optional Prefix**
  * **Key Compression**
  * **2..4 Timestamp Byte Length**
  * **Base32/Base64/Hex/String/Bytes Output Options**

`muid` utilizes a unique solution for a compressed (2 byte instead of 4 byte ) 
version of time, but the decision to use the 2 byte or 4 byte version is left
to the developers using the library. 

Provides the ability to either create deterministic id keys, or conceal
the id system or the details within the key. Using SHA3 Shake256's 
unique ability to create variable left tokens from any size message. 
Ensuring our keys are any desired byte or string length. The library is also
built so that it is easy to extend this hashing functionality by adding any
number of hashing algorithms. 

Ability to add a checksum, both a 1 byte and 4 byte version. Checksums are
supplied by either crc32 or adler32, and additionally an ultra simple an
version requiring a single byte by simply adding byte values and dividing 
by the number of bytes to give a single 1 byte checksum that requires 
no dependencies.

```
  2..4 Bytes  + 2..3 Bytes  + 2 Bytes + 0..N Bytes + 1..4 Bytes
  (Timestamp)  (MachineID)     (PID)     (Random)    (Checksum)
```

The resulting `muid` is minimum 8 bytes, default is 10 bytes, and can be as 
large as 32 bytes. Providing additional functionality over existing libraries, 
with far less resource usage, overhead and can approach 0% probability of 
collision. 

Machine ID is 2 bytes of the `crc32` checksum pulled from the system, otherwise
it falls back on random bytes which is perfectly fine for almost all use cases. 

The resulting Base32/Hex (default) or Hex only string is 20 bytes. 

The resulting URL-Safe base64 string is 16 bytes. 


Ability to easily prefix any id is already present. 

*A basic deterministic ID system based on a given seed is also planned. In
addition, the ability to hide the ID system using 2-way hashing, and various 
options for encoding and formatting will be expanded on soon.*


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



