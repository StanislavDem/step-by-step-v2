package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal // встроенная структура Personal из пакета personaldata
}

// Метод Parse парсит строку формата "678,0h50m" и записывает данные в поля DaySteps
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
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
	ds.Steps = steps

	// длительность
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("invalid duration value: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration must be greater than 0")
	}
	ds.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает строку с данными о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	// дистанция
	distKm := spentenergy.Distance(ds.Steps, ds.Height)

	// калории (для ходьбы)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// формируем строку
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distKm,
		calories,
	)

	return result, nil
}