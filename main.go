package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/shivasaxena/PMS7003"
)

func main() {

	deviceName := flag.String("device", "/dev/ttyAMA0", "Device Name")
	fullScreen := flag.Bool("fullScreen", false, "Boolean value to run application in full screen")
	flag.Parse()

	myApp := app.New()
	myWindow := myApp.NewWindow("Air Quality Monitor")
	device, e := PMS7003.Open(*deviceName, PMS7003.ActiveMode)
	if e != nil {
		panic(e)
	}
	defer device.Close()

	textValues := createUI(myWindow)

	go func() {

		for range time.Tick(time.Second) {

			sensorValue, err := device.Read()

			if err != nil {
				panic(err)
			}

			updateSensorValues(textValues, sensorValue)

		}
	}()
	myWindow.SetFullScreen(*fullScreen)
	myWindow.ShowAndRun()

}

func updateSensorValues(textValues []*canvas.Text, sensorValue PMS7003.PMS7003SensorValue) {

	updateTextValue(textValues[0], sensorValue.PM10Atmospheric)
	updateTextValue(textValues[1], sensorValue.PM25Atmospheric)
	updateTextValue(textValues[2], sensorValue.PM100Atmospheric)

}
func updateTextValue(textValue *canvas.Text, value uint16) {
	textValue.Text = fmt.Sprint(value) + " μ/m³"
	textValue.Color = getColorValue(value)
	textValue.Refresh()
}

func getColorValue(value uint16) color.NRGBA {

	if value >= 0 && value <= 50 {
		return color.NRGBA{R: 0, G: 228, B: 0, A: 255} // Green
	} else if value >= 51 && value <= 100 {
		return color.NRGBA{R: 255, G: 255, B: 0, A: 255} // Yellow
	} else if value >= 101 && value <= 150 {
		return color.NRGBA{R: 255, G: 126, B: 0, A: 255} // Orange
	} else if value >= 151 && value <= 200 {
		return color.NRGBA{R: 255, G: 0, B: 0, A: 255} // Red
	} else if value >= 201 && value <= 300 {
		return color.NRGBA{R: 153, G: 0, B: 76, A: 255} // Purple
	} else {
		return color.NRGBA{R: 126, G: 0, B: 35, A: 255} // Maroon
	}
}

func createUI(myWindow fyne.Window) []*canvas.Text {
	ppm10TextLable := canvas.NewText("PPM 1.0", theme.Color(theme.ColorNameForeground))
	ppm25TextLable := canvas.NewText("PPM 2.5", theme.Color(theme.ColorNameForeground))
	ppm100TextLable := canvas.NewText("PPM 10.0", theme.Color(theme.ColorNameForeground))

	ppm10TextLable.TextSize = 75
	ppm25TextLable.TextSize = 100
	ppm100TextLable.TextSize = 75

	ppm10TextValue := canvas.NewText("NaN μ/m³", theme.Color(theme.ColorNamePrimary))
	ppm25TextValue := canvas.NewText("NaN μ/m³", theme.Color(theme.ColorNamePrimary))
	ppm100TextValue := canvas.NewText("NaN μ/m³", theme.Color(theme.ColorNamePrimary))

	ppm10TextValue.TextSize = 50
	ppm25TextValue.TextSize = 75
	ppm100TextValue.TextSize = 50

	row1 := container.New(layout.NewCustomPaddedHBoxLayout(0), layout.NewSpacer(), ppm25TextLable, layout.NewSpacer())
	row2 := container.New(layout.NewCustomPaddedHBoxLayout(0), layout.NewSpacer(), ppm25TextValue, layout.NewSpacer())

	row3 := container.New(layout.NewCustomPaddedHBoxLayout(50), layout.NewSpacer(), ppm10TextLable, layout.NewSpacer(), ppm100TextLable, layout.NewSpacer())
	row4 := container.New(layout.NewCustomPaddedHBoxLayout(50), layout.NewSpacer(), ppm10TextValue, layout.NewSpacer(), ppm100TextValue, layout.NewSpacer())

	line := canvas.NewLine(color.Gray16{0xEEEE})
	line.StrokeWidth = 1
	line.Position1 = fyne.NewPos(0, 100)
	line.Position2 = fyne.NewPos(0, 0)

	fullScreenToggleIcon := widget.NewToolbarAction(theme.ViewFullScreenIcon(), func() {
		myWindow.SetFullScreen(!myWindow.FullScreen())
	})

	toolbar := widget.NewToolbar(
		fullScreenToggleIcon,
		widget.NewToolbarSpacer(),
	)

	toolBarBorder := container.NewBorder(toolbar, nil, nil, nil)

	myWindow.SetContent(container.New(layout.NewCustomPaddedVBoxLayout(50), toolBarBorder, row1, row2, line,
		row3, row4, layout.NewSpacer()))
	myWindow.SetFullScreen(false)

	f11Key := &desktop.CustomShortcut{KeyName: fyne.KeyF11, Modifier: fyne.KeyModifierControl}

	myWindow.Canvas().AddShortcut(f11Key, func(shortcut fyne.Shortcut) {
		log.Println("Presses Ctrl + F11")
		myWindow.SetFullScreen(!myWindow.FullScreen())
	})

	return []*canvas.Text{ppm10TextValue, ppm25TextValue, ppm100TextValue}

}
