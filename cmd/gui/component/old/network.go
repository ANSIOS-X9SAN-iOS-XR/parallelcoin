package component

import (
	"fmt"

	"gioui.org/layout"

	"github.com/p9c/pod/cmd/gui/rcd"
	"github.com/p9c/pod/pkg/gui/gelook"
	"github.com/p9c/pod/pkg/rpc/btcjson"
)

func PeersList(rc *rcd.RcVar, gtx *layout.Context, th *gelook.DuoUItheme) func() {
	c := th.Colors["Dark"]
	return func() {
		rc.Network.PeersList.Layout(gtx, len(rc.Network.Peers), func(i int) {
			t := rc.Network.Peers[i]
			th.DuoUIline(gtx, 0, 0, 1,
				th.Colors["Hint"])()
			layout.Flex{
				Spacing: layout.SpaceBetween,
			}.Layout(gtx,
				layout.Rigid(Label(gtx, th, th.Fonts["Mono"], 14,
					c, fmt.Sprint(t.ID))),
				layout.Flexed(1, peerDetails(gtx, th, i, t)),
				layout.Rigid(Label(gtx, th, th.Fonts["Mono"], 14,
					c, t.Addr)))
		})
	}
}

func peerDetails(gtx *layout.Context, th *gelook.DuoUItheme, i int, t *btcjson.GetPeerInfoResult) func() {
	prim := th.Fonts["Primary"]
	c := th.Colors["Dark"]
	return func() {
		layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceAround,
		}.Layout(gtx,
			layout.Rigid(Label(gtx, th, prim, 12, c, t.AddrLocal)),
			layout.Rigid(Label(gtx, th, prim, 12, c, t.Services)),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.RelayTxes))),
			// layout.Rigid(Label(gtx, th, prim, 12, c
			//fmt.Sprint(t.LastSend))),
			// layout.Rigid(Label(gtx, th, prim, 12, c
			//fmt.Sprint(t.LastRecv))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.BytesSent))),
			layout.Rigid(Label(gtx, th, prim, 12, c,
				fmt.Sprint(t.BytesRecv))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.ConnTime))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.TimeOffset))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.PingTime))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.PingWait))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.Version))),
			layout.Rigid(Label(gtx, th, prim, 12, c, t.SubVer)),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.Inbound))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.StartingHeight))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.CurrentHeight))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.BanScore))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.FeeFilter))),
			layout.Rigid(Label(gtx, th, prim, 12, c, fmt.Sprint(t.SyncNode))))
	}
}
