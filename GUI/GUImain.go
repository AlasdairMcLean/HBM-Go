package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"

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
	da, err := gtk.DrawingAreaNew() // make a new gtk drawing area where the image will go
	if err != nil {
		panic(err)
	}
	box.PackStart(da, true, true, 0) // 'pack' the drawing area from the top of the screen
	win.Add(box)                     // add the box that holds the drawing area to the window
	win.ShowAll()                    // reveal the window

	file := getfile()
	// as an example, read TSMEP1.csv
	pts, err := hbmutil.ReadMEP(file, 0, 1, 2, 3) //use the csv parse library in hbmutil to return an array of points and corresponding mep amplitudes
	if err != nil {
		panic(err)
	}
	ptsi := pts.ToMati() // round the float64 pts down to int
	xcol := 0
	ycol := 1
	a := hbmutil.MattoImgi(ptsi, xcol, ycol, 256, 256) // use the integer-based points-to-image function in the hbmutil package to return a 256 by 256 px image
	img := a.ToMatff()                                 // return the image to float64 as Cairo needs
	MEPs := img.Getcol(4)                              // get the 4th column
	img.Scale(float64(1.0 / hbmutil.Maxffl(MEPs)))     // scale the MEPs down to the [0,1] range that gtk's drawing area requires
	hbmplot.Drawim(da, img)                            // draws image on the drawing area using hbmutil package
	canvsc := 1
	keyMap := map[uint]func(){
		gdk.KEY_equal: func() {
			canvsc++
			hbmplot.ExpandCanvas(da, img, canvsc)
		},
		gdk.KEY_minus: func() {
			canvsc--
			hbmplot.ExpandCanvas(da, img, canvsc)
		},
		gdk.KEY_Insert: func() {
			if xcol < 2 {
				xcol++
			} else {
				xcol = 0
			}
			a = hbmutil.MattoImgi(ptsi, xcol, ycol, 256, 256)
			img = a.ToMatff()
			MEPs = img.Getcol(4)
			img.Scale(float64(1.0 / hbmutil.Maxffl(MEPs)))
			hbmplot.Drawim(da, img)
		},
		gdk.KEY_Delete: func() {
			if ycol < 2 {
				ycol++
			} else {
				ycol = 0
			}
			a = hbmutil.MattoImgi(ptsi, xcol, ycol, 256, 256)
			img = a.ToMatff()
			MEPs = img.Getcol(4)
			img.Scale(float64(1.0 / hbmutil.Maxffl(MEPs)))
			hbmplot.Drawim(da, img)
		},
	}

	win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		if zoom, found := keyMap[keyEvent.KeyVal()]; found {
			zoom()
			win.QueueDraw()
		}
	})

	gtk.Main() // start the main loop, waiting for user to close the window
}

//setupWindow will perform all the startup boilerplate for the main window of the application
func setupWindow(title string, wid, hei int) (*gtk.Window, *gtk.Box) {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL) //create a new toplevel window
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	gtk.WindowSetDefaultIconFromFile("./brainimg.png") // use the brainimg.png file as an icon
	win.SetTitle(title)                                // sets title to that specified as input to setupWindow function
	win.Connect("destroy", func() {                    //destroy the window if the user clicks the x
		gtk.MainQuit()
	})
	win.SetDefaultSize(wid, hei)                     // set the default size to that specified in setupWindow input
	win.SetPosition(gtk.WIN_POS_CENTER)              // centers the window on the screen
	mainbox := makevGTKBox()                         // make a new box for packing elements
	setupBoxToolBar(mainbox, "File", "Edit", "Help") // make an example toolbar
	return win, mainbox
}

func setupBoxToolBar(box *gtk.Box, label ...string) { //sets up the example toolbar. will have callback functions later.
	guitoolbar, err := gtk.ToolbarNew() // create a new empty gtk toolbar object
	if err != nil {
		panic(err)
	}

	for i, v := range label { //for each label specified in setupBoxToolBar input
		newbutton, _ := gtk.ToolButtonNew(nil, v) //create the new button with label equal to string
		guitoolbar.Insert(newbutton, i)           // insert that button into the toolbar
	}
	box.PackStart(guitoolbar, false, false, 0) //pack all the buttons in at the left side of the bar
}

func makevGTKBox() *gtk.Box { //boilerplate code for a new box to hold all the objects
	newbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1) //make a vertical box (means that startpack will pack from the top)
	if err != nil {
		panic(err)
	}

	return newbox
}

func getfile() string {
	dlgwin, err := gtk.WindowNew(gtk.WINDOW_POPUP) //create a new toplevel window
	if err != nil {
		panic("Unable to create window:", err)
	}
	dlgwin.SetTitle("Pick a csv file") // sets title to that specified as input to setupWindow function
	dlgwin.Connect("destroy", func() { //destroy the window if the user clicks the x
		gtk.MainQuit()
	})
	dlgwin.SetDefaultSize(WINWID/2, WINHEI/2) // set the default size to that specified in setupWindow input
	dlgwin.SetPosition(gtk.WIN_POS_CENTER)
	filechooser, _ := gtk.FileChooserDialogNewWith2Buttons(
		"Open...",
		dlgwin,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		"Cancel",
		gtk.RESPONSE_DELETE_EVENT,
		"Open",
		gtk.RESPONSE_ACCEPT)
	filter, _ := gtk.FileFilterNew()
	filter.AddPattern("*.csv")
	filter.SetName(".csv")
	filechooser.AddFilter(filter)

	switcher := filechooser.Run()
	filename := filechooser.GetFilename()
	filechooser.Destroy()
	if switcher != -3 {

		//Nothing more to do here
		filename = ""
		return ""
	}
	return filename
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
