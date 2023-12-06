package internal

func Part1(filename string) (*Challenge, *int) {
	Input := Challenge{}
	Input.init()

	multiplied := 1
	for _, r := range Input.Races {
		r.getMinMaxMilliseconds()
		multiplied *= r.TotalPossibleWays
	}

	return &Input, &multiplied
}

func Part2(Input *Challenge) *int {
	lowest := -1
	return &lowest
}

type Challenge struct {
	Races []Race
}

func (c *Challenge) init() {
	c.Races = []Race{
		{Duration: 44, DistanceRecord: 208},
		{Duration: 80, DistanceRecord: 1581},
		{Duration: 65, DistanceRecord: 1050},
		{Duration: 72, DistanceRecord: 1102},
		//{Duration: 7, DistanceRecord: 9},
		//{Duration: 15, DistanceRecord: 40},
		//{Duration: 30, DistanceRecord: 200},
	}
}

type Race struct {
	Duration          int
	DistanceRecord    int
	MinSpeed          int
	MaxSpeed          int
	TotalPossibleWays int
}

func (r *Race) getMinMaxMilliseconds() {
	r.MinSpeed = -1
	r.MaxSpeed = -1
	for i := 0; i <= r.Duration; i++ {
		speed := i
		remainingDur := r.Duration - i
		distanceTRaveled := speed * remainingDur
		if distanceTRaveled <= r.DistanceRecord {
			// continue to next iteration if current doesn't beat record
			continue
		}
		r.TotalPossibleWays += 1
		if r.MinSpeed == -1 || speed < r.MinSpeed {
			// update minimum speed
			r.MinSpeed = speed
		}

		if r.MaxSpeed == -1 || speed > r.MaxSpeed {
			// update maximum speed
			r.MaxSpeed = speed
		}
	}
	//r.TotalPossibleWays = r.MaxSpeed - r.MinSpeed
}
