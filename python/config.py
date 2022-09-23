import json
import dateutil.parser
import pprint


from typing import Dict

def parse(data: Dict, indent = 0) -> Dict:
    # where possible, parse time
    for key, val in data.items():
        if type(val) is dict:
            data[key] = parse(val, indent + 2)
        # noinspection PyBroadException
        try:
            parsed_time = dateutil.parser.isoparse(val)
            data[key] = parsed_time
            print(f'{" " * indent}Parsed "{key}" as datetime in ISO8601 format')
        except:
            ...
    print(f'Config parsed as:\n{pprint.pformat(data)}')
    return data

def parse_config(path: str) -> Dict:
    with open(path) as f:
        data = json.load(f)
        return parse(data)

