package main

import (
	"flag"
	"fmt"
	"github.com/saesh/mandelbrot/pkg/colors"
	g "github.com/saesh/mandelbrot/pkg/generator"
	"time"
)

func main() {
	mb := g.Mandelbrot{}

	var output string
	var maxRoutines int

	// Считываем все аргументы командной строки
	flag.IntVar(&mb.Height, "h", 1000, "Sets output image height")
	flag.IntVar(&mb.Width, "w", 1000, "Sets output image width.")
	flag.IntVar(&mb.MaxIterations, "i", 100, "Sets maximum iterations")
	flag.IntVar(&mb.Colors, "clr", colors.Hue, "Sets color")
	flag.Float64Var(&mb.X, "x", 0, "Sets x value")
	flag.Float64Var(&mb.Y, "y", 0, "Sets y value")
	flag.Float64Var(&mb.R, "r", 4, "Sets radius")
	flag.StringVar(&output, "output", "mandelbrot.jpeg", "Sets output image.")
	flag.IntVar(&maxRoutines, "async", 0, "Sets maximum amount of coroutines used. If set to 0, then executes synchronously.")
	flag.Parse()

	// Записываем время начала рендера
	start := time.Now()

	if maxRoutines > 0 {
		// Если в параметрах можно указано количество асинхронных операций,
		// то используем указанное количество
		mb.RenderWithMaxRoutines(maxRoutines)
	} else {
		// В противном случае используем стандартную реализацию
		mb.Render()
	}

	// Вычисляем кол-во миллисекунд, затраченно на рендер
	milliseconds := time.Now().Sub(start).Milliseconds()
	fmt.Println("Time elapsed:", milliseconds, "ms")

	// Выводим картинку
	_ = mb.WriteJpeg(output, 90)
}
