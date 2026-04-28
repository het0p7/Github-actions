import unittest

class TestBasicLogic(unittest.TestCase):
    
    def test_basic_math(self):
        """Test for basic arithmetic (2+2=4)"""
        self.assertEqual(2 + 2, 4)

    def test_hello_message(self):
        """Test for a simple string comparison"""
        message = "Hello"
        self.assertEqual(message, "Hello")

    def test_multiplication(self):
        """Test for multiplication (5*5=25)"""
        self.assertEqual(5 * 5, 25)

    def test_division(self):
        """Test for division (10/2=5)"""
        self.assertEqual(10 / 2, 5)

    def test_string_concatenation(self):
        """Test for string concatenation"""
        s1 = "Hello"
        s2 = "World"
        self.assertEqual(s1 + " " + s2, "Hello World")

    def test_list_length(self):
        """Test for list length"""
        items = ["apple", "banana", "cherry"]
        self.assertEqual(len(items), 3)

if __name__ == '__main__':
    unittest.main()
