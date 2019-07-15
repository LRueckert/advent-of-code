import unittest
import main


class ResultTestcase(unittest.TestCase):
    def test_getResultA(self):
        self.assertEqual(main.calculateResultA(["R2", "L3"]), 5)
        self.assertEqual(main.calculateResultA(["R2", "R2", "R2"]), 2)
        self.assertEqual(main.calculateResultA(["R5", "L5", "R5", "R3"]), 12)

    def test_getResultB(self):
        self.assertEqual(main.calculateResultB(["R8", "R4", "R4", "R8"]), 4)
