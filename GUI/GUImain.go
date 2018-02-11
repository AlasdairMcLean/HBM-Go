package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"

	"../hbmplot"
	"../hbmutil"
)

const (
	//WINWID is the width of the window by default
	WINWID = 800
	//WINHEI is the height of the window by default
	WINHEI = 600
)

func main() {
	gtk.Init(nil)

	// gui boilerplate
	win, box := setupWindow("HBM plot", 800, 600)
	da, err := gtk.DrawingAreaNew()
	if err != nil {
		panic(err)
	}

	box.PackStart(da, true, true, 0)
	win.Add(box)
	win.ShowAll()
	file := "TSMEP1.csv"
	pts, err := hbmutil.ReadMEP(file, 4)
	if err != nil {
		panic(err)
	}
	ptsi := pts.ToMati()
	a := hbmutil.PtstoImgi(ptsi, 256, 256)
	img := a.ToMatff()
	MEPs := img.Getcol(4)
	img.Scale(float64(1.0 / hbmutil.Max(MEPs)))
	hbmplot.Drawim(da, img)

	gtk.Main()
}

func setupWindow(title string, wid, hei int) (*gtk.Window, *gtk.Box) {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	gtk.WindowSetDefaultIconFromFile("./brainimg.png")
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(wid, hei)
	wid = hei + 1
	win.SetPosition(gtk.WIN_POS_CENTER)
	mainbox := makevGTKBox()
	setupBoxToolBar(mainbox, "File", "Edit", "Help")
	return win, mainbox
}

func setupBoxToolBar(box *gtk.Box, label ...string) {
	guitoolbar, err := gtk.ToolbarNew()
	if err != nil {
		panic(err)
	}

	for i, v := range label {
		newbutton, _ := gtk.ToolButtonNew(nil, v)
		guitoolbar.Insert(newbutton, i)
	}
	box.PackStart(guitoolbar, false, false, 0)
}

func makevGTKBox() *gtk.Box {
	newbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	if err != nil {
		panic(err)
	}

	return newbox
}

/*import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/gillesdemey/go-dicom"
)

func main() {

	bytes, err := ioutil.ReadFile("TAIZA_GRESPAN_SANTOS0100.dcm")
	if err != nil {
		panic(err)
	}

	parser, err := dicom.NewParser()
	if err != nil {
		panic(err)
	}

	data, c := parser.Parse(bytes)

	dcm := &dicom.DicomFile{}
	gw := new(sync.WaitGroup)
	dcm.Discard(c, gw)
	gw.Wait()
	for _, elem := range data.Elements {
		fmt.Printf("%+v\n", &elem)
	}
	pdata := data.Elements[52]
	//pdatain := pdata.Value
	//pdatain2 := (pdatain[0]).(string)
	pdatain2 := pdata.Value
	pdatain3 := pdatain2[0].(string)
	pdatain4 := []int32(pdatain3)
	fmt.Println("---")
	fmt.Println(pdatain4)
	fmt.Println("---")
	fmt.Println(len(pdatain4))
	fmt.Println(len(pdatain2))
	fmt.Println("---")
	var joined []string
	for i := 0; i < len(pdatain4); i++ {
		strings.Join(joined, string(pdatain4[i]))
	}
	fmt.Println(joined)
}
*/
