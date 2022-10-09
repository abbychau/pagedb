# PageDB

PageDB is a fast database makes use of data offsets in msgpack-encoded files to retrieve data, so to allow fast deserialization of data to JSON, and at the same time, allow partial data retrieval without deserializing the whole file, or loading the whole file into memory. It is not only memory efficient, but also fast, because it does not care about the data that we don't need by skipping them using the offset hints in each of the msgpack-encoded data.

## Technical details

It is written in Golang.

## Development

1. Install [Golang](https://golang.org/doc/install)
2. Clone this repository
3. Run `go build` to build the binary

## Status

This is a work in progress. 