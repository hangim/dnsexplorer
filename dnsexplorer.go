package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var (
	server_list chan string
	port_list   []int
	wg          sync.WaitGroup
)

func check(server string, port int) bool {
	cmd := exec.Command("dig",
		"baidu.com",
		fmt.Sprintf("@%s", server),
		fmt.Sprintf("-p%d", port),
		// "+tcp",
		"+time=2",
		"+tries=3")
	return cmd.Run() == nil
}

func main() {
	server_list = make(chan string)
	port_list = []int{53, 54, 80, 443, 1053, 5353, 27015} // modify if necessary

	for i := 0; i < 10; i++ {
		go func() {
			for {
				server := <-server_list
				for _, port := range port_list {
					if check(server, port) {
						fmt.Printf("%s#%d\n", server, port)
						break
					}
				}
				wg.Done()
			}
		}()
	}

	scanner := bufio.NewReader(os.Stdin)
	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
		server := strings.TrimSpace(line)
		if len(server) == 0 {
			continue
		}
		wg.Add(1)
		server_list <- server
	}

	wg.Wait()
}
