

package loaders


import(
    //"log"
    //"github.com/FreeFlightSim/theairlineproject-server/go-airline/airline"
    //"github.com/FreeFlightSim/theairlineproject-server/go-airline/airline"
    //"github.com/FreeFlightSim/theairlineproject-server/go-airline/finance"
)



var DATA_PATH = "../Data"



func GetDataPath(file_name string) string {
    
    return DATA_PATH + file_name
}
