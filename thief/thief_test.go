package thief

import (
	"fmt"
	"testing"
)

func TestCorrectCheckJsonFilter(t *testing.T) {
	var example = []byte(`{"is_cyber": true, "warning_type": "WARNING_AVERT_EYES",
	"warning_start": "2025-05-05T20:22:11.279Z",
	"positions": [
		[
		[
			209.43679046630862,
			149.66827011108398
		],
		[
			526.738899230957,
			525.7927265167236
		]
		]
	],
	"image": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAoAAAAHgCAYAAAA10dzkAAAAAXNSR0IArs4c6QAAIABJREFUeF6MvQmPZWlyHRZ3e0tuVVlb19bV6/RMT3OGs4gjmuJAAmxTFEyQBgiCAmwLpuAFhA3Zf9AQYMOmZRKgRHIWTk+vVV3dtWUtub337maccyK+777spq0c1GRX1sv37v3u90WcOHEiovhf/9UfjserE3ty8sxenh/bpu9t0/U2dqXZUFrRl1b2Zk1R2HJe2e2712x5MLPZXmM7+ztWV5VVZWWVlVaMpZVjaQ3/N7faatuct3a23lhblnb/8WP79acPbNN11nYrK8rBqnK0sRhttMLGsTAbSxvH0fqhNSt6w4eXtVlRmNk4WjEWVgyV1UNtu83Sduu5Xd6d2c3ru7a3W9lit7L9yztmpZmVhQ3jaG3X4pf5390wWDtU9vzlyr569MrOV2Zt11vXb2wYNjZax/+uq9L6vrWx76wuce+1Xbu6Zzdv7dtyt7Cq6axuKqurmZXFjhXjnpXjjlXjwoqxsqHvrSgGK8vRur63th/tfD3ap589sadPTmwcKms3A+/JRuNX3/c2joMVRWlV1dh8vrR",
	"ХУЙ": "asdasd"
  }`)

	flag, err := CheckJsonFilter(example)

	fmt.Printf("Результат фильтра: %v\n", flag)

	if err != nil {
		t.Errorf("Тест не пройден, ошибка при десериализации")
	}

	if !flag {
		t.Errorf("Тест не пройден, результаты не равны")

	}

}

func TestIncorrectJsonFilter(t *testing.T) {
	var example = []byte(`{"normal": true, "testovka": "Проверка функции не по нужным полям"}`)

	flag, err := CheckJsonFilter(example)

	fmt.Printf("Результат фильтра: %v\n", flag)

	if err != nil {
		t.Errorf("Тест не пройден, ошибка при десериализации")
	}

	if flag {
		t.Errorf("Тест не пройден, результаты не равны")
	}

}
