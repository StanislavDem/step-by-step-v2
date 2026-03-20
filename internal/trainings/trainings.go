package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal // встроенная структура Personal из пакета personaldata
}

// Метод Parse парсит строку формата "3456,Ходьба,3h00m" и и записывает данные в поля Training
func (t *Training) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("invalid data string format")
	}

	// шаги
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid steps value: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps count must be greater than 0")
	}
	t.Steps = steps

	// тип тренировки
	t.TrainingType = parts[1]

	// длительность
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("invalid duration value: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration must be greater than 0")
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает строку с данными о тренировке
func (t Training) ActionInfo() (string, error) {
	// дистанция из пакета spentenergy
	distKm := spentenergy.Distance(t.Steps, t.Height)

	// средняя скорость из пакета spentenergy
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	// калории
	var calories float64
	var err error

	switch strings.ToLower(t.TrainingType) {
	case "бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	// формируем строку
	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distKm,
		speed,
		calories,
	)

	return result, nil
}