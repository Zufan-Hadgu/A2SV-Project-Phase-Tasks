package main

import(
	"fmt"
)
type grades struct{
	subject string
	grade int
}

type gradeAverage struct{
	element []grades
}

func (gv gradeAverage) average()float64{
	total := 0
	for _, g := range gv.element{
		total += g.grade
	}
	return float64(total)/float64(len(gv.element))

}

func main(){
	var name string
	var Courses int
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter the number of courses you take: ")
	fmt.Scan(&Courses)

	AllGrades := make([]grades,Courses)

	for i := 0; i < Courses; i++{
		var g grades
		fmt.Printf("Enter subject #%d",i+1)
		fmt.Scan(&g.subject)	

		for {
			fmt.Printf("Enter Grade for the %s subject",g.subject)
			fmt.Scan(&g.grade)

			if g.grade >= 0 && g.grade <=100{
				break	
			}
			fmt.Printf("Enter valid grade for between 0 and 100 for  %s subject",g.subject)		
		}
		
		AllGrades = append(AllGrades, g)
	}

	fmt.Printf("\n%s's Grade Report:\n", name)
	for _, g := range AllGrades {
		fmt.Printf("Subject: %s, Grade: %d\n", g.subject, g.grade)
	}

	gv:= gradeAverage{element: AllGrades}
	fmt.Printf("The average of you grade is %.2f",gv.average())

}