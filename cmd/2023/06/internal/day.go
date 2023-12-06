package internal

func Part1() *int {
	Input := Challenge{}
	Input.initPart1()

	multiplied := 1
	for _, r := range Input.Races {
		r.getMinMaxMilliseconds()
		multiplied *= r.TotalPossibleWays
	}

	return &multiplied
}

func Part2() *int {
	Input := Challenge{}
	Input.initPart2()

	multiplied := 1
	for _, r := range Input.Races {
		r.getMinMaxMilliseconds()
		multiplied *= r.TotalPossibleWays
	}

	return &multiplied
}

type Challenge struct {
	Races []Race
}

func (c *Challenge) initPart1() {
	c.Races = []Race{
		// Part 1 race numbers
		{Duration: 44, DistanceRecord: 208},
		{Duration: 80, DistanceRecord: 1581},
		{Duration: 65, DistanceRecord: 1050},
		{Duration: 72, DistanceRecord: 1102},
		// Part 1 race number examples
		//{Duration: 7, DistanceRecord: 9},
		//{Duration: 15, DistanceRecord: 40},
		//{Duration: 30, DistanceRecord: 200},
	}
}

func (c *Challenge) initPart2() {
	c.Races = []Race{
		// Part 2 race number
		{Duration: 44806572, DistanceRecord: 208158110501102},
		// Part 2 race number example
		//{Duration: 71530, DistanceRecord: 940200},
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
