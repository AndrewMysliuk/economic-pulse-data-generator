package core

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
)

func BuildHistoryFromDir(outputDir string, windowDays int) (schema.History, error) {
	files, err := filepath.Glob(filepath.Join(outputDir, "*.json"))
	if err != nil {
		return schema.History{}, err
	}

	type rec struct {
		path string
		date string
		t    time.Time
	}
	var list []rec
	for _, p := range files {
		base := filepath.Base(p)
		if len(base) < 15 || !strings.HasSuffix(base, ".json") {
			continue
		}
		d := base[:10]
		t, e := time.Parse("2006-01-02", d)
		if e == nil {
			list = append(list, rec{path: p, date: d, t: t})
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].t.Before(list[j].t) })

	h := schema.NewHistory(windowDays)

	for _, r := range list {
		var day schema.DailyData
		b, err := os.ReadFile(r.path)
		if err != nil {
			return schema.History{}, err
		}
		if err := json.Unmarshal(b, &day); err != nil {
			return schema.History{}, err
		}

		if day.Date == "" {
			day.Date = r.date
		}

		h.AppendDay(day)
	}

	h.TrimToWindow()
	return h, nil
}

func UpdateHistoryIncremental(historyPath, outputDir string) error {
	const windowDays = 180

	h, err := schema.LoadHistory(historyPath)
	if err != nil {
		h, err = BuildHistoryFromDir(outputDir, windowDays)
		if err != nil {
			return err
		}
		return schema.SaveHistory(h, historyPath)
	}

	today := time.Now().UTC().Format("2006-01-02")
	dayFile := filepath.Join(outputDir, today+".json")
	b, err := os.ReadFile(dayFile)
	if err != nil {
		return nil
	}

	var day schema.DailyData
	if err := json.Unmarshal(b, &day); err != nil {
		return err
	}

	h.AppendDay(day)
	h.WindowDays = windowDays
	h.TrimToWindow()

	return schema.SaveHistory(h, historyPath)
}
