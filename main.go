package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Получение IP-адреса
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Ошибка при получении IP-адреса:", err)
		os.Exit(1)
	}

	fmt.Println("Информация о сервере:")

	// Перебор всех доступных адресов и вывод IP
	for _, addr := range addrs {
		// Исключаем случаи, где нет IP
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("IP-адрес:", ipNet.IP.String())
			}
		}
	}

	// Получение и вывод имени хоста
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Имя хоста:", hostname)
	} else {
		fmt.Println("Ошибка при получении имени хоста:", err)
	}
}
