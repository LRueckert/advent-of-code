import unittest
import main

class ResultTestcase(unittest.TestCase):
    def test_getResultA(self):
        file = "testA"
        self.assertEqual(main.getResult("A", file), 1514)


    def test_getResultB(self):
        file = "testB"
        self.assertEqual(main.decryptRoom("qzmt-zixmtkozy-ivhz-343"), "very encrypted name")
