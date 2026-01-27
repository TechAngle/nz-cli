package main

import (
	"flag"
	"fmt"
	"nz-cli/internal/client"
	"nz-cli/internal/commons"
	"nz-cli/internal/visuals"
	"os"
)

const (
	/*
		fr, idk which i should use here, but i guess it should be negative.
		If it positive, then such ID will be invalid, so let it be -100 at least.

		P.S. I'll hate nz devs (i am already are, whatever) if they would add negatives as IDs.
	*/
	INVALID_ID = -100
)

func main() {
	/*
		I tried to separate different flags (like additional and main) with different styles here.
		But I'll need to create more 'distinctive' color palette (or find another one).
	*/

	// login flags
	login := flag.Bool("login", false, visuals.FourthStyleBold.Render("Login to the system. Should be set with --username and --password flags"))
	username := flag.String("username", "", visuals.FourthStyleBold.Render("Username. Required if -login argument is set"))
	password := flag.String("password", "", visuals.FourthStyleBold.Render("Password. Required if -login argument is set"))

	// additional flags
	startDate := flag.String("start-date", commons.TODAY_DATE, visuals.FourthStyleBold.Render(commons.StartDateArgUsage))
	endDate := flag.String("end-date", commons.TODAY_DATE, visuals.FourthStyleBold.Render(commons.EndDateArgUsage))
	subjectId := flag.Int("subject-id", INVALID_ID, visuals.FourthStyleBold.Render(commons.SubjectIdArgUsage))

	// client flags
	diary := flag.Bool("diary", false, visuals.MainStyleBold.Render("Show diary"))
	grades := flag.Bool("grades", false, visuals.MainStyleBold.Render("Show grades"))
	perfomance := flag.Bool("perf", false, visuals.MainStyleBold.Render("Show perfomance"))

	flag.Parse()

	// Initializating client
	client, err := client.NewClient()
	if err != nil {
		// TODO: Replace it with something safer
		panic(err)
	}

	if *login {
		err = client.Login(*username, *password)
		if err != nil {
			panic(err)
		}
		return // hehe, i dont wanna go further, WRITE YOUR NEXT COMMAND AFTER AUTH :3
	}

	// checking for dates one more time ;)
	// just to be sure if someone stupid would set empty dates
	if *startDate == "" {
		fmt.Println(visuals.ErrorStyle.Render("Start Date is invalid!"))
		os.Exit(1)
	}
	if *endDate == "" {
		fmt.Println(visuals.ErrorStyle.Render("End Date is invalid!"))
		os.Exit(1)
	}

	// TODO: Add other api functions and flags here
	switch true {
	// -diary flag
	case *diary:
		err = client.Diary(*startDate, *endDate)

	// -grades flag
	case *grades: // TODO: Test it when i'll get any grades
		if *subjectId == INVALID_ID {
			err = fmt.Errorf("invalid subject id: %v", err)
			break
		}

		err = client.Grades(*startDate, *endDate, *subjectId)

	// -perfomance flag
	case *perfomance:
		err = client.Perfomance(*startDate, *endDate)
	}

	// if any error occurred we'll just close program with status code 1, why not.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
