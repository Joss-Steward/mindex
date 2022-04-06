package analyzer

import "time"

type Media struct {
	title    string
	release  time.Time
	director string
	genres   []string
}
