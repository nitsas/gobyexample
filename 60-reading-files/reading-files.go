// Before running this, do the following in your shell:
// echo "hello world\nsecond line" > /tmp/gobyexample-reading-files.txt

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func clearBuf(buf []byte) {
	for i := range buf {
		buf[i] = 0
	}
}

func main() {
	filename := "/tmp/gobyexample-reading-files.txt"
	data, err := os.ReadFile(filename)
	check(err)
	fmt.Println("data, err := os.ReadFile(filename)")
	fmt.Printf("string(data): %#v\n", string(data))
	fmt.Println()

	buf := make([]byte, 30)
	fmt.Println("buf := make([]byte, 30)")
	f, err := os.Open(filename)
	defer f.Close()
	check(err)
	numRead, err := f.Read(buf)
	check(err)
	fmt.Println("numRead, err := f.Read(buf)")
	fmt.Printf("numRead: %d, err: %e, string(buf): %#v\n", numRead, err, string(buf))
	fmt.Println()

	// If we read again without rewinding, we'll get an EOF panic.

	pos, err := f.Seek(6, io.SeekStart)
	check(err)
	// (f *File) Seek(offset int64, whence int)
	// whence means where to seek from:
	// - io.SeekStart == 0: beginning of file,
	// - io.SeekCurrent == 1: current position,
	// - io.SeekEnd == 2: end of file
	fmt.Println("f.Seek(6, io.SeekStart) -> f.Read(buf)")
	clearBuf(buf)
	numRead, err = f.Read(buf)
	fmt.Printf("numRead: %d, pos: %d, err: %e, string(buf): %#v\n", numRead, pos, err, string(buf))
	fmt.Println()

	// If we seek past the end of the file and read, we'll get an EOF panic.

	pos, err = f.Seek(-5, io.SeekCurrent)
	check(err)
	fmt.Println("f.Seek(-5, io.SeekCurrent) -> f.Read(buf)")
	clearBuf(buf)
	numRead, err = f.Read(buf)
	fmt.Printf("numRead: %d, pos: %d, err: %e, string(buf): %#v\n", numRead, pos, err, string(buf))
	fmt.Println()

	// io.ReadAtLeast(Reader, []byte, int)

	pos, err = f.Seek(-10, io.SeekEnd)
	check(err)
	fmt.Println("f.Seek(-10, io.SeekEnd) -> io.ReadAtLeast(f, buf, 12)")
	clearBuf(buf)
	// If we ask to read more than is available, it reads everything
	// but get an UnexpectedEOF at the same call.
	numRead, err = io.ReadAtLeast(f, buf, 12)
	fmt.Printf("numRead: %d, pos: %d, err: %e, string(buf): %#v\n", numRead, pos, err, string(buf))
	fmt.Println()

	// bufio.NewReader(io.Reader)

	pos, err = f.Seek(0, io.SeekStart)
	fmt.Println("f.Seek(0, io.SeekStart)")
	br := bufio.NewReader(f)
	// The underlying buffer has the default size of bufio.defaultBufSize: 4096.
	// We can also choose a buffer size via bufio.NewReaderSize(io.Reader, int)
	fmt.Println("br := bufio.NewReader(f)")
	fmt.Println("br.Buffered():", br.Buffered())
	bt, err := br.ReadByte()
	fmt.Printf("br.ReadByte(): %#v, err: %e\n", string(bt), err)
	fmt.Println("br.Buffered():", br.Buffered())
	bts, err := br.Peek(br.Buffered())
	fmt.Printf("br.Peek(br.Buffered()): %#v, err: %e\n", string(bts), err)
	bts, err = br.ReadBytes(byte('\n'))
	fmt.Printf("br.ReadBytes(byte('\\n')): %#v, err: %e\n", string(bts), err)
	fmt.Println("br.Buffered():", br.Buffered())
	bts, err = br.Peek(br.Buffered())
	fmt.Printf("br.Peek(br.Buffered()): %#v, err: %e\n", string(bts), err)

	// bufio also has the following types:
	// - Scanner: Splits a Reader on a delimiter and scans through delimited parts of the contents.
	// - Writer: Implements buffering for an io.Writer object.
	// - ReadWriter: Stores pointers to a Reader and a Writer. Implements io.ReadWriter.
}
