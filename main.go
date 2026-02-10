package main

//Simple to-do list
import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type task struct {
	name      string
	message   string
	time      time.Time
	id        int
	completed bool
}

func help() {
	fmt.Println("Add - эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("List - эта команда позволяет получить полный список всех задач")
	fmt.Println("Del - эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("Done - эта команда позволяет отменить задачу как выполненную")
	fmt.Println("Exit - эта команда позволяет завершить выполнение программы")
}
func add(NewTask task, m map[int]task) int {
	m[NewTask.id] = NewTask
	return NewTask.id
}
func list(m map[int]task) {
	if len(m) == 0 {
		fmt.Println("У вас нет текущих задач.")
	} else {
		for number, i := range m {
			fmt.Println("Задача с номером", number, ":", i.name, i.message, i.time.Format("2006-01-02 15:04:05"), i.completed)
		}
	}

}
func del(name string, m map[int]task) {
	for id, task := range m {
		if task.name == name {
			delete(m, id)
			fmt.Println("Удалили задачу(дело) с названием: ", name)
			return
		}
	}
}
func done(name string, m map[int]task) {
	for id, task := range m {
		if task.name == name {
			task.completed = true
			m[id] = task
			return
		}
	}
}
func exit() {
}
func printTask(t task) {
	fmt.Printf("Task: %s\nMessage: %s\nTime: %s\nDone: %t", t.name, t.message, t.time.Format("2006-01-02 15:04:05"), t.completed)
	fmt.Println()
}
func main() {
	m := make(map[int]task)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите команду: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Вы ничего не написали.")
			return
		}
		text := scanner.Text()
		if len(text) == 0 {
			fmt.Println("Вы ничего не написали.")
		}
		if text == "Help" {
			help()
		} else if text == "Exit" {
			fmt.Println("Вы завершили использование программы.")
			break
		} else if text == "List" {
			list(m)
		} else if text == "Add" {
			fmt.Print("Введите название задачи: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Print("Введите описание задачи: ")
			scanner.Scan()
			message := scanner.Text()
			newTask := task{
				name:      name,
				message:   message,
				time:      time.Now(),
				id:        len(m) + 1,
				completed: false,
			}
			add(newTask, m)
			fmt.Println("Задача успешно добавлена!")
		} else if text == "Done" {
			fmt.Print("Введите название задачи: ")
			scanner.Scan()
			name := scanner.Text()
			done(name, m)
			fmt.Println("Задача успешно добавлена!")
		} else {
			fmt.Println("Неизвестная команда. Воспользуйтесь командой Help.")
		}
	}
}
