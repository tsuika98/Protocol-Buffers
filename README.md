# Protocol-Buffers 
This repo showcases a simple example of using protocol-buffers in a Go application to create an instance of a type, marshal it, and then unmarshal it. 

## Protocol Buffers Compilation

Compile .proto file into generated Go source code

```bash
 src> protoc --go_out=. *.proto
```

## Running Go Example

Method 1: Generate binary using go build and run separately 

```bash
 src> go build .
 src> ./src
```
Method 2: Use go run 
```bash
 src> go run .
```

## Example output
```bash
PS C:\Users\tsuika\Documents\Projects\Protocol-Buffers\src> go run .
2020/11/16 23:16:06 Person before marshal
2020/11/16 23:16:06 Name: tsuika
2020/11/16 23:16:06 Age: 22
2020/11/16 23:16:06 SocialFollowing: youtube:2500  twitter:500

2020/11/16 23:16:06 Person after marshalling: [10 6 116 115 117 105 107 97 16 22 26 6 8 196 19 16 244 3]

2020/11/16 23:16:06 Person after unmarshalling
2020/11/16 23:16:06 Name: tsuika
2020/11/16 23:16:06 Age: 22
2020/11/16 23:16:06 SocialFollowing: youtube:2500  twitter:500
```
