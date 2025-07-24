package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

const (
	ntpServer  = "ntp1.stratum1.ru"
	timeFormat = time.RFC3339
)

// getNTPTime возвращает текущее время с NTP-сервера или ошибку.
func getNTPTime(server string) (time.Time, error) {
	return ntp.Time(server)
}

func main() {
	// Получаем время с NTP-сервера
	ntpTime, err := getNTPTime(ntpServer)
	if err != nil {
		// Выводим ошибку в STDERR и завершаем с ненулевым кодом
		log.Printf("Ошибка получения времени: %v", err)
		os.Exit(1)
	}

	// Выводим точное время
	fmt.Printf("Точное время (NTP): %v\n", ntpTime.Format(timeFormat))
}
