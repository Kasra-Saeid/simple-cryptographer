package app

import (
	"image/color"
	"log"
	"simple_cryptographer/config"
	"simple_cryptographer/internal/decryption"
	"simple_cryptographer/internal/encryption"
	"simple_cryptographer/pkg/cryptography"
	external_file "simple_cryptographer/pkg/external_files"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	ENCRYPT = "encrypt"
	DECRYPT = "decrypt"
)

func Run(cfg *config.Config) {
	secret := cfg.SecretKey
	cryptoPkg := cryptography.New(secret)
	txtFileRepo, err := external_file.NewTextFile("cache/phrase.txt")
	if err != nil {
		log.Fatalln(err)
	}
	enc := encryption.NewEncryption(cryptoPkg, *txtFileRepo)
	dec := decryption.NewDecryption(cryptoPkg, *txtFileRepo)
	a := app.New()
	w := a.NewWindow("simple cryptographer")
	w.Resize(fyne.NewSize(1024, 800))
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("enter text...")

	selectedCryptoOption := binding.NewString()

	output := canvas.NewText("reault", color.White)
	outPutContainer := container.NewVScroll(output)
	copyButton := widget.NewButtonWithIcon("copy", theme.ContentCopyIcon(), func() {
		w.Clipboard().SetContent(output.Text)
	})

	resultSection := container.NewHSplit(outPutContainer, copyButton)
	execButton := widget.NewButton(ENCRYPT, func() {
		selectedOption, err := selectedCryptoOption.Get()
		if err != nil {
			log.Fatalln(err)
		}
		if selectedOption == ENCRYPT {

			output.Text, err = enc.GetService().ToEncryptedText(input.Text)
			if err != nil {
				log.Fatal(err)
			}

			err = enc.GetService().SaveToFile(input.Text, output.Text)

			if err != nil {
				log.Fatal(err)
			}
		} else {

			output.Text, err = dec.GetService().ToDecryptedText(input.Text)
			if err != nil {
				log.Fatal(err)
			}

			err = enc.GetService().SaveToFile(input.Text, output.Text)

			if err != nil {
				log.Fatal(err)
			}
		}

	})

	cryptoOption := widget.NewRadioGroup([]string{ENCRYPT, DECRYPT}, func(s string) {
		selectedCryptoOption.Set(s)
		execButton.SetText(s)
	})

	cryptoOption.SetSelected(ENCRYPT)

	con := container.NewVBox(cryptoOption, input, resultSection, execButton)
	w.SetContent(con)
	w.Show()
	a.Run()
	defer txtFileRepo.CloseFile()
}
