import requests

list = [
    "12010270623176500",
    "12010270617506496",
    "12010270628526502",
    "12010270628596503",
    "12010270702116506",
    "12010270702466508",
    "12010270704006512",
    "12010270727426516",
    "12010270728286518",
    "12010270623276501",
    "12010270702026505",
    "12010270702216507",
    "12010270703386510",
    "12010270717346514",
    "12010270741026520",
    "12010270804286530",
    "12010270804396531",
    "12010270814266542",
    "12010270803306527",
    "12010270814126541",
    "12010270814576543",
    "12010270815136544",
    "12010270823216553",
    "12010270818566548",
    "12010270831556558",
    "12010270836116561",
    "12010270842106574",
    "12010270842186575",
    "12010270843096576",
    "12010270836456564",
]


def main():
    for trade_no in list:
        url = (
            "https://base.url/cm_prod/maintain/fix?trade_no=%s" % trade_no)
        payload = {}
        headers = {
            'token': 'klbbPKSON7OiusF9ZaJEreHq2Tfp3P4I32NWJ2fmqvrBDODRxti2ISjQ2pb0Kj0jqTPtrLyyjyAiqaoSGMzG5G17ABxvnkGiTg1kVDLXgxo=',
            'Authorization': 'Basic YWRtaW46YWRtaW4=',
            'Cookie': 'BIGipServerbase.url-2443=293730826.35593.0000'
        }
        response = requests.request("GET", url, headers=headers, data=payload)
        print("%s,%s" % (trade_no, response.text.encode('utf8')))


if __name__ == "__main__":
    main()
