package main

import (
	"fmt"

	"./turing"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {

	var input *turing.Tape
	var output *turing.Tape
	var cnt int
	var print string
	var machine *turing.Machine

	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Turing Machine!")
	window.SetIconName("gtk-dialog-info")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")
	window.Container.SetBorderWidth(80)

	vbox := gtk.NewVBox(false, 5)
	valign := gtk.NewAlignment(0, 1, 0, 0)
	vbox.Container.Add(valign)
	drawingArea := gtk.NewDrawingArea()
	drawingArea.SetSizeRequest(300, 500)
	var pixmap *gdk.Pixmap
	var gc *gdk.GC
	var draw *gdk.Drawable
	drawingArea.Connect("configure-event", func() {
		if pixmap != nil {
			pixmap.Unref()
		}
		var allocation *gtk.Allocation
		allocation = drawingArea.GetAllocation()
		draw = drawingArea.GetWindow().GetDrawable()
		pixmap = gdk.NewPixmap(draw, allocation.Width, allocation.Height, 24)
		gc = gdk.NewGC(pixmap.GetDrawable())
		gc.SetRgbFgColor(gdk.NewColor("white"))
		pixmap.GetDrawable().DrawRectangle(gc, true, 0, 0, -1, -1)

		gc.SetRgbFgColor(gdk.NewColor("black"))
		gc.SetRgbBgColor(gdk.NewColor("white"))
	}, nil)
	drawingArea.Connect("expose-event", func() {
		if pixmap != nil {
			drawingArea.GetWindow().GetDrawable().DrawDrawable(gc, pixmap.GetDrawable(), 0, 0, 0, 0, -1, -1)
		}
	})
	vbox.Container.Add(drawingArea)

	textBox := gtk.NewTextView()
	textBox.Container.SetSizeRequest(500, 30)

	comboboxentry := gtk.NewComboBoxTextWithEntry()

	t := gtk.NewEntry()

	window.Container.Add(vbox)
	hbox := gtk.NewHBox(false, 3)

	hbox.Container.Add(textBox)

	hbox.Container.Add(comboboxentry)

	hbox.Container.Add(t)
	loadBtn := gtk.NewButtonWithLabel("LOAD")
	loadBtn.SetSizeRequest(100, 30)

	hbox.Container.Add(loadBtn)
	clearBtn := gtk.NewButtonWithLabel("CLEAR")
	clearBtn.Clicked(func() {
		t.DeleteText(0, 1000)
	})
	clearBtn.SetSizeRequest(100, 30)
	hbox.Container.Add(clearBtn)
	halign := gtk.NewAlignment(1, 0, 0, 0)
	halign.Container.Add(hbox)
	vbox.PackStart(halign, false, false, 0)

	comboboxentry.AppendText("Sort")
	comboboxentry.AppendText("Increment")
	comboboxentry.AppendText("Addition")
	comboboxentry.Connect("changed", func() {
		println("value:", comboboxentry.GetActiveText())

		if comboboxentry.GetActiveText() == "Increment" {

			machine = turing.NewMachine([]turing.Rule{
				{State: "q0", Symbol: '1', Write: '1', Motion: turing.Right, Next: "q0"},
				{State: "q0", Symbol: ' ', Write: '1', Motion: turing.Stay, Next: "qf"},
			})

		}
		if comboboxentry.GetActiveText() == "Sort" {
			machine = turing.NewMachine([]turing.Rule{
				// Moving right, first b→B;s1
				{State: "s0", Symbol: 'a', Write: 'a', Motion: turing.Right, Next: "s0"},
				{State: "s0", Symbol: 'b', Write: 'B', Motion: turing.Right, Next: "s1"},
				{State: "s0", Symbol: ' ', Write: ' ', Motion: turing.Left, Next: "se"},
				// Conintue right to end of tape → s2
				{State: "s1", Symbol: 'a', Write: 'a', Motion: turing.Right, Next: "s1"},
				{State: "s1", Symbol: 'b', Write: 'b', Motion: turing.Right, Next: "s1"},
				{State: "s1", Symbol: ' ', Write: ' ', Motion: turing.Left, Next: "s2"},
				// Continue left over b.  a→b;s3, B→b;se
				{State: "s2", Symbol: 'a', Write: 'b', Motion: turing.Left, Next: "s3"},
				{State: "s2", Symbol: 'b', Write: 'b', Motion: turing.Left, Next: "s2"},
				{State: "s2", Symbol: 'B', Write: 'b', Motion: turing.Left, Next: "se"},
				// Continue left until B→a;s0
				{State: "s3", Symbol: 'a', Write: 'a', Motion: turing.Left, Next: "s3"},
				{State: "s3", Symbol: 'b', Write: 'b', Motion: turing.Left, Next: "s3"},
				{State: "s3", Symbol: 'B', Write: 'a', Motion: turing.Right, Next: "s0"},
				// Move to tape start → halt
				{State: "se", Symbol: 'a', Write: 'a', Motion: turing.Left, Next: "se"},
				{State: "se", Symbol: ' ', Write: ' ', Motion: turing.Right, Next: "see"},
			})

		}
		if comboboxentry.GetActiveText() == "Addition" {

			machine = turing.NewMachine([]turing.Rule{
				{State: "k0", Symbol: '1', Write: '1', Motion: turing.Right, Next: "k0"},
				{State: "k0", Symbol: '*', Write: '1', Motion: turing.Left, Next: "k1"},
				{State: "k0", Symbol: ' ', Write: ' ', Motion: turing.Left, Next: "ke"},

				{State: "k1", Symbol: '1', Write: '1', Motion: turing.Left, Next: "k1"},
				{State: "k1", Symbol: ' ', Write: ' ', Motion: turing.Right, Next: "k2"},

				//{State: "k2", Symbol: '1', Write: ' ', Motion: turing.Right, Next: "k3"},

				//{State: "k3", Symbol: '1', Write: ' ', Motion: turing.Right, Next: "ke"},
				//{State: "ke", Symbol: ' ', Write: ' ', Motion: turing.Right, Next: "kee"},
			})

		}
		loadBtn.Clicked(func() {
			input = turing.NewTape(' ', 0, []turing.Symbol(t.GetText()))
			cnt, output = machine.Run(input)
			//fmt.Println("Resulting tape:", output)
			draw.DrawRectangle(gc, true, 10, 20, 20, 30)
		})

		loadBtn.Connect("clicked", func() {
			print = fmt.Sprintf("Turing machine halts after %d operations \n Resulting tape: %s", cnt, output)
			textBox.GetBuffer().SetText(print)
			fmt.Println(print)

		})

	})

	window.SetSizeRequest(1300, 690)
	window.ShowAll()
	gtk.Main()
}
