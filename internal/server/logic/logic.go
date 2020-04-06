package logic

var pastFourScores = []float64{5.0, 4.0, 3.0, 2.0}

func GetSize() float64 {
	oldScores := pastFourScores[0] + pastFourScores[1]
	newScores := pastFourScores[2] + pastFourScores[3]
	diff := newScores - oldScores
	if diff > 0 {
		size := 600 + diff*60.0
		if size < 2000 {
			return size
		}
		return 2000.0
	}
	if diff > -5.0 && dif <= 0 {
		return 100.0 + 18.0*diff
	}
	return 10.0
}

func SetScore() bool {
	pastFourScores = append(pastFourScores, a)
	pastFourScores = pastFourScores[1:]
	return true
}
