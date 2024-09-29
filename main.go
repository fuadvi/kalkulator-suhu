package main

import "fmt"

type Celcius struct {
	Suhu float64
}

type Fahrenheit struct {
	Suhu float64
}

type Kelvin struct {
	Suhu float64
}

func (c Celcius) toCelcius() float64 {
	return c.Suhu
}

func (c Celcius) toFahrenheit() float64 {
	return ((9.0 / 5.0) * c.Suhu) + 32
}

func (c Celcius) toKelvin() float64 {
	return c.Suhu + 273.15
}

func (f Fahrenheit) toFahrenheit() float64 {
	return f.Suhu
}

func (f Fahrenheit) toCelcius() float64 {
	return (f.Suhu - 32) * (5.0 / 9.0)
}

func (f Fahrenheit) toKelvin() float64 {
	return (f.Suhu + 459.67) * (5.0 / 9.0)
}

func (k Kelvin) toCelcius() float64 {
	return k.Suhu - 273.15
}

func (k Kelvin) toFahrenheit() float64 {
	return (k.Suhu * (9.0 / 5.0)) - 459.67
}

func (k Kelvin) toKelvin() float64 {
	return k.Suhu
}

type hitungSuhu interface {
	toCelcius() float64
	toKelvin() float64
	toFahrenheit() float64
}

func main() {

	fmt.Println("pilih suhu awal")
	fmt.Println("1. Calcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("Masukan suhu awal yang di inginkan:")
	var suhuAwal int
	fmt.Scanf("%d", &suhuAwal)

	for suhuAwal < 1 || suhuAwal > 6 {
		fmt.Println("Suhu Awal tidak valid, masukan suhu awal yang diinginkan")
		fmt.Scanf("%d", &suhuAwal)
	}

	fmt.Println("pilih suhu akhir")
	fmt.Println("1. Calcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("Masukan suhu akhir yang di inginkan:")

	var suhuAkhir int
	fmt.Scanf("%d", &suhuAkhir)

	for suhuAkhir < 1 || suhuAkhir > 6 {
		fmt.Println("Suhu akhir tidak valid, masukan suhu akhir yang diinginkan")
		fmt.Scanf("%d", &suhuAkhir)
	}

	var suhu float64
	fmt.Println("masukan nilai suhu: ")
	fmt.Scanf("%f", &suhu)

	var interfaceSuhu hitungSuhu

	switch suhuAwal {
	case 1:
		interfaceSuhu = Celcius{suhu}
	case 2:
		interfaceSuhu = Fahrenheit{suhu}
	case 3:
		interfaceSuhu = Kelvin{suhu}
	}

	var suhuFinal float64
	switch suhuAkhir {
	case 1:
		suhuFinal = interfaceSuhu.toCelcius()
	case 2:
		suhuFinal = interfaceSuhu.toFahrenheit()
	case 3:
		suhuFinal = interfaceSuhu.toKelvin()
	}

	fmt.Printf("suhu akhir yang di dapat adalah %.2f", suhuFinal)
}
