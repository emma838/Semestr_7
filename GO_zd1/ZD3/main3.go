package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type car struct {
	mpg          float64 // spalanie (miles per gallon)
	cylinders    int     // liczba cylindrów
	displacement float64 // pojemność
	horsepower   float64 // moc
	weight       float64 // masa
	acceleration float64 // przyśpieszenie
	year         int     // rocznik
	origin       int     // pochodzenie
	name         string  // nazwa
}

//zd7-8:

// Funkcja porównująca dwa obiekty car
func compare(first *car, second *car) float64 {
	if first == second {
		return 0.0
	}
	difference := 1.0
	difference *= 1.0 - math.Abs(first.mpg-second.mpg)/40
	difference *= 1.0 - math.Abs(first.horsepower-second.horsepower)/300
	difference *= 1.0 - math.Abs(first.weight-second.weight)/5000
	difference *= 1.0 - math.Abs(first.acceleration-second.acceleration)/30
	return difference
}

// Metoda porównująca obiekt z innym obiektem car
func (this *car) compare(other *car) float64 {
	if this == other {
		return 0.0
	}
	difference := 1.0
	difference *= 1.0 - math.Abs(this.mpg-other.mpg)/40
	difference *= 1.0 - math.Abs(this.horsepower-other.horsepower)/300
	difference *= 1.0 - math.Abs(this.weight-other.weight)/5000
	difference *= 1.0 - math.Abs(this.acceleration-other.acceleration)/30
	return difference
}

// skanowanie
func loadCars() []*car {
	cars := []*car{}
	file, err := os.Open("cars.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		c := car{}
		c.mpg, _ = strconv.ParseFloat(line[0], 64)
		c.cylinders, _ = strconv.Atoi(line[1])
		c.displacement, _ = strconv.ParseFloat(line[2], 64)
		c.horsepower, _ = strconv.ParseFloat(line[3], 64)
		c.weight, _ = strconv.ParseFloat(line[4], 64)
		c.acceleration, _ = strconv.ParseFloat(line[5], 64)
		c.year, _ = strconv.Atoi(line[6])
		c.origin, _ = strconv.Atoi(line[7])
		c.name = line[8]
		cars = append(cars, &c)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cars
}

// Zadanie 11: Funkcja znajduje najbardziej podobny samochód do podanego target
func findMostSimilar(target *car, cars []*car) *car {
	var mostSimilar *car
	maxSimilarity := -1.0

	for _, candidate := range cars {
		if candidate == target {
			continue // Pomijamy porównanie samochodu z samym sobą
		}
		similarity := target.compare(candidate)
		if similarity > maxSimilarity {
			maxSimilarity = similarity
			mostSimilar = candidate
		}
	}

	return mostSimilar
}

// zd3:
func main() {
	car0 := car{18, 8, 307, 130, 3504, 12, 70, 1, "chevrolet malibu"}
	car1 := car{13, 8, 351, 158, 4363, 13, 73, 1, "ford ltd"}
	car2 := car{29, 4, 98, 83, 2219, 16.5, 74, 2, "audi fox"}
	car3 := car{20, 6, 232, 100, 2914, 16, 75, 1, "amc gremlin"}
	car4 := car{33, 4, 91, 53, 1795, 17.4, 76, 3, "honda civic"}
	car5 := car{23.2, 4, 156, 105, 2745, 16.7, 78, 1, "plymouth sapporo"}

	cars := []car{car0, car1, car2, car3, car4, car5}
	// fmt.Println(cars)

	//zd. 4: Zmiana w tablicy cars nie wpływa na obiekt car4 ponieważ cars jest tablicą wartości, a nie wskaźników, cars[4] jest kopią obiektu car4, a zmiana cars[4].name nie zmienia oryginalnego car4.
	cars[4].name = "test"
	fmt.Println("4. Po zmianie nazwy w tablicy wartości cars:")
	fmt.Println(cars[4])
	fmt.Println(car4)

	//zd 5: Nie trzeba tworzyć konstruktora do inicjalizacji struktury, ponieważ struktury mogą być inicjalizowane za pomocą literałów, jak pokazano powyżej.
	// dla bardziej złożonych struktur lub inicjalizacji z wartościami domyślnymi można napisać funkcję pomocniczą (konstruktor), która zwróci zainicjalizowany obiekt.

	//zd6:
	carsPointers := []*car{&car0, &car1, &car2, &car3, &car4, &car5}
	carsPointers[4].name = "test pointer"
	fmt.Println("6. Po zmianie nazwy w tablicy wskaźników carsPointers:")
	fmt.Println(*carsPointers[4])
	fmt.Println(car4)

	// //zd7-8: funkcja i metoda zwracają te same wyniki, pomimo różnej implementacji
	wynik1 := compare(&car0, &car1)
	wynik2 := car0.compare(&car1)

	fmt.Println("8. Wynik porównania przez funkcję:", wynik1)
	fmt.Println("8. Wynik porównania przez metodę:", wynik2)

	// Zadanie 9:
	target := carsPointers[2] // Trzeci samochód (car2) w tablicy wskaźników
	var mostSimilar *car
	maxSimilarity := -1.0

	for _, candidate := range carsPointers {
		if candidate == target {
			continue // Pomijamy porównanie samochodu z samym sobą
		}
		similarity := target.compare(candidate)
		if similarity > maxSimilarity {
			maxSimilarity = similarity
			mostSimilar = candidate
		}
	}

	fmt.Printf("9. Samochód docelowy: %+v\n", *target)
	fmt.Printf("9. Najbardziej podobny samochód: %+v\n", *mostSimilar)

	// Zadanie 10
	fileCars := loadCars()
	if len(fileCars) < 3 {
		log.Fatal("Plik cars.txt musi zawierać co najmniej 3 samochody.")
	}

	fileTarget := fileCars[2] // Trzeci samochód
	var fileMostSimilar *car
	fileMaxSimilarity := -1.0

	for _, candidate := range fileCars {
		if candidate == fileTarget {
			continue // Pomijamy porównanie samochodu z samym sobą
		}
		similarity := fileTarget.compare(candidate)
		if similarity > fileMaxSimilarity {
			fileMaxSimilarity = similarity
			fileMostSimilar = candidate
		}
	}

	fmt.Printf("10. Samochód docelowy z pliku: %+v\n", *fileTarget)
	fmt.Printf("10. Najbardziej podobny samochód z pliku: %+v\n", *fileMostSimilar)

	//zd11:
	fileTarget2 := fileCars[2] // Trzeci samochód
	fileMostSimilar2 := findMostSimilar(fileTarget2, fileCars)
	fmt.Printf("11. Samochód docelowy z pliku: %+v\n", *fileTarget2)
	fmt.Printf("11. Najbardziej podobny samochód z pliku: %+v\n", *fileMostSimilar2)

}
