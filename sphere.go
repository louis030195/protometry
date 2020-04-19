package protometry

func (s *Sphere) Fit(other Volume) bool {
	return false
}

func (s *Sphere) Intersects(other Volume) bool {
	return false
}

func (s *Sphere) Average(other Volume) Volume {
	return nil
}

func (s *Sphere) Mutate(rate float64) Volume {
	return nil
}
