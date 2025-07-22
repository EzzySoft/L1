package main

import "fmt"

// Интерфейс для телефона с USB-C разъёмом
type USBCPhone interface {
	ChargeWithUSBC()
}

// Интерфейс для телефона с microUSB разъёмом
type MicroUSBPhone interface {
	ChargeWithMicroUSB()
}

// Новый телефон (например, Google Pixel) с USB-C
type GooglePixel struct{}

func (p *GooglePixel) ChargeWithUSBC() {
	fmt.Println("Google Pixel заряжается через USB-C")
}

// Старый телефон (например, Samsung) с microUSB
type OldSamsung struct{}

func (s *OldSamsung) ChargeWithMicroUSB() {
	fmt.Println("Samsung заряжается через microUSB")
}

// Зарядка USB-C (современная)
type USBCCharger struct{}

func (c *USBCCharger) ChargeWithUSBC() {
	fmt.Println("Заряжаю через USB-C")
}

// Зарядка microUSB (старая)
type MicroUSBCharger struct{}

func (c *MicroUSBCharger) ChargeWithMicroUSB() {
	fmt.Println("Заряжаю через microUSB")
}

// Адаптер: USB-C -> microUSB
type USBCToMicroUSBAdapter struct {
	Charger *USBCCharger
}

func (a *USBCToMicroUSBAdapter) ChargeWithMicroUSB() {
	fmt.Println("Адаптер преобразует USB-C в microUSB...")
	a.Charger.ChargeWithUSBC()
}

// Адаптер: microUSB -> USB-C
type MicroUSBToUSBCAdapter struct {
	Charger *MicroUSBCharger
}

func (a *MicroUSBToUSBCAdapter) ChargeWithUSBC() {
	fmt.Println("Адаптер преобразует microUSB в USB-C...")
	a.Charger.ChargeWithMicroUSB()
}

func main() {
	// Новый телефон — Google Pixel
	pixel := &GooglePixel{}
	// Старый телефон — Samsung
	samsung := &OldSamsung{}
	// Новая зарядка USB-C
	usbcCharger := &USBCCharger{}
	// Старая зарядка microUSB
	microUSBCharger := &MicroUSBCharger{}

	// Заряжаем Pixel через родную зарядку
	fmt.Println("Google Pixel, родная зарядка:")
	pixel.ChargeWithUSBC()

	// Заряжаем Samsung через родную зарядку
	fmt.Println("\nSamsung, родная зарядка:")
	samsung.ChargeWithMicroUSB()

	// Заряжаем Samsung через USB-C зарядку с адаптером
	fmt.Println("\nSamsung с USB-C зарядкой через адаптер:")
	adapterUSBC := &USBCToMicroUSBAdapter{Charger: usbcCharger}
	adapterUSBC.ChargeWithMicroUSB()
	samsung.ChargeWithMicroUSB() // для наглядности (заряжается)

	// Заряжаем Pixel через старую microUSB зарядку с адаптером
	fmt.Println("\nGoogle Pixel с microUSB зарядкой через адаптер:")
	adapterMicroUSB := &MicroUSBToUSBCAdapter{Charger: microUSBCharger}
	adapterMicroUSB.ChargeWithUSBC()
	pixel.ChargeWithUSBC() // для наглядности (заряжается)
}
