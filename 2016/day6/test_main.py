import unittest
import main

class ResultTestcase(unittest.TestCase):
    def test_getResultA(self):
        file = "testA"
        self.assertEqual(main.getResult("A", file), "easter")


    def test_getResultB(self):
        file = "testB"
        self.assertEqual(main.getResult("B", file), "advent")
