package techUI

import (
	"fmt"
	"time"

	"github.com/nkarakotova/lim-console/registry"
	"github.com/nkarakotova/lim-core/errors/menuErrors"
	"github.com/nkarakotova/lim-core/models"
)

func printTrainings(a *registry.AppServiceFields) error {

	var s_year, s_day, e_year, e_day int
	var s_month, e_month time.Month

	fmt.Printf("Введите начальную дату (в формате YYYY.MM.dd): ")
	_, err := fmt.Scanf("%d.%d.%d", &s_year, &s_month, &s_day)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Введите конечную дату (в формате YYYY.MM.dd): ")
	_, err = fmt.Scanf("%d.%d.%d", &e_year, &e_month, &e_day)
	if err != nil {
		fmt.Println(err)
		return err
	}

	trainings, err := a.TrainingService.GetAllBetweenDateTime(time.Date(s_year, s_month, s_day, 0, 0, 0, 0, time.UTC), time.Date(e_year, e_month, e_day, 0, 0, 0, 0, time.UTC))
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(trainings) == 0 {
		fmt.Println("На данные даты ещё не поставлены тренеровки!")
		return menuErrors.ErrorMenu
	}

	fmt.Printf("Тренировки в данные даты:")
	for _, t := range trainings {
		fmt.Printf("\nid: %d\nНазвание: %s\nДата и время: %s\n\n", t.ID, t.Name, t.DateTime)
	}

	return nil
}

func printTrainingsOnWeek(a *registry.AppServiceFields) error {
	trainings, err := a.TrainingService.GetAllBetweenDateTime(time.Now(), time.Now().AddDate(0, 0, 7))
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(trainings) == 0 {
		fmt.Println("На неделю ещё не поставлены тренеровки!")
		return menuErrors.ErrorMenu
	}

	fmt.Printf("Тренировки на неделе:")
	for _, t := range trainings {
		fmt.Printf("\nid: %d\nНазвание: %s\nДата и время: %s\n\n", t.ID, t.Name, t.DateTime)
	}

	return nil
}

func printAllDirections(a *registry.AppServiceFields) error {
	directions, err := a.DirectionService.GetAll()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(directions) == 0 {
		fmt.Println("Направления ещё не добавлены!")
		return menuErrors.ErrorMenu
	}

	var gender string

	fmt.Printf("Все направления: ")
	for _, d := range directions {
		switch d.AcceptableGender {
		case models.Unknown:
			gender = "любой"
		case models.Male:
			gender = "мужской"
		case models.Female:
			gender = "женский"
		default:
			fmt.Println("Некорректно заданный пол!")
			return menuErrors.ErrorMenu
		}

		fmt.Printf("\nНазвание: %s\nОписание: %s\nДопустимый пол: %s\n\n", d.Name, d.Description, gender)
	}
	return nil
}

func getDirection(a *registry.AppServiceFields) (*models.Direction, error) {
	printAllDirections(a)
	var directionName string
	fmt.Printf("Введите название направления: ")
	_, err := fmt.Scanf("%s", &directionName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	direction, err := a.DirectionService.GetByName(directionName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return direction, nil
}

func printCoaches(a *registry.AppServiceFields) error {
	direction, err := getDirection(a)
	if err != nil {
		return err
	}

	coaches, err := a.CoachService.GetAllByDirection(direction.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(coaches) == 0 {
		fmt.Println("По данному направлению нет тренеров!")
		return menuErrors.ErrorMenu
	}

	fmt.Printf("Тренера по данному направлению: ")
	for _, c := range coaches {
		fmt.Printf("\nИмя: %s\nОписание: %s\n\n", c.Name, c.Description)
	}

	return nil
}

func printCoachesByDirection(a *registry.AppServiceFields, direction *models.Direction) error {
	coaches, err := a.CoachService.GetAllByDirection(direction.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(coaches) == 0 {
		fmt.Println("По данному направлению нет тренеров!")
		return menuErrors.ErrorMenu
	}

	fmt.Printf("Тренера по данному направлению: ")
	for _, c := range coaches {
		fmt.Printf("\nИмя: %s\nОписание: %s\n\n", c.Name, c.Description)
	}

	return nil
}
