package component

import (
	"fmt"
	"gioui.org/layout"

	"github.com/p9c/pod/cmd/gui/model"
	"github.com/p9c/pod/cmd/gui/rcd"
	"github.com/p9c/pod/pkg/gui/gel"
	"github.com/p9c/pod/pkg/gui/gelook"
)

//func TransactionsList(rc *rcd.RcVar, gtx *layout.Context, th *gelook.DuoUItheme) func() {
//	return func() {
//		transactionsPanel := th.DuoUIpanel()
//		transactionsPanel.PanelObject = rc.History.Txs.Txs
//		transactionsPanel.ScrollBar = th.ScrollBar()
//		transactionsPanelElement.PanelObjectsNumber = len(rc.History.Txs.Txs)
//		transactionsPanel.Layout(gtx, transactionsPanelElement, func(i int, in interface{}) {
//			txs := in.([]model.DuoUItransactionExcerpt)
//			t := txs[i]
//			th.DuoUIline(gtx, 0, 0, 1, th.Colors["Hint"])()
//			for t.Link.Clicked(gtx) {
//				rc.ShowPage = fmt.Sprintf("TRANSACTION %s", t.TxID)
//				rc.GetSingleTx(t.TxID)()
//				//SetPage(rc, txPage(rc, gtx, th, t.TxID))
//			}
//			width := gtx.Constraints.Width.Max
//			button := th.DuoUIbutton("", "", "", "", "", "", "", "", 0, 0, 0, 0, 0, 0, 0, 0)
//			button.InsideLayout(gtx, t.Link, func() {
//				gtx.Constraints.Width.Min = width
//				layout.Flex{
//					Spacing: layout.SpaceBetween,
//				}.Layout(gtx,
//					layout.Rigid(txsDetails(gtx, th, i, &t)),
//					layout.Rigid(Label(gtx, th, th.Fonts["Mono"], 12, th.Colors["Secondary"], fmt.Sprintf("%0.8f", t.Amount))))
//			})
//		})
//	}
//}

func TxsDetails(gtx *layout.Context, th *gelook.DuoUItheme, i int, t *model.DuoUItransactionExcerpt) func() {
	return func() {
		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(i))),
			layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.TxID)),
			layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprintf("%0.8f", t.Amount))),
			layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.Category)),
			layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.Time)),
		)
	}
}

func TransactionsFilter(rc *rcd.RcVar, gtx *layout.Context, th *gelook.DuoUItheme) func() {
	cats := rc.History.Categories
	return func() {
		layout.Flex{}.Layout(gtx,
			layout.Rigid(txsFilterItem(gtx, th, "ALL", cats.AllTxs)),
			layout.Rigid(txsFilterItem(gtx, th, "MINTED", cats.MintedTxs)),
			layout.Rigid(txsFilterItem(gtx, th, "IMATURE", cats.ImmatureTxs)),
			layout.Rigid(txsFilterItem(gtx, th, "SENT", cats.SentTxs)),
			layout.Rigid(txsFilterItem(gtx, th, "RECEIVED", cats.ReceivedTxs)))
		switch c := true; c {
		case cats.AllTxs.Checked(gtx):
			rc.History.Category = "all"
		case cats.MintedTxs.Checked(gtx):
			rc.History.Category = "generate"
		case cats.ImmatureTxs.Checked(gtx):
			rc.History.Category = "immature"
		case cats.SentTxs.Checked(gtx):
			rc.History.Category = "sent"
		case cats.ReceivedTxs.Checked(gtx):
			rc.History.Category = "received"
		}
	}
}

func txsFilterItem(gtx *layout.Context, th *gelook.DuoUItheme, id string, c *gel.CheckBox) func() {
	return func() {
		th.DuoUIcheckBox(id, th.Colors["Light"], th.Colors["Light"]).Layout(gtx, c)
	}
}
