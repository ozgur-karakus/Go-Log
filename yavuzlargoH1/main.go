package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Println("Lütfen giriş türünüzü seçin:")
		fmt.Println("0 - Admin Girişi")
		fmt.Println("1 - Öğrenci Girişi")

		var choice int
		fmt.Print("Seçenek: ")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			adminLogin()
		case 1:
			studentLogin()
		default:
			fmt.Println("Geçersiz seçenek!")
		}
	}
}

func logLogin(username string, success bool) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Log dosyasına erişim hatası:", err)
		return
	}
	defer file.Close()

	logStatus := "Başarılı"
	if !success {
		logStatus = "Başarısız"
	}

	logEntry := fmt.Sprintf("Kullanıcı Adı: %s\nGiriş Tarihi ve Saati: %s\nGiriş Durumu: %s\n\n", username, time.Now().Format("2006-01-02 15:04:05"), logStatus)

	_, err = file.WriteString(logEntry)
	if err != nil {
		fmt.Println("Log dosyasına yazma hatası:", err)
	}
}

func adminLogin() {
	fmt.Println("Admin Girişi")
	var username, password string
	attempts := 5

	for attempts > 0 {
		fmt.Print("Kullanıcı Adı: ")
		fmt.Scanln(&username)

		if username != "admin" {
			attempts--
			fmt.Printf("Geçersiz kullanıcı adı. Kalan hakkınız: %d\n", attempts)
			logLogin(username, false)
			continue
		}

		fmt.Print("Şifre: ")
		fmt.Scanln(&password)

		if password != "admin" {
			attempts--
			fmt.Printf("Hatalı şifre. Kalan hakkınız: %d\n", attempts)
			logLogin(username, false)
			continue
		}

		fmt.Println("..........................Giriş başarılı...............................")
		logLogin(username, true)
		adminMenu()
		return
	}

	fmt.Println("Giriş hakkınız tükendi.")
	return
}

func adminMenu() {
	fmt.Println("Admin menüsü : ")
	for {
		fmt.Println("Lütfen bir işlem seçin:")
		fmt.Println("0 - Logları Görüntüle")
		fmt.Println("1 - Çıkış Yap")

		var choice int
		fmt.Print("Seçenek: ")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			viewLogs()
		case 1:
			fmt.Println("Çıkış yapılıyor.")
			return
		default:
			fmt.Println("Geçersiz seçenek!")
		}
	}
}

func studentLogin() {
	fmt.Println("Öğrenci Girişi")
	var username, password string
	attempts := 5

	for attempts > 0 {
		fmt.Print("Kullanıcı Adı: ")
		fmt.Scanln(&username)

		if username != "root" {
			attempts--
			fmt.Printf("Geçersiz kullanıcı adı. Kalan hakkınız: %d\n", attempts)
			logLogin(username, false)
			continue
		}

		fmt.Print("Şifre: ")
		fmt.Scanln(&password)

		if password != "root" {
			attempts--
			fmt.Printf("Hatalı şifre. Kalan hakkınız: %d\n", attempts)
			logLogin(username, false)
			continue
		}

		fmt.Println("..................................Giriş başarılı..............................")
		logLogin(username, true)

		return
	}

	fmt.Println("Giriş hakkınız tükendi.")
	return
}

func viewLogs() {
	file, err := os.Open("logs.txt")
	if err != nil {
		fmt.Println("Log dosyasına erişim hatası:", err)
		return
	}
	defer file.Close()

	fmt.Println("Log Kayıtları:")
	data := make([]byte, 1024)
	for {
		n, err := file.Read(data)
		if err != nil {
			break
		}
		fmt.Print(string(data[:n]))
	}
}
