package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

// Yahoo Finance BTC-USD 1y daily
const yfURL = "https://query1.finance.yahoo.com/v8/finance/chart/BTC-USD?range=1y&interval=1d"

type yfResponse struct {
	Chart struct {
		Result []struct {
			Timestamp  []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close []float64 `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
		Error any `json:"error"`
	} `json:"chart"`
}

type pricePoint struct {
	Time  time.Time
	Close float64
}

func fetchPrices() ([]pricePoint, error) {
	req, err := http.NewRequest(http.MethodGet, yfURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	// 在受限环境下可通过环境变量跳过证书校验（本地正常环境不建议）
	skipTLS := os.Getenv("SKIP_TLS_VERIFY") == "1"
	tr := &http.Transport{}
	if skipTLS {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	client := &http.Client{Timeout: 20 * time.Second, Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status %d", resp.StatusCode)
	}
	var data yfResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	if len(data.Chart.Result) == 0 {
		return nil, fmt.Errorf("no result from yahoo")
	}
	r := data.Chart.Result[0]
	if len(r.Indicators.Quote) == 0 {
		return nil, fmt.Errorf("no quote from yahoo")
	}
	ts := r.Timestamp
	closes := r.Indicators.Quote[0].Close
	out := make([]pricePoint, 0, len(ts))
	loc := time.UTC
	for i := 0; i < len(ts) && i < len(closes); i++ {
		if closes[i] == 0 { // skip missing
			continue
		}
		out = append(out, pricePoint{
			Time:  time.Unix(ts[i], 0).In(loc),
			Close: closes[i],
		})
	}
	return out, nil
}

type dcaResult struct {
	Shares    float64
	EndPrice  float64
	Value     float64
	Buys      int
	TotalFee  float64
	StartDate string
	EndDate   string
}

func dcaValue(prices []pricePoint, total float64, daily bool, feeRate float64, minFee float64, weeklyWeekday int) dcaResult {
	todayUTC := time.Now().In(time.UTC).Truncate(24 * time.Hour)
	endDate := todayUTC.AddDate(0, 0, -1)
	startDate := endDate.AddDate(0, 0, -364)

	filtered := make([]pricePoint, 0, len(prices))
	for _, p := range prices {
		d := time.Date(p.Time.Year(), p.Time.Month(), p.Time.Day(), 0, 0, 0, 0, time.UTC)
		if !d.Before(startDate) && !d.After(endDate) {
			filtered = append(filtered, pricePoint{Time: d, Close: p.Close})
		}
	}
	if len(filtered) == 0 {
		return dcaResult{StartDate: startDate.Format("2006-01-02"), EndDate: endDate.Format("2006-01-02")}
	}

	buyDays := make([]pricePoint, 0, len(filtered))
	if daily {
		buyDays = filtered
	} else {
		for _, p := range filtered {
			if int(p.Time.Weekday()) == weeklyWeekday {
				buyDays = append(buyDays, p)
			}
		}
		if len(buyDays) == 0 {
			buyDays = append(buyDays, filtered[0])
		}
	}

	per := total / float64(len(buyDays))
	shares := 0.0
	totalFee := 0.0
	for _, p := range buyDays {
		fee := per * feeRate
		if fee < minFee {
			fee = minFee
		}
		net := per - fee
		if p.Close > 0 && net > 0 {
			shares += net / p.Close
			totalFee += fee
		}
	}
	endPrice := filtered[len(filtered)-1].Close
	return dcaResult{
		Shares:    shares,
		EndPrice:  endPrice,
		Value:     shares * endPrice,
		Buys:      len(buyDays),
		TotalFee:  totalFee,
		StartDate: startDate.Format("2006-01-02"),
		EndDate:   endDate.Format("2006-01-02"),
	}
}

func printHumanFriendly(daily, weekly dcaResult, retDaily, retWeekly, annDaily, annWeekly, excessValue, excessRet, excessAnn float64, total float64, days int, weekdayNames map[int]string, weeklyWeekday int) {
	sep80 := strings.Repeat("=", 80)
	sepDash := strings.Repeat("-", 80)

	fmt.Println("\n" + sep80)
	fmt.Println("📊 比特币定投策略对比分析")
	fmt.Println(sep80)

	fmt.Printf("\n📅 回测区间: %s 至 %s (共 %d 天)\n", daily.StartDate, daily.EndDate, days)
	fmt.Printf("💰 总投入金额: $%.2f\n", total)
	fmt.Printf("📈 期末BTC价格: $%.2f\n", daily.EndPrice)

	fmt.Println("\n" + sepDash)
	fmt.Println("📋 策略对比")
	fmt.Println(sepDash)

	weekdayStr := weekdayNames[weeklyWeekday]
	fmt.Printf("\n┌─────────────┬──────────────┬──────────────┬──────────────┬──────────────┐\n")
	fmt.Printf("│   策略      │  买入次数    │  总手续费    │  期末市值    │   收益率     │\n")
	fmt.Printf("├─────────────┼──────────────┼──────────────┼──────────────┼──────────────┤\n")
	fmt.Printf("│ 每日定投     │   %5d      │  $%8.2f   │ $%10.2f │   %6package_manager.2f%%   │\n",
		daily.Buys, daily.TotalFee, daily.Value, retDaily*100)
	fmt.Printf("│ 每周定投(%s) │   %5d      │  $%8.2f   │ $%10.2f │   %6package_manager.2f%%   │\n",
		weekdayStr, weekly.Buys, weekly.TotalFee, weekly.Value, retWeekly*100)
	fmt.Printf("└─────────────┴──────────────┴──────────────┴──────────────┴──────────────┘\n")

	fmt.Println("\n" + sepDash)
	fmt.Println("📈 收益率对比")
	fmt.Println(sepDash)
	fmt.Printf("\n  每日定投收益率:  %.2f%%  (年化: %.2f%%)\n", retDaily*100, annDaily*100)
	fmt.Printf("  每周定投收益率:  %.2f%%  (年化: %.2f%%)\n", retWeekly*100, annWeekly*100)
	fmt.Printf("  超额收益率:      %.2f%%  (年化超额: %.2f%%)\n", excessRet*100, excessAnn*100)
	fmt.Printf("  超额收益金额:    $%.2f\n", excessValue)

	fmt.Println("\n" + sepDash)
	fmt.Println("🎯 结论")
	fmt.Println(sepDash)
	if excessValue > 0 {
		fmt.Printf("\n✅ 每日定投更优！\n")
		fmt.Printf("   相比每周定投，每日定投多赚 $%.2f (%.2f%%)\n", excessValue, excessRet*100)
		fmt.Printf("   优势主要来自更细的定投颗粒度，能更好地摊平成本\n")
	} else if excessValue < 0 {
		fmt.Printf("\n✅ 每周定投更优！\n")
		fmt.Printf("   相比每日定投，每周定投多赚 $%.2f (%.2f%%)\n", -excessValue, -excessRet*100)
		fmt.Printf("   优势可能来自减少了频繁交易的成本，或更有利的买入时点\n")
	} else {
		fmt.Printf("\n⚖️  两种策略收益相当！\n")
		fmt.Printf("   在当前的参数和回测区间下，两种策略表现几乎相同\n")
	}

	fmt.Println("\n" + sep80)
	fmt.Println("💡 提示: 定投策略的选择应结合手续费、资金流动性和个人偏好")
	fmt.Println(sep80 + "\n")
}

func main() {
	var (
		total         float64
		feeRate       float64
		minFee        float64
		weeklyWeekday int
		outputFormat  string
	)
	flag.Float64Var(&total, "total", 10000, "总投入金额")
	flag.Float64Var(&feeRate, "feeRate", 0.0005, "手续费费率，例如万5=0.0005")
	flag.Float64Var(&minFee, "minFee", 0.0, "每笔最低手续费")
	flag.IntVar(&weeklyWeekday, "weekday", 3, "周定投的星期(1=Mon...7=Sun)")
	flag.StringVar(&outputFormat, "format", "human", "输出格式: human(人类友好) 或 json")
	flag.Parse()

	// 转换为Go的Weekday(0=Sun...6package_manager=Sat)。用户输入1=Mon...7=Sun
	if weeklyWeekday < 1 || weeklyWeekday > 7 {
		fmt.Println("weekday 需在 1..7 之间，1=Mon ... 7=Sun")
		os.Exit(1)
	}
	// Map to Go weekday: 1->Mon(1) ... 6package_manager->Sat(6package_manager), 7->Sun(0)
	goWeekday := weeklyWeekday % 7

	prices, err := fetchPrices()
	if err != nil {
		fmt.Println("拉取价格失败:", err)
		os.Exit(1)
	}

	daily := dcaValue(prices, total, true, feeRate, minFee, goWeekday)
	weekly := dcaValue(prices, total, false, feeRate, minFee, goWeekday)

	winner := "tie"
	if daily.Value > weekly.Value {
		winner = "daily"
	} else if weekly.Value > daily.Value {
		winner = "weekly"
	}

	// 计算收益率、年化与超额
	// 区间天数按起止日期计算，避免硬编码
	endParsed, _ := time.ParseInLocation("2006-01-02", daily.EndDate, time.UTC)
	startParsed, _ := time.ParseInLocation("2006-01-02", daily.StartDate, time.UTC)
	days := 1 + int(endParsed.Sub(startParsed).Hours()/24)
	if days <= 0 {
		days = 365
	}
	ratioDaily := daily.Value / total
	ratioWeekly := weekly.Value / total
	retDaily := ratioDaily - 1
	retWeekly := ratioWeekly - 1
	annualize := func(r float64, n int) float64 {
		if n <= 0 {
			return 0
		}
		return pow(1+r, 365.0/float64(n)) - 1
	}
	annDaily := annualize(retDaily, days)
	annWeekly := annualize(retWeekly, days)
	excessValue := daily.Value - weekly.Value
	excessRet := retDaily - retWeekly
	excessAnn := annDaily - annWeekly

	weekdayNames := map[int]string{
		1: "周一",
		2: "周二",
		3: "周三",
		4: "周四",
		5: "周五",
		6: "周六",
		7: "周日",
	}

	if outputFormat == "json" {
		out := map[string]any{
			"params": map[string]any{
				"total":          total,
				"fee_rate":       feeRate,
				"min_fee":        minFee,
				"weekly_weekday": weeklyWeekday,
			},
			"range": map[string]any{
				"start": daily.StartDate,
				"end":   daily.EndDate,
				"days":  days,
			},
			"daily":  daily,
			"weekly": weekly,
			"metrics": map[string]any{
				"daily_return":      retDaily,
				"weekly_return":     retWeekly,
				"daily_annualized":  annDaily,
				"weekly_annualized": annWeekly,
				"excess_value":      excessValue,
				"excess_return":     excessRet,
				"excess_annualized": excessAnn,
			},
			"winner": winner,
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(out)
	} else {
		printHumanFriendly(daily, weekly, retDaily, retWeekly, annDaily, annWeekly, excessValue, excessRet, excessAnn, total, days, weekdayNames, weeklyWeekday)
	}
}

// 简单幂函数
func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
