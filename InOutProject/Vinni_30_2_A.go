/* 30.	В базе данных лесного колледжа содержатся сведения об успеваемости студентов.
Структура входных файлов
Students.txt (ИД_Студента Группа Студент)
1;Г1;Тигра
2;Г2;Винни_Пух
3;Г1;Кролик
…
Session.txt (ИД_Студента Дисциплина Оценки)
1;Физика;3 4 5
2;Пчеловодство;5 5 5
2;Русский язык;3 3
3;Русский язык;5 5 5 4
1;Химия;3 3 5

Сформировать список студентов с любимыми дисциплинами, исходя из среднего
балла (гарантируется, что для каждого студента только одна любимая
дисциплина). Список упорядочить по названию группы и по фамилии студента.
Структура выходного файла out.txt
Г1 Кролик Русский язык
Г1 Тигра Химия
Г2 Винни_Пух Пчеловодство

Решение. Носонова У.А.
Начало решения 08.04.2024 17:30 - 20:00
Окончание решения 09.04.2024 15:30 - 16:30
Время 03:30
*/



package main

import (
	"bufio"    // Буферизированный ввод-вывод
	"fmt"      // Форматированный ввод-вывод
	"os"       // Операции с файлами и системой
	"sort"     // Сортировка
	"strconv"  // Конвертация строк в числа
	"strings"  // Работа со строками
)

// Student - структура, представляющая информацию о студенте
type Student struct {
	ID       int    
	Group    string 
	LastName string 
	Name     string 
}

// Session - структура, представляющая информацию о сессии студента
type Session struct {
	StudentID  int    
	Subject    string 
	Grades     []int  
}


func main() {
	// Чтение информации о студентах из файла Students.txt
	// Вызов функции readStudents
	students, err := readStudents("Students.txt")
	if err != nil {
		return
	}

	// Чтение информации о сессиях студентов из файла Session.txt
	// Вызов функции readSessions
	sessions, err := readSessions("Session.txt")
	if err != nil {
		return
	}

    // пустой слайс строк favoriteSubjects, для хранения информации о студентах и дисциплинах
	var favoriteSubjects []string

	// Поиск любимых дисциплин для каждого студента
	for _, student := range students {
		var maxAvgGrade float64 = -1
		var favoriteSubject string

		for _, session := range sessions {
			if session.StudentID == student.ID {
				avgGrade := average(session.Grades)
				if avgGrade > maxAvgGrade {
					maxAvgGrade = avgGrade
					favoriteSubject = session.Subject
				}
			}
		}

		// Добавление информации о студенте и его любимой дисциплине в список
		if favoriteSubject != "" {
			favoriteSubjects = append(favoriteSubjects, fmt.Sprintf("%s %s %s", student.Group, student.LastName, favoriteSubject))
		}
	}

	// Сортировка списка любимых дисциплин по группе и фамилии студента
	sort.Strings(favoriteSubjects)

	// Запись результата в файл out.txt
	outFile, err := os.Create("out.txt")
	if err != nil {
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for _, line := range favoriteSubjects {
		fmt.Fprintf(writer, "%s\n", line)
	}
	writer.Flush()

	fmt.Println("Результат записан в файл out.txt")
}

// readStudents читает информацию о студентах из файла и возвращает список студентов
func readStudents(filename string) ([]Student, error) {
	// Open - ткрытие файла для чтения
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Закрытие файла после завершения функции

	var students []Student
	scanner := bufio.NewScanner(file) // Инициализация сканера для чтения файла построчно
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";") // Разделение строки на части по разделителю ";"
		if len(parts) != 3 {
			continue // Пропуск строки, если не хватает частей
		}
		id, err := strconv.Atoi(parts[0]) // Atoi - преобразование строки в число для идентификатора студента
		if err != nil {
			return nil, err
		}
		student := Student{
			ID:       id,
			Group:    parts[1],
			LastName: parts[2],
			Name:     parts[2],
		}
		students = append(students, student) // Добавление студента в список студентов
	}

	return students, scanner.Err() // Возврат списка студентов и возможной ошибки сканера
}

// readSessions читает информацию о сессиях студентов из файла и возвращает список сессий
func readSessions(filename string) ([]Session, error) {
	// Открытие файла для чтения
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Закрытие файла после завершения функции

	var sessions []Session
	scanner := bufio.NewScanner(file) // Инициализация сканера для чтения файла построчно
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";") // Разделение строки на части по разделителю ";"
		if len(parts) < 3 {
			continue // Пропуск строки, если не хватает частей
		}
		id, err := strconv.Atoi(parts[0]) // Преобразование строки в число для идентификатора студента
		if err != nil {
			return nil, err
		}
		gradesStr := strings.Split(parts[2], " ") // Разделение строки с оценками на части по пробелу
		var grades []int
		for _, gradeStr := range gradesStr {
			grade, err := strconv.Atoi(gradeStr) 
			if err != nil {
				return nil, err
			}
			grades = append(grades, grade) // Добавление оценки в список оценок
		}
		// Создание новой сессии студента на основе извлеченных данных
		session := Session{
			StudentID: id,         // Присвоение идентификатора студента
            Subject:   parts[1],   // Присвоение дисциплины сессии
            Grades:    grades,     // Присвоение списка оценок студента по данной сессии
        }
		sessions = append(sessions, session) // Добавление сессии в список сессий
	}

	return sessions, scanner.Err() // Возврат списка сессий и возможной ошибки сканера
}


// average вычисляет средний балл на основе списка оценок
func average(grades []int) float64 {
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}
