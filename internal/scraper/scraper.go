package scraper

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"

	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/schema/enum/metric_unit"
	"github.com/AndrewMysliuk/economic-pulse-data-generator/internal/utils/parsed_links"
)

func Scrape(countryISO string) (schema.CountryMetrics, error) {
	cfg, ok := parsed_links.Countries[countryISO]
	if !ok {
		return schema.CountryMetrics{}, fmt.Errorf("unknown country %q", countryISO)
	}

	browserCtx, cancelBrowser, err := newBrowser()
	if err != nil {
		return schema.CountryMetrics{}, err
	}
	defer cancelBrowser()

	runCtx, cancelRun := context.WithTimeout(browserCtx, 5*time.Minute)
	defer cancelRun()

	result := schema.InitEmptyCountryMetrics()
	r := newRunner(runCtx, &result)

	r.scrape("PolicyRate", cfg.PolicyRate)
	r.scrape("Inflation", cfg.Inflation)
	r.scrape("Unemployment", cfg.Unemployment)
	r.scrape("PMI", cfg.PMI)
	r.scrape("EquityIndex", cfg.EquityIndex)

	for pair, ml := range cfg.Currencies {
		r.scrape("FX "+pair, ml)
	}

	r.scrape("Bond10Y", cfg.Bond10Y)

	result.PolicyRate.ComputeAverage()
	result.Inflation.ComputeAverage()
	result.Unemployment.ComputeAverage()
	result.PMI.ComputeAverage()
	result.EquityIndex.ComputeAverage()
	result.CurrencyIndex.ComputeAverage()
	result.BondYield10Y.ComputeAverage()

	return result, nil
}

// ------------------------- Browser bootstrap -------------------------

func newBrowser() (context.Context, context.CancelFunc, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)

	browserCtx, browserCancel := chromedp.NewContext(allocCtx)

	if err := configureNetwork(browserCtx); err != nil {
		allocCancel()
		browserCancel()
		return nil, nil, err
	}

	cancel := func() {
		browserCancel()
		allocCancel()
	}
	return browserCtx, cancel, nil
}

func configureNetwork(ctx context.Context) error {
	const ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(c context.Context) error {
			if err := network.Enable().Do(c); err != nil {
				return err
			}

			if err := emulation.SetUserAgentOverride(ua).Do(c); err != nil {
				return err
			}

			if err := network.SetExtraHTTPHeaders(network.Headers{
				"Accept-Language": "en-US,en;q=0.9",
				"DNT":             "1",
			}).Do(c); err != nil {
				return err
			}

			if err := network.SetBlockedURLs([]string{
				"*.png", "*.jpg", "*.jpeg", "*.gif", "*.webp", "*.svg",
				"*.woff", "*.woff2", "*.ttf", "*.otf", "*.mp4", "*.webm",
			}).Do(c); err != nil {
				return err
			}

			if err := emulation.SetLocaleOverride().Do(c); err != nil {
				return err
			}
			if err := emulation.SetTimezoneOverride("Europe/Sofia").Do(c); err != nil {
				return err
			}

			if err := emulation.SetDeviceMetricsOverride(1440, 900, 1.0, false).Do(c); err != nil {
				return err
			}
			return nil
		}),
	)
}

// ------------------------- Runner (rate-limit + retry) -------------------------

type runner struct {
	root   context.Context
	result *schema.CountryMetrics
	asOf   string
}

func newRunner(root context.Context, dst *schema.CountryMetrics) *runner {
	return &runner{
		root:   root,
		result: dst,
		asOf:   time.Now().UTC().Format(time.RFC3339),
	}
}

func (r *runner) scrape(name string, ml parsed_links.MetricLink) {
	if ml.URL == "" || ml.Selector == "" {
		return
	}
	var val float64
	err := withRetry(3, 700*time.Millisecond, func() error {
		tabCtx, tabCancel := chromedp.NewContext(r.root)
		defer tabCancel()

		ctx, cancel := context.WithTimeout(tabCtx, 30*time.Second)
		defer cancel()

		v, err := fetchValueChromedp(ctx, ml.URL, ml.Selector)
		if err != nil {
			return err
		}
		val = v
		return nil
	})
	if err != nil {
		log.Printf("%-15s ERROR: %v", name, err)
		return
	}

	switch {
	case name == "PolicyRate":
		r.result.PolicyRate.Sources = []schema.MetricSource{
			r.src(val, metric_unit.RatePct, ml.URL, "Policy rate"),
		}

	case name == "Inflation":
		r.result.Inflation.Sources = []schema.MetricSource{
			r.src(val, metric_unit.Percent, ml.URL, "CPI YoY"),
		}

	case name == "Unemployment":
		r.result.Unemployment.Sources = []schema.MetricSource{
			r.src(val, metric_unit.Percent, ml.URL, "Unemployment"),
		}

	case name == "PMI":
		r.result.PMI.Sources = []schema.MetricSource{
			r.src(val, metric_unit.Index, ml.URL, "PMI / Business confidence"),
		}

	case name == "EquityIndex":
		r.result.EquityIndex.Sources = []schema.MetricSource{
			r.src(val, metric_unit.Index, ml.URL, "Equity index"),
		}
		r.result.EquityIndex.AsOf = r.asOf

	case name == "Bond10Y":
		r.result.BondYield10Y.Sources = []schema.MetricSource{
			r.src(val, metric_unit.Percent, ml.URL, "10Y Gov Yield"),
		}
		r.result.BondYield10Y.AsOf = r.asOf

	case strings.HasPrefix(name, "FX "):
		pair := strings.TrimPrefix(name, "FX ")
		fx := schema.FxRate{
			Pair:       pair,
			Value:      val,
			AsOf:       r.asOf,
			SourceUrl:  ml.URL,
			SourceName: lastPathSegment(ml.URL),
		}
		_ = fx.Validate()
		r.result.FxRates = append(r.result.FxRates, fx)
	}

	sleepJitter(700*time.Millisecond, 200*time.Millisecond, 400*time.Millisecond)
}

func (r *runner) src(val float64, unit metric_unit.MetricUnit, url, name string) schema.MetricSource {
	v := val
	return schema.MetricSource{
		Value:      &v,
		Date:       r.asOf[:10],
		Unit:       unit,
		SourceUrl:  url,
		SourceName: name,
	}
}

func withRetry(attempts int, baseBackoff time.Duration, fn func() error) error {
	backoff := baseBackoff
	for i := 0; i < attempts; i++ {
		if err := fn(); err != nil {
			if i == attempts-1 {
				return fmt.Errorf("retry failed: %w", err)
			}
			time.Sleep(backoff + time.Duration(rand.Intn(400))*time.Millisecond)
			backoff *= 2
			continue
		}
		return nil
	}
	return nil
}

func sleepJitter(base, minJ, maxJ time.Duration) {
	j := time.Duration(int64(minJ) + rand.Int63n(int64(maxJ)))
	time.Sleep(base + j)
}

// ------------------------- Single metric fetch -------------------------

func fetchValueChromedp(ctx context.Context, url, selector string) (float64, error) {
	var text string
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.Text(selector, &text, chromedp.ByQuery),
	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		return 0, err
	}
	text = strings.TrimSpace(text)
	if text == "" {
		return 0, fmt.Errorf("empty text for selector")
	}
	return normalizeNumber(text)
}

// ------------------------- Number parsing -------------------------

func normalizeNumber(s string) (float64, error) {
	repl := strings.NewReplacer(
		"%", "", " ", "", "\u00A0", "", "\u202F", "", "\u2009", "",
	)
	s = repl.Replace(s)

	var b strings.Builder
	for _, r := range s {
		switch r {
		case '-', '.', ',', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			b.WriteRune(r)
		default:
		}
	}
	t := b.String()

	if strings.Count(t, ",") == 1 && strings.Count(t, ".") > 1 {
		t = strings.ReplaceAll(t, ".", "")
		t = strings.ReplaceAll(t, ",", ".")
	} else {
		t = strings.ReplaceAll(t, ",", "")
	}
	val, err := strconv.ParseFloat(t, 64)
	if err != nil {
		return 0, fmt.Errorf("parse float from %q: %w", s, err)
	}
	return val, nil
}

// ------------------------- URL helpers -------------------------

func lastPathSegment(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	return path.Base(u.Path)
}
