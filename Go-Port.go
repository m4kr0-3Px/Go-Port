package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("\n" + "\033[32m" + "------------Welcome To My Script------------" + "\033[0m" + "\n")

	fmt.Println("\033[31m" + "1-Scan all ports\n\n2-Scan specific port\n\n3-Scan well-known ports" + "\033[0m" + "\n")
	var process int
	fmt.Print("Enter your option: ")
	fmt.Scan(&process)
	switch process {

	case 1:
		scann_all()

	case 2:
		scan_specific()
	case 3:
		well_known_ports()
	default:
		fmt.Println("\033[31m" + "Unfortunately wrong option!" + "\033[0m")
		break

	}

}

func scan_specific() {
	loop := true
	for loop {
		var question string
		fmt.Print("\n" + "\033[33m" + "Continue to scan?(y/n): " + "\033[0m")
		fmt.Scan(&question)

		if question == "n" || question == "N" {
			loop = false

		} else if question == "y" || question == "Y" {

			var address string
			fmt.Print("\n" + "\033[31m" + "Enter the addres: " + "\033[0m")
			fmt.Scan(&address)
			var port int
			fmt.Print("\n" + "\033[31m" + "Enter the port addres: " + "\033[0m")
			fmt.Scan(&port)

			if port < 0 || port > 65535 {
				fmt.Println("\n" + "\033[31m" + "There is no such port" + "\033[0m")
				break
			} else {
				if address[:7] == "http://" {
					fmt.Println("\n" + "\033[35m" + "Scanning is starting,please wait..." + "\033[0m" + "\n")
					port_no := strconv.Itoa(port)

					conn, err := net.DialTimeout("tcp", address[8:]+":"+port_no, 3*time.Second)

					if err != nil {
						fmt.Println("\033[31m" + "Port is closed" + "\033[0m")
					} else if *&conn != nil {
						fmt.Println("\033[32m" + "Port is open-----> " + port_no + "\033[0m")

					} else {
						fmt.Println("\033[31m" + "Result is didn't return" + "\033[0m")

					}

				} else if address[:8] == "https://" {
					fmt.Println("\n+\033[35m" + "Scanning is starting,please wait..." + "\033[0m" + "\n")
					port_no := strconv.Itoa(port)

					conn, err := net.DialTimeout("tcp", address[8:]+":"+port_no, 3*time.Second)

					if err != nil {
						fmt.Println("\033[31m" + "Port is closed" + "\033[0m")
					} else if *&conn != nil {
						fmt.Println("\033[32m" + "Port is open-----> " + port_no + "\033[0m")

					} else {
						fmt.Println("\033[31m" + "Result is didn't return" + "\033[0m")

					}
				} else {
					fmt.Println("\n" + "\033[35m" + "Scanning is starting,please wait..." + "\033[0m" + "\n")
					port_no := strconv.Itoa(port)

					conn, err := net.DialTimeout("tcp", address+":"+port_no, 3*time.Second)

					if err != nil {
						fmt.Println("\033[31m" + "Port is closed" + "\033[0m")
					} else if *&conn != nil {
						fmt.Println("\033[32m" + "Port is open-----> " + port_no + "\033[0m")

					} else {
						fmt.Println("\033[31m" + "Result is didn't return" + "\033[0m")

					}

				}

			}

		} else {
			fmt.Println("\033[31m" + "Wrong asnwer!" + "\033[0m")
		}

	}

}

func scann_all() {

	var address string
	fmt.Print("\n" + "\033[33m" + "Enter the address: " + "\033[0m")
	fmt.Scan(&address)
	fmt.Println("\n" + "\033[36m" + "Scanning is starting,please wait..." + "\033[0m" + "\n")
	if address[:7] == "http://" {
		for i := 1; i <= 65535; i++ {
			port := strconv.Itoa(i)
			conn, err := net.DialTimeout("tcp", address[7:]+":"+port, 3*time.Second)

			if err != nil {
				continue
			} else if *&conn != nil {
				fmt.Println("\033[32m" + "Port is open-----> " + port + "\033[0m")
			} else {
				fmt.Println("\033[31m" + "Result is didn't return" + "\033[0m")
			}
			//defer conn.Close()

		}

	} else if address[:8] == "https://" {

		for i := 1; i < 65535; i++ {
			port := strconv.Itoa(i)
			conn, err := net.DialTimeout("tcp", address[8:]+":"+port, 3*time.Second)

			if err != nil {
				continue
			} else if *&conn != nil {
				fmt.Println("\033[32m" + "Port is open-----> " + port + "\033[0m")
			} else {
				fmt.Println("\033[31m" + "Result is didn't return" + "\033[0m")
			}
			//defer conn.Close()

		}

	} else {

		for i := 1; i <= 65535; i++ {
			port := strconv.Itoa(i)
			conn, err := net.DialTimeout("tcp", address+":"+port, 3*time.Second)

			if err != nil {
				continue
			} else if *&conn != nil {
				fmt.Println("\033[32m" + "Port is open-----> " + port + "\033[0m")
			} else {
				fmt.Println("\033[31m" + "Result is didn't return" + "\033[0m")
			}
			//defer conn.Close()

		}

	}

}
func well_known_ports() {
	var address string
	fmt.Print("\n" + "\033[33m" + "Enter the address: " + "\033[0m")
	fmt.Scan(&address)
	fmt.Println("\n" + "\033[36m" + "Scanning is starting,please wait..." + "\033[0m" + "\n")
	known_ports := map[int]string{

		20:   "FTP data transfer",
		21:   "FTP command control",
		22:   "SSH (Secure Shell)",
		23:   "Telnet",
		25:   "SMTP (Simple Mail Transfer Protocol)",
		53:   "DNS (Domain Name System)",
		80:   "HTTP (Hypertext Transfer Protocol)",
		110:  " POP3 (Mail Office Protocol, version 3)",
		119:  " NNTP (Network News Transfer Protocol)",
		123:  " NTP (Network Time Protocol)",
		143:  " IMAP (Internet Message Access Protocol)",
		161:  " SNMP (Simple Network Management Protocol)",
		194:  " IRC (Internet Relay Chat)",
		443:  " HTTPS (Secure HTTP)",
		465:  " SMTPS (SSL over SMTP)",
		514:  " Syslog",
		993:  " IMAPS (SSL over IMAP)",
		995:  " POP3S (SSL over POP3)",
		1723: "PPTP (Point-to-Point Tunneling Protocol)",
		3306: "MySQL Database Service",
	}

	if address[:7] == "http://" {
		for port_nums, port_name := range known_ports {

			ports := strconv.Itoa(port_nums)
			port_con, err := net.DialTimeout("tcp", address[7:]+":"+ports, 3*time.Second)

			if err != nil {
				continue
			} else if *&port_con != nil {
				fmt.Println("\033[32m" + "Port is open----->" + "\033[0m" + ports + "\033[35m" + "    =======>    " + "\033[0m" + port_name)
			} else {
				fmt.Println("\033[31m" + "Result is didn't return" + "\033[0m")
			}

		}

	} else if address[:8] == "https://" {
		for port_nums, port_name := range known_ports {

			ports := strconv.Itoa(port_nums)
			port_con, err := net.DialTimeout("tcp", address[8:]+":"+ports, 3*time.Second)

			if err != nil {
				continue
			} else if *&port_con != nil {
				fmt.Println("\033[32m" + "Port is open----->" + "\033[0m" + ports + "\033[35m" + "    =======>    " + "\033[0m" + port_name)
			} else {
				fmt.Println("\033[31m" + "Result is didn't return" + "\033[0m")
			}

		}

	} else {
		for port_nums, port_name := range known_ports {

			ports := strconv.Itoa(port_nums)
			port_con, err := net.DialTimeout("tcp", address+":"+ports, 3*time.Second)

			if err != nil {
				continue
			} else if *&port_con != nil {
				fmt.Println("\033[32m" + "Port is open----->" + "\033[0m" + ports + "\033[35m" + "    =======>    " + "\033[0m" + port_name)
			} else {
				fmt.Println("\033[31m" + "Result is didn't return" + "\033[0m")
			}

		}

	}

}
