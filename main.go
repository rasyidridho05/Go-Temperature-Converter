package main

import "fmt"

func main() {
	var option, celcius, fahrenheit, reamur, kelvin float32

	fmt.Println("Option List : \n 1. Input Celcius \n 2. Input Fahrenheit \n 3. Input Reamur \n 4. Input Kelvin")
	fmt.Print("Input your Option :")
	fmt.Scanln(&option)

	if option == 1 {
		fmt.Printf("Input Temperature in Celcius :")
		fmt.Scanln(&celcius)

		fahrenheit = (9.0 / 5.0 * celcius) + 32
		reamur = 4.0 / 5.0 * celcius
		kelvin = celcius + 273.15

		fmt.Println("===================================")
		fmt.Printf("\nCelcius    : %.2f°C", celcius)
		fmt.Printf("\nFahrenheit : %.2f°F", fahrenheit)
		fmt.Printf("\nReamur     : %.2f°R", reamur)
		fmt.Printf("\nKelvin     : %.2f°K", kelvin)
	} else if option == 2 {
		fmt.Printf("Input Temperature in Fahrenheit :")
		fmt.Scanln(&fahrenheit)

		celcius = (5.0 / 9.0 * fahrenheit) - 32
		reamur = (4.0 / 9.0 * fahrenheit) - 32
		kelvin = celcius + 273.15

		fmt.Println("===================================")
		fmt.Printf("\nCelcius    : %.2f°C", celcius)
		fmt.Printf("\nFahrenheit : %.2f°F", fahrenheit)
		fmt.Printf("\nReamur     : %.2f°R", reamur)
		fmt.Printf("\nKelvin     : %.2f°K", kelvin)
	} else if option == 3 {
		fmt.Printf("Input Temperature in Reamur :")
		fmt.Scanln(&reamur)

		celcius = 5.0 / 4.0 * reamur
		fahrenheit = (9.0 / 4.0 * reamur) + 32
		kelvin = celcius + 273.15

		fmt.Println("===================================")
		fmt.Printf("\nCelcius    : %.2f°C", celcius)
		fmt.Printf("\nFahrenheit : %.2f°F", fahrenheit)
		fmt.Printf("\nReamur     : %.2f°R", reamur)
		fmt.Printf("\nKelvin     : %.2f°K", kelvin)
	} else if option == 4 {
		fmt.Printf("Input Temperature in Celcius :")
		fmt.Scanln(&kelvin)

		celcius = kelvin - 273.15
		fahrenheit = (9.0 / 5.0 * celcius) + 32
		reamur = (4.0 / 5.0 * celcius)
		
		fmt.Println("===================================")
		fmt.Printf("\nCelcius    : %.2f°C", celcius)
		fmt.Printf("\nFahrenheit : %.2f°F", fahrenheit)
		fmt.Printf("\nReamur     : %.2f°R", reamur)
		fmt.Printf("\nKelvin     : %.2f°K", kelvin)
	} else {
		fmt.Print("Input doesn't match with the Options")
	}

}
