package port

import (
	"net"
	"strconv"
	"time"
)

var scan_me = map[int]string{ // maps each key with a value for example in this case it maps 80 to http.
	80:  "http",
	443: "https", // the ports to scan
}

func scanport(hostname string, port int) bool { // refering to the each variable with their variable type. 'bool' is for the return statement.
	address := hostname + ":" + strconv.Itoa(port)               // 'strconv.Itoa' is for converting Integers to Strings
	conn, err := net.DialTimeout("tcp", address, 20*time.Second) // the type of connection TCP/UDP, the host URL/IP ADDRESS with the port, the max time for each connection.
	if err != nil {
		return false
	}
	defer conn.Close() // 'defer' wait until the entire function is done and then execute the given command.
	return true
}
