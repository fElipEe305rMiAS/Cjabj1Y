// 代码生成时间: 2025-10-12 22:39:45
 * To maintain clarity, the code is divided into clear sections with appropriate
 * comments and documentation. Error handling is implemented where necessary.
 *
 * The server adheres to Go best practices and is structured to ensure
 * maintainability and extensibility.
 */

package main

import (
    "fmt"
    "log"
    "net"
    "net/rpc"
    "os"
)

// Define the structure for the RPC service.
type Arithmetic int

// Define the methods that can be called remotely.
func (t *Arithmetic) Add(args []int, reply *[2]int) error {
    *reply = [2]int{args[0] + args[1], 0}
    return nil
}

// Define the methods that can be called remotely.
func (t *Arithmetic) Mul(args []int, reply *[2]int) error {
    *reply = [2]int{args[0] * args[1], 0}
    return nil
}

// Set up the RPC server.
func main() {
    var server *rpc.Server
    var a, b int

    // Instantiate the RPC service.
    arithmetic := new(Arithmetic)

    // Register the service with the RPC server.
    server = rpc.NewServer()
    server.Register(arithmetic)

    // Set up a TCP listener.
    listener, e := net.Listen("tcp", "localhost:1234")
    if e != nil {
        log.Fatal("Error listening: ", e)
    }

    fmt.Println("Listening on localhost:1234")

    // Accept connections in a loop.
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Error accepting: 