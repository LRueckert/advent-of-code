import unittest
import main

class ResultTestcase(unittest.TestCase):
    def test_getResultA(self):
        input = "abc"
        self.assertEqual(main.getResult("A", input), "18f47a30")


    def test_getResultB(self):
        input = "abc"
        self.assertEqual(main.getResult("B", input), "05ace8e3")
