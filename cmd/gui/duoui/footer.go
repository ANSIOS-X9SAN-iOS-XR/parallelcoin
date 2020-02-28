package duoui

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/p9c/pod/cmd/gui/mvc/controller"
	"github.com/p9c/pod/cmd/gui/mvc/model"
	"github.com/p9c/pod/cmd/gui/mvc/theme"
	"github.com/p9c/pod/pkg/util/interrupt"
)

var (
	buttonLog      = new(controller.Button)
	buttonQuit     = new(controller.Button)
	buttonSettings = new(controller.Button)
	buttonNetwork  = new(controller.Button)
	buttonBlocks   = new(controller.Button)
	buttonConsole  = new(controller.Button)
	buttonHelp     = new(controller.Button)
	cornerNav      = &layout.List{
		Axis: layout.Horizontal,
	}
	footerNav = &layout.List{
		Axis: layout.Horizontal,
	}
	footerMenuItemWidth             int = 48
	footerMenuItemHeight            int = 48
	footerMenuItemTextSize          int = 16
	footerMenuItemIconSize          int = 32
	footerMenuItemPaddingVertical   int = 8
	footerMenuItemPaddingHorizontal int = 8
)

func (ui *DuoUI) DuoUIfooter() func() {
	return func() {
		cs := ui.ly.Context.Constraints
		theme.DuoUIdrawRectangle(ui.ly.Context, cs.Width.Max, 64, ui.ly.Theme.Color.Dark, [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})

		layout.Flex{Spacing: layout.SpaceBetween}.Layout(ui.ly.Context,
			layout.Rigid(func() {
				layout.UniformInset(unit.Dp(0)).Layout(ui.ly.Context, func() {
					cornerButtons := []func(){
						func() {
							layout.UniformInset(unit.Dp(0)).Layout(ui.ly.Context, func() {
								var closeMeniItem theme.DuoUIbutton
								closeMeniItem = ui.ly.Theme.DuoUIbutton("", "", "", ui.ly.Theme.Color.Dark, "closeIcon", currentCurrentPageColor(ui.rc.ShowPage, "CLOSE", ui.ly.Theme.Color.Light, ui.ly.Theme.Color.Primary), footerMenuItemTextSize, footerMenuItemIconSize, footerMenuItemWidth, footerMenuItemHeight, footerMenuItemPaddingVertical, footerMenuItemPaddingHorizontal)
								for buttonQuit.Clicked(ui.ly.Context) {
									ui.rc.Dialog.Show = true
									ui.rc.Dialog = &model.DuoUIdialog{
										Show: true,
										Ok: func() {
											interrupt.Request()
											// TODO make this close the window or at least switch to a shutdown screen
											ui.rc.Dialog.Show = false
										},
										Close: func() {
											interrupt.RequestRestart()
											// TODO make this close the window or at least switch to a shutdown screen
											ui.rc.Dialog.Show = false
										},
										Cancel: func() { ui.rc.Dialog.Show = false },
										Title:  "Are you sure?",
										Text:   "Confirm ParallelCoin close",
									}
								}
								closeMeniItem.IconLayout(ui.ly.Context, buttonQuit)
							})
						},

						func() {
							var logMenuItem theme.DuoUIbutton
							logMenuItem = ui.ly.Theme.DuoUIbutton("", "", "", ui.ly.Theme.Color.Dark, "traceIcon", currentCurrentPageColor(ui.rc.ShowPage, "LOG", ui.ly.Theme.Color.Light, ui.ly.Theme.Color.Primary), footerMenuItemTextSize, footerMenuItemIconSize, footerMenuItemWidth, footerMenuItemHeight, footerMenuItemPaddingVertical, footerMenuItemPaddingHorizontal)

							for buttonLog.Clicked(ui.ly.Context) {
								ui.rc.ShowPage = "LOG"
							}
							logMenuItem.IconLayout(ui.ly.Context, buttonLog)
						},
					}
					cornerNav.Layout(ui.ly.Context, len(cornerButtons), func(i int) {
						layout.UniformInset(unit.Dp(0)).Layout(ui.ly.Context, cornerButtons[i])
					})
				})
			}),
			layout.Rigid(func() {
				layout.UniformInset(unit.Dp(0)).Layout(ui.ly.Context, func() {
					navButtons := []func(){
						ui.footerMenuButton(ui.ly.Pages.Theme["NETWORK"], "CONNECTIONS: "+fmt.Sprint(ui.rc.Status.Node.ConnectionCount), "", buttonNetwork),
						ui.footerMenuButton(ui.ly.Pages.Theme["EXPLORER"], "BLOCKS: "+fmt.Sprint(ui.rc.Status.Node.BlockCount), "", buttonBlocks),
						ui.footerMenuButton(ui.ly.Pages.Theme["MINER"], "", "helpIcon", buttonHelp),
						ui.footerMenuButton(ui.ly.Pages.Theme["CONSOLE"], "", "consoleIcon", buttonConsole),
						ui.footerMenuButton(ui.ly.Pages.Theme["SETTINGS"], "", "settingsIcon", buttonSettings),
					}
					footerNav.Layout(ui.ly.Context, len(navButtons), func(i int) {
						layout.UniformInset(unit.Dp(0)).Layout(ui.ly.Context, navButtons[i])
					})
				})
			}),
		)
	}
}

func (ui *DuoUI) footerMenuButton(page *theme.DuoUIpage, text, icon string, footerButton *controller.Button) func() {
	return func() {
		layout.UniformInset(unit.Dp(0)).Layout(ui.ly.Context, func() {
			var footerMenuItem theme.DuoUIbutton
			if icon != "" {
				footerMenuItem = ui.ly.Theme.DuoUIbutton("", "", "", ui.ly.Theme.Color.Dark, icon, currentCurrentPageColor(ui.rc.ShowPage, page.Title, navItemIconColor, ui.ly.Theme.Color.Primary), footerMenuItemTextSize, footerMenuItemIconSize, footerMenuItemWidth, footerMenuItemHeight, footerMenuItemPaddingVertical, footerMenuItemPaddingHorizontal)
				for footerButton.Clicked(ui.ly.Context) {
					ui.rc.ShowPage = page.Title
					setPage(ui.rc, ui.ly.Pages.Theme[page.Title])
					//page.Command()
				}
				footerMenuItem.IconLayout(ui.ly.Context, footerButton)
			} else {
				footerMenuItem = ui.ly.Theme.DuoUIbutton(ui.ly.Theme.Font.Primary, text, currentCurrentPageColor(ui.rc.ShowPage, page.Title, ui.ly.Theme.Color.Light, ui.ly.Theme.Color.Primary), "", "", "", footerMenuItemTextSize, footerMenuItemIconSize, 0, footerMenuItemHeight, footerMenuItemPaddingVertical, 0)
				for footerButton.Clicked(ui.ly.Context) {
					ui.rc.ShowPage = page.Title
					setPage(ui.rc, ui.ly.Pages.Theme[page.Title])
					//page.Command()
				}
				footerMenuItem.Layout(ui.ly.Context, footerButton)
			}
		})
	}
}
