package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Gui struct {
	ui            *ebitenui.UI
	updatesPerSec *int
}

func NewGui() *Gui {
	// Root container fills the screen
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	// Panel: small container in the top-right
	menu := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x44, 0x44, 0x44, 0xcc})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(250, 0),
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true}),
			widget.GridLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
				Left:   20,
				Right:  20,
			}),
			widget.GridLayoutOpts.Spacing(0, 20))),
	)

	updatesPerSec := 30

	sc := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(3),
			widget.GridLayoutOpts.Stretch([]bool{false, true}, nil), // label not stretched, slider stretched
			widget.GridLayoutOpts.Spacing(10, 0),
		)),
	)

	var sliderValLable *widget.Label

	// Slider inside the menu
	slider := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(1, 60),
		widget.SliderOpts.InitialCurrent(30),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			&widget.ButtonImage{
				Idle:    image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Hover:   image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Pressed: image.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
			},
		),
		widget.SliderOpts.PageSizeFunc(func() int { return 1 }),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			updatesPerSec = args.Current
			sliderValLable.Label = fmt.Sprintf("%d", updatesPerSec)
		}),
	)

	font, err := loadFontFace()
	if err != nil {
		panic(fmt.Sprintf("unable to load font face: %v", err))
	}

	sliderLabel := widget.NewLabel(
		widget.LabelOpts.TextOpts(widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}))),
		widget.LabelOpts.Text("Speed", font, &widget.LabelColor{
			Idle:     color.White,
			Disabled: color.Gray{150},
		}),
	)

	sliderValLable = widget.NewLabel(
		widget.LabelOpts.TextOpts(widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}))),
		widget.LabelOpts.Text(fmt.Sprintf("%d", slider.Current), font, &widget.LabelColor{
			Idle:     color.White,
			Disabled: color.Gray{150},
		}),
	)

	sc.AddChild(sliderLabel)
	sc.AddChild(slider)
	sc.AddChild(sliderValLable)

	menu.AddChild(sc)
	rootContainer.AddChild(menu)

	return &Gui{
		ui: &ebitenui.UI{
			Container: rootContainer,
		},
		updatesPerSec: &updatesPerSec,
	}
}

func loadFontFace() (text.Face, error) {
	fontFile, err := embeddedAssets.Open("assets/fonts/notosans-bold.ttf")
	if err != nil {
		return nil, err
	}

	s, err := text.NewGoTextFaceSource(fontFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &text.GoTextFace{
		Source: s,
		Size:   24,
	}, nil
}

func (gui *Gui) Update() {
	gui.ui.Update()
}

func (gui *Gui) Draw(img *ebiten.Image) {
	gui.ui.Draw(img)
}
