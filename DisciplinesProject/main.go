package main

import (
	"encoding/csv"  // Пакет для работы с CSV файлами
	"fmt"           // Пакет для форматированного вывода
	"log"           // Пакет для логирования ошибок
	"os"            // Пакет для работы с операционной системой, включая файловую систему
	"path/filepath" // Пакет для работы с путями к файлам и директориям
	"strings"       // Пакет для работы со строками
)

func main() {
	// Открываем файл Disciplines.csv
	file, err := os.Open("Disciplines.csv")
	if err != nil {
		log.Fatalf("Не удалось открыть файл: %v", err)
	}
	defer file.Close()

	// Определяем путь к общему каталогу
	commonDir := "CatalogDiscipline" // Имя общего каталога
	err = os.MkdirAll(commonDir, 0755)
	if err != nil {
		log.Fatalf("Не удалось создать общий каталог %s: %v", commonDir, err)
	}

	// Создаем ридер для CSV файла с явным указанием разделителя ";"
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Читаем CSV файл
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Ошибка чтения CSV файла: %v", err)
	}

	// Перебираем записи и создаем каталоги
	for _, record := range records {
		if len(record) != 2 {
			log.Fatalf("Неверный формат записи в файле: %v", record)
		}

		// Удаляем лишние пробелы из названия дисциплины
		disciplineName := strings.TrimSpace(record[1])

		// Создаем каталог для дисциплины в общем каталоге
		disciplineDir := filepath.Join(commonDir, disciplineName)
		err := os.Mkdir(disciplineDir, 0755)
		if err != nil {
			log.Fatalf("Не удалось создать каталог %s: %v", disciplineDir, err)
		}
		fmt.Printf("Каталог %s успешно создан\n", disciplineDir)
	}
}
