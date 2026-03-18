package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	// Проверка входных данных
	if steps <= 0 {
		return 0, errors.New("steps count must be greater than 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than 0")
	}

	// Рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты
	durationMinutes := duration.Minutes()

	// Расчёт калорий (weight * meanSpeed * durationInMinutes) / minInH
	calories := (weight * speed * durationMinutes) / float64(minInH)

	// Умножаем на корректирующий коэффициент walkingCaloriesCoefficient
	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	
	// Проверка входных данных
	if steps <= 0 {
		return 0, errors.New("steps count must be greater than 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than 0")
	}

	// Рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)
	
	// Переводим продолжительность в минуты
	durationMinutes := duration.Minutes()

	// Расчёт калорий (weight * meanSpeed * durationInMinutes) / minInH
	calories := (weight * speed * durationMinutes) / float64(minInH)

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	
	// Проверка
	if duration <= 0 {
		return 0
	}
	// Вычисляем дистанцию в км.
	distKm := Distance(steps, height)
	durationHours := duration.Hours()

	// Вычисляем среднюю скорость
	return distKm / durationHours
}

func Distance(steps int, height float64) float64 {
	
	// Рассчитываем длину шага опираясь на рост пользователя
	stepLen := height * stepLengthCoefficient

	// Умножаем количество шагов на длину шага в м.
	distanceMeters := float64(steps) * stepLen

	// Переводим м. в км.
	return distanceMeters / mInKm
}
