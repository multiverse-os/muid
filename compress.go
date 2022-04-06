package muid

type CompressionType uint8

const (
  NoCompression CompressionType = iota
  Snappy
  Zstd
)

// TODO: And the Id Compress() chainable 

func compress(compressionType CompressionType) {
	//CompressedStream := snappy.NewWriter(Data_writer)
}


