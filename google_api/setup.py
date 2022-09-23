from setuptools import setup
from Cython.Build import cythonize
import sys

setup(
    ext_modules = cythonize("main.pyx")
)