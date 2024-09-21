package techUI

import (
	"fmt"
	"time"

	"github.com/nkarakotova/lim-console/registry"
	"github.com/nkarakotova/lim-core/errors/menuErrors"
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

func printCoaches(a *registry.AppServiceFields) error {
	coaches, err := a.CoachService.GetAll()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(coaches) == 0 {
		fmt.Println("Тренеров нет!")
		return menuErrors.ErrorMenu
	}

	fmt.Printf("Тренера: ")
	for _, c := range coaches {
		fmt.Printf("\nИмя: %s\n\n", c.Name)
	}

	return nil
}
