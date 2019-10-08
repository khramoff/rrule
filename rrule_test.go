package rrule

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var now = time.Date(2018, 8, 25, 9, 8, 7, 6, time.UTC) // it's a saturday

var cases = []struct {
	Name     string
	String   string
	RRule    RRule
	Dates    []string
	Terminal bool

	NoBenchmark            bool
	NoTest                 bool
	NoTeambitionComparison bool
}{
	{
		Name: "simple secondly",
		RRule: RRule{
			Frequency: Secondly,
			Count:     3,
			Dtstart:   now,
		},
		Dates:    []string{"2018-08-25T09:08:07Z", "2018-08-25T09:08:08Z", "2018-08-25T09:08:09Z"},
		Terminal: true,
	},
	{
		Name: "simple minutely",
		RRule: RRule{
			Frequency: Minutely,
			Count:     3,
			Dtstart:   now,
		},
		Dates:    []string{"2018-08-25T09:08:07Z", "2018-08-25T09:09:07Z", "2018-08-25T09:10:07Z"},
		Terminal: true,
	},

	{
		Name: "simple hourly",
		RRule: RRule{
			Frequency: Hourly,
			Count:     3,
			Dtstart:   now,
		},
		Dates:    []string{"2018-08-25T09:08:07Z", "2018-08-25T10:08:07Z", "2018-08-25T11:08:07Z"},
		Terminal: true,
	},

	{
		Name: "simple daily",
		RRule: RRule{
			Frequency: Daily,
			Count:     3,
			Dtstart:   now,
		},
		Dates:    []string{"2018-08-25T09:08:07Z", "2018-08-26T09:08:07Z", "2018-08-27T09:08:07Z"},
		Terminal: true,
	},

	{
		Name: "secondly setpos",
		RRule: RRule{
			Frequency: Secondly,
			Count:     4,
			Dtstart:   now,
			BySeconds: []int{1, 2, 3},
			ByMonths:  []time.Month{time.August, time.September},
			BySetPos:  []int{1, 3, -1},
		},
		Dates:    []string{"2018-08-25T09:09:01Z", "2018-08-25T09:09:02Z", "2018-08-25T09:09:03Z", "2018-08-25T09:10:01Z"},
		Terminal: true,
	},
	{
		Name: "minutely setpos",
		RRule: RRule{
			Frequency: Minutely,
			Count:     4,
			Dtstart:   now,
			BySeconds: []int{1, 2, 3},
			ByMonths:  []time.Month{time.August, time.September},
			BySetPos:  []int{1, 3, -1},
		},
		Dates:    []string{"2018-08-25T09:09:01Z", "2018-08-25T09:09:03Z", "2018-08-25T09:10:01Z", "2018-08-25T09:10:03Z"},
		Terminal: true,
	},

	{
		Name: "hourly setpos",
		RRule: RRule{
			Frequency: Hourly,
			Count:     4,
			Dtstart:   now,
			ByMinutes: []int{1, 2, 3},
			ByMonths:  []time.Month{time.August, time.September},
			BySetPos:  []int{1, 3, -1},
		},
		Dates:    []string{"2018-08-25T10:01:07Z", "2018-08-25T10:03:07Z", "2018-08-25T11:01:07Z", "2018-08-25T11:03:07Z"},
		Terminal: true,
	},

	{
		Name: "daily setpos",
		RRule: RRule{
			Frequency: Daily,
			Count:     4,
			Dtstart:   now,
			ByHours:   []int{1, 2, 3},
			ByMonths:  []time.Month{time.August, time.September},
			BySetPos:  []int{1, 3, -1},
		},
		Dates:    []string{"2018-08-26T01:08:07Z", "2018-08-26T03:08:07Z", "2018-08-27T01:08:07Z", "2018-08-27T03:08:07Z"},
		Terminal: true,
	},
	{
		Name:   "weekly setpos",
		String: "FREQ=WEEKLY;COUNT=4;BYHOUR=1,2,3;BYMONTH=8,9;BYSETPOS=1,3,-1",
		RRule: RRule{
			Frequency: Weekly,
			Count:     4,
			Dtstart:   now,
			ByHours:   []int{1, 2, 3},
			ByMonths:  []time.Month{time.August, time.September},
			BySetPos:  []int{1, 3, -1},
		},
		Dates:    []string{"2018-09-01T01:08:07Z", "2018-09-01T03:08:07Z", "2018-09-08T01:08:07Z", "2018-09-08T03:08:07Z"},
		Terminal: true,
	},

	{
		Name: "monthly setpos",
		RRule: RRule{
			Frequency:  Monthly,
			ByWeekdays: []QualifiedWeekday{{N: 0, WD: time.Monday}, {N: 0, WD: time.Tuesday}, {N: 0, WD: time.Wednesday}, {N: 0, WD: time.Thursday}, {N: 0, WD: time.Friday}, {N: 0, WD: time.Saturday}, {N: 0, WD: time.Sunday}},
			Count:      4,
			Dtstart:    now,
			ByMonths:   []time.Month{time.August, time.September},
			BySetPos:   []int{1, 3, -1},
		},
		Dates:    []string{"2018-08-31T09:08:07Z", "2018-09-01T09:08:07Z", "2018-09-03T09:08:07Z", "2018-09-30T09:08:07Z"},
		Terminal: true,
	},

	{
		Name: "yearly setpos",
		RRule: RRule{
			Frequency:  Yearly,
			ByWeekdays: []QualifiedWeekday{{N: 0, WD: time.Monday}, {N: 0, WD: time.Tuesday}, {N: 0, WD: time.Wednesday}, {N: 0, WD: time.Thursday}, {N: 0, WD: time.Friday}, {N: 0, WD: time.Saturday}, {N: 0, WD: time.Sunday}},
			Count:      4,
			Dtstart:    now,
			ByMonths:   []time.Month{time.August, time.September},
			BySetPos:   []int{1, 3, -1},
		},
		String:   "FREQ=YEARLY;COUNT=4;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYMONTH=8,9;BYSETPOS=1,3,-1",
		Dates:    []string{"2018-09-30T09:08:07Z", "2019-08-01T09:08:07Z", "2019-08-03T09:08:07Z", "2019-09-30T09:08:07Z"},
		Terminal: true,
	},

	{
		Name: "daily until",
		RRule: RRule{
			Frequency: Daily,
			Until:     time.Date(2018, 8, 30, 0, 0, 0, 0, time.UTC),
			Dtstart:   now,
		},
		Dates:    []string{"2018-08-25T09:08:07Z", "2018-08-26T09:08:07Z", "2018-08-27T09:08:07Z", "2018-08-28T09:08:07Z", "2018-08-29T09:08:07Z"},
		Terminal: true,
		String:   "FREQ=DAILY;UNTIL=20180830T000000Z",
	},

	{
		Name: "daily until floating",
		RRule: RRule{
			Frequency:     Daily,
			Until:         time.Date(2018, 8, 30, 0, 0, 0, 0, time.UTC),
			UntilFloating: true,
			Dtstart:       now,
		},
		Dates:    []string{"2018-08-25T09:08:07Z", "2018-08-26T09:08:07Z", "2018-08-27T09:08:07Z", "2018-08-28T09:08:07Z", "2018-08-29T09:08:07Z"},
		Terminal: true,
		String:   "FREQ=DAILY;UNTIL=20180830T000000",
	},

	{
		Name: "simple monthly",
		RRule: RRule{
			Frequency: Monthly,
			Count:     3,
			Dtstart:   now,
		},
		Dates:    []string{"2018-08-25T09:08:07Z", "2018-09-25T09:08:07Z", "2018-10-25T09:08:07Z"},
		Terminal: true,
	},

	{
		Name: "long monthly",
		RRule: RRule{
			Frequency: Monthly,
			Count:     300,
			Dtstart:   now,
		},
		Terminal: true,
		NoTest:   true,
	},

	{
		Name: "monthly by weekday",
		RRule: RRule{
			Frequency:  Monthly,
			Count:      3,
			Dtstart:    now,
			ByWeekdays: []QualifiedWeekday{{N: 1, WD: time.Tuesday}},
		},
		Dates:    []string{"2018-09-04T09:08:07Z", "2018-10-02T09:08:07Z", "2018-11-06T09:08:07Z"},
		Terminal: true,
	},

	{
		Name: "simple weekly",
		RRule: RRule{
			Frequency: Weekly,
			Count:     3,
			Dtstart:   now,
		},
		Dates:    []string{"2018-08-25T09:08:07Z", "2018-09-01T09:08:07Z", "2018-09-08T09:08:07Z"},
		Terminal: true,
	},

	{
		Name:   "weekly by weekday",
		String: "FREQ=WEEKLY;COUNT=3;BYDAY=TU",
		RRule: RRule{
			Frequency:  Weekly,
			Count:      3,
			Dtstart:    now,
			ByWeekdays: []QualifiedWeekday{{WD: time.Tuesday}},
		},
		Dates:    []string{"2018-08-28T09:08:07Z", "2018-09-04T09:08:07Z", "2018-09-11T09:08:07Z"},
		Terminal: true,
	},

	{
		Name:   "yearly by weekday",
		String: "FREQ=YEARLY;COUNT=4;BYDAY=TU,35WE,-17MO",
		RRule: RRule{
			Frequency:  Yearly,
			Count:      4,
			Dtstart:    now,
			ByWeekdays: []QualifiedWeekday{{WD: time.Tuesday}, {N: 35, WD: time.Wednesday}, {N: -17, WD: time.Monday}},
		},
		Dates:    []string{"2018-08-28T09:08:07Z", "2018-08-29T09:08:07Z", "2018-09-04T09:08:07Z", "2018-09-10T09:08:07Z"},
		Terminal: true,

		// I'm not sure if I'm reading the spec wrong or if they have a bug, but they return no results.
		// lib-recur agrees with my implementation.
		NoTeambitionComparison: true,
	},

	{
		Name: "monthly by monthday",
		RRule: RRule{
			Frequency:   Monthly,
			Count:       3,
			Dtstart:     now,
			ByMonthDays: []int{10},
		},
		Dates:    []string{"2018-09-10T09:08:07Z", "2018-10-10T09:08:07Z", "2018-11-10T09:08:07Z"},
		Terminal: true,
	},

	{
		Name: "daily daylight savings",
		RRule: RRule{
			Frequency: Daily,
			Count:     3,
			Dtstart:   time.Date(2018, time.November, 03, 01, 00, 00, 00, NewYork()),
		},
		Dates:    []string{"2018-11-03T01:00:00-04:00", "2018-11-04T01:00:00-04:00", "2018-11-05T01:00:00-05:00"},
		Terminal: true,
	},

	{
		Name:   "no daily daylight savings",
		String: "FREQ=DAILY;COUNT=3",
		RRule: RRule{
			Frequency: Daily,
			Count:     3,
			Dtstart:   time.Date(2018, time.November, 03, 01, 00, 00, 00, Phoenix()),
		},
		Dates:    []string{"2018-11-03T01:00:00-07:00", "2018-11-04T01:00:00-07:00", "2018-11-05T01:00:00-07:00"},
		Terminal: true,
	},

	{
		Name: "half-hourly daylight savings",
		RRule: RRule{
			Frequency: Hourly,
			Count:     6,
			Dtstart:   time.Date(2018, time.November, 04, 00, 30, 00, 00, NewYork()),
			ByMinutes: []int{0, 30},
		},
		Dates:                  []string{"2018-11-04T00:30:00-04:00", "2018-11-04T01:00:00-04:00", "2018-11-04T01:30:00-04:00", "2018-11-04T01:00:00-05:00", "2018-11-04T01:30:00-05:00", "2018-11-04T02:00:00-05:00"},
		Terminal:               true,
		NoTeambitionComparison: true,
	},

	{
		Name: "rfc: Monthly on the first Friday until December 24, 1997",
		RRule: RRule{
			Frequency:  Monthly,
			Until:      time.Date(1997, time.December, 24, 0, 0, 0, 0, time.UTC),
			ByWeekdays: []QualifiedWeekday{{WD: time.Friday, N: 1}},
			Dtstart:    time.Date(1997, time.September, 5, 9, 0, 0, 0, NewYork()),
		},
		String:   "FREQ=MONTHLY;UNTIL=19971224T000000Z;BYDAY=1FR",
		Dates:    []string{"1997-09-05T09:00:00-04:00", "1997-10-03T09:00:00-04:00", "1997-11-07T09:00:00-05:00", "1997-12-05T09:00:00-05:00"},
		Terminal: true,
	},

	{
		Name: "end of time",
		RRule: RRule{
			Frequency: Yearly,
			Dtstart:   time.Date(219248495, time.December, 7, 0, 0, 0, 0, time.UTC),
		},
		String: "FREQ=YEARLY",
		Dates: []string{
			"219248495-12-07T00:00:00Z",
			"219248496-12-07T00:00:00Z",
			"219248497-12-07T00:00:00Z",
			"219248498-12-07T00:00:00Z",
		},
		NoTeambitionComparison: true,
	},

	{
		Name: "leap day monthly omit",
		RRule: RRule{
			Frequency: Monthly,
			Dtstart:   time.Date(2019, time.August, 29, 0, 0, 0, 0, time.UTC),
			Interval:  6,
			Count:     4,
		},
		String:   "FREQ=MONTHLY;COUNT=4;INTERVAL=6",
		Terminal: true,
		Dates: []string{
			"2019-08-29T00:00:00Z",
			"2020-02-29T00:00:00Z",
			"2020-08-29T00:00:00Z",
			"2021-08-29T00:00:00Z",
		},
	},

	{
		Name: "leap day monthly prev",
		RRule: RRule{
			Frequency:       Monthly,
			Dtstart:         time.Date(2019, time.August, 29, 0, 0, 0, 0, time.UTC),
			Interval:        6,
			Count:           4,
			InvalidBehavior: PrevInvalid,
		},
		String:   "FREQ=MONTHLY;COUNT=4;INTERVAL=6;SKIP=BACKWARD;RSCALE=GREGORIAN",
		Terminal: true,
		Dates: []string{
			"2019-08-29T00:00:00Z",
			"2020-02-29T00:00:00Z",
			"2020-08-29T00:00:00Z",
			"2021-02-28T00:00:00Z",
		},
		NoTeambitionComparison: true,
	},

	{
		Name: "leap day monthly next",
		RRule: RRule{
			Frequency:       Monthly,
			Dtstart:         time.Date(2019, time.August, 29, 0, 0, 0, 0, time.UTC),
			Interval:        6,
			Count:           4,
			InvalidBehavior: NextInvalid,
		},
		String:   "FREQ=MONTHLY;COUNT=4;INTERVAL=6;SKIP=FORWARD;RSCALE=GREGORIAN",
		Terminal: true,
		Dates: []string{
			"2019-08-29T00:00:00Z",
			"2020-02-29T00:00:00Z",
			"2020-08-29T00:00:00Z",
			"2021-03-01T00:00:00Z",
		},
		NoTeambitionComparison: true,
	},

	{
		Name: "leap year day 366 omit",
		RRule: RRule{
			Frequency:  Yearly,
			Dtstart:    time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC),
			Count:      5,
			ByYearDays: []int{366},
		},
		String:   "FREQ=YEARLY;COUNT=5;BYYEARDAY=366",
		Terminal: true,
		Dates: []string{
			"2016-12-31T00:00:00Z",
			"2020-12-31T00:00:00Z",
			"2024-12-31T00:00:00Z",
			"2028-12-31T00:00:00Z",
			"2032-12-31T00:00:00Z",
		},
	},

	{
		Name: "leap year day 366 next",
		RRule: RRule{
			Frequency:       Yearly,
			Dtstart:         time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC),
			Count:           5,
			ByYearDays:      []int{366},
			InvalidBehavior: NextInvalid,
		},
		String:   "FREQ=YEARLY;COUNT=5;BYYEARDAY=366;SKIP=FORWARD;RSCALE=GREGORIAN",
		Terminal: true,
		Dates: []string{
			"2016-12-31T00:00:00Z",
			"2018-01-01T00:00:00Z",
			"2019-01-01T00:00:00Z",
			"2020-01-01T00:00:00Z",
			"2020-12-31T00:00:00Z",
		},
		NoTeambitionComparison: true,
	},

	{
		Name: "leap year day 366 prev",
		RRule: RRule{
			Frequency:       Yearly,
			Dtstart:         time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC),
			Count:           5,
			ByYearDays:      []int{366},
			InvalidBehavior: PrevInvalid,
		},
		String:   "FREQ=YEARLY;COUNT=5;BYYEARDAY=366;SKIP=BACKWARD;RSCALE=GREGORIAN",
		Terminal: true,
		Dates: []string{
			"2016-12-31T00:00:00Z",
			"2017-12-31T00:00:00Z",
			"2018-12-31T00:00:00Z",
			"2019-12-31T00:00:00Z",
			"2020-12-31T00:00:00Z",
		},
		NoTeambitionComparison: true,
	},

	{
		Name: "rfc weekno",
		RRule: RRule{
			Frequency:     Yearly,
			Dtstart:       time.Date(1997, 5, 12, 9, 0, 0, 0, NewYork()),
			Count:         3,
			ByWeekNumbers: []int{20},
			ByWeekdays:    []QualifiedWeekday{{WD: time.Monday}},
		},
		String: "FREQ=YEARLY;COUNT=3;BYDAY=MO;BYWEEKNO=20",
		Dates: []string{
			"1997-05-12T09:00:00-04:00",
			"1998-05-11T09:00:00-04:00",
			"1999-05-17T09:00:00-04:00",
		},
	},
}

func MustRRule(str string) RRule {
	r, err := ParseRRule(str)
	if err != nil {
		panic(err)
	}
	return r
}

func NewYork() *time.Location {
	return mustLoadLoc("America/New_York")
}

func Phoenix() *time.Location {
	return mustLoadLoc("America/Phoenix")
}

func mustLoadLoc(loc string) *time.Location {
	ny, err := time.LoadLocation(loc)
	if ny == nil {
		errStr := "not found"
		if err != nil {
			errStr = err.Error()
		}

		panic("error loading IANA tzdata, which is required for daylight savings tests: " + errStr)
	}
	return ny
}

func TestRRule(t *testing.T) {
	for _, tc := range cases {
		if tc.NoTest {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			if tc.String != "" {
				t.Log(tc.String)

				parsed, err := ParseRRule(tc.String)
				require.NoError(t, err)
				require.NotNil(t, parsed)

				// unset dtstart for the comparisons, because it's only used operationally.
				// it's set on the test cases because we need it to run them.
				dtstart := tc.RRule.Dtstart
				tc.RRule.Dtstart = time.Time{}
				assert.Equal(t, tc.String, tc.RRule.String(), "RRule does not render to the correct string")
				assert.Equal(t, tc.RRule, parsed)

				tc.RRule.Dtstart = dtstart.Truncate(time.Second) // restore dtstart, but truncate it because rrule only operates at second.
			}

			dates := All(tc.RRule.Iterator(), 0)
			assert.Equal(t, tc.Dates, rfcAll(dates))
		})
	}
}

func BenchmarkRRule(b *testing.B) {
	for _, tc := range cases {
		if tc.NoBenchmark {
			continue
		}

		b.Run(tc.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				All(tc.RRule.Iterator(), 0)
			}
		})
	}
}

func rfcAll(times []time.Time) []string {
	strs := make([]string, len(times))
	for i, t := range times {
		strs[i] = t.Format(time.RFC3339)
	}
	return strs
}
