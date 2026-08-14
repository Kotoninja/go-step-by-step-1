package main

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	return Report{User: user, ReportID: 1, Date: reportDate}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	reportes := []Report{}

	for _, user := range users {
		reportes = append(reportes, CreateReport(user, reportDate))
	}

	return reportes
}

