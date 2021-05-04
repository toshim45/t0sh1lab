from collections import defaultdict

INVALID_BIN_SHELF = "Format kode bin/rak salah"
SHELF_NOT_EXIST = "Bin asal belum mempunyai rak"
invalid_bin_test_case = dict(
        positive_test_invalid_bin_format=["IDCW1-BOX-07663", INVALID_BIN_SHELF],
        positive_test_invalid_shelf_format=["IDCW1-2-AU-3", INVALID_BIN_SHELF],
        positive_test_bin_has_no_shelf=["IDCW1-BOX-043548", SHELF_NOT_EXIST],
    )
cases = invalid_bin_test_case.items()
test_name = cases[0]
bin_input = cases[1][0]
expected_message = cases[1][1]

print("test-name: {0} \nbin-input:{1} \nexpected-message:{2}\n".format(test_name, bin_input, expected_message))

zones = defaultdict(dict)
points = [(1,'A'),(2,'B'),(3,'A')]

for i,p in enumerate(points):
   zones[p[1]][p[0]]=p

print('zones: ', zones)
