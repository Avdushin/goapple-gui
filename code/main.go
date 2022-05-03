package main

import (
	"image/color"
	"log"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("GOAPPLE-GUI")
	myWindow.Resize(fyne.NewSize(400, 400))

	// set header icon
	hicon, _ := fyne.LoadResourceFromPath("/usr/local/share/applications/src/assets/icons/logo.jpg")
	myApp.SetIcon(hicon)

	// header menu

	about := fyne.NewMenuItem("About", func() {
		url := "https://github.com/avdushin"
		exec.Command("xdg-open", url).Start()
	})

	Theme := fyne.NewMenuItem("Theme", nil)

	// Themes menu
	dark := fyne.NewMenuItem("Dark Theme", func() { myApp.Settings().SetTheme(theme.DarkTheme()) })
	light := fyne.NewMenuItem("Light Theme", func() { myApp.Settings().SetTheme(theme.LightTheme()) })

	Theme.ChildMenu = fyne.NewMenu(
		"",
		dark,
		light,
	)

	FileMenu := fyne.NewMenu("File", about)
	SettingsMenu := fyne.NewMenu("Settings", Theme)

	// main menu
	menu := fyne.NewMainMenu(FileMenu, SettingsMenu)
	// setup menu
	myWindow.SetMainMenu(menu)
	// progress bar
	infinite := widget.NewProgressBarInfinite()
	infinite.Stop()

	// set error's text color
	errorColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}

	errorLabel2 := canvas.NewText(" ERROR: Select your distro!", errorColor)
	errorLabel2.Hide()

	errorLabel3 := widget.NewLabel("IF THE TERMINAL WINDOW DOES NOT APPEAR - SELECT ANOTHER TERMINAL FROM THE LIST")
	errorLabel3.Wrapping = fyne.TextWrapBreak
	errorLabel3.Hide()

	selectDistro := widget.NewSelect(
		[]string{"Manjaro", "Solus", "Fedora"}, func(distro string) {
			log.Printf("Selected distro : %s", distro)
		})

	// stopBTN := widget.NewButton("STOP", func() {
	// 	kill := exec.Command("killall", selectTerm.Selected, "sh", selectDistro.Selected)
	// 	kill.Stdout = os.Stdout
	// 	kill.Stdin = os.Stdin
	// 	kill.Stderr = os.Stderr
	// 	kill.Run()
	// })

	setupBTN := widget.NewButton("SETUP", func() {
		switch selectDistro.Selected {
		case "Manjaro":
			infinite.Start()
			c := exec.Command("st", "-f", "30", "sh", "/usr/local/share/applications/src/distros/Manjaro/Manjaro.sh")
			c.Stdout = os.Stdout
			c.Stdin = os.Stdin
			c.Stderr = os.Stderr
			c.Run()
		case "Solus":
			infinite.Start()
			c := exec.Command("st", "-f", "30", "sh", "/usr/local/share/applications/src/distros/Solus/Solus.sh")
			c.Stdout = os.Stdout
			c.Stdin = os.Stdin
			c.Stderr = os.Stderr
			c.Run()
		case "Fedora":
			infinite.Start()
			c := exec.Command("st", "-f", "30", "sh", "/usr/local/share/applications/src/distros/Fedora/Fedora.sh")
			c.Stdout = os.Stdout
			c.Stdin = os.Stdin
			c.Stderr = os.Stderr
			c.Run()
		}
		// IsDistro Selected check
		switch selectDistro.Selected {
		case "":
			infinite.Stop()
			errorLabel2.Show()
		default:
			errorLabel2.Hide()
			errorLabel3.Show()
		}
	})

	// app label
	label2 := widget.NewLabel("Please select distro:")

	myWindow.SetContent(container.NewVBox(
		label2,
		selectDistro,
		errorLabel2,
		setupBTN,
		infinite,
		errorLabel3,
	))
	myWindow.ShowAndRun()
}
