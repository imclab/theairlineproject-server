

package loaders

import(
    //"fmt"
    "os"
)




// For not the Data is a symbolic link
var DATA_PATH = "Data"



func GetDataPath(file_name string) string {
    
    return DATA_PATH + file_name
}





func GetDataFileBytes(file_path string) ([]byte, error) {
    
    data_path := GetDataPath(file_path)
    
    file, err := os.Open(data_path)
    defer file.Close()
    if err != nil {
        return nil, err
    }
    
    fi, err := file.Stat()
    filelen := fi.Size()
    buf := make([]byte, filelen)
    
    read, err := file.Read(buf)
    if read != int(filelen) {
        return nil, err
    }
    return buf, err    
}

