package coding
 
import (
    "compress/gzip"
    "compress/zlib"
    "bytes"
    "io/ioutil"

)

// compress gzip and zlib
var (
    Compr   = ComprNon
    Decompr = ComprNon
)

func SetComprGzip(){
    Compr   = ComprGzip
    Decompr = DecomprGzip
}

func SetComprZlib(){
    Compr   = ComprZlib
    Decompr = DecomprZlib
}

func SetComprNon(){
    Compr   = ComprNon
    Decompr = ComprNon
}


func ComprNon(in []byte)(bs []byte,err error){
    return in,nil
}

//gzip
func ComprGzip(in []byte)(bs []byte,err error){
    var b bytes.Buffer
    w := gzip.NewWriter(&b)
    defer w.Close()
    _,err = w.Write(in)
    if err != nil {
        return
    }
    err = w.Flush()
    return b.Bytes(),err
}

func DecomprGzip(in []byte)(bs []byte,err error){
    b := bytes.NewBuffer(in)
    r, err := gzip.NewReader(b)
    if err != nil {
        return
    }
    defer r.Close()
    bs,_ = ioutil.ReadAll(r)
    return
}


// zlib
func ComprZlib(in []byte)(bs []byte,err error){
    var b bytes.Buffer
    w := zlib.NewWriter(&b)
    defer w.Close()
    _,err = w.Write(in)
    if err != nil {
        return
    }
    err = w.Flush()
    return b.Bytes(),err
}

func DecomprZlib(in []byte)(bs []byte,err error){
    b := bytes.NewBuffer(in)
    r, err := zlib.NewReader(b)
    if err != nil {
        return
    }
    defer r.Close()
    bs,_ = ioutil.ReadAll(r)
    return
}