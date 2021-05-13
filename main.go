package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "log"
)

type charInfo struct {
    validChars      []byte
    invalidChars    []byte
}

type fileInfo struct {
    filename    string
    filesize    int64
    data        []byte
    c_i         []charInfo
    index       int
    indexes     []int
}

func setup(filename string) *fileInfo {
    info := fileInfo{ filename: filename }
    
    file, err := os.Stat(filename)
    if err != nil {
        log.Fatal("File does not exist")
    }
    info.filesize = file.Size()
    
    data, err := ioutil.ReadFile(info.filename)
    if err != nil {
        log.Fatal("Error reading the file")
    }
    
    info.data = data
    return &info
}

func (info *fileInfo) gather_info() {
    c_i := charInfo{ }
    for i := 0; i < int(info.filesize); i++ {
        if info.data[i] < 0x41 || info.data[i] > 0x7a {
            if !(info.data[i] == 0x20 || info.data[i] == 0x0a) {
                c_i.invalidChars = append(c_i.invalidChars, info.data[i])
                info.indexes = append(info.indexes, i)
            } else {
                c_i.invalidChars = append(c_i.invalidChars, info.data[i])
                c_i.validChars = append(c_i.validChars, info.data[i])
            }
        } else {
            c_i.validChars = append(c_i.validChars, info.data[i])
        }
    }
    
    if len(c_i.invalidChars) > 0 || len(c_i.validChars) > 0 {
        info.c_i = append(info.c_i, c_i)
        info.index += 1
    }
}

func (info *fileInfo) print_valid() {
    for i := 0; i < int(info.filesize); i++ {
        if !(i >= len(info.c_i[info.index - 1].validChars) - 1) {
            fmt.Print(string(info.c_i[info.index - 1].validChars[i]))
        }
    }
}
func (info *fileInfo) print_invalid() {
    for i := 0; i < int(info.filesize) - 20; i++ {
        if !(i >= len(info.c_i[info.index - 1].invalidChars) -1) {
            fmt.Print(string(info.c_i[info.index - 1].invalidChars[i]))
        }
    }
}

func main() {
    info := setup("main.go")
    info.gather_info()
    info.print_invalid()
    //fmt.Println(string(info.c_i[info.index - 1].invalidChars))
}
