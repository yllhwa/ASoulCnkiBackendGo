with open("test.des", "rb") as f:
    while True:
        key = int.from_bytes(f.read(6), "big")
        off = int.from_bytes(f.read(6), "big")