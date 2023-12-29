package helpers

func GetDummyData(index int) (string, int, string) {
	names := []string{
		"Butterfly Stretch",
		"Malasana/Yogi Squat",
		"Happy Baby Pose",
		"Reclined Bound Angle Pose",
		"Frog Pose",
		"Supine Figure Four",
		"Half Pigeon",
		"Double Pigeon",
		"Low Lunge",
		"Crescent Lunge",
		"Camel Pose",
		"Dancer Pose",
		"Supported Back Bend",
		"Supported Bridge",
		"Hip Pry",
		"Hero Pose With Block",
	}

	duration := []int{
		20,
		30,
		40,
		20,
		30,
		50,
		30,
		50,
		40,
		40,
		30,
		30,
		40,
		60,
		30,
		50,
	}

	pictures := []string{
		"image/yoga-1.jpg",
		"image/yoga-2.jpg",
		"image/yoga-3.jpg",
		"image/yoga-4.jpg",
		"image/yoga-5.jpg",
		"image/yoga-6.jpg",
		"image/yoga-7.jpg",
		"image/yoga-8.jpg",
		"image/yoga-9.jpg",
		"image/yoga-10.jpg",
		"image/yoga-11.jpg",
		"image/yoga-12.jpg",
		"image/yoga-13.jpg",
		"image/yoga-14.jpg",
		"image/yoga-15.jpg",
		"image/yoga-16.jpg",
	}

	return names[index], duration[index], pictures[index]
}
