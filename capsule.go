package protometry


func (c *Capsule) Fit(other Volume) bool {
	return false
}

func (c *Capsule) Intersects(other Volume) bool {
	return false
}

func (c *Capsule) Average(other Volume) Volume {
	return nil
}

func (c *Capsule) Mutate(rate float64) Volume {
	return nil
}
