import json
from pprint import pprint
from typing import List

import config
import generate_user
import csv


def read_in_col(path: str) -> List:
    names: List
    with open(path) as f:
        csvreader = csv.reader(f)
        _names = list(csvreader)

    names: List = []
    for namel in _names:
        names.append(namel[0])
    return names


if __name__ == "__main__":
    conf = config.parse_config("./mock-data-config.json")

    fnames = read_in_col("fnames.csv")
    lnames = read_in_col("lnames.csv")
    pictures = read_in_col("image_urls.txt")
    gen = generate_user.UserGenerator(conf)
    print()
    print(json.dumps(gen.generate(fnames, lnames, profile_pics=pictures), indent=2))
    print()
    print(json.dumps(gen.generate(fnames, lnames, profile_pics=pictures), indent=2))
    print()
    print(json.dumps(gen.generate(fnames, lnames, profile_pics=pictures), indent=2))
    print()
    print(json.dumps(gen.generate(fnames, lnames, profile_pics=pictures), indent=2))

