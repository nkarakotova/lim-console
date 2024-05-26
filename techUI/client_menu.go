package techUI

import (
	"fmt"

	"github.com/nkarakotova/lim-console/registry"
	"github.com/nkarakotova/lim-core/errors/menuErrors"
	"github.com/nkarakotova/lim-core/models"
)

func loginClient(a *registry.AppServiceFields) (*models.Client, error) {
	var telephone string
	fmt.Printf("Введите телефон: ")
	_, err := fmt.Scanf("%s", &telephone)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var password string
	fmt.Printf("Введите пароль: ")
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client, err := a.ClientService.Login(telephone, password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return client, nil
}

func createClient(a *registry.AppServiceFields) (*models.Client, error) {
	var name string
	fmt.Printf("Введите имя: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var telephone string
	fmt.Printf("Введите телефон: ")
	_, err = fmt.Scanf("%s", &telephone)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var password string
	fmt.Printf("Введите пароль: ")
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var mail string
	fmt.Printf("Введите почту: ")
	_, err = fmt.Scanf("%s", &mail)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var age uint16
	fmt.Printf("Введите возраст: ")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gender models.Gender
	fmt.Printf("Введите пол (0 - не указано, 1 - мужской, 2 - женский): ")
	_, err = fmt.Scanf("%d", &gender)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := &models.Client{Name: name,
		Telephone: telephone,
		Mail:      mail,
		Password:  password,
		Age:       age,
		Gender:    gender}

	err = a.ClientService.Create(client)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return client, nil
}

func printInfoClient(client *models.Client) {
	var gender string

	switch client.Gender {
	case models.Unknown:
		gender = "не указан"
	case models.Male:
		gender = "мужской"
	case models.Female:
		gender = "женский"
	default:
		fmt.Println("Некорректно заданный пол!")
	}
	fmt.Printf("\nИмя: %s\nТелефон: %s\nПочта: %s\nВозраст: %d\nПол: %s\n\n", client.Name, client.Telephone, client.Mail, client.Age, gender)
}

func printSubscription(a *registry.AppServiceFields, client *models.Client) {
	if client.SubscriptionID == 0 {
		fmt.Println("Нет абонемента!")
		return
	}

	subscription, err := a.SubscriptionService.GetByID(client.SubscriptionID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nДата начала действия: %s\nДата конца действия: %s\nОставшееся количество тренировок: %d\n\n", subscription.StartDate, subscription.EndDate, subscription.RemainingTrainingsNum)
}

func createAssigenment(a *registry.AppServiceFields, client *models.Client) {
	err := printTrainingsOnWeek(a)
	if err != nil {
		return
	}

	var tID uint64
	fmt.Printf("Выберете тренировку и введите её id: ")
	_, err = fmt.Scanf("%d", &tID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = a.ClientService.СreateAssignment(client.ID, tID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешно создана!")
}

func printAssigenment(a *registry.AppServiceFields, client *models.Client) error {
	trainings, err := a.TrainingService.GetAllByClient(client.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(trainings) == 0 {
		fmt.Println("Ещё нет записей на тренировки!")
		return menuErrors.ErrorMenu
	}

	fmt.Printf("Тренировки, на которые есть запись:")
	for _, t := range trainings {
		fmt.Printf("\nid: %d\nНазвание: %s\nДата и время: %s\n\n", t.ID, t.Name, t.DateTime)
	}

	return nil
}

func deleteAssigenment(a *registry.AppServiceFields, client *models.Client) {
	err := printAssigenment(a, client)
	if err != nil {
		return
	}

	var tID uint64
	fmt.Printf("Выберете тренировку и введите её id: ")
	_, err = fmt.Scanf("%d", &tID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = a.ClientService.DeleteAssignment(client.ID, tID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешно удалена!")
}

const client_loop_string = `Меню клиента: 
	0 -- выйти
	1 -- посмотреть информацию о себе
	2 -- посмотреть расписание на неделю
	3 -- посмотреть расписание на выбраный промежуток времени 
	4 -- посмотреть направления
	5 -- посмотреть тренеров по направлению
	6 -- посмотреть информацию по абонементу
	7 -- записаться на тренировку
	8 -- отменить запись на тренировку
	9 -- посмотреть тренировки, на которые есть запись на неделю
	10 -- посмотреть тренировки, на которые есть запись на выбранный промежуток времени
Выберите действие: `

func clientMenu(a *registry.AppServiceFields, client *models.Client) {
	var num int = 1

	for num != 0 {
		fmt.Printf("\n\n%s", client_loop_string)

		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Printf("\nПункт меню введён некорректно!\n\n")
			continue
		}

		if num == 0 {
			return
		}

		switch num {
		case 0:
			return
		case 1:
			printInfoClient(client)
		case 2:
			printTrainingsOnWeek(a)
		case 3:
			printTrainings(a)
		case 4:
			printAllDirections(a)
		case 5:
			printCoaches(a)
		case 6:
			printSubscription(a, client)
		case 7:
			createAssigenment(a, client)
		case 8:
			deleteAssigenment(a, client)
		case 9:
			printAssigenment(a, client)
		default:
			fmt.Printf("\nНеверный пункт меню!\n\n")
		}
	}
}
