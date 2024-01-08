package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	"temperature-service/models"
)

type Printer struct {
	w *tabwriter.Writer
}

func New() *Printer {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
	return &Printer{
		w: w,
	}
}

func (p *Printer) CityHeader() {
	fmt.Fprintln(p.w, "Name\tTempC\tTempF\tBeach vacation ready?\tSki vacation ready?")
}

func (p *Printer) CityDetails(c models.CityTemp) {
	fmt.Fprintf(p.w, "%v\t%v\t%v\t%v\t%v\n", c.Name(), c.TempC(), c.TempF(), c.BeachVacationReady(), c.SkiVacationReady())
}

func (p *Printer) Cleanup() {
	p.w.Flush()
}
