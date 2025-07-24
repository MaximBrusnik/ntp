package ntp

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

const (
	timeFormat = time.RFC3339
)

// PrintNTPTime печатает текущее время с NTP-сервера или ошибку.
func PrintNTPTime(server string) {
	ntpTime, err := ntp.Time(server)
	if err != nil {
		// Выводим ошибку в STDERR и завершаем с ненулевым кодом
		log.Printf("Ошибка получения времени: %v", err)
		os.Exit(1)
	}
	// Выводим точное время
	fmt.Printf("Точное время (NTP): %v\n", ntpTime.Format(timeFormat))
}
