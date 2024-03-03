package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("./controller/logo.jpeg", props.Rect{
					Center:  true,
					Percent: 75,
				})

				if err != nil {
					fmt.Println("Image file was not loaded 😱 - ", err)
				}

			})
		})
	})
}

func buildFruitList(m pdf.Maroto) {
	headings := getHeadings()
	// contents := data.FruitList(20)
	contents := [][]string{{"Wheel", "Red and juicy", "200"}, {"RY", "Orange and juicy", "70"}}
	purpleColor := getPurpleColor()

	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Car Part", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())
	m.TableList(headings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Left,
		AlternatedBackground: &purpleColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	m.Row(20, func() {
		m.ColSpace(7)
		m.Col(2, func() {
			m.Text("Total:", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("$ XXXX.00", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})
}

func buildFooter(m pdf.Maroto) {
	begin := time.Now()
	m.SetAliasNbPages("{nb}")
	m.SetFirstPageNb(1)

	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(6, func() {
				m.Text(begin.Format("02/01/2006"), props.Text{
					Top:   10,
					Size:  8,
					Color: getGreyColor(),
					Align: consts.Left,
				})
			})

			m.Col(6, func() {
				m.Text("Page "+strconv.Itoa(m.GetCurrentPage())+" of {nb}", props.Text{
					Top:   10,
					Size:  8,
					Style: consts.Italic,
					Color: getGreyColor(),
					Align: consts.Right,
				})
			})

		})
	})
}

func getHeadings() []string {
	return []string{"Fruit", "Description", "Price"}
}

// Colours

func getPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getGreyColor() color.Color {
	return color.Color{
		Red:   206,
		Green: 206,
		Blue:  206,
	}
}
