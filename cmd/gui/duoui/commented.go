package duoui


// func (ui *DuoUI) DuoUItoastSys() {
//	layout.Align(layout.NE).Layout(ui.ly.Context, func() {
//		listToasts.Layout(ui.ly.Context, len(ui.rc.Toasts), func(i int) {
//			layout.UniformInset(unit.Dp(16)).Layout(ui.ly.Context, ui.rc.Toasts[i])
//		})
//	})
// }

// func (ui *DuoUI) toastAdd() {
//	ui.rc.Toasts = append(ui.rc.Toasts, func() {
//		//iconOK, _ := view.NewDuoUIicon(icons.NavigationCheck)
//		theme.DuoUIdrawRectangle(ui.ly.Context, 418, 160, "aa000000", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
//		theme.DuoUIdrawRectangle(ui.ly.Context, 408, 150, ui.ly.Theme.Color.Primary, [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
//		layout.Flex{
//			Axis:      layout.Vertical,
//			Alignment: layout.Middle,
//		}.Layout(ui.ly.Context,
//			layout.Rigid(func() {
//				layout.Flex{
//					Axis:      layout.Horizontal,
//					Alignment: layout.Middle,
//				}.Layout(ui.ly.Context,
//					layout.Rigid(func() {
//						layout.Align(layout.Center).Layout(ui.ly.Context, func() {
//							layout.Inset{Top: unit.Dp(24), Bottom: unit.Dp(8), Left: unit.Dp(0), Right: unit.Dp(4)}.Layout(ui.ly.Context, func() {
//								cur := ui.ly.Theme.H4("TOAST MESSAGE!!!")
//								cur.Color = ui.ly.Theme.Color.Light
//								cur.Alignment = text.Start
//								cur.Layout(ui.ly.Context)
//							})
//						})
//					}),
//				)
//			}),
//			layout.Rigid(func() {
//				layout.Flex{
//					Axis:      layout.Horizontal,
//					Alignment: layout.Middle,
//				}.Layout(ui.ly.Context)
//				//layout.Rigid(toastButton("OK", "ffcfcfcf", "ff308030", "ffcfcfcf", duo, rc, buttonToastOK, iconOK)),
//
//			}),
//		)
//	})
//	go func(ui *DuoUI) {
//		time.Sleep(3 * time.Second)
//		ui.rc.Toasts[len(ui.rc.Toasts)-1] = nil // or the zero value of T
//		ui.rc.Toasts = ui.rc.Toasts[:len(ui.rc.Toasts)-1]
//		op.InvalidateOp{}.Add(ui.ly.Context.Ops)
//	}(ui)
// }

