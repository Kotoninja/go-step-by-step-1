package main

import (
	"errors"
	"fmt"
	"sort"
)

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Company struct {
	workers []worker
}

type worker struct {
	name       string
	position   string
	salary     uint
	experience uint
}

func getPositionRank(position string) int {
	switch position {
	case "директор":
		return 1
	case "зам. директора":
		return 2
	case "начальник цеха":
		return 3
	case "мастер":
		return 4
	case "рабочий":
		return 5
	default:
		return 6
	}
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if salary == 0 {
		return errors.New("salary must be greater than 0")
	}

	c.workers = append(c.workers, worker{name: name, position: position, salary: salary, experience: experience})
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	sort.Slice(c.workers, func(i, j int) bool {
		incomeI := c.workers[i].salary
		incomeJ := c.workers[j].salary

		if incomeI != incomeJ {
			return incomeI > incomeJ
		}

		return getPositionRank(c.workers[i].position) < getPositionRank(c.workers[j].position)
	})

	result := make([]string, len(c.workers))
	for i, w := range c.workers {
		income := w.salary
		result[i] = fmt.Sprintf("%s — %d — %s", w.name, income, w.position)
	}
	return result, nil
}
