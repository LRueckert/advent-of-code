import unittest
import main


class ResultTestcase(unittest.TestCase):
    def test_getResultA(self):
        self.assertEqual(main.calculateResultA("ADVENT"), 6)
        self.assertEqual(main.calculateResultA("A(1x5)BC"), 7)
        self.assertEqual(main.calculateResultA("(3x3)XYZ"), 9)
        self.assertEqual(main.calculateResultA("A(2x2)BCD(2x2)EFG"), 11)
        self.assertEqual(main.calculateResultA("(6x1)(1x3)A"), 6)
        self.assertEqual(main.calculateResultA("X(8x2)(3x3)ABCY"), 18)

    def test_getResultB(self):
        file = "testB"
        self.assertEqual(main.calculateResultB("(3x3)XYZ"), 9)
        self.assertEqual(main.calculateResultB("X(8x2)(3x3)ABCY"), 20)
        self.assertEqual(main.calculateResultB("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"), 445)
        self.assertEqual(main.calculateResultB("(27x12)(20x12)(13x14)(7x10)(1x12)A"), 241920)
