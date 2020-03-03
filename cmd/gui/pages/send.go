package pages

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/p9c/pod/cmd/gui/component"
	"github.com/p9c/pod/cmd/gui/model"
	"github.com/p9c/pod/cmd/gui/rcd"
	"github.com/p9c/pod/pkg/gui/clipboard"
	"github.com/p9c/pod/pkg/gui/controller"
	"github.com/p9c/pod/pkg/gui/theme"
	"strconv"
)

var (
	layautList = &layout.List{
		Axis: layout.Vertical,
	}
	address           string
	amount            float64
	passPharse        string
	addressLineEditor = &controller.Editor{
		SingleLine: true,
	}
	amountLineEditor = &controller.Editor{
		SingleLine: true,
	}
	passLineEditor = &controller.Editor{
		SingleLine: true,
	}
	buttonPasteAddress = new(controller.Button)
	buttonPasteAmount  = new(controller.Button)
	buttonSend         = new(controller.Button)
)

func Send(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme) *theme.DuoUIpage {
	return th.DuoUIpage("SEND", 10, func() {}, func() {}, sendBody(rc, gtx, th), func() {})
}

func sendBody(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme) func() {
	return func() {
		layout.Flex{}.Layout(gtx,
			layout.Rigid(func() {
				cs := gtx.Constraints
				theme.DuoUIdrawRectangle(gtx, cs.Width.Max, 180, th.Colors["Light"], [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})

				widgets := []func(){
					func() {
						layout.Flex{}.Layout(gtx,
							layout.Flexed(1, component.Editor(gtx, th, addressLineEditor, "DUO address", func(e controller.SubmitEvent) {
								address = e.Text
							})),
							layout.Rigid(component.Button(gtx, th, buttonPasteAddress, th.Fonts["Primary"], 12, th.Colors["ButtonText"], th.Colors["ButtonBg"], "PASTE ADDRESS", func() {
								addressLineEditor.SetText(clipboard.Get())
							})))
					},
					func() {
						layout.Flex{}.Layout(gtx,
							layout.Flexed(1, component.Editor(gtx, th, amountLineEditor, "DUO Amount", func(e controller.SubmitEvent) {
								f, err := strconv.ParseFloat(e.Text, 64)
								if err != nil {
									amount = f
									amountLineEditor.SetText("")
								}
							})),
							layout.Rigid(component.Button(gtx, th, buttonPasteAmount, th.Fonts["Primary"], 12, th.Colors["ButtonText"], th.Colors["ButtonBg"], "PASTE AMOUNT", func() {
								amountLineEditor.SetText(clipboard.Get())
							})))
					},
					func() {
						layout.Flex{}.Layout(gtx,
							layout.Rigid(component.Button(gtx, th, buttonSend, th.Fonts["Primary"], 12, th.Colors["ButtonText"], th.Colors["ButtonBg"], "SEND", func() {
								rc.Dialog.Show = true
								rc.Dialog = &model.DuoUIdialog{
									Show: true,
									Ok:   rc.DuoSend(passPharse, address, amount),
									Close: func() {

									},
									CustomField: func() {
										layout.Flex{}.Layout(gtx,
											layout.Flexed(1, component.Editor(gtx, th, passLineEditor, "Enter your password", func(e controller.SubmitEvent) {
												passPharse = e.Text
											})))
									},
									Cancel: func() { rc.Dialog.Show = false },
									Title:  "Are you sure?",
									Text:   "Confirm ParallelCoin send",
								}
							})))
					},
				}
				layautList.Layout(gtx, len(widgets), func(i int) {
					layout.UniformInset(unit.Dp(8)).Layout(gtx, widgets[i])
				})
			}))
	}
}
