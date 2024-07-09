soal 1

1.
CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255)
);

2.
INSERT INTO users (id, name, email) VALUES
(1, 'John Doe', 'johndoe@example.com'),
(2, 'Jane Smith', 'janesmith@example.com'),
(3, 'Bob Johnson', 'bobjohnson@example.com');

3.
CREATE TABLE orders (
    id INT PRIMARY KEY,
    user_id INT,
    amount DECIMAL(10,2),
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

4.
INSERT INTO orders (id, user_id, amount, created_at) VALUES
(1, 1, 100.00, '2022-01-02 10:30:00'),
(2, 2, 50.00, '2022-01-03 09:00:00'),
(3, 1, 150.00, '2022-01-04 14:15:00'),
(4, 3, 200.00, '2022-01-05 17:45:00'),
(5, 2, 75.00, '2022-01-06 11:20:00');

5.
CREATE TABLE order_items (
    id INT PRIMARY KEY,
    order_id INT,
    product_name VARCHAR(255),
    price DECIMAL(10,2),
    quantity INT,
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

6.
INSERT INTO order_items (id, order_id, product_name, price, quantity) VALUES
(1, 1, 'T-Shirt', 25.00, 2),
(2, 1, 'Jeans', 50.00, 1),
(3, 2, 'Socks', 10.00, 5),
(4, 3, 'Shoes', 75.00, 2),
(5, 4, 'Jacket', 100.00, 1),
(6, 5, 'Sweater', 25.00, 3);

--
SELECT u.name AS user_name, SUM(oi.price * oi.quantity) AS total_spent
FROM users u
JOIN orders o ON u.id = o.user_id
JOIN order_items oi ON o.id = oi.order_id
WHERE o.created_at >= '2022-01-01'
GROUP BY u.name
HAVING SUM(oi.price * oi.quantity) >= 1000;


-- soal ke 2

1. 
CREATE DATABASE IF NOT EXISTS first;
USE first;

2.
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255)
);

3.
INSERT INTO users (name, email) VALUES 
('John Doe', 'john.doe@example.com'),
('Jane Smith', 'jane.smith@example.com');

4.
CREATE DATABASE IF NOT EXISTS second;
USE second;

5.
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT,
    amount NUMERIC(10,2),
    created_at TIMESTAMP
);

6.
INSERT INTO orders (user_id, amount, created_at) VALUES 
(1, 100.00, NOW()),
(2, 200.00, NOW()),
(1, 50.00, NOW());

--
SELECT u.name AS user_name, o.amount, o.created_at
FROM first.users u
JOIN second.orders o ON u.id = o.user_id
ORDER BY u.id, o.created_at;


soal 3

1.
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	
	numsMap := make(map[int]int)

	
	for i, num := range nums {
	
		complement := target - num

		
		if idx, found := numsMap[complement]; found {
			
			return []int{idx, i}
		}

		
		numsMap[num] = i
	}

	
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println("Output:", result)
}

2.
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println("Output:", result)
}


3.
package main

import (
	"fmt"
	"strings"
)

func findSubstring(s string, words []string) []int {
	var result []int
	wordCount := len(words)
	if wordCount == 0 {
		return result
	}

	wordLength := len(words[0])
	totalLength := wordCount * wordLength

	if len(s) < totalLength {
		return result
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i <= len(s)-totalLength; i++ {
		substr := s[i : i+totalLength]
		if matchSubstring(substr, wordMap, wordLength) {
			result = append(result, i)
		}
	}

	return result
}

func matchSubstring(substr string, wordMap map[string]int, wordLength int) bool {
	tempMap := make(map[string]int)

	for i := 0; i < len(substr); i += wordLength {
		word := substr[i : i+wordLength]
		if count, found := wordMap[word]; found {
			tempMap[word]++
			if tempMap[word] > count {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "word"}

	result := findSubstring(s, words)
	fmt.Println("Output:", result)
}

