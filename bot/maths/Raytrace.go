package maths

func RayTraceBlocks(start, end Vec3d[float64]) []Vec3d[float64] {
	var result []Vec3d[float64]
	diff := end.Sub(start)
	distance := diff.Length()
	if distance == 0 {
		return result
	}
	for i := 0; i < int(distance); i++ {
		result = append(result, start.Add(diff.MulScalar(float64(i)/distance)))
	}
	return result
}
