package model

import "sort"

type Sample struct {
	ID    string
	Name  string
	Value float64
}

type Summary struct {
	Count  int
	Sum    float64
	Max    float64
	Min    float64
	Failed int
}

func ValidSample(s *Sample) bool {
	return s != nil && s.ID != "" && s.Name != ""
}

func SortSamples(ss []*Sample) []*Sample {
	sort.SliceStable(ss, func(i, j int) bool { return ss[i].ID < ss[j].ID })
	return ss
}

func BuildBuckets(ss []*Sample, size int) [][]*Sample {
	if size <= 0 {
		size = 1
	}
	out := make([][]*Sample, 0, (len(ss)+size-1)/size)
	for i := 0; i < len(ss); i += size {
		end := i + size
		if end > len(ss) {
			end = len(ss)
		}
		b := make([]*Sample, end-i)
		copy(b, ss[i:end])
		out = append(out, b)
	}
	return out
}

func MergeSummary(dst Summary, src Summary) Summary {
	if src.Count == 0 && src.Failed == 0 {
		return dst
	}
	if dst.Count == 0 {
		dst.Max = src.Max
		dst.Min = src.Min
	} else {
		if src.Max > dst.Max {
			dst.Max = src.Max
		}
		if src.Min < dst.Min {
			dst.Min = src.Min
		}
	}
	dst.Count += src.Count
	dst.Sum += src.Sum
	dst.Failed += src.Failed
	return dst
}
