package main

type actor struct {
	name      string
	photo     string
	jobTitle  string
	bio       string
	birthDate string
	TopMovies []movies
}

type movies struct {
	Title string
	Year  string
}

func crawl(month int, day int) {

}

func main() {
	crawl(*month, *day)
}
